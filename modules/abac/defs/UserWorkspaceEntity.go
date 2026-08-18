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

// The base class definition for userWorkspaceEntity
type UserWorkspaceEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this record belongs to.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId          emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserPermissions      []string                `gorm:"-" json:"userPermissions" sql:"-" yaml:"userPermissions"`
	RolePermission       []interface{}           `gorm:"-" json:"rolePermission" sql:"-" yaml:"rolePermission"`
	WorkspacePermissions []string                `gorm:"-" json:"workspacePermissions" sql:"-" yaml:"workspacePermissions"`
	CreatedAt            abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt            abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *UserWorkspaceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetUserWorkspaceEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user this record belongs to.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name: prefix + "user-permissions",
			Type: "slice",
		},
		{
			Name: prefix + "role-permission",
			Type: "slice",
		},
		{
			Name: prefix + "workspace-permissions",
			Type: "slice",
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
func CastUserWorkspaceEntityFromCli(c emigo.CliCastable) UserWorkspaceEntity {
	data := UserWorkspaceEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-permissions") {
		emigo.InflatePossibleSlice(c.String("user-permissions"), &data.UserPermissions)
	}
	if c.IsSet("role-permission") {
		emigo.InflatePossibleSlice(c.String("role-permission"), &data.RolePermission)
	}
	if c.IsSet("workspace-permissions") {
		emigo.InflatePossibleSlice(c.String("workspace-permissions"), &data.WorkspacePermissions)
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

// UserWorkspaceEntityCreateFn creates a new UserWorkspaceEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func UserWorkspaceEntityCreateFn(tx *gorm.DB, dto *UserWorkspaceEntity) (*UserWorkspaceEntity, error) {
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

// UserWorkspaceEntityUpdateFn applies a partial update to the UserWorkspaceEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// UserWorkspaceEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func UserWorkspaceEntityUpdateFn(tx *gorm.DB, uniqueId string, input UserWorkspaceOptionalDto) (*UserWorkspaceEntity, error) {
	var entity UserWorkspaceEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
		}
		if input.UserPermissions.IsSet() {
			changes["UserPermissions"] = input.UserPermissions
		}
		if input.RolePermission.IsSet() {
			changes["RolePermission"] = input.RolePermission
		}
		if input.WorkspacePermissions.IsSet() {
			changes["WorkspacePermissions"] = input.WorkspacePermissions
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
	var updated UserWorkspaceEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// UserWorkspaceEntityGetFn looks up a single UserWorkspaceEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func UserWorkspaceEntityGetFn(tx *gorm.DB, uniqueId string) (*UserWorkspaceEntity, error) {
	var entity UserWorkspaceEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// UserWorkspaceEntityBrowseFn returns UserWorkspaceEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func UserWorkspaceEntityBrowseFn(tx *gorm.DB, qs UserWorkspaceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*UserWorkspaceEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&UserWorkspaceEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*UserWorkspaceEntity
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

// UserWorkspaceEntityAwareDeleteAffected reports one relation of UserWorkspaceEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on UserWorkspaceEntity itself, so deleting UserWorkspaceEntity doesn't cascade into them.
type UserWorkspaceEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// UserWorkspaceEntityAwareDeletePreview is the result of UserWorkspaceEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts UserWorkspaceEntityAwareDeleteFn would delete/clear
// alongside the UserWorkspaceEntity row(s) themselves.
type UserWorkspaceEntityAwareDeletePreview struct {
	Message  string                                   `json:"message"`
	Affected []UserWorkspaceEntityAwareDeleteAffected `json:"affected"`
}

// UserWorkspaceEntityAwareDeletePreviewFn looks up the UserWorkspaceEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// UserWorkspaceEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling UserWorkspaceEntityAwareDeleteFn.
func UserWorkspaceEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*UserWorkspaceEntityAwareDeletePreview, error) {
	var rows []*UserWorkspaceEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &UserWorkspaceEntityAwareDeletePreview{Message: "No matching UserWorkspaceEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []UserWorkspaceEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d UserWorkspaceEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &UserWorkspaceEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// UserWorkspaceEntityAwareDeleteFn deletes the UserWorkspaceEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation UserWorkspaceEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func UserWorkspaceEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*UserWorkspaceEntity
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
		return tx.Where("id IN ?", ids).Delete(&UserWorkspaceEntity{}).Error
	})
}

// UserWorkspaceEntityActionsSig bundles the actions available for UserWorkspaceEntity. Extend this (and
// UserWorkspaceEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type UserWorkspaceEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *UserWorkspaceEntity) (*UserWorkspaceEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input UserWorkspaceOptionalDto) (*UserWorkspaceEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*UserWorkspaceEntity, error)
	Browse             func(tx *gorm.DB, qs UserWorkspaceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*UserWorkspaceEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*UserWorkspaceEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var UserWorkspaceEntityActions UserWorkspaceEntityActionsSig = UserWorkspaceEntityActionsSig{
	Create:             UserWorkspaceEntityCreateFn,
	Update:             UserWorkspaceEntityUpdateFn,
	Get:                UserWorkspaceEntityGetFn,
	Browse:             UserWorkspaceEntityBrowseFn,
	AwareDeletePreview: UserWorkspaceEntityAwareDeletePreviewFn,
	AwareDelete:        UserWorkspaceEntityAwareDeleteFn,
}
