package walletdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletProviderConfigOptionalDto
type WalletProviderConfigOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
	ProviderType emigo.Nullable[string] `json:"providerType" validate:"required" yaml:"providerType"`
	// Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
	Region emigo.Nullable[string] `json:"region" validate:"required" yaml:"region"`
	// Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
	IsEnabled emigo.Nullable[bool] `json:"isEnabled" yaml:"isEnabled"`
	// Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
	Config complexes.JSON `json:"config" yaml:"config"`
}

func (x *WalletProviderConfigOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletProviderConfigOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "provider-type",
			Type:        "string?",
			Description: "Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. \"stripe\", \"przelewy24\", \"blik\", \"zarinpal\". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.",
		},
		{
			Name:        prefix + "region",
			Type:        "string?",
			Description: "Which region this config applies to - an ISO-3166-ish code (e.g. \"PL\", \"IR\", \"US\") or \"global\", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.",
		},
		{
			Name:        prefix + "is-enabled",
			Type:        "bool?",
			Description: "Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.",
		},
		{
			Name:        prefix + "config",
			Type:        "complex",
			Description: "Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.",
		},
	}
}
func CastWalletProviderConfigOptionalDtoFromCli(c emigo.CliCastable) WalletProviderConfigOptionalDto {
	data := WalletProviderConfigOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("provider-type") {
		emigo.ParseNullable(c.String("provider-type"), &data.ProviderType)
	}
	if c.IsSet("region") {
		emigo.ParseNullable(c.String("region"), &data.Region)
	}
	if c.IsSet("is-enabled") {
		emigo.ParseNullable(c.String("is-enabled"), &data.IsEnabled)
	}
	if c.IsSet("config") {
		if u, ok := any(&data.Config).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("config")))
		}
	}
	return data
}
