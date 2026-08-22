// Package zarinpal implements wallet.GatewayAdapter (modules/wallet/GatewayAdapter.go)
// against ZarinPal's REST API v4 (https://zarinpal-lab.github.io/API-Docs/) - the major
// Iranian payment gateway (IPG). Self-contained - the only thing it shares with the
// other provider packages is the GatewayAdapter interface itself.
//
// Register a matching walletGateway row (code "zarinpal") and inject this adapter via
// WalletModuleConfig.Gateways at wallet.WalletModuleSetup(...) time - see
// cmd/nima-server/main.go.
//
// ZarinPal's callback shape is genuinely different from Stripe/Przelewy24: there is no
// signed server-to-server webhook at all. ZarinPal redirects the payer's browser via GET
// to callback_url?Authority=...&Status=OK|NOK, and the merchant confirms authenticity by
// calling ZarinPal's own verify.json synchronously (passing back the amount it originally
// requested, as an integrity check) rather than checking a shared-secret signature.
// GatewayWebhookHandler (modules/wallet/GatewayWebhookImplementation.go) already
// normalizes that GET's query string into the same (rawBody []byte, headers http.Header)
// shape every VerifyWebhook sees, so this adapter still only implements the one
// interface. What's ZarinPal-specific is that VerifyWebhook here needs the original
// amount to call verify.json, which isn't in the callback params - rather than adding an
// interface method just for this, VerifyWebhook resolves the walletPaymentAttempt itself
// (fireback.GetDbRef(), by gateway_reference = Authority) the same way
// GatewayWebhookHandler does elsewhere. That's a small, deliberate layering compromise,
// not an interface change.
package zarinpal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

const (
	defaultBaseURL = "https://api.zarinpal.com"
	sandboxBaseURL = "https://sandbox.zarinpal.com"
)

// Config holds this adapter's credentials/endpoints. Any zero field falls back to the
// matching environment variable when New is called.
type Config struct {
	MerchantId string // ZARINPAL_MERCHANT_ID (a 36-character UUID-like string)
	Sandbox    bool   // ZARINPAL_SANDBOX
	BaseURL    string // defaults to the real/sandbox ZarinPal host; override only for tests
}

func envOrDefault(value, envKey string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envKey)
}

// Adapter implements wallet.GatewayAdapter for ZarinPal.
type Adapter struct {
	cfg Config
}

// New builds a ZarinPal Adapter. Missing Config fields fall back to environment
// variables, so wiring `zarinpal.New(zarinpal.Config{})` unconditionally in main.go is
// safe even when ZarinPal isn't configured - InitiatePayment then just fails with a
// clear error instead of the process refusing to start.
func New(cfg Config) *Adapter {
	cfg.MerchantId = envOrDefault(cfg.MerchantId, "ZARINPAL_MERCHANT_ID")
	if !cfg.Sandbox {
		cfg.Sandbox = os.Getenv("ZARINPAL_SANDBOX") == "true"
	}
	if cfg.BaseURL == "" {
		if cfg.Sandbox {
			cfg.BaseURL = sandboxBaseURL
		} else {
			cfg.BaseURL = defaultBaseURL
		}
	}
	return &Adapter{cfg: cfg}
}

func (a *Adapter) Code() string { return "zarinpal" }

// startPayHost mirrors BaseURL's sandbox/production choice for the hosted payment page,
// which lives under www.zarinpal.com (production) or sandbox.zarinpal.com (sandbox) -
// not under api.zarinpal.com/sandbox.zarinpal.com like the REST calls.
func (a *Adapter) startPayHost() string {
	if a.cfg.Sandbox {
		return "https://sandbox.zarinpal.com"
	}
	return "https://www.zarinpal.com"
}

