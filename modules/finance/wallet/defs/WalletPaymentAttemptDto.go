package walletdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletPaymentAttemptDto
type WalletPaymentAttemptDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The wallet this attempt would credit/debit if it succeeds.
	Wallet emigo.OneNullable[WalletDto] `json:"wallet" validate:"required" yaml:"wallet"`
	// The gateway this attempt is routed through.
	Gateway emigo.OneNullable[WalletGatewayDto] `json:"gateway" validate:"required" yaml:"gateway"`
	// What this attempt is for.
	Purpose string `json:"purpose" validate:"required,oneof=topup purchase withdrawal" yaml:"purpose"`
	// Requested amount as a minor-units string, in currency.
	Amount string `json:"amount" validate:"required" yaml:"amount"`
	// Currency code for amount - must match the wallet's currency.
	Currency string `json:"currency" validate:"required" yaml:"currency"`
	// Current lifecycle state of this attempt.
	Status string `json:"status" validate:"required,oneof=pending requires_action succeeded failed cancelled expired" yaml:"status"`
	// The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
	GatewayReference emigo.Nullable[string] `json:"gatewayReference" yaml:"gatewayReference"`
	// Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
	IdempotencyKey string `json:"idempotencyKey" validate:"required" yaml:"idempotencyKey"`
	// When this attempt was created.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
	// Human-readable reason, populated when status is "failed".
	FailureReason emigo.Nullable[string] `json:"failureReason" yaml:"failureReason"`
	// The raw request sent to the gateway when initiating this attempt.
	RawRequest complexes.JSON `json:"rawRequest" yaml:"rawRequest"`
	// The raw response/init payload received back from the gateway.
	RawResponse complexes.JSON `json:"rawResponse" yaml:"rawResponse"`
	// The ledger entry that was created once this attempt succeeded. Empty until then.
	WalletTransaction emigo.OneNullable[WalletTransactionDto] `json:"walletTransaction" yaml:"walletTransaction"`
	// When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
	ExpiresAt complexes.XDate `json:"expiresAt" yaml:"expiresAt"`
	// When this attempt reached a terminal status. Empty until then.
	CompletedAt complexes.XDate `json:"completedAt" yaml:"completedAt"`
	// Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
	ReturnUrl emigo.Nullable[string] `json:"returnUrl" yaml:"returnUrl"`
}

func (x *WalletPaymentAttemptDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletPaymentAttemptDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "wallet",
			Type:        "one?",
			Description: "The wallet this attempt would credit/debit if it succeeds.",
		},
		{
			Name:        prefix + "gateway",
			Type:        "one?",
			Description: "The gateway this attempt is routed through.",
		},
		{
			Name:        prefix + "purpose",
			Type:        "string",
			Description: "What this attempt is for.",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Requested amount as a minor-units string, in currency.",
		},
		{
			Name:        prefix + "currency",
			Type:        "string",
			Description: "Currency code for amount - must match the wallet's currency.",
		},
		{
			Name:        prefix + "status",
			Type:        "string",
			Description: "Current lifecycle state of this attempt.",
		},
		{
			Name:        prefix + "gateway-reference",
			Type:        "string?",
			Description: "The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.",
		},
		{
			Name:        prefix + "idempotency-key",
			Type:        "string",
			Description: "Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.",
		},
		{
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "When this attempt was created.",
		},
		{
			Name:        prefix + "failure-reason",
			Type:        "string?",
			Description: "Human-readable reason, populated when status is \"failed\".",
		},
		{
			Name:        prefix + "raw-request",
			Type:        "complex",
			Description: "The raw request sent to the gateway when initiating this attempt.",
		},
		{
			Name:        prefix + "raw-response",
			Type:        "complex",
			Description: "The raw response/init payload received back from the gateway.",
		},
		{
			Name:        prefix + "wallet-transaction",
			Type:        "one?",
			Description: "The ledger entry that was created once this attempt succeeded. Empty until then.",
		},
		{
			Name:        prefix + "expires-at",
			Type:        "complex",
			Description: "When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.",
		},
		{
			Name:        prefix + "completed-at",
			Type:        "complex",
			Description: "When this attempt reached a terminal status. Empty until then.",
		},
		{
			Name:        prefix + "return-url",
			Type:        "string?",
			Description: "Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).",
		},
	}
}
func CastWalletPaymentAttemptDtoFromCli(c emigo.CliCastable) WalletPaymentAttemptDto {
	data := WalletPaymentAttemptDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("wallet") {
		data.Wallet = emigo.CapturePossibleOneNullable(CastWalletDtoFromCli, "wallet", c)
	}
	if c.IsSet("gateway") {
		data.Gateway = emigo.CapturePossibleOneNullable(CastWalletGatewayDtoFromCli, "gateway", c)
	}
	if c.IsSet("purpose") {
		data.Purpose = c.String("purpose")
	}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("gateway-reference") {
		emigo.ParseNullable(c.String("gateway-reference"), &data.GatewayReference)
	}
	if c.IsSet("idempotency-key") {
		data.IdempotencyKey = c.String("idempotency-key")
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("failure-reason") {
		emigo.ParseNullable(c.String("failure-reason"), &data.FailureReason)
	}
	if c.IsSet("raw-request") {
		if u, ok := any(&data.RawRequest).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("raw-request")))
		}
	}
	if c.IsSet("raw-response") {
		if u, ok := any(&data.RawResponse).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("raw-response")))
		}
	}
	if c.IsSet("wallet-transaction") {
		data.WalletTransaction = emigo.CapturePossibleOneNullable(CastWalletTransactionDtoFromCli, "wallet-transaction", c)
	}
	if c.IsSet("expires-at") {
		if u, ok := any(&data.ExpiresAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("expires-at")))
		}
	}
	if c.IsSet("completed-at") {
		if u, ok := any(&data.CompletedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("completed-at")))
		}
	}
	if c.IsSet("return-url") {
		emigo.ParseNullable(c.String("return-url"), &data.ReturnUrl)
	}
	return data
}
