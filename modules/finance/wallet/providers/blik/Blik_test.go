package blik

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

func TestInitiatePayment_MissingConfigFailsClearly(t *testing.T) {
	a := New(Config{}) // no env vars set in this test process
	attempt := &walletdefs.WalletPaymentAttemptEntity{UniqueId: "attempt-1", Amount: "1000", Currency: "PLN"}
	if _, err := a.InitiatePayment(t.Context(), attempt, nil); err == nil {
		t.Fatalf("expected a clear configuration error, got nil")
	}
}

func TestInitiatePayment_BuildsExpectedRequest(t *testing.T) {
	var gotAuth string
	var gotBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"transactionId": "txn_abc",
			"redirectUrl":   "https://acquirer.example.com/pay/txn_abc",
		})
	}))
	defer server.Close()

	a := New(Config{BaseURL: server.URL, ApiKey: "key123", MerchantId: "m1"})
	attempt := &walletdefs.WalletPaymentAttemptEntity{UniqueId: "attempt-1", Amount: "1000", Currency: "PLN"}

	result, err := a.InitiatePayment(t.Context(), attempt, nil)
	if err != nil {
		t.Fatalf("InitiatePayment: %v", err)
	}
	if result.GatewayReference != "txn_abc" || result.RedirectUrl != "https://acquirer.example.com/pay/txn_abc" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if gotAuth != "Bearer key123" {
		t.Fatalf("unexpected Authorization header: %q", gotAuth)
	}
	if gotBody["orderId"] != "attempt-1" {
		t.Fatalf("unexpected orderId: %v", gotBody["orderId"])
	}
}

func sign(apiKey, orderId, transactionId, status string) string {
	mac := hmac.New(sha256.New, []byte(apiKey))
	mac.Write([]byte(orderId + "." + transactionId + "." + status))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestVerifyWebhook_ValidSignatureSucceeds(t *testing.T) {
	a := New(Config{BaseURL: "http://unused.invalid", ApiKey: "key123", MerchantId: "m1"})
	body, _ := json.Marshal(map[string]string{
		"transactionId": "txn_abc",
		"orderId":       "attempt-1",
		"status":        "SUCCESS",
		"signature":     sign("key123", "attempt-1", "txn_abc", "SUCCESS"),
	})

	event, err := a.VerifyWebhook(t.Context(), body, http.Header{})
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if !event.Succeeded || event.GatewayReference != "txn_abc" {
		t.Fatalf("unexpected event: %+v", event)
	}
}

func TestVerifyWebhook_TamperedStatusRejected(t *testing.T) {
	a := New(Config{BaseURL: "http://unused.invalid", ApiKey: "key123", MerchantId: "m1"})
	// Signature computed over "FAILED" but status field says "SUCCESS" - must fail.
	body, _ := json.Marshal(map[string]string{
		"transactionId": "txn_abc",
		"orderId":       "attempt-1",
		"status":        "SUCCESS",
		"signature":     sign("key123", "attempt-1", "txn_abc", "FAILED"),
	})

	if _, err := a.VerifyWebhook(t.Context(), body, http.Header{}); err == nil {
		t.Fatalf("expected signature mismatch error, got nil")
	}
}
