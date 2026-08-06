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

// The base class definition for gsmProviderEntity
type GsmProviderEntity struct {
	Id               int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId         string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	ApiKey           string `json:"apiKey" yaml:"apiKey"`
	MainSenderNumber string `json:"mainSenderNumber" validate:"required" yaml:"mainSenderNumber"`
	Type             string `json:"type" validate:"required" yaml:"type"`
	InvokeUrl        string `json:"invokeUrl" yaml:"invokeUrl"`
	InvokeBody       string `json:"invokeBody" yaml:"invokeBody"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *GsmProviderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmProviderEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "api-key",
			Type: "string",
		},
		{
			Name: prefix + "main-sender-number",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "enum",
		},
		{
			Name: prefix + "invoke-url",
			Type: "string",
		},
		{
			Name: prefix + "invoke-body",
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
func CastGsmProviderEntityFromCli(c emigo.CliCastable) GsmProviderEntity {
	data := GsmProviderEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("api-key") {
		data.ApiKey = c.String("api-key")
	}
	if c.IsSet("main-sender-number") {
		data.MainSenderNumber = c.String("main-sender-number")
	}
	if c.IsSet("invoke-url") {
		data.InvokeUrl = c.String("invoke-url")
	}
	if c.IsSet("invoke-body") {
		data.InvokeBody = c.String("invoke-body")
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
// GsmProviderEntityCreateFn creates a new GsmProviderEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func GsmProviderEntityCreateFn(tx *gorm.DB, dto *GsmProviderEntity) (*GsmProviderEntity, error) {
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

// GsmProviderEntityUpdateFn applies a partial update to the GsmProviderEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// GsmProviderEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func GsmProviderEntityUpdateFn(tx *gorm.DB, uniqueId string, input GsmProviderOptionalDto) (*GsmProviderEntity, error) {
	var entity GsmProviderEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.ApiKey.IsSet() {
			changes["ApiKey"] = input.ApiKey
		}
		if input.MainSenderNumber.IsSet() {
			changes["MainSenderNumber"] = input.MainSenderNumber
		}
		if input.Type.IsSet() {
			changes["Type"] = input.Type
		}
		if input.InvokeUrl.IsSet() {
			changes["InvokeUrl"] = input.InvokeUrl
		}
		if input.InvokeBody.IsSet() {
			changes["InvokeBody"] = input.InvokeBody
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
	var updated GsmProviderEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// GsmProviderEntityGetFn looks up a single GsmProviderEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func GsmProviderEntityGetFn(tx *gorm.DB, uniqueId string) (*GsmProviderEntity, error) {
	var entity GsmProviderEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GsmProviderEntityBrowseFn returns GsmProviderEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func GsmProviderEntityBrowseFn(tx *gorm.DB, qs GsmProviderBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*GsmProviderEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&GsmProviderEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*GsmProviderEntity
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

// GsmProviderEntityAwareDeleteAffected reports one relation of GsmProviderEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on GsmProviderEntity itself, so deleting GsmProviderEntity doesn't cascade into them.
type GsmProviderEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// GsmProviderEntityAwareDeletePreview is the result of GsmProviderEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts GsmProviderEntityAwareDeleteFn would delete/clear
// alongside the GsmProviderEntity row(s) themselves.
type GsmProviderEntityAwareDeletePreview struct {
	Message  string                                 `json:"message"`
	Affected []GsmProviderEntityAwareDeleteAffected `json:"affected"`
}

// GsmProviderEntityAwareDeletePreviewFn looks up the GsmProviderEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// GsmProviderEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling GsmProviderEntityAwareDeleteFn.
func GsmProviderEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*GsmProviderEntityAwareDeletePreview, error) {
	var rows []*GsmProviderEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &GsmProviderEntityAwareDeletePreview{Message: "No matching GsmProviderEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []GsmProviderEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d GsmProviderEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &GsmProviderEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// GsmProviderEntityAwareDeleteFn deletes the GsmProviderEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation GsmProviderEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func GsmProviderEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*GsmProviderEntity
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
		return tx.Where("id IN ?", ids).Delete(&GsmProviderEntity{}).Error
	})
}

// GsmProviderEntityActionsSig bundles the actions available for GsmProviderEntity. Extend this (and
// GsmProviderEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type GsmProviderEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *GsmProviderEntity) (*GsmProviderEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input GsmProviderOptionalDto) (*GsmProviderEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*GsmProviderEntity, error)
	Browse             func(tx *gorm.DB, qs GsmProviderBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*GsmProviderEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*GsmProviderEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var GsmProviderEntityActions GsmProviderEntityActionsSig = GsmProviderEntityActionsSig{
	Create:             GsmProviderEntityCreateFn,
	Update:             GsmProviderEntityUpdateFn,
	Get:                GsmProviderEntityGetFn,
	Browse:             GsmProviderEntityBrowseFn,
	AwareDeletePreview: GsmProviderEntityAwareDeletePreviewFn,
	AwareDelete:        GsmProviderEntityAwareDeleteFn,
}
