package wallet

import (
	"errors"
	"net/http"
	"time"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/complexes"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Purchase.go is the one place a wallet's balance is ever mutated. Every path that moves
// money - the purchase/adjustBalance HTTP actions, a topup's gatewayWebhook credit, and
// any other nima module's in-process call - goes through applyLedgerEntry, so there is
// exactly one concurrency-safe, idempotent, precision-safe implementation to reason about.

// PurchaseInput is the argument to Purchase - the in-process Go API other nima modules
// import and call directly to debit a wallet for a sale, with no HTTP round-trip needed.
// This is also what the purchase HTTP action (PurchaseImplementation.go) calls under the
// hood - both share this one implementation.
type PurchaseInput struct {
	WalletUniqueId string
	Amount         string // minor-units decimal string, see Money.go
	ReferenceType  string // free-form caller identity, e.g. "course-purchase"
	ReferenceId    string
	IdempotencyKey string
	Note           string
	CreatedBy      string // userId of the actor, or "" for a pure-system call
}

// Purchase debits WalletUniqueId by Amount, recording a walletTransaction with
// reason="purchase". Safe to call concurrently against the same wallet (a DB row lock
// serializes concurrent purchases - see applyLedgerEntry) and safe to retry with the same
// IdempotencyKey (a retry returns the original ledger entry instead of double-debiting).
// Fails with a *fireback.IError (insufficient balance / wallet not active / not found) if
// the debit can't be applied.
func Purchase(in PurchaseInput) (*walletdefs.WalletTransactionEntity, error) {
	return applyLedgerEntry(fireback.GetDbRef(), ledgerChange{
		WalletUniqueId: in.WalletUniqueId,
		Direction:      "debit",
		Amount:         in.Amount,
		Reason:         "purchase",
		ReferenceType:  in.ReferenceType,
		ReferenceId:    in.ReferenceId,
		IdempotencyKey: in.IdempotencyKey,
		Note:           in.Note,
		CreatedBy:      in.CreatedBy,
	})
}

// ledgerChange is the internal, direction-agnostic argument to applyLedgerEntry. Exported
// entry points (Purchase, plus AdjustBalanceAction/gatewayWebhook's credit) each fix
// Direction/Reason and forward the rest.
type ledgerChange struct {
	WalletUniqueId string
	Direction      string // "credit" or "debit"
	Amount         string
	Reason         string
	ReferenceType  string
	ReferenceId    string
	IdempotencyKey string
	Note           string
	CreatedBy      string
}

// applyLedgerEntry is the single concurrency-safe, idempotent mutation path for a
// wallet's balance:
//  1. If a walletTransaction with this IdempotencyKey already exists, return it unchanged
//     (safe retry) instead of re-applying the change.
//  2. Inside one DB transaction, take a row lock on the wallet (SELECT ... FOR UPDATE) so
//     concurrent calls against the *same* wallet serialize instead of racing.
//  3. Validate status/amount using math/big (Money.go) - never float - and compute the
//     new balance.
//  4. Update the wallet's balance and insert the walletTransaction row together, in the
//     same transaction, so the ledger and the balance can never drift apart.
//  5. If two concurrent callers both raced past step 1 with the same IdempotencyKey, the
//     unique index on walletTransaction.idempotencyKey rejects the second insert; that
//     case is caught and turned back into "return the existing row" too.
func applyLedgerEntry(db *gorm.DB, change ledgerChange) (*walletdefs.WalletTransactionEntity, error) {
	if existing, err := findByIdempotencyKey(db, change.IdempotencyKey); err != nil {
		return nil, err
	} else if existing != nil {
		return existing, nil
	}

	var result *walletdefs.WalletTransactionEntity
	txErr := db.Transaction(func(tx *gorm.DB) error {
		var w walletdefs.WalletEntity
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&w, "unique_id = ?", change.WalletUniqueId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return walletNotFoundError(change.WalletUniqueId)
			}
			return err
		}

		if w.Status != "active" {
			return walletNotActiveError(w.Status)
		}

		var newBalance string
		switch change.Direction {
		case "credit":
			nb, err := AddAmounts(w.Balance, change.Amount)
			if err != nil {
				return invalidAmountError(err)
			}
			newBalance = nb
		case "debit":
			nb, ok, err := SubAmounts(w.Balance, change.Amount)
			if err != nil {
				return invalidAmountError(err)
			}
			if !ok {
				return insufficientBalanceError(w.Balance, change.Amount)
			}
			newBalance = nb
		default:
			return &fireback.IError{
				HttpCode: http.StatusInternalServerError,
				Message:  fireback.ErrorItem{"en": "invalid ledger direction: " + change.Direction},
			}
		}

		if err := tx.Model(&w).Updates(map[string]any{
			"balance": newBalance,
			"version": w.Version + 1,
		}).Error; err != nil {
			return err
		}

		entry := &walletdefs.WalletTransactionEntity{
			WalletId:       w.Id,
			Direction:      change.Direction,
			Amount:         change.Amount,
			BalanceAfter:   newBalance,
			Reason:         change.Reason,
			IdempotencyKey: change.IdempotencyKey,
			CreatedAt:      complexes.XDate(time.Now().Format("2006-01-02")),
		}
		if change.ReferenceType != "" {
			entry.ReferenceType = emigo.NullableOf(change.ReferenceType)
		}
		if change.ReferenceId != "" {
			entry.ReferenceId = emigo.NullableOf(change.ReferenceId)
		}
		if change.Note != "" {
			entry.Note = emigo.NullableOf(change.Note)
		}
		if change.CreatedBy != "" {
			entry.CreatedBy = emigo.NullableOf(change.CreatedBy)
		}
		if err := tx.Create(entry).Error; err != nil {
			return err
		}
		result = entry
		return nil
	})

	if txErr != nil {
		// Belt-and-suspenders: two concurrent callers can both pass the step-1 check
		// with the same IdempotencyKey before either commits: the loser's insert fails
		// the unique constraint on walletTransaction.idempotencyKey. Re-check for the
		// now-committed winner's row instead of surfacing that as a hard failure - this
		// intentionally doesn't inspect the driver-specific error code (Postgres/MySQL/
		// SQLite all phrase it differently), it just trusts the unique index and
		// re-queries.
		if existing, findErr := findByIdempotencyKey(db, change.IdempotencyKey); findErr == nil && existing != nil {
			return existing, nil
		}
		if ierr, ok := txErr.(*fireback.IError); ok {
			return nil, ierr
		}
		return nil, fireback.GormErrorToIError(txErr)
	}
	return result, nil
}

