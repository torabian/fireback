package walletdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

// The base class definition for walletConfigEntity
type WalletConfigEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
	MaxWalletsPerUser int64 `json:"maxWalletsPerUser" yaml:"maxWalletsPerUser"`
	// Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
	MaxWalletsPerWorkspace int64 `json:"maxWalletsPerWorkspace" yaml:"maxWalletsPerWorkspace"`
	// Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
	MaxWalletsPerUserPerCurrency emigo.Nullable[int64] `json:"maxWalletsPerUserPerCurrency" yaml:"maxWalletsPerUserPerCurrency"`
	// Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
	MaxWalletsPerWorkspacePerCurrency emigo.Nullable[int64] `json:"maxWalletsPerWorkspacePerCurrency" yaml:"maxWalletsPerWorkspacePerCurrency"`
}

func (x *WalletConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletConfigEntityCliFlags(prefix string) []emigo.CliFlag {
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
func CastWalletConfigEntityFromCli(c emigo.CliCastable) WalletConfigEntity {
	data := WalletConfigEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
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

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletConfigEntityCreateFn creates a new WalletConfigEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletConfigEntityCreateFn(tx *gorm.DB, dto *WalletConfigEntity) (*WalletConfigEntity, error) {
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

// WalletConfigEntityUpdateFn applies a partial update to the WalletConfigEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletConfigEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletConfigEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletConfigOptionalDto) (*WalletConfigEntity, error) {
	var entity WalletConfigEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.MaxWalletsPerUser.IsSet() {
			changes["MaxWalletsPerUser"] = input.MaxWalletsPerUser
		}
		if input.MaxWalletsPerWorkspace.IsSet() {
			changes["MaxWalletsPerWorkspace"] = input.MaxWalletsPerWorkspace
		}
		if input.MaxWalletsPerUserPerCurrency.IsSet() {
			changes["MaxWalletsPerUserPerCurrency"] = input.MaxWalletsPerUserPerCurrency
		}
		if input.MaxWalletsPerWorkspacePerCurrency.IsSet() {
			changes["MaxWalletsPerWorkspacePerCurrency"] = input.MaxWalletsPerWorkspacePerCurrency
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
	var updated WalletConfigEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletConfigEntityGetFn looks up a single WalletConfigEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletConfigEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletConfigEntity, error) {
	var entity WalletConfigEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletConfigEntityBrowseFn returns WalletConfigEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletConfigEntityBrowseFn(tx *gorm.DB, qs WalletConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletConfigEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletConfigEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletConfigEntity
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

/**
 * Query parameters for WalletConfigBrowseAction
 */
// Query wrapper with private fields
type WalletConfigBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func WalletConfigBrowseActionQueryFromString(rawQuery string) WalletConfigBrowseActionQuery {
	v := WalletConfigBrowseActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func WalletConfigBrowseActionQueryFromHttp(r *http.Request) WalletConfigBrowseActionQuery {
	return WalletConfigBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q WalletConfigBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q WalletConfigBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *WalletConfigBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *WalletConfigBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

// WalletConfigEntityAwareDeleteAffected reports one relation of WalletConfigEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletConfigEntity itself, so deleting WalletConfigEntity doesn't cascade into them.
type WalletConfigEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletConfigEntityAwareDeletePreview is the result of WalletConfigEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletConfigEntityAwareDeleteFn would delete/clear
// alongside the WalletConfigEntity row(s) themselves.
type WalletConfigEntityAwareDeletePreview struct {
	Message  string                                  `json:"message"`
	Affected []WalletConfigEntityAwareDeleteAffected `json:"affected"`
}

// WalletConfigEntityAwareDeletePreviewFn looks up the WalletConfigEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletConfigEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletConfigEntityAwareDeleteFn.
func WalletConfigEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletConfigEntityAwareDeletePreview, error) {
	var rows []*WalletConfigEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletConfigEntityAwareDeletePreview{Message: "No matching WalletConfigEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletConfigEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletConfigEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletConfigEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletConfigEntityAwareDeleteFn deletes the WalletConfigEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletConfigEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletConfigEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletConfigEntity
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
		return tx.Where("id IN ?", ids).Delete(&WalletConfigEntity{}).Error
	})
}

// WalletConfigEntityActionsSig bundles the actions available for WalletConfigEntity. Extend this (and
// WalletConfigEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletConfigEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletConfigEntity) (*WalletConfigEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletConfigOptionalDto) (*WalletConfigEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletConfigEntity, error)
	Browse             func(tx *gorm.DB, qs WalletConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletConfigEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletConfigEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletConfigEntityActions WalletConfigEntityActionsSig = WalletConfigEntityActionsSig{
	Create:             WalletConfigEntityCreateFn,
	Update:             WalletConfigEntityUpdateFn,
	Get:                WalletConfigEntityGetFn,
	Browse:             WalletConfigEntityBrowseFn,
	AwareDeletePreview: WalletConfigEntityAwareDeletePreviewFn,
	AwareDelete:        WalletConfigEntityAwareDeleteFn,
}
