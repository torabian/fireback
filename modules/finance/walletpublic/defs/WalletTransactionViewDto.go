package walletpublicdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletTransactionViewDto
type WalletTransactionViewDto struct {
	// Unique id of this ledger entry.
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// "credit" or "debit".
	Direction string `json:"direction" yaml:"direction"`
	// Magnitude of the change, as a positive minor-units string.
	Amount string `json:"amount" yaml:"amount"`
	// Wallet balance immediately after this entry, as a minor-units string.
	BalanceAfter string `json:"balanceAfter" yaml:"balanceAfter"`
	// What kind of event produced this entry.
	Reason string `json:"reason" yaml:"reason"`
	// Free-form name of the module/feature that caused this entry.
	ReferenceType emigo.Nullable[string] `json:"referenceType" yaml:"referenceType"`
	// Id within referenceType this entry relates to.
	ReferenceId emigo.Nullable[string] `json:"referenceId" yaml:"referenceId"`
	// Optional human-readable note.
	Note emigo.Nullable[string] `json:"note" yaml:"note"`
	// When this entry was recorded.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
}

func (x *WalletTransactionViewDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletTransactionViewDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Unique id of this ledger entry.",
		},
		{
			Name:        prefix + "direction",
			Type:        "string",
			Description: "\"credit\" or \"debit\".",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Magnitude of the change, as a positive minor-units string.",
		},
		{
			Name:        prefix + "balance-after",
			Type:        "string",
			Description: "Wallet balance immediately after this entry, as a minor-units string.",
		},
		{
			Name:        prefix + "reason",
			Type:        "string",
			Description: "What kind of event produced this entry.",
		},
		{
			Name:        prefix + "reference-type",
			Type:        "string?",
			Description: "Free-form name of the module/feature that caused this entry.",
		},
		{
			Name:        prefix + "reference-id",
			Type:        "string?",
			Description: "Id within referenceType this entry relates to.",
		},
		{
			Name:        prefix + "note",
			Type:        "string?",
			Description: "Optional human-readable note.",
		},
		{
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "When this entry was recorded.",
		},
	}
}
func CastWalletTransactionViewDtoFromCli(c emigo.CliCastable) WalletTransactionViewDto {
	data := WalletTransactionViewDto{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("direction") {
		data.Direction = c.String("direction")
	}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("balance-after") {
		data.BalanceAfter = c.String("balance-after")
	}
	if c.IsSet("reason") {
		data.Reason = c.String("reason")
	}
	if c.IsSet("reference-type") {
		emigo.ParseNullable(c.String("reference-type"), &data.ReferenceType)
	}
	if c.IsSet("reference-id") {
		emigo.ParseNullable(c.String("reference-id"), &data.ReferenceId)
	}
	if c.IsSet("note") {
		emigo.ParseNullable(c.String("note"), &data.Note)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	return data
}
