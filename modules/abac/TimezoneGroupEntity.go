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

// The base class definition for timezoneGroupEntity
type TimezoneGroupEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Title which is shown to the user and allows them to select.
	Title string `json:"title" yaml:"title"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *TimezoneGroupEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTimezoneGroupEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "title",
			Type:        "string",
			Description: "Title which is shown to the user and allows them to select.",
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
func CastTimezoneGroupEntityFromCli(c emigo.CliCastable) TimezoneGroupEntity {
	data := TimezoneGroupEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
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
// TimezoneGroupEntityCreateFn creates a new TimezoneGroupEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func TimezoneGroupEntityCreateFn(tx *gorm.DB, dto *TimezoneGroupEntity) (*TimezoneGroupEntity, error) {
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

// TimezoneGroupEntityUpdateFn applies a partial update to the TimezoneGroupEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// TimezoneGroupEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func TimezoneGroupEntityUpdateFn(tx *gorm.DB, uniqueId string, input TimezoneGroupOptionalDto) (*TimezoneGroupEntity, error) {
	var entity TimezoneGroupEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
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
	var updated TimezoneGroupEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// TimezoneGroupEntityGetFn looks up a single TimezoneGroupEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func TimezoneGroupEntityGetFn(tx *gorm.DB, uniqueId string) (*TimezoneGroupEntity, error) {
	var entity TimezoneGroupEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// TimezoneGroupEntityBrowseFn returns TimezoneGroupEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func TimezoneGroupEntityBrowseFn(tx *gorm.DB, qs TimezoneGroupBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TimezoneGroupEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&TimezoneGroupEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*TimezoneGroupEntity
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

// TimezoneGroupEntityAwareDeleteAffected reports one relation of TimezoneGroupEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on TimezoneGroupEntity itself, so deleting TimezoneGroupEntity doesn't cascade into them.
type TimezoneGroupEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// TimezoneGroupEntityAwareDeletePreview is the result of TimezoneGroupEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts TimezoneGroupEntityAwareDeleteFn would delete/clear
// alongside the TimezoneGroupEntity row(s) themselves.
type TimezoneGroupEntityAwareDeletePreview struct {
	Message  string                                   `json:"message"`
	Affected []TimezoneGroupEntityAwareDeleteAffected `json:"affected"`
}

// TimezoneGroupEntityAwareDeletePreviewFn looks up the TimezoneGroupEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// TimezoneGroupEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling TimezoneGroupEntityAwareDeleteFn.
func TimezoneGroupEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*TimezoneGroupEntityAwareDeletePreview, error) {
	var rows []*TimezoneGroupEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &TimezoneGroupEntityAwareDeletePreview{Message: "No matching TimezoneGroupEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []TimezoneGroupEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d TimezoneGroupEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &TimezoneGroupEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// TimezoneGroupEntityAwareDeleteFn deletes the TimezoneGroupEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation TimezoneGroupEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func TimezoneGroupEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*TimezoneGroupEntity
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
		return tx.Where("id IN ?", ids).Delete(&TimezoneGroupEntity{}).Error
	})
}

// TimezoneGroupEntityActionsSig bundles the actions available for TimezoneGroupEntity. Extend this (and
// TimezoneGroupEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type TimezoneGroupEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *TimezoneGroupEntity) (*TimezoneGroupEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input TimezoneGroupOptionalDto) (*TimezoneGroupEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*TimezoneGroupEntity, error)
	Browse             func(tx *gorm.DB, qs TimezoneGroupBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TimezoneGroupEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*TimezoneGroupEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var TimezoneGroupEntityActions TimezoneGroupEntityActionsSig = TimezoneGroupEntityActionsSig{
	Create:             TimezoneGroupEntityCreateFn,
	Update:             TimezoneGroupEntityUpdateFn,
	Get:                TimezoneGroupEntityGetFn,
	Browse:             TimezoneGroupEntityBrowseFn,
	AwareDeletePreview: TimezoneGroupEntityAwareDeletePreviewFn,
	AwareDelete:        TimezoneGroupEntityAwareDeleteFn,
}
