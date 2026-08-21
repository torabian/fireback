package messagingdefs

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

// The base class definition for emailProviderEntity
type EmailProviderEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.
	Type string `json:"type" validate:"required" yaml:"type"`
	// Give the email provider configuration a name, which makes it easier later to query.
	Title string `json:"title" yaml:"title"`
	// JSON field which contains api keys, or other kind of configuration based on the type of the email provider.
	Config complexes.JSON `json:"config" yaml:"config"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailProviderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEmailProviderEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "type",
			Type:        "enum",
			Description: "Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.",
		},
		{
			Name:        prefix + "title",
			Type:        "string",
			Description: "Give the email provider configuration a name, which makes it easier later to query.",
		},
		{
			Name:        prefix + "config",
			Type:        "complex",
			Description: "JSON field which contains api keys, or other kind of configuration based on the type of the email provider.",
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
func CastEmailProviderEntityFromCli(c emigo.CliCastable) EmailProviderEntity {
	data := EmailProviderEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("config") {
		if u, ok := any(&data.Config).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("config")))
		}
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
// EmailProviderEntityCreateFn creates a new EmailProviderEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func EmailProviderEntityCreateFn(tx *gorm.DB, dto *EmailProviderEntity) (*EmailProviderEntity, error) {
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

// EmailProviderEntityUpdateFn applies a partial update to the EmailProviderEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// EmailProviderEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func EmailProviderEntityUpdateFn(tx *gorm.DB, uniqueId string, input EmailProviderOptionalDto) (*EmailProviderEntity, error) {
	var entity EmailProviderEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Type.IsSet() {
			changes["Type"] = input.Type
		}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		changes["Config"] = input.Config
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
	var updated EmailProviderEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// EmailProviderEntityGetFn looks up a single EmailProviderEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func EmailProviderEntityGetFn(tx *gorm.DB, uniqueId string) (*EmailProviderEntity, error) {
	var entity EmailProviderEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// EmailProviderEntityBrowseFn returns EmailProviderEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func EmailProviderEntityBrowseFn(tx *gorm.DB, qs EmailProviderBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*EmailProviderEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&EmailProviderEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*EmailProviderEntity
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

// EmailProviderEntityAwareDeleteAffected reports one relation of EmailProviderEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on EmailProviderEntity itself, so deleting EmailProviderEntity doesn't cascade into them.
type EmailProviderEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// EmailProviderEntityAwareDeletePreview is the result of EmailProviderEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts EmailProviderEntityAwareDeleteFn would delete/clear
// alongside the EmailProviderEntity row(s) themselves.
type EmailProviderEntityAwareDeletePreview struct {
	Message  string                                   `json:"message"`
	Affected []EmailProviderEntityAwareDeleteAffected `json:"affected"`
}

// EmailProviderEntityAwareDeletePreviewFn looks up the EmailProviderEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// EmailProviderEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling EmailProviderEntityAwareDeleteFn.
func EmailProviderEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*EmailProviderEntityAwareDeletePreview, error) {
	var rows []*EmailProviderEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &EmailProviderEntityAwareDeletePreview{Message: "No matching EmailProviderEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []EmailProviderEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d EmailProviderEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &EmailProviderEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// EmailProviderEntityAwareDeleteFn deletes the EmailProviderEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation EmailProviderEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func EmailProviderEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*EmailProviderEntity
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
		return tx.Where("id IN ?", ids).Delete(&EmailProviderEntity{}).Error
	})
}

// EmailProviderEntityActionsSig bundles the actions available for EmailProviderEntity. Extend this (and
// EmailProviderEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type EmailProviderEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *EmailProviderEntity) (*EmailProviderEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input EmailProviderOptionalDto) (*EmailProviderEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*EmailProviderEntity, error)
	Browse             func(tx *gorm.DB, qs EmailProviderBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*EmailProviderEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*EmailProviderEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var EmailProviderEntityActions EmailProviderEntityActionsSig = EmailProviderEntityActionsSig{
	Create:             EmailProviderEntityCreateFn,
	Update:             EmailProviderEntityUpdateFn,
	Get:                EmailProviderEntityGetFn,
	Browse:             EmailProviderEntityBrowseFn,
	AwareDeletePreview: EmailProviderEntityAwareDeletePreviewFn,
	AwareDelete:        EmailProviderEntityAwareDeleteFn,
}
