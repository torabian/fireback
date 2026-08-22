package walletdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for walletOptionalDto
type WalletOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
	OwnerType emigo.Nullable[string] `json:"ownerType" validate:"required,oneof=user workspace" yaml:"ownerType"`
	// Unique id of the owning user. Set only when ownerType is "user".
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// Unique id of the owning workspace. Set only when ownerType is "workspace".
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
	Currency emigo.Nullable[string] `json:"currency" validate:"required" yaml:"currency"`
	// Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
	Balance emigo.Nullable[string] `json:"balance" yaml:"balance"`
	// "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
	Status emigo.Nullable[string] `json:"status" validate:"required,oneof=active frozen closed" yaml:"status"`
	// Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
	Label emigo.Nullable[string] `json:"label" yaml:"label"`
	// Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
	IsDefault emigo.Nullable[bool] `json:"isDefault" yaml:"isDefault"`
	// Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
	Version emigo.Nullable[int64] `json:"version" yaml:"version"`
}

func (x *WalletOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "owner-type",
			Type:        "string?",
			Description: "Who owns this wallet: \"user\" (userId is set, workspaceId empty) or \"workspace\" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "Unique id of the owning user. Set only when ownerType is \"user\".",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "Unique id of the owning workspace. Set only when ownerType is \"workspace\".",
		},
		{
			Name:        prefix + "currency",
			Type:        "string?",
			Description: "Currency code this wallet holds (must match an active walletCurrency.code, e.g. \"USD\" or \"BTC\"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.",
		},
		{
			Name:        prefix + "balance",
			Type:        "string?",
			Description: "Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.",
		},
		{
			Name:        prefix + "status",
			Type:        "string?",
			Description: "\"active\" wallets can be topped up and purchased from. \"frozen\" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). \"closed\" is a terminal, non-reactivatable state.",
		},
		{
			Name:        prefix + "label",
			Type:        "string?",
			Description: "Optional owner-given nickname for the wallet (e.g. \"Main\", \"Savings\") - purely cosmetic, shown in the owner-facing UI.",
		},
		{
			Name:        prefix + "is-default",
			Type:        "bool?",
			Description: "Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.",
		},
		{
			Name:        prefix + "version",
			Type:        "int64?",
			Description: "Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal (\"has this wallet ever moved\").",
		},
	}
}
func CastWalletOptionalDtoFromCli(c emigo.CliCastable) WalletOptionalDto {
	data := WalletOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("owner-type") {
		emigo.ParseNullable(c.String("owner-type"), &data.OwnerType)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("currency") {
		emigo.ParseNullable(c.String("currency"), &data.Currency)
	}
	if c.IsSet("balance") {
		emigo.ParseNullable(c.String("balance"), &data.Balance)
	}
	if c.IsSet("status") {
		emigo.ParseNullable(c.String("status"), &data.Status)
	}
	if c.IsSet("label") {
		emigo.ParseNullable(c.String("label"), &data.Label)
	}
	if c.IsSet("is-default") {
		emigo.ParseNullable(c.String("is-default"), &data.IsDefault)
	}
	if c.IsSet("version") {
		emigo.ParseNullable(c.String("version"), &data.Version)
	}
	return data
}
