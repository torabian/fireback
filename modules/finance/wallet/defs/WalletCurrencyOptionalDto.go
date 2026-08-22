package walletdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for walletCurrencyOptionalDto
type WalletCurrencyOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
	Code emigo.Nullable[string] `json:"code" validate:"required" yaml:"code"`
	// Display name, e.g. "US Dollar" or "Bitcoin".
	Name emigo.Nullable[string] `json:"name" validate:"required" yaml:"name"`
	// Whether this is a fiat or crypto currency.
	Kind emigo.Nullable[string] `json:"kind" validate:"required,oneof=fiat crypto" yaml:"kind"`
	// Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
	Decimals emigo.Nullable[int] `json:"decimals" validate:"required" yaml:"decimals"`
	// Optional display symbol, e.g. "$" or "₿".
	Symbol emigo.Nullable[string] `json:"symbol" yaml:"symbol"`
	// Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
	IsActive emigo.Nullable[bool] `json:"isActive" yaml:"isActive"`
}

func (x *WalletCurrencyOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletCurrencyOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "code",
			Type:        "string?",
			Description: "Currency code, e.g. \"USD\", \"EUR\", \"BTC\", \"ETH\". Unique across all currencies.",
		},
		{
			Name:        prefix + "name",
			Type:        "string?",
			Description: "Display name, e.g. \"US Dollar\" or \"Bitcoin\".",
		},
		{
			Name:        prefix + "kind",
			Type:        "string?",
			Description: "Whether this is a fiat or crypto currency.",
		},
		{
			Name:        prefix + "decimals",
			Type:        "int?",
			Description: "Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so \"10050\" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.",
		},
		{
			Name:        prefix + "symbol",
			Type:        "string?",
			Description: "Optional display symbol, e.g. \"$\" or \"₿\".",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool?",
			Description: "Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.",
		},
	}
}
func CastWalletCurrencyOptionalDtoFromCli(c emigo.CliCastable) WalletCurrencyOptionalDto {
	data := WalletCurrencyOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("code") {
		emigo.ParseNullable(c.String("code"), &data.Code)
	}
	if c.IsSet("name") {
		emigo.ParseNullable(c.String("name"), &data.Name)
	}
	if c.IsSet("kind") {
		emigo.ParseNullable(c.String("kind"), &data.Kind)
	}
	if c.IsSet("decimals") {
		emigo.ParseNullable(c.String("decimals"), &data.Decimals)
	}
	if c.IsSet("symbol") {
		emigo.ParseNullable(c.String("symbol"), &data.Symbol)
	}
	if c.IsSet("is-active") {
		emigo.ParseNullable(c.String("is-active"), &data.IsActive)
	}
	return data
}
