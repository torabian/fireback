// Package blik implements wallet.GatewayAdapter (modules/wallet/GatewayAdapter.go) for
// BLIK, the Polish mobile payment scheme. Self-contained - the only thing it shares with
// the other provider packages is the GatewayAdapter interface itself.
//
// IMPORTANT: unlike Stripe/Przelewy24/ZarinPal, BLIK has no vendor-neutral,
// direct-to-merchant REST API a small business can just sign up and call - every
// merchant integration goes through some acquirer/PSP's own gateway (Przelewy24, PayU,
// Adyen, Nuvei, Worldline, Stripe itself, ...), each with its own request/response
// shape. This package therefore implements a *generic* BLIK-acquirer adapter: submit
// merchant/order/amount, get back a redirect (or the BLIK 6-digit-code flow via the same
// endpoint shape most acquirers converge on), verify a webhook via HMAC. Point Config at
// whichever acquirer contract the business actually holds - BaseURL has no sensible
// universal default and must be set explicitly (unlike the other three adapters). If
// that acquirer turns out to be Przelewy24 itself, using the przelewy24 package's own
// BLIK payment method there instead is equally valid - this package exists for the case
// where BLIK is contracted through a separate acquirer.
//
// Register a matching walletGateway row (code "blik") and inject this adapter via
// WalletModuleConfig.Gateways at wallet.WalletModuleSetup(...) time - see
// cmd/nima-server/main.go.
package blik

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

// Config holds this adapter's credentials/endpoints. Any zero field falls back to the
// matching environment variable when New is called. Unlike the other three provider
// packages, BaseURL has no built-in default - it must point at whichever acquirer the
// business has actually contracted with.
type Config struct {
	BaseURL       string // BLIK_ACQUIRER_BASE_URL - required, no default (see package doc)
	ApiKey        string // BLIK_API_KEY - HMAC signing secret shared with the acquirer
	MerchantId    string // BLIK_MERCHANT_ID
	PublicBaseURL string // WALLET_PUBLIC_BASE_URL - this server's own public URL, used to build notifyUrl
}

func envOrDefault(value, envKey string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envKey)
}

// Adapter implements wallet.GatewayAdapter for a generic BLIK acquirer.
type Adapter struct {
	cfg Config
}

// New builds a BLIK Adapter. Missing Config fields fall back to environment variables,
// so wiring `blik.New(blik.Config{})` unconditionally in main.go is safe even when no
// acquirer is configured - InitiatePayment then just fails with a clear error instead of
// the process refusing to start.
func New(cfg Config) *Adapter {
	cfg.BaseURL = envOrDefault(cfg.BaseURL, "BLIK_ACQUIRER_BASE_URL")
	cfg.ApiKey = envOrDefault(cfg.ApiKey, "BLIK_API_KEY")
	cfg.MerchantId = envOrDefault(cfg.MerchantId, "BLIK_MERCHANT_ID")
	cfg.PublicBaseURL = envOrDefault(cfg.PublicBaseURL, "WALLET_PUBLIC_BASE_URL")
	return &Adapter{cfg: cfg}
}

func (a *Adapter) Code() string { return "blik" }

func (a *Adapter) configured() error {
	if a.cfg.BaseURL == "" || a.cfg.ApiKey == "" || a.cfg.MerchantId == "" {
		return fmt.Errorf("blik: BLIK_ACQUIRER_BASE_URL/BLIK_API_KEY/BLIK_MERCHANT_ID not fully configured - point this at an actual acquirer contract, see package doc comment")
	}
	return nil
}

func (a *Adapter) notifyURL() string {
	return strings.TrimRight(a.cfg.PublicBaseURL, "/") + "/wallet/gateway/blik/webhook"
}

// InitiatePayment starts a BLIK transaction with the configured acquirer.
func (a *Adapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*wallet.GatewayInitResult, error) {
	if err := a.configured(); err != nil {
		return nil, err
	}

	returnUrlPtr, _ := attempt.ReturnUrl.Get()
	returnUrl := ""
	if returnUrlPtr != nil {
		returnUrl = *returnUrlPtr
	}

	payload := map[string]any{
		"merchantId": a.cfg.MerchantId,
		"orderId":    attempt.UniqueId,
		"amount":     attempt.Amount,
		"currency":   attempt.Currency,
		"returnUrl":  returnUrl,
		"notifyUrl":  a.notifyURL(),
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(a.cfg.BaseURL, "/")+"/transactions", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.cfg.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("blik: request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("blik: start transaction failed (%d): %s", resp.StatusCode, string(respBody))
	}

	var parsed struct {
		TransactionId string `json:"transactionId"`
		RedirectUrl   string `json:"redirectUrl"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("blik: could not parse response: %w", err)
	}
	if parsed.TransactionId == "" {
		return nil, fmt.Errorf("blik: response had no transactionId: %s", string(respBody))
	}

	return &wallet.GatewayInitResult{
		GatewayReference: parsed.TransactionId,
		RedirectUrl:      parsed.RedirectUrl,
		RawRequest:       string(body),
		RawResponse:      json.RawMessage(respBody),
	}, nil
}

type webhookBody struct {
	TransactionId string `json:"transactionId"`
	OrderId       string `json:"orderId"`
	Status        string `json:"status"` // e.g. "SUCCESS" / "FAILED"
	Signature     string `json:"signature"`
}

// VerifyWebhook validates a generic HMAC-SHA256(ApiKey, "orderId.transactionId.status")
// signature. Confirm this matches the actual acquirer's documented scheme before relying
// on it in production - it's the common shape across several acquirers, not a
// spec-guaranteed one (see package doc comment).
func (a *Adapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*wallet.GatewayEvent, error) {
	if err := a.configured(); err != nil {
		return nil, err
	}
	var body webhookBody
	if err := json.Unmarshal(rawBody, &body); err != nil {
		return nil, fmt.Errorf("blik: could not parse webhook body: %w", err)
	}
	if body.TransactionId == "" {
		return nil, fmt.Errorf("blik: webhook missing transactionId")
	}

	mac := hmac.New(sha256.New, []byte(a.cfg.ApiKey))
	mac.Write([]byte(body.OrderId + "." + body.TransactionId + "." + body.Status))
	expected := hex.EncodeToString(mac.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(strings.ToLower(expected)), []byte(strings.ToLower(body.Signature))) != 1 {
		return nil, fmt.Errorf("blik: signature mismatch")
	}

	succeeded := strings.EqualFold(body.Status, "SUCCESS")
	failureReason := ""
	if !succeeded {
		failureReason = "blik: status=" + body.Status
	}

	return &wallet.GatewayEvent{
		GatewayReference: body.TransactionId,
		EventType:        "transaction." + strings.ToLower(body.Status),
		Succeeded:        succeeded,
		FailureReason:    failureReason,
		Payload:          json.RawMessage(rawBody),
	}, nil
}
