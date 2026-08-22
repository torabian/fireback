package walletdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletTransactionOptionalDto
type WalletTransactionOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The wallet this ledger entry belongs to.
	Wallet emigo.OneNullable[WalletDto] `json:"wallet" validate:"required" yaml:"wallet"`
	// "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
	Direction emigo.Nullable[string] `json:"direction" validate:"required,oneof=credit debit" yaml:"direction"`
	// Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
	Amount emigo.Nullable[string] `json:"amount" validate:"required" yaml:"amount"`
	// Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
	BalanceAfter emigo.Nullable[string] `json:"balanceAfter" validate:"required" yaml:"balanceAfter"`
	// What kind of event produced this ledger entry.
	Reason emigo.Nullable[string] `json:"reason" validate:"required,oneof=topup purchase refund adjustment transfer_in transfer_out fee chargeback" yaml:"reason"`
	// Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
	ReferenceType emigo.Nullable[string] `json:"referenceType" yaml:"referenceType"`
	// Id within referenceType this entry relates to (e.g. the order id in the calling module).
	ReferenceId emigo.Nullable[string] `json:"referenceId" yaml:"referenceId"`
	// Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
	IdempotencyKey emigo.Nullable[string] `json:"idempotencyKey" validate:"required" yaml:"idempotencyKey"`
	// Optional human-readable note, e.g. why an adjustment was made.
	Note emigo.Nullable[string] `json:"note" yaml:"note"`
	// Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
	CreatedBy emigo.Nullable[string] `json:"createdBy" yaml:"createdBy"`
	// When this ledger entry was recorded.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
	// Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
	Metadata complexes.JSON `json:"metadata" yaml:"metadata"`
}

func (x *WalletTransactionOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletTransactionOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "wallet",
			Type:        "one?",
			Description: "The wallet this ledger entry belongs to.",
		},
		{
			Name:        prefix + "direction",
			Type:        "string?",
			Description: "\"credit\" increases the wallet's balance, \"debit\" decreases it. amount is always a positive minor-units string; direction carries the sign.",
		},
		{
			Name:        prefix + "amount",
			Type:        "string?",
			Description: "Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.",
		},
		{
			Name:        prefix + "balance-after",
			Type:        "string?",
			Description: "Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.",
		},
		{
			Name:        prefix + "reason",
			Type:        "string?",
			Description: "What kind of event produced this ledger entry.",
		},
		{
			Name:        prefix + "reference-type",
			Type:        "string?",
			Description: "Free-form identifier of the calling module/feature that caused this entry, e.g. \"course-purchase\" - lets other nima modules tag their own purchases without editing this module's yaml.",
		},
		{
			Name:        prefix + "reference-id",
			Type:        "string?",
			Description: "Id within referenceType this entry relates to (e.g. the order id in the calling module).",
		},
		{
			Name:        prefix + "idempotency-key",
			Type:        "string?",
			Description: "Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.",
		},
		{
			Name:        prefix + "note",
			Type:        "string?",
			Description: "Optional human-readable note, e.g. why an adjustment was made.",
		},
		{
			Name:        prefix + "created-by",
			Type:        "string?",
			Description: "Unique id of the user who triggered this entry, or a fixed marker such as \"system\" or the gateway code, for entries with no human actor.",
		},
		{
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "When this ledger entry was recorded.",
		},
		{
			Name:        prefix + "metadata",
			Type:        "complex",
			Description: "Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.",
		},
	}
}
func CastWalletTransactionOptionalDtoFromCli(c emigo.CliCastable) WalletTransactionOptionalDto {
	data := WalletTransactionOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("wallet") {
		data.Wallet = emigo.CapturePossibleOneNullable(CastWalletDtoFromCli, "wallet", c)
	}
	if c.IsSet("direction") {
		emigo.ParseNullable(c.String("direction"), &data.Direction)
	}
	if c.IsSet("amount") {
		emigo.ParseNullable(c.String("amount"), &data.Amount)
	}
	if c.IsSet("balance-after") {
		emigo.ParseNullable(c.String("balance-after"), &data.BalanceAfter)
	}
	if c.IsSet("reason") {
		emigo.ParseNullable(c.String("reason"), &data.Reason)
	}
	if c.IsSet("reference-type") {
		emigo.ParseNullable(c.String("reference-type"), &data.ReferenceType)
	}
	if c.IsSet("reference-id") {
		emigo.ParseNullable(c.String("reference-id"), &data.ReferenceId)
	}
	if c.IsSet("idempotency-key") {
		emigo.ParseNullable(c.String("idempotency-key"), &data.IdempotencyKey)
	}
	if c.IsSet("note") {
		emigo.ParseNullable(c.String("note"), &data.Note)
	}
	if c.IsSet("created-by") {
		emigo.ParseNullable(c.String("created-by"), &data.CreatedBy)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("metadata") {
		if u, ok := any(&data.Metadata).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("metadata")))
		}
	}
	return data
}
