package wallet

import (
	"fmt"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
)

// testDB connects a throwaway sqlite file DB (fireback forces a single connection for
// sqlite - see DatabaseConnection.go's CreateDatabasePool - so this also exercises how
// applyLedgerEntry behaves when concurrent callers contend for that one connection, not
// just for the wallet row) and AutoMigrates every wallet entity, bypassing the full
// CLI/HTTP auth stack entirely - these tests call applyLedgerEntry/Purchase directly,
// the same engine every action (purchase/adjustBalance/gatewayWebhook) shares.
// DirectConnectToDb also sets fireback's package-global db ref, so fireback.GetDbRef()
// (which Purchase() itself calls) resolves to this same connection.
func testDB(t *testing.T) *gorm.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "wallet-test.db")
	db, err := fireback.DirectConnectToDb(fireback.Config{
		DbVendor:   "sqlite",
		DbName:     path,
		DbLogLevel: "silent",
	})
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}
	// DirectConnectToDb alone doesn't apply the sqlite-specific pragmas
	// CreateDatabasePool normally would (single connection + busy_timeout) - without
	// these, concurrent writers get an immediate "database is locked" error instead of
	// queueing, which the concurrency test below would otherwise misreport as a
	// business-logic failure.
	if sqlDb, err := db.DB(); err == nil {
		sqlDb.SetMaxOpenConns(1)
	}
	db.Exec("PRAGMA busy_timeout = 5000")

	// AutoMigrate can't be used here: every entity's generated UniqueId column
	// declares `default:gen_random_uuid()` (a Postgres/pgcrypto function - see any
	// generated *Entity.go, this isn't wallet-specific), which isn't valid SQLite
	// DEFAULT syntax and fails at CREATE TABLE time. This repo's entity modules are
	// evidently only ever exercised against real Postgres (see docker-compose.yml) -
	// sqlite is a supported fireback *driver*, but not something these generated
	// entities are portable to as-is. So: hand-create just the two tables these tests
	// touch, with SQLite-safe DDL (no server-side default - UniqueId is set in Go
	// before every insert instead), matching GORM's default snake_case naming exactly.
	ddl := []string{
		`CREATE TABLE wallet_entities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			unique_id TEXT UNIQUE DEFAULT (lower(hex(randomblob(16)))),
			owner_type TEXT,
			user_id TEXT,
			workspace_id TEXT,
			currency TEXT,
			balance TEXT,
			status TEXT,
			label TEXT,
			is_default BOOLEAN,
			version INTEGER
		)`,
		`CREATE TABLE wallet_transaction_entities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			unique_id TEXT UNIQUE DEFAULT (lower(hex(randomblob(16)))),
			direction TEXT,
			amount TEXT,
			balance_after TEXT,
			reason TEXT,
			reference_type TEXT,
			reference_id TEXT,
			idempotency_key TEXT UNIQUE,
			note TEXT,
			created_by TEXT,
			metadata TEXT,
			created_at DATETIME,
			wallet_id INTEGER
		)`,
	}
	for _, stmt := range ddl {
		if err := db.Exec(stmt).Error; err != nil {
			t.Fatalf("create table: %v\n%s", err, stmt)
		}
	}
	return db
}

func mustCreateWallet(t *testing.T, db *gorm.DB, balance string) *walletdefs.WalletEntity {
	t.Helper()
	w := &walletdefs.WalletEntity{
		// The test schema has no server-side default for unique_id (see testDB's
		// comment) - set it explicitly, same effect a real Postgres gen_random_uuid()
		// default would have.
		UniqueId:  uuid.NewString(),
		OwnerType: "user",
		Currency:  "USD",
		Balance:   balance,
		Status:    "active",
	}
	w.UserId.Set(strPtr("u1"))
	created, err := walletdefs.WalletEntityActions.Create(db, w)
	if err != nil {
		t.Fatalf("create wallet: %v", err)
	}
	return created
}

func TestPurchase_SequentialDebitsAndInsufficientBalance(t *testing.T) {
	db := testDB(t)
	w := mustCreateWallet(t, db, "10000")

	if _, err := applyLedgerEntry(db, ledgerChange{
		WalletUniqueId: w.UniqueId, Direction: "debit", Amount: "2500",
		Reason: "purchase", IdempotencyKey: "k1",
	}); err != nil {
		t.Fatalf("first purchase: %v", err)
	}
	if _, err := applyLedgerEntry(db, ledgerChange{
		WalletUniqueId: w.UniqueId, Direction: "debit", Amount: "2500",
		Reason: "purchase", IdempotencyKey: "k2",
	}); err != nil {
		t.Fatalf("second purchase: %v", err)
	}

	var reloaded walletdefs.WalletEntity
	if err := db.First(&reloaded, "unique_id = ?", w.UniqueId).Error; err != nil {
		t.Fatalf("reload wallet: %v", err)
	}
	if reloaded.Balance != "5000" {
		t.Fatalf("balance = %q, want %q", reloaded.Balance, "5000")
	}

	// Insufficient balance must be rejected and must NOT touch the balance.
	_, err := applyLedgerEntry(db, ledgerChange{
		WalletUniqueId: w.UniqueId, Direction: "debit", Amount: "999999",
		Reason: "purchase", IdempotencyKey: "k3",
	})
	if err == nil {
		t.Fatalf("expected insufficient balance error, got nil")
	}
	ierr, ok := err.(*fireback.IError)
	if !ok || ierr.Message["$"] != "wallet.errors.insufficientBalance" {
		t.Fatalf("expected insufficientBalance IError, got %#v", err)
	}
	if err := db.First(&reloaded, "unique_id = ?", w.UniqueId).Error; err != nil {
		t.Fatalf("reload wallet: %v", err)
	}
	if reloaded.Balance != "5000" {
		t.Fatalf("balance changed after rejected purchase: got %q, want %q", reloaded.Balance, "5000")
	}
}

