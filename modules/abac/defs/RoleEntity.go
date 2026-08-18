//go:build !wasm

package abacdefs

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
	"gorm.io/gorm"
)

// The base class definition for roleEntity
type RoleEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Name     string `json:"name" validate:"required,omitempty,min=1,max=200" yaml:"name"`
	// The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
	CapabilitiesListId complexes.JSON          `json:"capabilitiesListId" yaml:"capabilitiesListId"`
	IsDeletable        emigo.Nullable[bool]    `gorm:"default:true" json:"isDeletable" yaml:"isDeletable"`
	IsUpdatable        emigo.Nullable[bool]    `gorm:"default:true" json:"isUpdatable" yaml:"isUpdatable"`
	WorkspaceId        emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId             emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt          abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt          abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RoleEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetRoleEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "capabilities-list-id",
			Type:        "complex",
			Description: "The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).",
		},
		{
			Name: prefix + "is-deletable",
			Type: "bool?",
		},
		{
			Name: prefix + "is-updatable",
			Type: "bool?",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
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
func CastRoleEntityFromCli(c emigo.CliCastable) RoleEntity {
	data := RoleEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("capabilities-list-id") {
		if u, ok := any(&data.CapabilitiesListId).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("capabilities-list-id")))
		}
	}
	if c.IsSet("is-deletable") {
		emigo.ParseNullable(c.String("is-deletable"), &data.IsDeletable)
	}
	if c.IsSet("is-updatable") {
		emigo.ParseNullable(c.String("is-updatable"), &data.IsUpdatable)
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

// RoleEntityCreateFn creates a new RoleEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func RoleEntityCreateFn(tx *gorm.DB, dto *RoleEntity) (*RoleEntity, error) {
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

// RoleEntityUpdateFn applies a partial update to the RoleEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// RoleEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func RoleEntityUpdateFn(tx *gorm.DB, uniqueId string, input RoleOptionalDto) (*RoleEntity, error) {
	var entity RoleEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Name.IsSet() {
			changes["Name"] = input.Name
		}
		changes["CapabilitiesListId"] = input.CapabilitiesListId
		if input.IsDeletable.IsSet() {
			changes["IsDeletable"] = input.IsDeletable
		}
		if input.IsUpdatable.IsSet() {
			changes["IsUpdatable"] = input.IsUpdatable
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
	var updated RoleEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// RoleEntityGetFn looks up a single RoleEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func RoleEntityGetFn(tx *gorm.DB, uniqueId string) (*RoleEntity, error) {
	var entity RoleEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// RoleEntityBrowseFn returns RoleEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func RoleEntityBrowseFn(tx *gorm.DB, qs RoleBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*RoleEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&RoleEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*RoleEntity
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

// RoleEntityAwareDeleteAffected reports one relation of RoleEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on RoleEntity itself, so deleting RoleEntity doesn't cascade into them.
type RoleEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// RoleEntityAwareDeletePreview is the result of RoleEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts RoleEntityAwareDeleteFn would delete/clear
// alongside the RoleEntity row(s) themselves.
type RoleEntityAwareDeletePreview struct {
	Message  string                          `json:"message"`
	Affected []RoleEntityAwareDeleteAffected `json:"affected"`
}

// RoleEntityAwareDeletePreviewFn looks up the RoleEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// RoleEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling RoleEntityAwareDeleteFn.
func RoleEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*RoleEntityAwareDeletePreview, error) {
	var rows []*RoleEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &RoleEntityAwareDeletePreview{Message: "No matching RoleEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []RoleEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d RoleEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &RoleEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// RoleEntityAwareDeleteFn deletes the RoleEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation RoleEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func RoleEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*RoleEntity
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
		return tx.Where("id IN ?", ids).Delete(&RoleEntity{}).Error
	})
}

// RoleEntityActionsSig bundles the actions available for RoleEntity. Extend this (and
// RoleEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type RoleEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *RoleEntity) (*RoleEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input RoleOptionalDto) (*RoleEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*RoleEntity, error)
	Browse             func(tx *gorm.DB, qs RoleBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*RoleEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*RoleEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var RoleEntityActions RoleEntityActionsSig = RoleEntityActionsSig{
	Create:             RoleEntityCreateFn,
	Update:             RoleEntityUpdateFn,
	Get:                RoleEntityGetFn,
	Browse:             RoleEntityBrowseFn,
	AwareDeletePreview: RoleEntityAwareDeletePreviewFn,
	AwareDelete:        RoleEntityAwareDeleteFn,
}
