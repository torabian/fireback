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

// The base class definition for walletProviderConfigEntity
type WalletProviderConfigEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
	ProviderType string `gorm:"uniqueIndex:idx_wallet_provider_config_type_region" json:"providerType" validate:"required" yaml:"providerType"`
	// Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
	Region string `gorm:"uniqueIndex:idx_wallet_provider_config_type_region" json:"region" validate:"required" yaml:"region"`
	// Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
	IsEnabled bool `json:"isEnabled" yaml:"isEnabled"`
	// Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
	Config complexes.JSON `json:"config" yaml:"config"`
}

func (x *WalletProviderConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletProviderConfigEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "provider-type",
			Type:        "string",
			Description: "Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. \"stripe\", \"przelewy24\", \"blik\", \"zarinpal\". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.",
		},
		{
			Name:        prefix + "region",
			Type:        "string",
			Description: "Which region this config applies to - an ISO-3166-ish code (e.g. \"PL\", \"IR\", \"US\") or \"global\", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.",
		},
		{
			Name:        prefix + "is-enabled",
			Type:        "bool",
			Description: "Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.",
		},
		{
			Name:        prefix + "config",
			Type:        "complex",
			Description: "Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.",
		},
	}
}
func CastWalletProviderConfigEntityFromCli(c emigo.CliCastable) WalletProviderConfigEntity {
	data := WalletProviderConfigEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("provider-type") {
		data.ProviderType = c.String("provider-type")
	}
	if c.IsSet("region") {
		data.Region = c.String("region")
	}
	if c.IsSet("is-enabled") {
		data.IsEnabled = bool(c.Bool("is-enabled"))
	}
	if c.IsSet("config") {
		if u, ok := any(&data.Config).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("config")))
		}
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletProviderConfigEntityCreateFn creates a new WalletProviderConfigEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletProviderConfigEntityCreateFn(tx *gorm.DB, dto *WalletProviderConfigEntity) (*WalletProviderConfigEntity, error) {
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

// WalletProviderConfigEntityUpdateFn applies a partial update to the WalletProviderConfigEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletProviderConfigEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletProviderConfigEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletProviderConfigOptionalDto) (*WalletProviderConfigEntity, error) {
	var entity WalletProviderConfigEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.ProviderType.IsSet() {
			changes["ProviderType"] = input.ProviderType
		}
		if input.Region.IsSet() {
			changes["Region"] = input.Region
		}
		if input.IsEnabled.IsSet() {
			changes["IsEnabled"] = input.IsEnabled
		}
		changes["Config"] = input.Config
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
	var updated WalletProviderConfigEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletProviderConfigEntityGetFn looks up a single WalletProviderConfigEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletProviderConfigEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletProviderConfigEntity, error) {
	var entity WalletProviderConfigEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletProviderConfigEntityBrowseFn returns WalletProviderConfigEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletProviderConfigEntityBrowseFn(tx *gorm.DB, qs WalletProviderConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletProviderConfigEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletProviderConfigEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletProviderConfigEntity
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

// WalletProviderConfigEntityAwareDeleteAffected reports one relation of WalletProviderConfigEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletProviderConfigEntity itself, so deleting WalletProviderConfigEntity doesn't cascade into them.
type WalletProviderConfigEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletProviderConfigEntityAwareDeletePreview is the result of WalletProviderConfigEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletProviderConfigEntityAwareDeleteFn would delete/clear
// alongside the WalletProviderConfigEntity row(s) themselves.
type WalletProviderConfigEntityAwareDeletePreview struct {
	Message  string                                          `json:"message"`
	Affected []WalletProviderConfigEntityAwareDeleteAffected `json:"affected"`
}

// WalletProviderConfigEntityAwareDeletePreviewFn looks up the WalletProviderConfigEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletProviderConfigEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletProviderConfigEntityAwareDeleteFn.
func WalletProviderConfigEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletProviderConfigEntityAwareDeletePreview, error) {
	var rows []*WalletProviderConfigEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletProviderConfigEntityAwareDeletePreview{Message: "No matching WalletProviderConfigEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletProviderConfigEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletProviderConfigEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletProviderConfigEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletProviderConfigEntityAwareDeleteFn deletes the WalletProviderConfigEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletProviderConfigEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletProviderConfigEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletProviderConfigEntity
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
		return tx.Where("id IN ?", ids).Delete(&WalletProviderConfigEntity{}).Error
	})
}

// WalletProviderConfigEntityActionsSig bundles the actions available for WalletProviderConfigEntity. Extend this (and
// WalletProviderConfigEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletProviderConfigEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletProviderConfigEntity) (*WalletProviderConfigEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletProviderConfigOptionalDto) (*WalletProviderConfigEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletProviderConfigEntity, error)
	Browse             func(tx *gorm.DB, qs WalletProviderConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletProviderConfigEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletProviderConfigEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletProviderConfigEntityActions WalletProviderConfigEntityActionsSig = WalletProviderConfigEntityActionsSig{
	Create:             WalletProviderConfigEntityCreateFn,
	Update:             WalletProviderConfigEntityUpdateFn,
	Get:                WalletProviderConfigEntityGetFn,
	Browse:             WalletProviderConfigEntityBrowseFn,
	AwareDeletePreview: WalletProviderConfigEntityAwareDeletePreviewFn,
	AwareDelete:        WalletProviderConfigEntityAwareDeleteFn,
}
