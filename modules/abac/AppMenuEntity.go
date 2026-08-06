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

// The base class definition for appMenuEntity
type AppMenuEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Label that will be visible to user
	Label string `json:"label" yaml:"label"`
	// Location that will be navigated in case of click or selection on ui
	Href string `json:"href" yaml:"href"`
	// Icon string address which matches the resources on the front-end apps.
	Icon string `json:"icon" yaml:"icon"`
	// Custom window location url matchers, for inner screens.
	ActiveMatcher string `json:"activeMatcher" yaml:"activeMatcher"`
	// The unique-id of the capability which is required for the menu to be visible.
	CapabilityId emigo.Nullable[string] `json:"capabilityId" yaml:"capabilityId"`
	// The unique-id of the parent menu item, for nested/tree menus.
	ParentId    emigo.Nullable[string]  `json:"parentId" yaml:"parentId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *AppMenuEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetAppMenuEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "label",
			Type:        "string",
			Description: "Label that will be visible to user",
		},
		{
			Name:        prefix + "href",
			Type:        "string",
			Description: "Location that will be navigated in case of click or selection on ui",
		},
		{
			Name:        prefix + "icon",
			Type:        "string",
			Description: "Icon string address which matches the resources on the front-end apps.",
		},
		{
			Name:        prefix + "active-matcher",
			Type:        "string",
			Description: "Custom window location url matchers, for inner screens.",
		},
		{
			Name:        prefix + "capability-id",
			Type:        "string?",
			Description: "The unique-id of the capability which is required for the menu to be visible.",
		},
		{
			Name:        prefix + "parent-id",
			Type:        "string?",
			Description: "The unique-id of the parent menu item, for nested/tree menus.",
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
func CastAppMenuEntityFromCli(c emigo.CliCastable) AppMenuEntity {
	data := AppMenuEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	if c.IsSet("href") {
		data.Href = c.String("href")
	}
	if c.IsSet("icon") {
		data.Icon = c.String("icon")
	}
	if c.IsSet("active-matcher") {
		data.ActiveMatcher = c.String("active-matcher")
	}
	if c.IsSet("capability-id") {
		emigo.ParseNullable(c.String("capability-id"), &data.CapabilityId)
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
// AppMenuEntityCreateFn creates a new AppMenuEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func AppMenuEntityCreateFn(tx *gorm.DB, dto *AppMenuEntity) (*AppMenuEntity, error) {
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

// AppMenuEntityUpdateFn applies a partial update to the AppMenuEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// AppMenuEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func AppMenuEntityUpdateFn(tx *gorm.DB, uniqueId string, input AppMenuOptionalDto) (*AppMenuEntity, error) {
	var entity AppMenuEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Label.IsSet() {
			changes["Label"] = input.Label
		}
		if input.Href.IsSet() {
			changes["Href"] = input.Href
		}
		if input.Icon.IsSet() {
			changes["Icon"] = input.Icon
		}
		if input.ActiveMatcher.IsSet() {
			changes["ActiveMatcher"] = input.ActiveMatcher
		}
		if input.CapabilityId.IsSet() {
			changes["CapabilityId"] = input.CapabilityId
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
	var updated AppMenuEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// AppMenuEntityGetFn looks up a single AppMenuEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func AppMenuEntityGetFn(tx *gorm.DB, uniqueId string) (*AppMenuEntity, error) {
	var entity AppMenuEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// AppMenuEntityBrowseFn returns AppMenuEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func AppMenuEntityBrowseFn(tx *gorm.DB, qs AppMenuBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*AppMenuEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&AppMenuEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*AppMenuEntity
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

// AppMenuEntityAwareDeleteAffected reports one relation of AppMenuEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on AppMenuEntity itself, so deleting AppMenuEntity doesn't cascade into them.
type AppMenuEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// AppMenuEntityAwareDeletePreview is the result of AppMenuEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts AppMenuEntityAwareDeleteFn would delete/clear
// alongside the AppMenuEntity row(s) themselves.
type AppMenuEntityAwareDeletePreview struct {
	Message  string                             `json:"message"`
	Affected []AppMenuEntityAwareDeleteAffected `json:"affected"`
}

// AppMenuEntityAwareDeletePreviewFn looks up the AppMenuEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// AppMenuEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling AppMenuEntityAwareDeleteFn.
func AppMenuEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*AppMenuEntityAwareDeletePreview, error) {
	var rows []*AppMenuEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &AppMenuEntityAwareDeletePreview{Message: "No matching AppMenuEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []AppMenuEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d AppMenuEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &AppMenuEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// AppMenuEntityAwareDeleteFn deletes the AppMenuEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation AppMenuEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func AppMenuEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*AppMenuEntity
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
		return tx.Where("id IN ?", ids).Delete(&AppMenuEntity{}).Error
	})
}

// AppMenuEntityActionsSig bundles the actions available for AppMenuEntity. Extend this (and
// AppMenuEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type AppMenuEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *AppMenuEntity) (*AppMenuEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input AppMenuOptionalDto) (*AppMenuEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*AppMenuEntity, error)
	Browse             func(tx *gorm.DB, qs AppMenuBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*AppMenuEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*AppMenuEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var AppMenuEntityActions AppMenuEntityActionsSig = AppMenuEntityActionsSig{
	Create:             AppMenuEntityCreateFn,
	Update:             AppMenuEntityUpdateFn,
	Get:                AppMenuEntityGetFn,
	Browse:             AppMenuEntityBrowseFn,
	AwareDeletePreview: AppMenuEntityAwareDeletePreviewFn,
	AwareDelete:        AppMenuEntityAwareDeleteFn,
}
