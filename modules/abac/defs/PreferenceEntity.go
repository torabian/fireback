//go:build !wasm

package abacdefs

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"gorm.io/gorm"
)

// The base class definition for preferenceEntity
type PreferenceEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Timezone string `json:"timezone" yaml:"timezone"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PreferenceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetPreferenceEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "timezone",
			Type: "string",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which created/owns the record.",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
	}
}
func CastPreferenceEntityFromCli(c emigo.CliCastable) PreferenceEntity {
	data := PreferenceEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("timezone") {
		data.Timezone = c.String("timezone")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("updated-at") {
		if u, ok := any(&data.UpdatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("updated-at")))
		}
	}
	return data
}

// PreferenceEntityCreateFn creates a new PreferenceEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func PreferenceEntityCreateFn(tx *gorm.DB, dto *PreferenceEntity) (*PreferenceEntity, error) {
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

// PreferenceEntityUpdateFn applies a partial update to the PreferenceEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// PreferenceEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func PreferenceEntityUpdateFn(tx *gorm.DB, uniqueId string, input PreferenceOptionalDto) (*PreferenceEntity, error) {
	var entity PreferenceEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Timezone.IsSet() {
			changes["Timezone"] = input.Timezone
		}
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
		}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		changes["CreatedAt"] = input.CreatedAt
		changes["UpdatedAt"] = input.UpdatedAt
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
	var updated PreferenceEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// PreferenceEntityGetFn looks up a single PreferenceEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func PreferenceEntityGetFn(tx *gorm.DB, uniqueId string) (*PreferenceEntity, error) {
	var entity PreferenceEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// PreferenceEntityBrowseFn returns PreferenceEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func PreferenceEntityBrowseFn(tx *gorm.DB, qs PreferenceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PreferenceEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&PreferenceEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*PreferenceEntity
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

// PreferenceEntityAwareDeleteAffected reports one relation of PreferenceEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on PreferenceEntity itself, so deleting PreferenceEntity doesn't cascade into them.
type PreferenceEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// PreferenceEntityAwareDeletePreview is the result of PreferenceEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts PreferenceEntityAwareDeleteFn would delete/clear
// alongside the PreferenceEntity row(s) themselves.
type PreferenceEntityAwareDeletePreview struct {
	Message  string                                `json:"message"`
	Affected []PreferenceEntityAwareDeleteAffected `json:"affected"`
}

// PreferenceEntityAwareDeletePreviewFn looks up the PreferenceEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// PreferenceEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling PreferenceEntityAwareDeleteFn.
func PreferenceEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*PreferenceEntityAwareDeletePreview, error) {
	var rows []*PreferenceEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &PreferenceEntityAwareDeletePreview{Message: "No matching PreferenceEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []PreferenceEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d PreferenceEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &PreferenceEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// PreferenceEntityAwareDeleteFn deletes the PreferenceEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation PreferenceEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func PreferenceEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*PreferenceEntity
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
		return tx.Where("id IN ?", ids).Delete(&PreferenceEntity{}).Error
	})
}

// PreferenceEntityActionsSig bundles the actions available for PreferenceEntity. Extend this (and
// PreferenceEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type PreferenceEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *PreferenceEntity) (*PreferenceEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input PreferenceOptionalDto) (*PreferenceEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*PreferenceEntity, error)
	Browse             func(tx *gorm.DB, qs PreferenceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PreferenceEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*PreferenceEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var PreferenceEntityActions PreferenceEntityActionsSig = PreferenceEntityActionsSig{
	Create:             PreferenceEntityCreateFn,
	Update:             PreferenceEntityUpdateFn,
	Get:                PreferenceEntityGetFn,
	Browse:             PreferenceEntityBrowseFn,
	AwareDeletePreview: PreferenceEntityAwareDeletePreviewFn,
	AwareDelete:        PreferenceEntityAwareDeleteFn,
}
