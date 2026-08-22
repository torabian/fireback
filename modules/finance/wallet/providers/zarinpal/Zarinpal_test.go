package zarinpal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

func TestInitiatePayment_RequiresReturnUrl(t *testing.T) {
	a := New(Config{MerchantId: "merchant-1", BaseURL: "http://unused.invalid"})
	attempt := &walletdefs.WalletPaymentAttemptEntity{UniqueId: "attempt-1", Amount: "50000"}
	if _, err := a.InitiatePayment(t.Context(), attempt, nil); err == nil {
		t.Fatalf("expected an error when returnUrl is missing")
	}
}

func TestInitiatePayment_BuildsExpectedRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		if body["merchant_id"] != "merchant-1" {
			t.Errorf("unexpected merchant_id: %v", body["merchant_id"])
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{"code": 100, "authority": "A00000000000000000000000000012345"},
		})
	}))
	defer server.Close()

	a := New(Config{MerchantId: "merchant-1", BaseURL: server.URL})
	attempt := &walletdefs.WalletPaymentAttemptEntity{UniqueId: "attempt-1", Amount: "50000"}
	attempt.ReturnUrl.Set(strPtr("https://nima.example.com/wallet/return"))

	result, err := a.InitiatePayment(t.Context(), attempt, nil)
	if err != nil {
		t.Fatalf("InitiatePayment: %v", err)
	}
	if result.GatewayReference != "A00000000000000000000000000012345" {
		t.Fatalf("unexpected gateway reference: %q", result.GatewayReference)
	}
	if result.RedirectUrl != "https://www.zarinpal.com/pg/StartPay/A00000000000000000000000000012345" {
		t.Fatalf("unexpected redirect url: %q", result.RedirectUrl)
	}
}

func TestVerifyWebhook_NonOkStatusFailsWithoutNetworkCall(t *testing.T) {
	a := New(Config{MerchantId: "merchant-1", BaseURL: "http://unused.invalid"})
	body := []byte(`{"Authority":"AUTH123","Status":"NOK"}`)

	event, err := a.VerifyWebhook(t.Context(), body, http.Header{})
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if event.Succeeded {
		t.Fatalf("expected Succeeded=false for Status=NOK")
	}
}

// TestVerifyWebhook_ResolvesAttemptAndVerifies exercises the full OK path: a real
// (sqlite, throwaway) DB holds the attempt row VerifyWebhook resolves by
// gateway_reference to get the original amount, and a fake ZarinPal server confirms it.
func TestVerifyWebhook_ResolvesAttemptAndVerifies(t *testing.T) {
	setupTestDB(t)

	var gotAmount float64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		gotAmount, _ = body["amount"].(float64)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{"code": 100, "ref_id": 998877},
		})
	}))
	defer server.Close()

	db := fireback.GetDbRef()
	if err := db.Exec(`INSERT INTO wallet_payment_attempt_entities (unique_id, gateway_reference, amount, currency, status) VALUES (?, ?, ?, ?, ?)`,
		"attempt-1", "AUTH123", "50000", "IRR", "pending").Error; err != nil {
		t.Fatalf("seed attempt: %v", err)
	}

	a := New(Config{MerchantId: "merchant-1", BaseURL: server.URL})
	body := []byte(`{"Authority":"AUTH123","Status":"OK"}`)

	event, err := a.VerifyWebhook(t.Context(), body, http.Header{})
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if !event.Succeeded {
		t.Fatalf("expected Succeeded=true, got %+v", event)
	}
	if gotAmount != 50000 {
		t.Fatalf("verify call did not use the attempt's original amount: got %v", gotAmount)
	}
}

func setupTestDB(t *testing.T) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "zarinpal-test.db")
	db, err := fireback.DirectConnectToDb(fireback.Config{DbVendor: "sqlite", DbName: path, DbLogLevel: "silent"})
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}
	if err := db.Exec(`CREATE TABLE wallet_payment_attempt_entities (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		unique_id TEXT UNIQUE,
		gateway_reference TEXT,
		amount TEXT,
		currency TEXT,
		status TEXT
	)`).Error; err != nil {
		t.Fatalf("create table: %v", err)
	}
}

func strPtr(s string) *string { return &s }