func TestPurchase_IdempotentRetryDoesNotDoubleDebit(t *testing.T) {
	db := testDB(t)
	w := mustCreateWallet(t, db, "1000")

	first, err := applyLedgerEntry(db, ledgerChange{
		WalletUniqueId: w.UniqueId, Direction: "debit", Amount: "100",
		Reason: "purchase", IdempotencyKey: "retry-me",
	})
	if err != nil {
		t.Fatalf("first call: %v", err)
	}
	second, err := applyLedgerEntry(db, ledgerChange{
		WalletUniqueId: w.UniqueId, Direction: "debit", Amount: "100",
		Reason: "purchase", IdempotencyKey: "retry-me",
	})
	if err != nil {
		t.Fatalf("retried call: %v", err)
	}
	if first.UniqueId != second.UniqueId {
		t.Fatalf("retry created a new ledger entry: first=%s second=%s", first.UniqueId, second.UniqueId)
	}

	var reloaded walletdefs.WalletEntity
	if err := db.First(&reloaded, "unique_id = ?", w.UniqueId).Error; err != nil {
		t.Fatalf("reload wallet: %v", err)
	}
	if reloaded.Balance != "900" {
		t.Fatalf("balance = %q, want %q (must be debited exactly once)", reloaded.Balance, "900")
	}
}

// TestPurchase_ConcurrentPurchasesNeverOverdraw is the concurrency test the wallet
// system exists to pass: N goroutines race to debit 100 from a wallet that can only
// afford 10 of them. Exactly 10 must succeed, the rest must fail with insufficient
// balance, the final balance must land at exactly 0 (never negative, never short), and
// the ledger must contain exactly one row per successful debit - no lost updates, no
// double-spends, regardless of goroutine scheduling.
func TestPurchase_ConcurrentPurchasesNeverOverdraw(t *testing.T) {
	db := testDB(t)
	w := mustCreateWallet(t, db, "1000")

	const attempts = 25
	const amount = 100

	var succeeded, failed int64
	var wg sync.WaitGroup
	wg.Add(attempts)
	for i := 0; i < attempts; i++ {
		i := i
		go func() {
			defer wg.Done()
			_, err := applyLedgerEntry(db, ledgerChange{
				WalletUniqueId: w.UniqueId,
				Direction:      "debit",
				Amount:         fmt.Sprintf("%d", amount),
				Reason:         "purchase",
				IdempotencyKey: fmt.Sprintf("concurrent-%d", i),
			})
			if err != nil {
				atomic.AddInt64(&failed, 1)
			} else {
				atomic.AddInt64(&succeeded, 1)
			}
		}()
	}
	wg.Wait()

	if succeeded != 10 {
		t.Fatalf("succeeded = %d, want 10", succeeded)
	}
	if failed != attempts-10 {
		t.Fatalf("failed = %d, want %d", failed, attempts-10)
	}

	var reloaded walletdefs.WalletEntity
	if err := db.First(&reloaded, "unique_id = ?", w.UniqueId).Error; err != nil {
		t.Fatalf("reload wallet: %v", err)
	}
	if reloaded.Balance != "0" {
		t.Fatalf("final balance = %q, want %q (overdraft or lost update under concurrency)", reloaded.Balance, "0")
	}

	var ledgerCount int64
	db.Model(&walletdefs.WalletTransactionEntity{}).Where("wallet_id = ?", reloaded.Id).Count(&ledgerCount)
	if ledgerCount != 10 {
		t.Fatalf("ledger row count = %d, want 10 (one per successful debit)", ledgerCount)
	}
}

func TestPurchase_ExportedFunctionMatchesEngine(t *testing.T) {
	db := testDB(t)
	w := mustCreateWallet(t, db, "500")

	entry, err := Purchase(PurchaseInput{
		WalletUniqueId: w.UniqueId,
		Amount:         "150",
		ReferenceType:  "course-purchase",
		ReferenceId:    "course-1",
		IdempotencyKey: "exported-fn-1",
	})
	if err != nil {
		t.Fatalf("Purchase: %v", err)
	}
	if entry.Reason != "purchase" || entry.Direction != "debit" || entry.BalanceAfter != "350" {
		t.Fatalf("unexpected ledger entry: %+v", entry)
	}
}