func findByIdempotencyKey(db *gorm.DB, idempotencyKey string) (*walletdefs.WalletTransactionEntity, error) {
	if idempotencyKey == "" {
		return nil, &fireback.IError{
			HttpCode: http.StatusBadRequest,
			Message:  fireback.ErrorItem{"en": "idempotencyKey is required"},
		}
	}
	var existing walletdefs.WalletTransactionEntity
	err := db.First(&existing, "idempotency_key = ?", idempotencyKey).Error
	if err == nil {
		return &existing, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

func walletNotFoundError(walletId string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusNotFound,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.notFound",
			"en": "Wallet not found.",
			"fa": "کیف پول یافت نشد.",
			"ru": "Кошелёк не найден.",
			"pl": "Nie znaleziono portfela.",
		},
		MessageParams: map[string]any{"walletId": walletId},
	}
}

func walletNotActiveError(status string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusConflict,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.notActive",
			"en": "This wallet is not active and cannot be topped up or spent from.",
			"fa": "این کیف پول فعال نیست و امکان شارژ یا برداشت از آن وجود ندارد.",
			"ru": "Этот кошелёк неактивен, пополнение и списание недоступны.",
			"pl": "Ten portfel jest nieaktywny i nie można go zasilić ani z niego wydawać.",
		},
		MessageParams: map[string]any{"status": status},
	}
}

func insufficientBalanceError(balance, amount string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusConflict,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.insufficientBalance",
			"en": "Insufficient wallet balance for this purchase.",
			"fa": "موجودی کیف پول برای این خرید کافی نیست.",
			"ru": "Недостаточно средств на кошельке для этой покупки.",
			"pl": "Niewystarczające środki na portfelu na ten zakup.",
		},
		MessageParams: map[string]any{"balance": balance, "amount": amount},
	}
}

func invalidAmountError(err error) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusBadRequest,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.invalidAmount",
			"en": "The amount is not a valid non-negative integer minor-units value.",
			"fa": "مبلغ یک عدد صحیح نامنفی معتبر (واحد خرد) نیست.",
			"ru": "Сумма не является допустимым неотрицательным целым числом (в минимальных единицах).",
			"pl": "Kwota nie jest prawidłową nieujemną liczbą całkowitą w jednostkach pomocniczych.",
		},
		MessageParams: map[string]any{"detail": err.Error()},
	}
}
