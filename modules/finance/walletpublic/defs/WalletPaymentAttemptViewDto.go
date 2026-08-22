package walletpublicdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletPaymentAttemptViewDto
type WalletPaymentAttemptViewDto struct {
	// Unique id of this attempt.
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// "topup", "purchase", or "withdrawal".
	Purpose string `json:"purpose" yaml:"purpose"`
	// Requested amount as a minor-units string.
	Amount string `json:"amount" yaml:"amount"`
	// Currency code for amount.
	Currency string `json:"currency" yaml:"currency"`
	// Current lifecycle state of this attempt.
	Status string `json:"status" yaml:"status"`
	// Code of the gateway this attempt is routed through.
	GatewayCode string `json:"gatewayCode" yaml:"gatewayCode"`
	// Human-readable reason, populated when status is "failed".
	FailureReason emigo.Nullable[string] `json:"failureReason" yaml:"failureReason"`
	// When this attempt was created.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
}

func (x *WalletPaymentAttemptViewDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletPaymentAttemptViewDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Unique id of this attempt.",
		},
		{
			Name:        prefix + "purpose",
			Type:        "string",
			Description: "\"topup\", \"purchase\", or \"withdrawal\".",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Requested amount as a minor-units string.",
		},
		{
			Name:        prefix + "currency",
			Type:        "string",
			Description: "Currency code for amount.",
		},
		{
			Name:        prefix + "status",
			Type:        "string",
			Description: "Current lifecycle state of this attempt.",
		},
		{
			Name:        prefix + "gateway-code",
			Type:        "string",
			Description: "Code of the gateway this attempt is routed through.",
		},
		{
			Name:        prefix + "failure-reason",
			Type:        "string?",
			Description: "Human-readable reason, populated when status is \"failed\".",
		},
		{
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "When this attempt was created.",
		},
	}
}
func CastWalletPaymentAttemptViewDtoFromCli(c emigo.CliCastable) WalletPaymentAttemptViewDto {
	data := WalletPaymentAttemptViewDto{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
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
	if c.IsSet("gateway-code") {
		data.GatewayCode = c.String("gateway-code")
	}
	if c.IsSet("failure-reason") {
		emigo.ParseNullable(c.String("failure-reason"), &data.FailureReason)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	return data
}
