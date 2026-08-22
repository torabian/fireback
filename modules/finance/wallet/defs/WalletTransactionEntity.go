package walletdefs

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/fireback/complexes"
	"gorm.io/gorm"
)

// The base class definition for walletTransactionEntity
type WalletTransactionEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The wallet this ledger entry belongs to.
	Wallet *WalletEntity `gorm:"foreignKey:WalletId;references:Id" json:"wallet" validate:"required" yaml:"wallet"`
	// "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
	Direction string `json:"direction" validate:"required,oneof=credit debit" yaml:"direction"`
	// Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
	Amount string `json:"amount" validate:"required" yaml:"amount"`
	// Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
	BalanceAfter string `json:"balanceAfter" validate:"required" yaml:"balanceAfter"`
	// What kind of event produced this ledger entry.
	Reason string `json:"reason" validate:"required,oneof=topup purchase refund adjustment transfer_in transfer_out fee chargeback" yaml:"reason"`
	// Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
	ReferenceType emigo.Nullable[string] `json:"referenceType" yaml:"referenceType"`
	// Id within referenceType this entry relates to (e.g. the order id in the calling module).
	ReferenceId emigo.Nullable[string] `json:"referenceId" yaml:"referenceId"`
	// Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
	IdempotencyKey string `gorm:"unique;not null" json:"idempotencyKey" validate:"required" yaml:"idempotencyKey"`
	// Optional human-readable note, e.g. why an adjustment was made.
	Note emigo.Nullable[string] `json:"note" yaml:"note"`
	// Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
	CreatedBy emigo.Nullable[string] `json:"createdBy" yaml:"createdBy"`
	// When this ledger entry was recorded.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
	// Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
	Metadata complexes.JSON `json:"metadata" yaml:"metadata"`
	WalletId int64          `gorm:"index" json:"-" yaml:"-"`
}

func (x *WalletTransactionEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletTransactionEntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name:        prefix + "wallet",
			Type:        "class",
			Description: "The wallet this ledger entry belongs to.",
		},
		{
			Name:        prefix + "direction",
			Type:        "string",
			Description: "\"credit\" increases the wallet's balance, \"debit\" decreases it. amount is always a positive minor-units string; direction carries the sign.",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.",
		},
		{
			Name:        prefix + "balance-after",
			Type:        "string",
			Description: "Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.",
		},
		{
			Name:        prefix + "reason",
			Type:        "string",
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
			Type:        "string",
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
		{
			Name: prefix + "wallet-id",
			Type: "int64",
		},
	}
}
func CastWalletTransactionEntityFromCli(c emigo.CliCastable) WalletTransactionEntity {
	data := WalletTransactionEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
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
	if c.IsSet("idempotency-key") {
		data.IdempotencyKey = c.String("idempotency-key")
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
	if c.IsSet("wallet-id") {
		data.WalletId = int64(c.Int64("wallet-id"))
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletTransactionEntityCreateFn creates a new WalletTransactionEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletTransactionEntityCreateFn(tx *gorm.DB, dto *WalletTransactionEntity) (*WalletTransactionEntity, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// WalletTransactionEntityUpdateFn applies a partial update to the WalletTransactionEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletTransactionEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletTransactionEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletTransactionOptionalDto) (*WalletTransactionEntity, error) {
	var entity WalletTransactionEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Wallet.IsSet() {
			if input.Wallet.Operation != "select" {
				return fmt.Errorf("wallet: updating a one/one? relation only supports the \"select\" operation (link to an existing row by its uniqueId), got %q", input.Wallet.Operation)
			}
			var selectorId string
			if s, ok := input.Wallet.Selector.(string); ok {
				selectorId = s
			}
			resolvedId, err := emigorm.ReconcileOne[WalletEntity](tx, input.Wallet.Operation, selectorId, nil)
			if err != nil {
				return err
			}
			changes["WalletId"] = resolvedId
		}
		if input.Direction.IsSet() {
			changes["Direction"] = input.Direction
		}
		if input.Amount.IsSet() {
			changes["Amount"] = input.Amount
		}
		if input.BalanceAfter.IsSet() {
			changes["BalanceAfter"] = input.BalanceAfter
		}
		if input.Reason.IsSet() {
			changes["Reason"] = input.Reason
		}
		if input.ReferenceType.IsSet() {
			changes["ReferenceType"] = input.ReferenceType
		}
		if input.ReferenceId.IsSet() {
			changes["ReferenceId"] = input.ReferenceId
		}
		if input.IdempotencyKey.IsSet() {
			changes["IdempotencyKey"] = input.IdempotencyKey
		}
		if input.Note.IsSet() {
			changes["Note"] = input.Note
		}
		if input.CreatedBy.IsSet() {
			changes["CreatedBy"] = input.CreatedBy
		}
		changes["CreatedAt"] = input.CreatedAt
		changes["Metadata"] = input.Metadata
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var updated WalletTransactionEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletTransactionEntityGetFn looks up a single WalletTransactionEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletTransactionEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletTransactionEntity, error) {
	var entity WalletTransactionEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletTransactionEntityBrowseFn returns WalletTransactionEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletTransactionEntityBrowseFn(tx *gorm.DB, qs WalletTransactionBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletTransactionEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletTransactionEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletTransactionEntity
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, qs.Sort), qs.Cursor), qs.StartIndex, qs.ItemsPerPage)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}
	return items, meta, nil
}

