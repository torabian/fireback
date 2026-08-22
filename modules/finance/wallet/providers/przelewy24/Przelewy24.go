// Package przelewy24 implements wallet.GatewayAdapter (modules/wallet/GatewayAdapter.go)
// against Przelewy24's REST API v1 (https://developers.przelewy24.pl/). Self-contained -
// the only thing it shares with the other provider packages is the GatewayAdapter
// interface itself.
//
// Register a matching walletGateway row (code "przelewy24") and inject this adapter via
// WalletModuleConfig.Gateways at wallet.WalletModuleSetup(...) time - see
// cmd/nima-server/main.go.
//
// NOTE ON SIGN CORRECTNESS: the registration sign (sessionId/merchantId/amount/currency/
// crc, SHA384 of the JSON-encoded struct) matches P24's published documentation. The
// *notification* sign (recomputed in VerifyWebhook) is reconstructed from P24's
// documented field list (merchantId/posId/sessionId/amount/originAmount/currency/
// orderId/methodId/statement/crc) in that order - P24's own docs note the exact field
// set/order differs per request type and don't fully spell out the notification one in
// public form, so this should be checked against a P24 sandbox transaction before
// relying on it in production (see developers.przelewy24.pl/html/en_calculate_sign.html).
package przelewy24

import (
	"context"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

const (
	defaultBaseURL = "https://secure.przelewy24.pl"
	sandboxBaseURL = "https://sandbox.przelewy24.pl"
)

// Config holds this adapter's credentials/endpoints. Any zero field falls back to the
// matching environment variable when New is called.
type Config struct {
	MerchantId    int    // P24_MERCHANT_ID
	PosId         int    // P24_POS_ID (often equal to MerchantId for simple integrations)
	ApiKey        string // P24_API_KEY - HTTP Basic auth password, distinct from CrcKey
	CrcKey        string // P24_CRC_KEY - shared secret folded into every sign
	Sandbox       bool   // P24_SANDBOX
	PublicBaseURL string // WALLET_PUBLIC_BASE_URL - this server's own public URL, used to build urlStatus
	BaseURL       string // defaults to the real/sandbox P24 host; override only for tests
}

func envOrDefault(value, envKey string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envKey)
}

func envIntOrDefault(value int, envKey string) int {
	if value != 0 {
		return value
	}
	n, _ := strconv.Atoi(os.Getenv(envKey))
	return n
}

// Adapter implements wallet.GatewayAdapter for Przelewy24.
type Adapter struct {
	cfg Config
}

