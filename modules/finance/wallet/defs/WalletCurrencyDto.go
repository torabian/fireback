package walletdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for walletCurrencyDto
type WalletCurrencyDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
	Code string `json:"code" validate:"required" yaml:"code"`
	// Display name, e.g. "US Dollar" or "Bitcoin".
	Name string `json:"name" validate:"required" yaml:"name"`
	// Whether this is a fiat or crypto currency.
	Kind string `json:"kind" validate:"required,oneof=fiat crypto" yaml:"kind"`
	// Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
	Decimals int `json:"decimals" validate:"required" yaml:"decimals"`
	// Optional display symbol, e.g. "$" or "₿".
	Symbol emigo.Nullable[string] `json:"symbol" yaml:"symbol"`
	// Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
	IsActive bool `json:"isActive" yaml:"isActive"`
}

func (x *WalletCurrencyDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletCurrencyDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "code",
			Type:        "string",
			Description: "Currency code, e.g. \"USD\", \"EUR\", \"BTC\", \"ETH\". Unique across all currencies.",
		},
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Display name, e.g. \"US Dollar\" or \"Bitcoin\".",
		},
		{
			Name:        prefix + "kind",
			Type:        "string",
			Description: "Whether this is a fiat or crypto currency.",
		},
		{
			Name:        prefix + "decimals",
			Type:        "int",
			Description: "Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so \"10050\" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.",
		},
		{
			Name:        prefix + "symbol",
			Type:        "string?",
			Description: "Optional display symbol, e.g. \"$\" or \"₿\".",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool",
			Description: "Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.",
		},
	}
}
func CastWalletCurrencyDtoFromCli(c emigo.CliCastable) WalletCurrencyDto {
	data := WalletCurrencyDto{}
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
	if c.IsSet("decimals") {
		data.Decimals = int(c.Int64("decimals"))
	}
	if c.IsSet("symbol") {
		emigo.ParseNullable(c.String("symbol"), &data.Symbol)
	}
	if c.IsSet("is-active") {
		data.IsActive = bool(c.Bool("is-active"))
	}
	return data
}
