package stripe

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

func signedHeader(secret, timestamp, body string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + body))
	sig := hex.EncodeToString(mac.Sum(nil))
	return "t=" + timestamp + ",v1=" + sig
}

func TestVerifyWebhook_ValidSignatureSucceeds(t *testing.T) {
	a := New(Config{WebhookSecret: "whsec_test"})
	body := `{"id":"evt_1","type":"payment_intent.succeeded","data":{"object":{"id":"pi_123"}}}`
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	headers := http.Header{"Stripe-Signature": []string{signedHeader("whsec_test", ts, body)}}

	event, err := a.VerifyWebhook(t.Context(), []byte(body), headers)
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if !event.Succeeded || event.GatewayReference != "pi_123" || event.ExternalEventId != "evt_1" {
		t.Fatalf("unexpected event: %+v", event)
	}
}

func TestVerifyWebhook_WrongSecretRejected(t *testing.T) {
	a := New(Config{WebhookSecret: "whsec_test"})
	body := `{"id":"evt_1","type":"payment_intent.succeeded","data":{"object":{"id":"pi_123"}}}`
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	headers := http.Header{"Stripe-Signature": []string{signedHeader("whsec_WRONG", ts, body)}}

	if _, err := a.VerifyWebhook(t.Context(), []byte(body), headers); err == nil {
		t.Fatalf("expected signature mismatch error, got nil")
	}
}

func TestVerifyWebhook_StaleTimestampRejected(t *testing.T) {
	a := New(Config{WebhookSecret: "whsec_test"})
	body := `{"id":"evt_1","type":"payment_intent.succeeded","data":{"object":{"id":"pi_123"}}}`
	staleTs := strconv.FormatInt(time.Now().Add(-1*time.Hour).Unix(), 10)
	headers := http.Header{"Stripe-Signature": []string{signedHeader("whsec_test", staleTs, body)}}

	if _, err := a.VerifyWebhook(t.Context(), []byte(body), headers); err == nil {
		t.Fatalf("expected stale-timestamp error, got nil")
	}
}

func TestVerifyWebhook_FailedPaymentReportsFailureReason(t *testing.T) {
	a := New(Config{WebhookSecret: "whsec_test"})
	body := `{"id":"evt_2","type":"payment_intent.payment_failed","data":{"object":{"id":"pi_456"}}}`
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	headers := http.Header{"Stripe-Signature": []string{signedHeader("whsec_test", ts, body)}}

	event, err := a.VerifyWebhook(t.Context(), []byte(body), headers)
	if err != nil {
		t.Fatalf("VerifyWebhook: %v", err)
	}
	if event.Succeeded || event.FailureReason == "" {
		t.Fatalf("unexpected event: %+v", event)
	}
}

// TestInitiatePayment_BuildsExpectedRequest exercises InitiatePayment against a fake
// Stripe server (httptest), verifying the request shape (auth header, form fields) and
// that the response is parsed correctly - no live Stripe credentials involved.
func TestInitiatePayment_BuildsExpectedRequest(t *testing.T) {
	var gotAuth, gotAmount, gotCurrency string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		gotAmount = r.PostForm.Get("amount")
		gotCurrency = r.PostForm.Get("currency")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"id":            "pi_test123",
			"client_secret": "pi_test123_secret_abc",
		})
	}))
	defer server.Close()

	a := New(Config{SecretKey: "sk_test_abc", BaseURL: server.URL})
	attempt := &walletdefs.WalletPaymentAttemptEntity{
		UniqueId: "attempt-1",
		Amount:   "1999",
		Currency: "USD",
	}

	result, err := a.InitiatePayment(t.Context(), attempt, nil)
	if err != nil {
		t.Fatalf("InitiatePayment: %v", err)
	}
	if result.GatewayReference != "pi_test123" || result.ClientSecret != "pi_test123_secret_abc" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if gotAuth == "" {
		t.Fatalf("expected Authorization header to be set")
	}
	if gotAmount != "1999" || gotCurrency != "usd" {
		t.Fatalf("amount/currency not passed through correctly: amount=%q currency=%q", gotAmount, gotCurrency)
	}
}
