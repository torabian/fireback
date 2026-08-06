package fireback

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for capabilityEntity
type CapabilityEntity struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId    string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

func (x *CapabilityEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCapabilityEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
	}
}
func CastCapabilityEntityFromCli(c emigo.CliCastable) CapabilityEntity {
	data := CapabilityEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// CapabilityEntityCreateFn creates a new CapabilityEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func CapabilityEntityCreateFn(tx *gorm.DB, dto *CapabilityEntity) (*CapabilityEntity, error) {
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

// CapabilityEntityUpdateFn applies a partial update to the CapabilityEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// CapabilityEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func CapabilityEntityUpdateFn(tx *gorm.DB, uniqueId string, input CapabilityOptionalDto) (*CapabilityEntity, error) {
	var entity CapabilityEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Name.IsSet() {
			changes["Name"] = input.Name
		}
		if input.Description.IsSet() {
			changes["Description"] = input.Description
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
	var updated CapabilityEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// CapabilityEntityGetFn looks up a single CapabilityEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func CapabilityEntityGetFn(tx *gorm.DB, uniqueId string) (*CapabilityEntity, error) {
	var entity CapabilityEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// CapabilityEntityBrowseFn returns CapabilityEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func CapabilityEntityBrowseFn(tx *gorm.DB, qs CapabilityBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*CapabilityEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&CapabilityEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*CapabilityEntity
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

// CapabilityEntityAwareDeleteAffected reports one relation of CapabilityEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on CapabilityEntity itself, so deleting CapabilityEntity doesn't cascade into them.
type CapabilityEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// CapabilityEntityAwareDeletePreview is the result of CapabilityEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts CapabilityEntityAwareDeleteFn would delete/clear
// alongside the CapabilityEntity row(s) themselves.
type CapabilityEntityAwareDeletePreview struct {
	Message  string                                `json:"message"`
	Affected []CapabilityEntityAwareDeleteAffected `json:"affected"`
}

// CapabilityEntityAwareDeletePreviewFn looks up the CapabilityEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// CapabilityEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling CapabilityEntityAwareDeleteFn.
func CapabilityEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*CapabilityEntityAwareDeletePreview, error) {
	var rows []*CapabilityEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &CapabilityEntityAwareDeletePreview{Message: "No matching CapabilityEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []CapabilityEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d CapabilityEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &CapabilityEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// CapabilityEntityAwareDeleteFn deletes the CapabilityEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation CapabilityEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func CapabilityEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*CapabilityEntity
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
		return tx.Where("id IN ?", ids).Delete(&CapabilityEntity{}).Error
	})
}

// CapabilityEntityActionsSig bundles the actions available for CapabilityEntity. Extend this (and
// CapabilityEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type CapabilityEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *CapabilityEntity) (*CapabilityEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input CapabilityOptionalDto) (*CapabilityEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*CapabilityEntity, error)
	Browse             func(tx *gorm.DB, qs CapabilityBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*CapabilityEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*CapabilityEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var CapabilityEntityActions CapabilityEntityActionsSig = CapabilityEntityActionsSig{
	Create:             CapabilityEntityCreateFn,
	Update:             CapabilityEntityUpdateFn,
	Get:                CapabilityEntityGetFn,
	Browse:             CapabilityEntityBrowseFn,
	AwareDeletePreview: CapabilityEntityAwareDeletePreviewFn,
	AwareDelete:        CapabilityEntityAwareDeleteFn,
}
