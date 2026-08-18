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

// The base class definition for emailConfirmationEntity
type EmailConfirmationEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this confirmation belongs to.
	UserId    emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Status    string                 `json:"status" yaml:"status"`
	Email     string                 `json:"email" yaml:"email"`
	Key       string                 `json:"key" yaml:"key"`
	ExpiresAt string                 `json:"expiresAt" yaml:"expiresAt"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailConfirmationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetEmailConfirmationEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Description: "The unique-id of the user this confirmation belongs to.",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "key",
			Type: "string",
		},
		{
			Name: prefix + "expires-at",
			Type: "string",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
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
func CastEmailConfirmationEntityFromCli(c emigo.CliCastable) EmailConfirmationEntity {
	data := EmailConfirmationEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("expires-at") {
		data.ExpiresAt = c.String("expires-at")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
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

// EmailConfirmationEntityCreateFn creates a new EmailConfirmationEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func EmailConfirmationEntityCreateFn(tx *gorm.DB, dto *EmailConfirmationEntity) (*EmailConfirmationEntity, error) {
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

// EmailConfirmationEntityUpdateFn applies a partial update to the EmailConfirmationEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// EmailConfirmationEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func EmailConfirmationEntityUpdateFn(tx *gorm.DB, uniqueId string, input EmailConfirmationOptionalDto) (*EmailConfirmationEntity, error) {
	var entity EmailConfirmationEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.Status.IsSet() {
			changes["Status"] = input.Status
		}
		if input.Email.IsSet() {
			changes["Email"] = input.Email
		}
		if input.Key.IsSet() {
			changes["Key"] = input.Key
		}
		if input.ExpiresAt.IsSet() {
			changes["ExpiresAt"] = input.ExpiresAt
		}
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
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
	var updated EmailConfirmationEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// EmailConfirmationEntityGetFn looks up a single EmailConfirmationEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func EmailConfirmationEntityGetFn(tx *gorm.DB, uniqueId string) (*EmailConfirmationEntity, error) {
	var entity EmailConfirmationEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// EmailConfirmationEntityBrowseFn returns EmailConfirmationEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func EmailConfirmationEntityBrowseFn(tx *gorm.DB, qs EmailConfirmationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*EmailConfirmationEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&EmailConfirmationEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*EmailConfirmationEntity
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

// EmailConfirmationEntityAwareDeleteAffected reports one relation of EmailConfirmationEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on EmailConfirmationEntity itself, so deleting EmailConfirmationEntity doesn't cascade into them.
type EmailConfirmationEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// EmailConfirmationEntityAwareDeletePreview is the result of EmailConfirmationEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts EmailConfirmationEntityAwareDeleteFn would delete/clear
// alongside the EmailConfirmationEntity row(s) themselves.
type EmailConfirmationEntityAwareDeletePreview struct {
	Message  string                                       `json:"message"`
	Affected []EmailConfirmationEntityAwareDeleteAffected `json:"affected"`
}

// EmailConfirmationEntityAwareDeletePreviewFn looks up the EmailConfirmationEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// EmailConfirmationEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling EmailConfirmationEntityAwareDeleteFn.
func EmailConfirmationEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*EmailConfirmationEntityAwareDeletePreview, error) {
	var rows []*EmailConfirmationEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &EmailConfirmationEntityAwareDeletePreview{Message: "No matching EmailConfirmationEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []EmailConfirmationEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d EmailConfirmationEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &EmailConfirmationEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// EmailConfirmationEntityAwareDeleteFn deletes the EmailConfirmationEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation EmailConfirmationEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func EmailConfirmationEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*EmailConfirmationEntity
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
		return tx.Where("id IN ?", ids).Delete(&EmailConfirmationEntity{}).Error
	})
}

// EmailConfirmationEntityActionsSig bundles the actions available for EmailConfirmationEntity. Extend this (and
// EmailConfirmationEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type EmailConfirmationEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *EmailConfirmationEntity) (*EmailConfirmationEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input EmailConfirmationOptionalDto) (*EmailConfirmationEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*EmailConfirmationEntity, error)
	Browse             func(tx *gorm.DB, qs EmailConfirmationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*EmailConfirmationEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*EmailConfirmationEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var EmailConfirmationEntityActions EmailConfirmationEntityActionsSig = EmailConfirmationEntityActionsSig{
	Create:             EmailConfirmationEntityCreateFn,
	Update:             EmailConfirmationEntityUpdateFn,
	Get:                EmailConfirmationEntityGetFn,
	Browse:             EmailConfirmationEntityBrowseFn,
	AwareDeletePreview: EmailConfirmationEntityAwareDeletePreviewFn,
	AwareDelete:        EmailConfirmationEntityAwareDeleteFn,
}
