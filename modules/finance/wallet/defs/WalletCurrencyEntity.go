package walletdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for walletCurrencyEntity
type WalletCurrencyEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
	Code string `gorm:"unique;not null;size:20" json:"code" validate:"required" yaml:"code"`
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

func (x *WalletCurrencyEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletCurrencyEntityCliFlags(prefix string) []emigo.CliFlag {
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
func CastWalletCurrencyEntityFromCli(c emigo.CliCastable) WalletCurrencyEntity {
	data := WalletCurrencyEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
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

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletCurrencyEntityCreateFn creates a new WalletCurrencyEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletCurrencyEntityCreateFn(tx *gorm.DB, dto *WalletCurrencyEntity) (*WalletCurrencyEntity, error) {
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

// WalletCurrencyEntityUpdateFn applies a partial update to the WalletCurrencyEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletCurrencyEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletCurrencyEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletCurrencyOptionalDto) (*WalletCurrencyEntity, error) {
	var entity WalletCurrencyEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Code.IsSet() {
			changes["Code"] = input.Code
		}
		if input.Name.IsSet() {
			changes["Name"] = input.Name
		}
		if input.Kind.IsSet() {
			changes["Kind"] = input.Kind
		}
		if input.Decimals.IsSet() {
			changes["Decimals"] = input.Decimals
		}
		if input.Symbol.IsSet() {
			changes["Symbol"] = input.Symbol
		}
		if input.IsActive.IsSet() {
			changes["IsActive"] = input.IsActive
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
	var updated WalletCurrencyEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletCurrencyEntityGetFn looks up a single WalletCurrencyEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletCurrencyEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletCurrencyEntity, error) {
	var entity WalletCurrencyEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletCurrencyEntityBrowseFn returns WalletCurrencyEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletCurrencyEntityBrowseFn(tx *gorm.DB, qs WalletCurrencyBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletCurrencyEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletCurrencyEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletCurrencyEntity
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

// WalletCurrencyEntityAwareDeleteAffected reports one relation of WalletCurrencyEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletCurrencyEntity itself, so deleting WalletCurrencyEntity doesn't cascade into them.
type WalletCurrencyEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletCurrencyEntityAwareDeletePreview is the result of WalletCurrencyEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletCurrencyEntityAwareDeleteFn would delete/clear
// alongside the WalletCurrencyEntity row(s) themselves.
type WalletCurrencyEntityAwareDeletePreview struct {
	Message  string                                    `json:"message"`
	Affected []WalletCurrencyEntityAwareDeleteAffected `json:"affected"`
}

// WalletCurrencyEntityAwareDeletePreviewFn looks up the WalletCurrencyEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletCurrencyEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletCurrencyEntityAwareDeleteFn.
func WalletCurrencyEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletCurrencyEntityAwareDeletePreview, error) {
	var rows []*WalletCurrencyEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletCurrencyEntityAwareDeletePreview{Message: "No matching WalletCurrencyEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletCurrencyEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletCurrencyEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletCurrencyEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletCurrencyEntityAwareDeleteFn deletes the WalletCurrencyEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletCurrencyEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletCurrencyEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletCurrencyEntity
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
		return tx.Where("id IN ?", ids).Delete(&WalletCurrencyEntity{}).Error
	})
}

// WalletCurrencyEntityActionsSig bundles the actions available for WalletCurrencyEntity. Extend this (and
// WalletCurrencyEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletCurrencyEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletCurrencyEntity) (*WalletCurrencyEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletCurrencyOptionalDto) (*WalletCurrencyEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletCurrencyEntity, error)
	Browse             func(tx *gorm.DB, qs WalletCurrencyBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletCurrencyEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletCurrencyEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletCurrencyEntityActions WalletCurrencyEntityActionsSig = WalletCurrencyEntityActionsSig{
	Create:             WalletCurrencyEntityCreateFn,
	Update:             WalletCurrencyEntityUpdateFn,
	Get:                WalletCurrencyEntityGetFn,
	Browse:             WalletCurrencyEntityBrowseFn,
	AwareDeletePreview: WalletCurrencyEntityAwareDeletePreviewFn,
	AwareDelete:        WalletCurrencyEntityAwareDeleteFn,
}
