package walletdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for walletEntity
type WalletEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
	OwnerType string `json:"ownerType" validate:"required,oneof=user workspace" yaml:"ownerType"`
	// Unique id of the owning user. Set only when ownerType is "user".
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// Unique id of the owning workspace. Set only when ownerType is "workspace".
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
	Currency string `json:"currency" validate:"required" yaml:"currency"`
	// Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
	Balance string `json:"balance" yaml:"balance"`
	// "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
	Status string `json:"status" validate:"required,oneof=active frozen closed" yaml:"status"`
	// Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
	Label emigo.Nullable[string] `json:"label" yaml:"label"`
	// Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
	IsDefault bool `json:"isDefault" yaml:"isDefault"`
	// Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
	Version int64 `json:"version" yaml:"version"`
}

func (x *WalletEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "owner-type",
			Type:        "string",
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
			Type:        "string",
			Description: "Currency code this wallet holds (must match an active walletCurrency.code, e.g. \"USD\" or \"BTC\"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.",
		},
		{
			Name:        prefix + "balance",
			Type:        "string",
			Description: "Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.",
		},
		{
			Name:        prefix + "status",
			Type:        "string",
			Description: "\"active\" wallets can be topped up and purchased from. \"frozen\" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). \"closed\" is a terminal, non-reactivatable state.",
		},
		{
			Name:        prefix + "label",
			Type:        "string?",
			Description: "Optional owner-given nickname for the wallet (e.g. \"Main\", \"Savings\") - purely cosmetic, shown in the owner-facing UI.",
		},
		{
			Name:        prefix + "is-default",
			Type:        "bool",
			Description: "Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.",
		},
		{
			Name:        prefix + "version",
			Type:        "int64",
			Description: "Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal (\"has this wallet ever moved\").",
		},
	}
}
func CastWalletEntityFromCli(c emigo.CliCastable) WalletEntity {
	data := WalletEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("owner-type") {
		data.OwnerType = c.String("owner-type")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
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
	if c.IsSet("version") {
		data.Version = int64(c.Int64("version"))
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletEntityCreateFn creates a new WalletEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletEntityCreateFn(tx *gorm.DB, dto *WalletEntity) (*WalletEntity, error) {
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

// WalletEntityUpdateFn applies a partial update to the WalletEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletOptionalDto) (*WalletEntity, error) {
	var entity WalletEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.OwnerType.IsSet() {
			changes["OwnerType"] = input.OwnerType
		}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
		}
		if input.Currency.IsSet() {
			changes["Currency"] = input.Currency
		}
		if input.Balance.IsSet() {
			changes["Balance"] = input.Balance
		}
		if input.Status.IsSet() {
			changes["Status"] = input.Status
		}
		if input.Label.IsSet() {
			changes["Label"] = input.Label
		}
		if input.IsDefault.IsSet() {
			changes["IsDefault"] = input.IsDefault
		}
		if input.Version.IsSet() {
			changes["Version"] = input.Version
		}
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
	var updated WalletEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletEntityGetFn looks up a single WalletEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletEntity, error) {
	var entity WalletEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletEntityBrowseFn returns WalletEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletEntityBrowseFn(tx *gorm.DB, qs WalletBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletEntity
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

// WalletEntityAwareDeleteAffected reports one relation of WalletEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletEntity itself, so deleting WalletEntity doesn't cascade into them.
type WalletEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletEntityAwareDeletePreview is the result of WalletEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletEntityAwareDeleteFn would delete/clear
// alongside the WalletEntity row(s) themselves.
type WalletEntityAwareDeletePreview struct {
	Message  string                            `json:"message"`
	Affected []WalletEntityAwareDeleteAffected `json:"affected"`
}

// WalletEntityAwareDeletePreviewFn looks up the WalletEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletEntityAwareDeleteFn.
func WalletEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletEntityAwareDeletePreview, error) {
	var rows []*WalletEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletEntityAwareDeletePreview{Message: "No matching WalletEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletEntityAwareDeleteFn deletes the WalletEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletEntity
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
		return tx.Where("id IN ?", ids).Delete(&WalletEntity{}).Error
	})
}

// WalletEntityActionsSig bundles the actions available for WalletEntity. Extend this (and
// WalletEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletEntity) (*WalletEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletOptionalDto) (*WalletEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletEntity, error)
	Browse             func(tx *gorm.DB, qs WalletBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletEntityActions WalletEntityActionsSig = WalletEntityActionsSig{
	Create:             WalletEntityCreateFn,
	Update:             WalletEntityUpdateFn,
	Get:                WalletEntityGetFn,
	Browse:             WalletEntityBrowseFn,
	AwareDeletePreview: WalletEntityAwareDeletePreviewFn,
	AwareDelete:        WalletEntityAwareDeleteFn,
}
