// Package stripe implements wallet.GatewayAdapter (modules/wallet/GatewayAdapter.go)
// against Stripe's Payment Intents API (https://docs.stripe.com/api/payment_intents).
// Self-contained - the only thing it shares with the other provider packages is the
// GatewayAdapter interface itself.
//
// Register a matching walletGateway row (code "stripe") and inject this adapter via
// WalletModuleConfig.Gateways at wallet.WalletModuleSetup(...) time - see
// cmd/nima-server/main.go.
package stripe

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

const defaultBaseURL = "https://api.stripe.com"

// Config holds this adapter's credentials/endpoints. Any zero field falls back to the
// matching environment variable when New is called.
type Config struct {
	// SecretKey is Stripe's API secret key (sk_test_.../sk_live_...). Falls back to
	// STRIPE_SECRET_KEY.
	SecretKey string
	// WebhookSecret is the signing secret configured on the Stripe Dashboard for the
	// /wallet/gateway/stripe/webhook endpoint. Falls back to STRIPE_WEBHOOK_SECRET.
	WebhookSecret string
	// BaseURL defaults to Stripe's real API host - override only for tests
	// (httptest.Server).
	BaseURL string
}

func envOrDefault(value, envKey string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envKey)
}

// Adapter implements wallet.GatewayAdapter for Stripe.
type Adapter struct {
	cfg Config
}

// New builds a Stripe Adapter. Missing Config fields fall back to environment
// variables, so wiring `stripe.New(stripe.Config{})` unconditionally in main.go is safe
// even when Stripe isn't configured for this installation - InitiatePayment then just
// fails with a clear error instead of the process refusing to start.
func New(cfg Config) *Adapter {
	cfg.SecretKey = envOrDefault(cfg.SecretKey, "STRIPE_SECRET_KEY")
	cfg.WebhookSecret = envOrDefault(cfg.WebhookSecret, "STRIPE_WEBHOOK_SECRET")
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}
	return &Adapter{cfg: cfg}
}

func (a *Adapter) Code() string { return "stripe" }

// InitiatePayment creates a Stripe PaymentIntent. Amount/currency are passed straight
// through from the attempt - Stripe's own amount convention is already integer
// minor-units, matching this module's convention exactly, so no conversion is needed.
func (a *Adapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*wallet.GatewayInitResult, error) {
	if a.cfg.SecretKey == "" {
		return nil, fmt.Errorf("stripe: STRIPE_SECRET_KEY is not configured")
	}

	form := url.Values{}
	form.Set("amount", attempt.Amount)
	form.Set("currency", strings.ToLower(attempt.Currency))
	form.Set("automatic_payment_methods[enabled]", "true")
	form.Set("metadata[walletPaymentAttemptId]", attempt.UniqueId)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.BaseURL+"/v1/payment_intents", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(a.cfg.SecretKey, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("stripe: request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("stripe: create payment intent failed (%d): %s", resp.StatusCode, string(body))
	}

	var parsed struct {
		Id           string `json:"id"`
		ClientSecret string `json:"client_secret"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("stripe: could not parse response: %w", err)
	}
	if parsed.Id == "" {
		return nil, fmt.Errorf("stripe: response had no payment intent id: %s", string(body))
	}

	return &wallet.GatewayInitResult{
		GatewayReference: parsed.Id,
		ClientSecret:     parsed.ClientSecret,
		RawRequest:       form.Encode(),
		RawResponse:      json.RawMessage(body),
	}, nil
}

// VerifyWebhook implements Stripe's documented webhook signature scheme
// (https://docs.stripe.com/webhooks#verify-manually): parses the Stripe-Signature
// header ("t=<unix>,v1=<hex>[,v0=...]"), recomputes hex(HMAC-SHA256(WebhookSecret,
// "{t}.{rawBody}")), and rejects anything that doesn't match v1 or falls outside a
// 5-minute tolerance window (replay protection).
func (a *Adapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*wallet.GatewayEvent, error) {
	if a.cfg.WebhookSecret == "" {
		return nil, fmt.Errorf("stripe: STRIPE_WEBHOOK_SECRET is not configured")
	}
	ts, v1, err := parseSignatureHeader(headers.Get("Stripe-Signature"))
	if err != nil {
		return nil, err
	}
	if err := verifySignature(a.cfg.WebhookSecret, ts, rawBody, v1); err != nil {
		return nil, err
	}

	var evt struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Data struct {
			Object struct {
				Id string `json:"id"`
			} `json:"object"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rawBody, &evt); err != nil {
		return nil, fmt.Errorf("stripe: could not parse event: %w", err)
	}

	failureReason := ""
	if evt.Type == "payment_intent.payment_failed" {
		failureReason = "stripe: payment_intent.payment_failed"
	}

	return &wallet.GatewayEvent{
		ExternalEventId:  evt.Id,
		EventType:        evt.Type,
		GatewayReference: evt.Data.Object.Id,
		Succeeded:        evt.Type == "payment_intent.succeeded",
		FailureReason:    failureReason,
		Payload:          json.RawMessage(rawBody),
	}, nil
}

// parseSignatureHeader parses "t=1614556800,v1=abc123,v0=..." into the timestamp and
// (first) v1 signature. Stripe may list more than one v1 entry during secret rotation;
// this keeps it simple and checks only the first.
func parseSignatureHeader(header string) (timestamp string, v1 string, err error) {
	if header == "" {
		return "", "", fmt.Errorf("stripe: missing Stripe-Signature header")
	}
	for _, part := range strings.Split(header, ",") {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "t":
			timestamp = kv[1]
		case "v1":
			if v1 == "" {
				v1 = kv[1]
			}
		}
	}
	if timestamp == "" || v1 == "" {
		return "", "", fmt.Errorf("stripe: malformed Stripe-Signature header")
	}
	return timestamp, v1, nil
}

func verifySignature(secret, timestamp string, rawBody []byte, v1 string) error {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("stripe: invalid timestamp in signature header")
	}
	if d := time.Since(time.Unix(ts, 0)); d > 5*time.Minute || d < -5*time.Minute {
		return fmt.Errorf("stripe: webhook timestamp outside tolerance (possible replay)")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + string(rawBody)))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(v1)) {
		return fmt.Errorf("stripe: signature mismatch")
	}
	return nil
}
