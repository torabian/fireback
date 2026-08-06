package abac

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"gorm.io/gorm"
)

// The base class definition for workspaceEntity
type WorkspaceEntity struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId    string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Description string `json:"description" yaml:"description"`
	Name        string `json:"name" validate:"required" yaml:"name"`
	// The unique-id of the workspace type which defines this workspace's role.
	TypeId string `json:"typeId" validate:"required" yaml:"typeId"`
	// The unique-id of the parent workspace, for nested/tree workspaces.
	ParentId    emigo.Nullable[string]  `json:"parentId" yaml:"parentId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:        prefix + "type-id",
			Type:        "string",
			Description: "The unique-id of the workspace type which defines this workspace's role.",
		},
		{
			Name:        prefix + "parent-id",
			Type:        "string?",
			Description: "The unique-id of the parent workspace, for nested/tree workspaces.",
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
func CastWorkspaceEntityFromCli(c emigo.CliCastable) WorkspaceEntity {
	data := WorkspaceEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("type-id") {
		data.TypeId = c.String("type-id")
	}
	if c.IsSet("parent-id") {
		emigo.ParseNullable(c.String("parent-id"), &data.ParentId)
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

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WorkspaceEntityCreateFn creates a new WorkspaceEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WorkspaceEntityCreateFn(tx *gorm.DB, dto *WorkspaceEntity) (*WorkspaceEntity, error) {
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

// WorkspaceEntityUpdateFn applies a partial update to the WorkspaceEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WorkspaceEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WorkspaceEntityUpdateFn(tx *gorm.DB, uniqueId string, input WorkspaceOptionalDto) (*WorkspaceEntity, error) {
	var entity WorkspaceEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Description.IsSet() {
			changes["Description"] = input.Description
		}
		if input.Name.IsSet() {
			changes["Name"] = input.Name
		}
		if input.TypeId.IsSet() {
			changes["TypeId"] = input.TypeId
		}
		if input.ParentId.IsSet() {
			changes["ParentId"] = input.ParentId
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
	var updated WorkspaceEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WorkspaceEntityGetFn looks up a single WorkspaceEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WorkspaceEntityGetFn(tx *gorm.DB, uniqueId string) (*WorkspaceEntity, error) {
	var entity WorkspaceEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WorkspaceEntityBrowseFn returns WorkspaceEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WorkspaceEntityBrowseFn(tx *gorm.DB, qs WorkspaceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WorkspaceEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WorkspaceEntity
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

// WorkspaceEntityAwareDeleteAffected reports one relation of WorkspaceEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WorkspaceEntity itself, so deleting WorkspaceEntity doesn't cascade into them.
type WorkspaceEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WorkspaceEntityAwareDeletePreview is the result of WorkspaceEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WorkspaceEntityAwareDeleteFn would delete/clear
// alongside the WorkspaceEntity row(s) themselves.
type WorkspaceEntityAwareDeletePreview struct {
	Message  string                               `json:"message"`
	Affected []WorkspaceEntityAwareDeleteAffected `json:"affected"`
}

// WorkspaceEntityAwareDeletePreviewFn looks up the WorkspaceEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WorkspaceEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WorkspaceEntityAwareDeleteFn.
func WorkspaceEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WorkspaceEntityAwareDeletePreview, error) {
	var rows []*WorkspaceEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WorkspaceEntityAwareDeletePreview{Message: "No matching WorkspaceEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WorkspaceEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WorkspaceEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WorkspaceEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WorkspaceEntityAwareDeleteFn deletes the WorkspaceEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WorkspaceEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WorkspaceEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WorkspaceEntity
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
		return tx.Where("id IN ?", ids).Delete(&WorkspaceEntity{}).Error
	})
}

// WorkspaceEntityActionsSig bundles the actions available for WorkspaceEntity. Extend this (and
// WorkspaceEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WorkspaceEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WorkspaceEntity) (*WorkspaceEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WorkspaceOptionalDto) (*WorkspaceEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WorkspaceEntity, error)
	Browse             func(tx *gorm.DB, qs WorkspaceBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WorkspaceEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WorkspaceEntityActions WorkspaceEntityActionsSig = WorkspaceEntityActionsSig{
	Create:             WorkspaceEntityCreateFn,
	Update:             WorkspaceEntityUpdateFn,
	Get:                WorkspaceEntityGetFn,
	Browse:             WorkspaceEntityBrowseFn,
	AwareDeletePreview: WorkspaceEntityAwareDeletePreviewFn,
	AwareDelete:        WorkspaceEntityAwareDeleteFn,
}