// New builds a Przelewy24 Adapter. Missing Config fields fall back to environment
// variables, so wiring `przelewy24.New(przelewy24.Config{})` unconditionally in main.go
// is safe even when P24 isn't configured - InitiatePayment then just fails with a clear
// error instead of the process refusing to start.
func New(cfg Config) *Adapter {
	cfg.MerchantId = envIntOrDefault(cfg.MerchantId, "P24_MERCHANT_ID")
	cfg.PosId = envIntOrDefault(cfg.PosId, "P24_POS_ID")
	cfg.ApiKey = envOrDefault(cfg.ApiKey, "P24_API_KEY")
	cfg.CrcKey = envOrDefault(cfg.CrcKey, "P24_CRC_KEY")
	cfg.PublicBaseURL = envOrDefault(cfg.PublicBaseURL, "WALLET_PUBLIC_BASE_URL")
	if !cfg.Sandbox {
		cfg.Sandbox = os.Getenv("P24_SANDBOX") == "true"
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

func (a *Adapter) Code() string { return "przelewy24" }

func (a *Adapter) configured() error {
	if a.cfg.MerchantId == 0 || a.cfg.PosId == 0 || a.cfg.ApiKey == "" || a.cfg.CrcKey == "" {
		return fmt.Errorf("przelewy24: P24_MERCHANT_ID/P24_POS_ID/P24_API_KEY/P24_CRC_KEY not fully configured")
	}
	return nil
}

func (a *Adapter) webhookURL() string {
	return strings.TrimRight(a.cfg.PublicBaseURL, "/") + "/wallet/gateway/przelewy24/webhook"
}

// registerSignPayload's field order (and therefore its JSON key order, which
// encoding/json preserves for structs) is exactly what gets hashed - see the package
// doc comment.
type registerSignPayload struct {
	SessionId  string `json:"sessionId"`
	MerchantId int    `json:"merchantId"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
	Crc        string `json:"crc"`
}

type notificationSignPayload struct {
	MerchantId   int    `json:"merchantId"`
	PosId        int    `json:"posId"`
	SessionId    string `json:"sessionId"`
	Amount       int64  `json:"amount"`
	OriginAmount int64  `json:"originAmount"`
	Currency     string `json:"currency"`
	OrderId      int64  `json:"orderId"`
	MethodId     int    `json:"methodId"`
	Statement    string `json:"statement"`
	Crc          string `json:"crc"`
}

type verifySignPayload struct {
	SessionId string `json:"sessionId"`
	OrderId   int64  `json:"orderId"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Crc       string `json:"crc"`
}

func sha384Hex(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	sum := sha512.Sum384(b)
	return hex.EncodeToString(sum[:]), nil
}

func constantTimeHexEqual(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(strings.ToLower(a)), []byte(strings.ToLower(b))) == 1
}

// InitiatePayment registers a transaction and returns the P24 hosted-payment-page
// redirect URL.
func (a *Adapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*wallet.GatewayInitResult, error) {
	if err := a.configured(); err != nil {
		return nil, err
	}
	amount, err := strconv.ParseInt(attempt.Amount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("przelewy24: invalid amount %q: %w", attempt.Amount, err)
	}

	sign, err := sha384Hex(registerSignPayload{
		SessionId:  attempt.UniqueId,
		MerchantId: a.cfg.MerchantId,
		Amount:     amount,
		Currency:   attempt.Currency,
		Crc:        a.cfg.CrcKey,
	})
	if err != nil {
		return nil, err
	}

	returnUrl, _ := attempt.ReturnUrl.Get()
	returnUrlValue := ""
	if returnUrl != nil {
		returnUrlValue = *returnUrl
	}

	// TODO: P24's register endpoint requires a customer email. Nothing upstream
	// (topup's in.fields) threads the wallet owner's email through to this layer yet -
	// left blank here until that's added; P24 may reject registrations with no email.
	payload := map[string]any{
		"merchantId":  a.cfg.MerchantId,
		"posId":       a.cfg.PosId,
		"sessionId":   attempt.UniqueId,
		"amount":      amount,
		"currency":    attempt.Currency,
		"description": "Wallet top-up " + attempt.UniqueId,
		"email":       "",
		"country":     "PL",
		"urlReturn":   returnUrlValue,
		"urlStatus":   a.webhookURL(),
		"sign":        sign,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.BaseURL+"/api/v1/transaction/register", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(strconv.Itoa(a.cfg.PosId), a.cfg.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("przelewy24: request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("przelewy24: register transaction failed (%d): %s", resp.StatusCode, string(respBody))
	}

	var parsed struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("przelewy24: could not parse response: %w", err)
	}
	if parsed.Data.Token == "" {
		return nil, fmt.Errorf("przelewy24: response had no token: %s", string(respBody))
	}

	return &wallet.GatewayInitResult{
		GatewayReference: attempt.UniqueId, // P24 has no id of its own until payment - sessionId (our UniqueId) is what the notification correlates on
		RedirectUrl:      a.cfg.BaseURL + "/trnRequest/" + parsed.Data.Token,
		RawRequest:       string(body),
		RawResponse:      json.RawMessage(respBody),
	}, nil
}

type notification struct {
	MerchantId   int    `json:"merchantId"`
	PosId        int    `json:"posId"`
	SessionId    string `json:"sessionId"`
	Amount       int64  `json:"amount"`
	OriginAmount int64  `json:"originAmount"`
	Currency     string `json:"currency"`
	OrderId      int64  `json:"orderId"`
	MethodId     int    `json:"methodId"`
	Statement    string `json:"statement"`
	Sign         string `json:"sign"`
}

// VerifyWebhook validates the notification signature P24 posts to urlStatus, then
// synchronously calls P24's own verify endpoint to finalize/capture the transaction -
// P24 requires this second call before a registered transaction is actually settled.
func (a *Adapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*wallet.GatewayEvent, error) {
	if err := a.configured(); err != nil {
		return nil, err
	}
	var n notification
	if err := json.Unmarshal(rawBody, &n); err != nil {
		return nil, fmt.Errorf("przelewy24: could not parse notification: %w", err)
	}

	expectedSign, err := sha384Hex(notificationSignPayload{
		MerchantId:   n.MerchantId,
		PosId:        n.PosId,
		SessionId:    n.SessionId,
		Amount:       n.Amount,
		OriginAmount: n.OriginAmount,
		Currency:     n.Currency,
		OrderId:      n.OrderId,
		MethodId:     n.MethodId,
		Statement:    n.Statement,
		Crc:          a.cfg.CrcKey,
	})
	if err != nil {
		return nil, err
	}
	if !constantTimeHexEqual(expectedSign, n.Sign) {
		return nil, fmt.Errorf("przelewy24: notification signature mismatch")
	}

	succeeded, verifyErr := a.verifyTransaction(ctx, n.SessionId, n.OrderId, n.Amount, n.Currency)
	failureReason := ""
	if verifyErr != nil {
		failureReason = verifyErr.Error()
	}

	return &wallet.GatewayEvent{
		GatewayReference: n.SessionId,
		EventType:        "transaction.notification",
		Succeeded:        succeeded,
		FailureReason:    failureReason,
		Payload:          json.RawMessage(rawBody),
	}, nil
}

// verifyTransaction calls PUT /api/v1/transaction/verify, the second required step to
// actually settle a P24 transaction after its notification's signature checks out.
func (a *Adapter) verifyTransaction(ctx context.Context, sessionId string, orderId, amount int64, currency string) (bool, error) {
	sign, err := sha384Hex(verifySignPayload{
		SessionId: sessionId,
		OrderId:   orderId,
		Amount:    amount,
		Currency:  currency,
		Crc:       a.cfg.CrcKey,
	})
	if err != nil {
		return false, err
	}
	payload := map[string]any{
		"merchantId": a.cfg.MerchantId,
		"posId":      a.cfg.PosId,
		"sessionId":  sessionId,
		"amount":     amount,
		"currency":   currency,
		"orderId":    orderId,
		"sign":       sign,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, a.cfg.BaseURL+"/api/v1/transaction/verify", strings.NewReader(string(body)))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(strconv.Itoa(a.cfg.PosId), a.cfg.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("przelewy24: verify request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 300 {
		return false, fmt.Errorf("przelewy24: verify failed (%d): %s", resp.StatusCode, string(respBody))
	}

	var parsed struct {
		Data struct {
			Status string `json:"status"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return false, fmt.Errorf("przelewy24: could not parse verify response: %w", err)
	}
	return parsed.Data.Status == "success", nil
}
