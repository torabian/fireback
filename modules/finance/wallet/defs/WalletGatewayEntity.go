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

// The base class definition for walletGatewayEntity
type WalletGatewayEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
	Code string `gorm:"unique;not null;size:50" json:"code" validate:"required" yaml:"code"`
	// Display name shown to wallet owners choosing a topup method.
	Name string `json:"name" validate:"required" yaml:"name"`
	// Whether this gateway settles in fiat or crypto.
	Kind string `json:"kind" validate:"required,oneof=fiat crypto" yaml:"kind"`
	// Whether wallet owners can currently start a topup through this gateway.
	IsActive bool `json:"isActive" yaml:"isActive"`
	// Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
	Config complexes.JSON `json:"config" yaml:"config"`
	// JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
	SupportedCurrencies complexes.JSON `json:"supportedCurrencies" yaml:"supportedCurrencies"`
}

func (x *WalletGatewayEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletGatewayEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Description: "Unique code identifying this gateway's GatewayAdapter implementation, e.g. \"stripe\", \"manual\", \"onchain_eth\".",
		},
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Display name shown to wallet owners choosing a topup method.",
		},
		{
			Name:        prefix + "kind",
			Type:        "string",
			Description: "Whether this gateway settles in fiat or crypto.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool",
			Description: "Whether wallet owners can currently start a topup through this gateway.",
		},
		{
			Name:        prefix + "config",
			Type:        "complex",
			Description: "Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.",
		},
		{
			Name:        prefix + "supported-currencies",
			Type:        "complex",
			Description: "JSON array of walletCurrency codes this gateway can top up (e.g. [\"USD\",\"EUR\"]).",
		},
	}
}
func CastWalletGatewayEntityFromCli(c emigo.CliCastable) WalletGatewayEntity {
	data := WalletGatewayEntity{}
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
	if c.IsSet("is-active") {
		data.IsActive = bool(c.Bool("is-active"))
	}
	if c.IsSet("config") {
		if u, ok := any(&data.Config).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("config")))
		}
	}
	if c.IsSet("supported-currencies") {
		if u, ok := any(&data.SupportedCurrencies).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("supported-currencies")))
		}
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletGatewayEntityCreateFn creates a new WalletGatewayEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletGatewayEntityCreateFn(tx *gorm.DB, dto *WalletGatewayEntity) (*WalletGatewayEntity, error) {
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

// WalletGatewayEntityUpdateFn applies a partial update to the WalletGatewayEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletGatewayEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletGatewayEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletGatewayOptionalDto) (*WalletGatewayEntity, error) {
	var entity WalletGatewayEntity
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
		if input.IsActive.IsSet() {
			changes["IsActive"] = input.IsActive
		}
		changes["Config"] = input.Config
		changes["SupportedCurrencies"] = input.SupportedCurrencies
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
	var updated WalletGatewayEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletGatewayEntityGetFn looks up a single WalletGatewayEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletGatewayEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletGatewayEntity, error) {
	var entity WalletGatewayEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletGatewayEntityBrowseFn returns WalletGatewayEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletGatewayEntityBrowseFn(tx *gorm.DB, qs WalletGatewayBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletGatewayEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletGatewayEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletGatewayEntity
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

// WalletGatewayEntityAwareDeleteAffected reports one relation of WalletGatewayEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletGatewayEntity itself, so deleting WalletGatewayEntity doesn't cascade into them.
type WalletGatewayEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletGatewayEntityAwareDeletePreview is the result of WalletGatewayEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletGatewayEntityAwareDeleteFn would delete/clear
// alongside the WalletGatewayEntity row(s) themselves.
type WalletGatewayEntityAwareDeletePreview struct {
	Message  string                                   `json:"message"`
	Affected []WalletGatewayEntityAwareDeleteAffected `json:"affected"`
}

// WalletGatewayEntityAwareDeletePreviewFn looks up the WalletGatewayEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletGatewayEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletGatewayEntityAwareDeleteFn.
func WalletGatewayEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletGatewayEntityAwareDeletePreview, error) {
	var rows []*WalletGatewayEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletGatewayEntityAwareDeletePreview{Message: "No matching WalletGatewayEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletGatewayEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletGatewayEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletGatewayEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletGatewayEntityAwareDeleteFn deletes the WalletGatewayEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletGatewayEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletGatewayEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletGatewayEntity
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
		return tx.Where("id IN ?", ids).Delete(&WalletGatewayEntity{}).Error
	})
}

// WalletGatewayEntityActionsSig bundles the actions available for WalletGatewayEntity. Extend this (and
// WalletGatewayEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletGatewayEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletGatewayEntity) (*WalletGatewayEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletGatewayOptionalDto) (*WalletGatewayEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletGatewayEntity, error)
	Browse             func(tx *gorm.DB, qs WalletGatewayBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletGatewayEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletGatewayEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletGatewayEntityActions WalletGatewayEntityActionsSig = WalletGatewayEntityActionsSig{
	Create:             WalletGatewayEntityCreateFn,
	Update:             WalletGatewayEntityUpdateFn,
	Get:                WalletGatewayEntityGetFn,
	Browse:             WalletGatewayEntityBrowseFn,
	AwareDeletePreview: WalletGatewayEntityAwareDeletePreviewFn,
	AwareDelete:        WalletGatewayEntityAwareDeleteFn,
}
