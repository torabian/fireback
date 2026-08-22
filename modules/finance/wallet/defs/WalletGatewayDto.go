package walletdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletGatewayDto
type WalletGatewayDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
	Code string `json:"code" validate:"required" yaml:"code"`
	// Display name shown to wallet owners choosing a topup method.
	Name string `json:"name" validate:"required" yaml:"name"`
	// Whether this gateway settles in fiat or crypto.
	Kind string `json:"kind" validate:"required,oneof=fiat crypto" yaml:"kind"`
	// Whether wallet owners can currently start a topup through this gateway.
	IsActive bool `json:"isActive" yaml:"isActive"`
	// Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
	Config complexes.JSON `json:"config" yaml:"config"`
	// JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
	SupportedCurrencies complexes.JSON `json:"supportedCurrencies" yaml:"supportedCurrencies"`
}

func (x *WalletGatewayDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletGatewayDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "code",
			Type:        "string",
			Description: "Unique code identifying this gateway's GatewayAdapter implementation, e.g. \"stripe\", \"manual\", \"onchain_eth\".",
		},
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Display name shown to wallet owners choosing a topup method.",
		},
		{
			Name:        prefix + "kind",
			Type:        "string",
			Description: "Whether this gateway settles in fiat or crypto.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool",
			Description: "Whether wallet owners can currently start a topup through this gateway.",
		},
		{
			Name:        prefix + "config",
			Type:        "complex",
			Description: "Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.",
		},
		{
			Name:        prefix + "supported-currencies",
			Type:        "complex",
			Description: "JSON array of walletCurrency codes this gateway can top up (e.g. [\"USD\",\"EUR\"]).",
		},
	}
}
func CastWalletGatewayDtoFromCli(c emigo.CliCastable) WalletGatewayDto {
	data := WalletGatewayDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("kind") {
		data.Kind = c.String("kind")
	}
	if c.IsSet("is-active") {
		data.IsActive = bool(c.Bool("is-active"))
	}
	if c.IsSet("config") {
		if u, ok := any(&data.Config).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("config")))
		}
	}
	if c.IsSet("supported-currencies") {
		if u, ok := any(&data.SupportedCurrencies).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("supported-currencies")))
		}
	}
	return data
}
