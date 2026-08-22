package przelewy24

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

func TestSha384Hex_Deterministic(t *testing.T) {
	sign1, err := sha384Hex(registerSignPayload{
		SessionId: "sess1", MerchantId: 12345, Amount: 1999, Currency: "PLN", Crc: "crc-secret",
	})
	if err != nil {
		t.Fatalf("sha384Hex: %v", err)
	}
	sign2, err := sha384Hex(registerSignPayload{
		SessionId: "sess1", MerchantId: 12345, Amount: 1999, Currency: "PLN", Crc: "crc-secret",
	})
	if sign1 != sign2 {
		t.Fatalf("sign not deterministic: %q vs %q", sign1, sign2)
	}
	if len(sign1) != 96 { // SHA-384 -> 48 bytes -> 96 hex chars
		t.Fatalf("unexpected sign length %d", len(sign1))
	}

	// Changing any field must change the sign.
	sign3, _ := sha384Hex(registerSignPayload{
		SessionId: "sess1", MerchantId: 12345, Amount: 2000, Currency: "PLN", Crc: "crc-secret",
	})
	if sign1 == sign3 {
		t.Fatalf("sign did not change when amount changed")
	}
}

func TestInitiatePayment_BuildsExpectedRequest(t *testing.T) {
	var gotAuth string
	var gotBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]string{"token": "tok_abc123"},
		})
	}))
	defer server.Close()

	a := New(Config{MerchantId: 1, PosId: 1, ApiKey: "apikey", CrcKey: "crc", BaseURL: server.URL, PublicBaseURL: "https://nima.example.com"})
	attempt := &walletdefs.WalletPaymentAttemptEntity{UniqueId: "attempt-1", Amount: "1999", Currency: "PLN"}

	result, err := a.InitiatePayment(t.Context(), attempt, nil)
	if err != nil {
		t.Fatalf("InitiatePayment: %v", err)
	}
	if result.RedirectUrl != server.URL+"/trnRequest/tok_abc123" {
		t.Fatalf("unexpected redirect url: %q", result.RedirectUrl)
	}
	if gotAuth == "" {
		t.Fatalf("expected Basic auth header")
	}
	if gotBody["sessionId"] != "attempt-1" || gotBody["currency"] != "PLN" {
		t.Fatalf("unexpected request body: %+v", gotBody)
	}
	if gotBody["sign"] == "" || gotBody["sign"] == nil {
		t.Fatalf("expected a sign field in request body")
	}
}

func TestVerifyWebhook_ValidSignatureAndVerifyCallSucceeds(t *testing.T) {
	verifyCalled := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			verifyCalled = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"data": map[string]string{"status": "success"}})
			return
		}
	}))
	defer server.Close()

	a := New(Config{MerchantId: 1, PosId: 1, ApiKey: "apikey", CrcKey: "crc-secret", BaseURL: server.URL})

	n := notification{MerchantId: 1, PosId: 1, SessionId: "attempt-1", Amount: 1999, OriginAmount: 1999, Currency: "PLN", OrderId: 555, MethodId: 1, Statement: "test"}
	sign, err := sha384Hex(notificationSignPayload{
		MerchantId: n.MerchantId, PosId: n.PosId, SessionId: n.SessionId, Amount: n.Amount,
		OriginAmount: n.OriginAmount, Currency: n.Currency, OrderId: n.OrderId, MethodId: n.MethodId,
		Statement: n.Statement, Crc: "crc-secret",
	})
	if err != nil {
		t.Fatalf("sha384Hex: %v", err)
	}
	n.Sign = sign
	body, _ := json.Marshal(n)

	event, err := a.VerifyWebhook(t.Context(), body, http.Header{})
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if !event.Succeeded {
		t.Fatalf("expected Succeeded=true, got event=%+v", event)
	}
	if !verifyCalled {
		t.Fatalf("expected the verify endpoint to be called")
	}
}

func TestVerifyWebhook_TamperedSignatureRejected(t *testing.T) {
	a := New(Config{MerchantId: 1, PosId: 1, ApiKey: "apikey", CrcKey: "crc-secret", BaseURL: "http://unused.invalid"})

	n := notification{MerchantId: 1, PosId: 1, SessionId: "attempt-1", Amount: 1999, Currency: "PLN", OrderId: 555, Sign: "deadbeef"}
	body, _ := json.Marshal(n)

	if _, err := a.VerifyWebhook(t.Context(), body, http.Header{}); err == nil {
		t.Fatalf("expected signature mismatch error, got nil")
	}
}