// InitiatePayment requests a payment Authority and returns ZarinPal's hosted-payment-page
// redirect URL. attempt.Amount is Rial - ZarinPal's own minor-unit-less currency, so no
// conversion is needed as long as the matching walletCurrency (IRR) declares decimals: 0.
func (a *Adapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*wallet.GatewayInitResult, error) {
	if a.cfg.MerchantId == "" {
		return nil, fmt.Errorf("zarinpal: ZARINPAL_MERCHANT_ID is not configured")
	}
	amount, err := strconv.ParseInt(attempt.Amount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("zarinpal: invalid amount %q: %w", attempt.Amount, err)
	}

	returnUrlPtr, _ := attempt.ReturnUrl.Get()
	returnUrl := ""
	if returnUrlPtr != nil {
		returnUrl = *returnUrlPtr
	}
	if returnUrl == "" {
		return nil, fmt.Errorf("zarinpal: returnUrl is required (ZarinPal redirects the browser back to it as callback_url)")
	}

	payload := map[string]any{
		"merchant_id":  a.cfg.MerchantId,
		"amount":       amount,
		"callback_url": returnUrl,
		"description":  "Wallet top-up " + attempt.UniqueId,
		"metadata":     map[string]string{"walletPaymentAttemptId": attempt.UniqueId},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.BaseURL+"/pg/v4/payment/request.json", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("zarinpal: request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		Data struct {
			Code      int    `json:"code"`
			Authority string `json:"authority"`
		} `json:"data"`
		Errors any `json:"errors"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("zarinpal: could not parse response: %w", err)
	}
	if parsed.Data.Code != 100 || parsed.Data.Authority == "" {
		return nil, fmt.Errorf("zarinpal: payment request failed (code %d): %s", parsed.Data.Code, string(respBody))
	}

	return &wallet.GatewayInitResult{
		GatewayReference: parsed.Data.Authority,
		RedirectUrl:      a.startPayHost() + "/pg/StartPay/" + parsed.Data.Authority,
		RawRequest:       string(body),
		RawResponse:      json.RawMessage(respBody),
	}, nil
}

// callbackParams is what GatewayWebhookHandler's GET-query normalization produces for
// ZarinPal's redirect: {"Authority":"...","Status":"OK"|"NOK"}.
type callbackParams struct {
	Authority string `json:"Authority"`
	Status    string `json:"Status"`
}

// VerifyWebhook handles ZarinPal's GET-redirect callback (see the package doc comment
// for why this looks different from a signed POST webhook): rejects immediately on
// Status != "OK", otherwise resolves the matching attempt to get the original amount and
// calls ZarinPal's verify.json to confirm - code 100 (fresh success) or 101 (already
// verified, e.g. a duplicate browser redirect) both count as success.
func (a *Adapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*wallet.GatewayEvent, error) {
	if a.cfg.MerchantId == "" {
		return nil, fmt.Errorf("zarinpal: ZARINPAL_MERCHANT_ID is not configured")
	}
	var cb callbackParams
	if err := json.Unmarshal(rawBody, &cb); err != nil {
		return nil, fmt.Errorf("zarinpal: could not parse callback: %w", err)
	}
	if cb.Authority == "" {
		return nil, fmt.Errorf("zarinpal: callback missing Authority")
	}
	if cb.Status != "OK" {
		return &wallet.GatewayEvent{
			GatewayReference: cb.Authority,
			EventType:        "payment.callback",
			Succeeded:        false,
			FailureReason:    "zarinpal: status=" + cb.Status,
			Payload:          json.RawMessage(rawBody),
		}, nil
	}

	var attempt walletdefs.WalletPaymentAttemptEntity
	if err := fireback.GetDbRef().First(&attempt, "gateway_reference = ?", cb.Authority).Error; err != nil {
		return nil, fmt.Errorf("zarinpal: could not resolve payment attempt for Authority %s: %w", cb.Authority, err)
	}
	amount, err := strconv.ParseInt(attempt.Amount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("zarinpal: invalid stored amount %q: %w", attempt.Amount, err)
	}

	verifyPayload := map[string]any{
		"merchant_id": a.cfg.MerchantId,
		"amount":      amount,
		"authority":   cb.Authority,
	}
	verifyBody, err := json.Marshal(verifyPayload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.BaseURL+"/pg/v4/payment/verify.json", strings.NewReader(string(verifyBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("zarinpal: verify request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		Data struct {
			Code  int    `json:"code"`
			RefId int64  `json:"ref_id"`
			Card  string `json:"card_pan"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("zarinpal: could not parse verify response: %w", err)
	}

	succeeded := parsed.Data.Code == 100 || parsed.Data.Code == 101
	failureReason := ""
	if !succeeded {
		failureReason = fmt.Sprintf("zarinpal: verify returned code %d", parsed.Data.Code)
	}

	return &wallet.GatewayEvent{
		GatewayReference: cb.Authority,
		EventType:        "payment.verify",
		Succeeded:        succeeded,
		FailureReason:    failureReason,
		Payload:          json.RawMessage(respBody),
	}, nil
}
