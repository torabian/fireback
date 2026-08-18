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

// The base class definition for regionalContentEntity
type RegionalContentEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
	Content string `json:"content" validate:"required" yaml:"content"`
	// Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
	Region string `gorm:"index:regional_content_index,unique" json:"region" validate:"required" yaml:"region"`
	// Optional subject line - only used for email-type content.
	Title string `json:"title" yaml:"title"`
	// Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
	LanguageId string `gorm:"index:regional_content_index,unique" json:"languageId" validate:"required" yaml:"languageId"`
	// Which kind of message this content is used for.
	KeyGroup string `gorm:"index:regional_content_index,unique" json:"keyGroup" validate:"required" yaml:"keyGroup"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RegionalContentEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetRegionalContentEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "content",
			Type:        "string",
			Description: "The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.",
		},
		{
			Name:        prefix + "region",
			Type:        "string",
			Description: "Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.",
		},
		{
			Name:        prefix + "title",
			Type:        "string",
			Description: "Optional subject line - only used for email-type content.",
		},
		{
			Name:        prefix + "language-id",
			Type:        "string",
			Description: "Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.",
		},
		{
			Name:        prefix + "key-group",
			Type:        "enum",
			Description: "Which kind of message this content is used for.",
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
func CastRegionalContentEntityFromCli(c emigo.CliCastable) RegionalContentEntity {
	data := RegionalContentEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("content") {
		data.Content = c.String("content")
	}
	if c.IsSet("region") {
		data.Region = c.String("region")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("language-id") {
		data.LanguageId = c.String("language-id")
	}
	if c.IsSet("key-group") {
		data.KeyGroup = c.String("key-group")
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

// RegionalContentEntityCreateFn creates a new RegionalContentEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func RegionalContentEntityCreateFn(tx *gorm.DB, dto *RegionalContentEntity) (*RegionalContentEntity, error) {
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

// RegionalContentEntityUpdateFn applies a partial update to the RegionalContentEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// RegionalContentEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func RegionalContentEntityUpdateFn(tx *gorm.DB, uniqueId string, input RegionalContentOptionalDto) (*RegionalContentEntity, error) {
	var entity RegionalContentEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Content.IsSet() {
			changes["Content"] = input.Content
		}
		if input.Region.IsSet() {
			changes["Region"] = input.Region
		}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		if input.LanguageId.IsSet() {
			changes["LanguageId"] = input.LanguageId
		}
		if input.KeyGroup.IsSet() {
			changes["KeyGroup"] = input.KeyGroup
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
	var updated RegionalContentEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// RegionalContentEntityGetFn looks up a single RegionalContentEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func RegionalContentEntityGetFn(tx *gorm.DB, uniqueId string) (*RegionalContentEntity, error) {
	var entity RegionalContentEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// RegionalContentEntityBrowseFn returns RegionalContentEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func RegionalContentEntityBrowseFn(tx *gorm.DB, qs RegionalContentBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*RegionalContentEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&RegionalContentEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*RegionalContentEntity
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

// RegionalContentEntityAwareDeleteAffected reports one relation of RegionalContentEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on RegionalContentEntity itself, so deleting RegionalContentEntity doesn't cascade into them.
type RegionalContentEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// RegionalContentEntityAwareDeletePreview is the result of RegionalContentEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts RegionalContentEntityAwareDeleteFn would delete/clear
// alongside the RegionalContentEntity row(s) themselves.
type RegionalContentEntityAwareDeletePreview struct {
	Message  string                                     `json:"message"`
	Affected []RegionalContentEntityAwareDeleteAffected `json:"affected"`
}

// RegionalContentEntityAwareDeletePreviewFn looks up the RegionalContentEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// RegionalContentEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling RegionalContentEntityAwareDeleteFn.
func RegionalContentEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*RegionalContentEntityAwareDeletePreview, error) {
	var rows []*RegionalContentEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &RegionalContentEntityAwareDeletePreview{Message: "No matching RegionalContentEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []RegionalContentEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d RegionalContentEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &RegionalContentEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// RegionalContentEntityAwareDeleteFn deletes the RegionalContentEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation RegionalContentEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func RegionalContentEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*RegionalContentEntity
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
		return tx.Where("id IN ?", ids).Delete(&RegionalContentEntity{}).Error
	})
}

// RegionalContentEntityActionsSig bundles the actions available for RegionalContentEntity. Extend this (and
// RegionalContentEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type RegionalContentEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *RegionalContentEntity) (*RegionalContentEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input RegionalContentOptionalDto) (*RegionalContentEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*RegionalContentEntity, error)
	Browse             func(tx *gorm.DB, qs RegionalContentBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*RegionalContentEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*RegionalContentEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var RegionalContentEntityActions RegionalContentEntityActionsSig = RegionalContentEntityActionsSig{
	Create:             RegionalContentEntityCreateFn,
	Update:             RegionalContentEntityUpdateFn,
	Get:                RegionalContentEntityGetFn,
	Browse:             RegionalContentEntityBrowseFn,
	AwareDeletePreview: RegionalContentEntityAwareDeletePreviewFn,
	AwareDelete:        RegionalContentEntityAwareDeleteFn,
}