// WalletTransactionEntityAwareDeleteAffected reports one relation of WalletTransactionEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletTransactionEntity itself, so deleting WalletTransactionEntity doesn't cascade into them.
type WalletTransactionEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletTransactionEntityAwareDeletePreview is the result of WalletTransactionEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletTransactionEntityAwareDeleteFn would delete/clear
// alongside the WalletTransactionEntity row(s) themselves.
type WalletTransactionEntityAwareDeletePreview struct {
	Message  string                                       `json:"message"`
	Affected []WalletTransactionEntityAwareDeleteAffected `json:"affected"`
}

// WalletTransactionEntityAwareDeletePreviewFn looks up the WalletTransactionEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletTransactionEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletTransactionEntityAwareDeleteFn.
func WalletTransactionEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletTransactionEntityAwareDeletePreview, error) {
	var rows []*WalletTransactionEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletTransactionEntityAwareDeletePreview{Message: "No matching WalletTransactionEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletTransactionEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletTransactionEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletTransactionEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletTransactionEntityAwareDeleteFn deletes the WalletTransactionEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletTransactionEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletTransactionEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletTransactionEntity
		if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
			return err
		}
		if len(rows) == 0 {
			return nil
		}
		ids := make([]int64, len(rows))
		for i := range rows {
			ids[i] = rows[i].Id
		}
		return tx.Where("id IN ?", ids).Delete(&WalletTransactionEntity{}).Error
	})
}

// WalletTransactionEntityActionsSig bundles the actions available for WalletTransactionEntity. Extend this (and
// WalletTransactionEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletTransactionEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletTransactionEntity) (*WalletTransactionEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletTransactionOptionalDto) (*WalletTransactionEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletTransactionEntity, error)
	Browse             func(tx *gorm.DB, qs WalletTransactionBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletTransactionEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletTransactionEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletTransactionEntityActions WalletTransactionEntityActionsSig = WalletTransactionEntityActionsSig{
	Create:             WalletTransactionEntityCreateFn,
	Update:             WalletTransactionEntityUpdateFn,
	Get:                WalletTransactionEntityGetFn,
	Browse:             WalletTransactionEntityBrowseFn,
	AwareDeletePreview: WalletTransactionEntityAwareDeletePreviewFn,
	AwareDelete:        WalletTransactionEntityAwareDeleteFn,
}
