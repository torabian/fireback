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

// The base class definition for tableViewSizingEntity
type TableViewSizingEntity struct {
	Id        int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId  string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	TableName string `json:"tableName" validate:"required" yaml:"tableName"`
	Sizes     string `json:"sizes" yaml:"sizes"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The time the record was created. Populated automatically by gorm.
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	// The time the record was last updated. Populated automatically by gorm.
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *TableViewSizingEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTableViewSizingEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "table-name",
			Type: "string",
		},
		{
			Name: prefix + "sizes",
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
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "The time the record was created. Populated automatically by gorm.",
		},
		{
			Name:        prefix + "updated-at",
			Type:        "complex",
			Description: "The time the record was last updated. Populated automatically by gorm.",
		},
	}
}
func CastTableViewSizingEntityFromCli(c emigo.CliCastable) TableViewSizingEntity {
	data := TableViewSizingEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("table-name") {
		data.TableName = c.String("table-name")
	}
	if c.IsSet("sizes") {
		data.Sizes = c.String("sizes")
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
// TableViewSizingEntityCreateFn creates a new TableViewSizingEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func TableViewSizingEntityCreateFn(tx *gorm.DB, dto *TableViewSizingEntity) (*TableViewSizingEntity, error) {
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

// TableViewSizingEntityUpdateFn applies a partial update to the TableViewSizingEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// TableViewSizingEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func TableViewSizingEntityUpdateFn(tx *gorm.DB, uniqueId string, input TableViewSizingOptionalDto) (*TableViewSizingEntity, error) {
	var entity TableViewSizingEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.TableName.IsSet() {
			changes["TableName"] = input.TableName
		}
		if input.Sizes.IsSet() {
			changes["Sizes"] = input.Sizes
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
	var updated TableViewSizingEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// TableViewSizingEntityGetFn looks up a single TableViewSizingEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func TableViewSizingEntityGetFn(tx *gorm.DB, uniqueId string) (*TableViewSizingEntity, error) {
	var entity TableViewSizingEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// TableViewSizingEntityBrowseFn returns TableViewSizingEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func TableViewSizingEntityBrowseFn(tx *gorm.DB, qs TableViewSizingBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TableViewSizingEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&TableViewSizingEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*TableViewSizingEntity
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

// TableViewSizingEntityAwareDeleteAffected reports one relation of TableViewSizingEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on TableViewSizingEntity itself, so deleting TableViewSizingEntity doesn't cascade into them.
type TableViewSizingEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// TableViewSizingEntityAwareDeletePreview is the result of TableViewSizingEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts TableViewSizingEntityAwareDeleteFn would delete/clear
// alongside the TableViewSizingEntity row(s) themselves.
type TableViewSizingEntityAwareDeletePreview struct {
	Message  string                                     `json:"message"`
	Affected []TableViewSizingEntityAwareDeleteAffected `json:"affected"`
}

// TableViewSizingEntityAwareDeletePreviewFn looks up the TableViewSizingEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// TableViewSizingEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling TableViewSizingEntityAwareDeleteFn.
func TableViewSizingEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*TableViewSizingEntityAwareDeletePreview, error) {
	var rows []*TableViewSizingEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &TableViewSizingEntityAwareDeletePreview{Message: "No matching TableViewSizingEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []TableViewSizingEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d TableViewSizingEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &TableViewSizingEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// TableViewSizingEntityAwareDeleteFn deletes the TableViewSizingEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation TableViewSizingEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func TableViewSizingEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*TableViewSizingEntity
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
		return tx.Where("id IN ?", ids).Delete(&TableViewSizingEntity{}).Error
	})
}

// TableViewSizingEntityActionsSig bundles the actions available for TableViewSizingEntity. Extend this (and
// TableViewSizingEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type TableViewSizingEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *TableViewSizingEntity) (*TableViewSizingEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input TableViewSizingOptionalDto) (*TableViewSizingEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*TableViewSizingEntity, error)
	Browse             func(tx *gorm.DB, qs TableViewSizingBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TableViewSizingEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*TableViewSizingEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var TableViewSizingEntityActions TableViewSizingEntityActionsSig = TableViewSizingEntityActionsSig{
	Create:             TableViewSizingEntityCreateFn,
	Update:             TableViewSizingEntityUpdateFn,
	Get:                TableViewSizingEntityGetFn,
	Browse:             TableViewSizingEntityBrowseFn,
	AwareDeletePreview: TableViewSizingEntityAwareDeletePreviewFn,
	AwareDelete:        TableViewSizingEntityAwareDeleteFn,
}
