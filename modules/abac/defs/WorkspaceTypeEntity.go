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

// The base class definition for workspaceTypeEntity
type WorkspaceTypeEntity struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId    string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Title       string `json:"title" validate:"required,omitempty,min=1,max=250" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	// Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
	Slug string `gorm:"unique" json:"slug" validate:"required,omitempty,min=2,max=50" yaml:"slug"`
	// The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
	RoleId      string                  `json:"roleId" validate:"required" yaml:"roleId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceTypeEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceTypeEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name:        prefix + "slug",
			Type:        "string",
			Description: "Unique, URL-safe identifier for this workspace type. Must start with \"/\", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:\"unique\" here is the DB-level backstop for the uniqueness half of that.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string",
			Description: "The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.",
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
func CastWorkspaceTypeEntityFromCli(c emigo.CliCastable) WorkspaceTypeEntity {
	data := WorkspaceTypeEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("slug") {
		data.Slug = c.String("slug")
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
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
// WorkspaceTypeEntityCreateFn creates a new WorkspaceTypeEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WorkspaceTypeEntityCreateFn(tx *gorm.DB, dto *WorkspaceTypeEntity) (*WorkspaceTypeEntity, error) {
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

// WorkspaceTypeEntityUpdateFn applies a partial update to the WorkspaceTypeEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WorkspaceTypeEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WorkspaceTypeEntityUpdateFn(tx *gorm.DB, uniqueId string, input WorkspaceTypeOptionalDto) (*WorkspaceTypeEntity, error) {
	var entity WorkspaceTypeEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		if input.Description.IsSet() {
			changes["Description"] = input.Description
		}
		if input.Slug.IsSet() {
			changes["Slug"] = input.Slug
		}
		if input.RoleId.IsSet() {
			changes["RoleId"] = input.RoleId
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
	var updated WorkspaceTypeEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WorkspaceTypeEntityGetFn looks up a single WorkspaceTypeEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WorkspaceTypeEntityGetFn(tx *gorm.DB, uniqueId string) (*WorkspaceTypeEntity, error) {
	var entity WorkspaceTypeEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WorkspaceTypeEntityBrowseFn returns WorkspaceTypeEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WorkspaceTypeEntityBrowseFn(tx *gorm.DB, qs WorkspaceTypeBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceTypeEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WorkspaceTypeEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WorkspaceTypeEntity
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

// WorkspaceTypeEntityAwareDeleteAffected reports one relation of WorkspaceTypeEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WorkspaceTypeEntity itself, so deleting WorkspaceTypeEntity doesn't cascade into them.
type WorkspaceTypeEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WorkspaceTypeEntityAwareDeletePreview is the result of WorkspaceTypeEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WorkspaceTypeEntityAwareDeleteFn would delete/clear
// alongside the WorkspaceTypeEntity row(s) themselves.
type WorkspaceTypeEntityAwareDeletePreview struct {
	Message  string                                   `json:"message"`
	Affected []WorkspaceTypeEntityAwareDeleteAffected `json:"affected"`
}

// WorkspaceTypeEntityAwareDeletePreviewFn looks up the WorkspaceTypeEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WorkspaceTypeEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WorkspaceTypeEntityAwareDeleteFn.
func WorkspaceTypeEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WorkspaceTypeEntityAwareDeletePreview, error) {
	var rows []*WorkspaceTypeEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WorkspaceTypeEntityAwareDeletePreview{Message: "No matching WorkspaceTypeEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WorkspaceTypeEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WorkspaceTypeEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WorkspaceTypeEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WorkspaceTypeEntityAwareDeleteFn deletes the WorkspaceTypeEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WorkspaceTypeEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WorkspaceTypeEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WorkspaceTypeEntity
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
		return tx.Where("id IN ?", ids).Delete(&WorkspaceTypeEntity{}).Error
	})
}

// WorkspaceTypeEntityActionsSig bundles the actions available for WorkspaceTypeEntity. Extend this (and
// WorkspaceTypeEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WorkspaceTypeEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WorkspaceTypeEntity) (*WorkspaceTypeEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WorkspaceTypeOptionalDto) (*WorkspaceTypeEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WorkspaceTypeEntity, error)
	Browse             func(tx *gorm.DB, qs WorkspaceTypeBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceTypeEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WorkspaceTypeEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WorkspaceTypeEntityActions WorkspaceTypeEntityActionsSig = WorkspaceTypeEntityActionsSig{
	Create:             WorkspaceTypeEntityCreateFn,
	Update:             WorkspaceTypeEntityUpdateFn,
	Get:                WorkspaceTypeEntityGetFn,
	Browse:             WorkspaceTypeEntityBrowseFn,
	AwareDeletePreview: WorkspaceTypeEntityAwareDeletePreviewFn,
	AwareDelete:        WorkspaceTypeEntityAwareDeleteFn,
}
