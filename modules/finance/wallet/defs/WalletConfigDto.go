package walletdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for walletConfigDto
type WalletConfigDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
	MaxWalletsPerUser int64 `json:"maxWalletsPerUser" yaml:"maxWalletsPerUser"`
	// Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
	MaxWalletsPerWorkspace int64 `json:"maxWalletsPerWorkspace" yaml:"maxWalletsPerWorkspace"`
	// Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
	MaxWalletsPerUserPerCurrency emigo.Nullable[int64] `json:"maxWalletsPerUserPerCurrency" yaml:"maxWalletsPerUserPerCurrency"`
	// Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
	MaxWalletsPerWorkspacePerCurrency emigo.Nullable[int64] `json:"maxWalletsPerWorkspacePerCurrency" yaml:"maxWalletsPerWorkspacePerCurrency"`
}

func (x *WalletConfigDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletConfigDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "max-wallets-per-user",
			Type:        "int64",
			Description: "Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.",
		},
		{
			Name:        prefix + "max-wallets-per-workspace",
			Type:        "int64",
			Description: "Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.",
		},
		{
			Name:        prefix + "max-wallets-per-user-per-currency",
			Type:        "int64?",
			Description: "Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.",
		},
		{
			Name:        prefix + "max-wallets-per-workspace-per-currency",
			Type:        "int64?",
			Description: "Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.",
		},
	}
}
func CastWalletConfigDtoFromCli(c emigo.CliCastable) WalletConfigDto {
	data := WalletConfigDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("max-wallets-per-user") {
		data.MaxWalletsPerUser = int64(c.Int64("max-wallets-per-user"))
	}
	if c.IsSet("max-wallets-per-workspace") {
		data.MaxWalletsPerWorkspace = int64(c.Int64("max-wallets-per-workspace"))
	}
	if c.IsSet("max-wallets-per-user-per-currency") {
		emigo.ParseNullable(c.String("max-wallets-per-user-per-currency"), &data.MaxWalletsPerUserPerCurrency)
	}
	if c.IsSet("max-wallets-per-workspace-per-currency") {
		emigo.ParseNullable(c.String("max-wallets-per-workspace-per-currency"), &data.MaxWalletsPerWorkspacePerCurrency)
	}
	return data
}
