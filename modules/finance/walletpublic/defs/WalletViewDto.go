package walletpublicdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for walletViewDto
type WalletViewDto struct {
	// Unique id of the wallet.
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// "user" or "workspace".
	OwnerType string `json:"ownerType" yaml:"ownerType"`
	// Owning workspace id, set only when ownerType is "workspace".
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// Currency code this wallet holds.
	Currency string `json:"currency" yaml:"currency"`
	// Current balance as a minor-units decimal string.
	Balance string `json:"balance" yaml:"balance"`
	// "active", "frozen", or "closed".
	Status string `json:"status" yaml:"status"`
	// Owner-given nickname, if any.
	Label emigo.Nullable[string] `json:"label" yaml:"label"`
	// Whether this is the owner's default wallet for its currency.
	IsDefault bool `json:"isDefault" yaml:"isDefault"`
}

func (x *WalletViewDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletViewDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Unique id of the wallet.",
		},
		{
			Name:        prefix + "owner-type",
			Type:        "string",
			Description: "\"user\" or \"workspace\".",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "Owning workspace id, set only when ownerType is \"workspace\".",
		},
		{
			Name:        prefix + "currency",
			Type:        "string",
			Description: "Currency code this wallet holds.",
		},
		{
			Name:        prefix + "balance",
			Type:        "string",
			Description: "Current balance as a minor-units decimal string.",
		},
		{
			Name:        prefix + "status",
			Type:        "string",
			Description: "\"active\", \"frozen\", or \"closed\".",
		},
		{
			Name:        prefix + "label",
			Type:        "string?",
			Description: "Owner-given nickname, if any.",
		},
		{
			Name:        prefix + "is-default",
			Type:        "bool",
			Description: "Whether this is the owner's default wallet for its currency.",
		},
	}
}
func CastWalletViewDtoFromCli(c emigo.CliCastable) WalletViewDto {
	data := WalletViewDto{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("owner-type") {
		data.OwnerType = c.String("owner-type")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	if c.IsSet("balance") {
		data.Balance = c.String("balance")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("label") {
		emigo.ParseNullable(c.String("label"), &data.Label)
	}
	if c.IsSet("is-default") {
		data.IsDefault = bool(c.Bool("is-default"))
	}
	return data
}
