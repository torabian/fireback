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

// The base class definition for phoneConfirmationEntity
type PhoneConfirmationEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this confirmation belongs to.
	UserId      emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Status      string                 `json:"status" yaml:"status"`
	PhoneNumber string                 `json:"phoneNumber" yaml:"phoneNumber"`
	Key         string                 `json:"key" yaml:"key"`
	ExpiresAt   string                 `json:"expiresAt" yaml:"expiresAt"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PhoneConfirmationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPhoneConfirmationEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "phone-number",
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
func CastPhoneConfirmationEntityFromCli(c emigo.CliCastable) PhoneConfirmationEntity {
	data := PhoneConfirmationEntity{}
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
	if c.IsSet("phone-number") {
		data.PhoneNumber = c.String("phone-number")
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

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// PhoneConfirmationEntityCreateFn creates a new PhoneConfirmationEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func PhoneConfirmationEntityCreateFn(tx *gorm.DB, dto *PhoneConfirmationEntity) (*PhoneConfirmationEntity, error) {
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

// PhoneConfirmationEntityUpdateFn applies a partial update to the PhoneConfirmationEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// PhoneConfirmationEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func PhoneConfirmationEntityUpdateFn(tx *gorm.DB, uniqueId string, input PhoneConfirmationOptionalDto) (*PhoneConfirmationEntity, error) {
	var entity PhoneConfirmationEntity
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
		if input.PhoneNumber.IsSet() {
			changes["PhoneNumber"] = input.PhoneNumber
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
	var updated PhoneConfirmationEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// PhoneConfirmationEntityGetFn looks up a single PhoneConfirmationEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func PhoneConfirmationEntityGetFn(tx *gorm.DB, uniqueId string) (*PhoneConfirmationEntity, error) {
	var entity PhoneConfirmationEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// PhoneConfirmationEntityBrowseFn returns PhoneConfirmationEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func PhoneConfirmationEntityBrowseFn(tx *gorm.DB, qs PhoneConfirmationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PhoneConfirmationEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&PhoneConfirmationEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*PhoneConfirmationEntity
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

// PhoneConfirmationEntityAwareDeleteAffected reports one relation of PhoneConfirmationEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on PhoneConfirmationEntity itself, so deleting PhoneConfirmationEntity doesn't cascade into them.
type PhoneConfirmationEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// PhoneConfirmationEntityAwareDeletePreview is the result of PhoneConfirmationEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts PhoneConfirmationEntityAwareDeleteFn would delete/clear
// alongside the PhoneConfirmationEntity row(s) themselves.
type PhoneConfirmationEntityAwareDeletePreview struct {
	Message  string                                       `json:"message"`
	Affected []PhoneConfirmationEntityAwareDeleteAffected `json:"affected"`
}

// PhoneConfirmationEntityAwareDeletePreviewFn looks up the PhoneConfirmationEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// PhoneConfirmationEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling PhoneConfirmationEntityAwareDeleteFn.
func PhoneConfirmationEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*PhoneConfirmationEntityAwareDeletePreview, error) {
	var rows []*PhoneConfirmationEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &PhoneConfirmationEntityAwareDeletePreview{Message: "No matching PhoneConfirmationEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []PhoneConfirmationEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d PhoneConfirmationEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &PhoneConfirmationEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// PhoneConfirmationEntityAwareDeleteFn deletes the PhoneConfirmationEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation PhoneConfirmationEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func PhoneConfirmationEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*PhoneConfirmationEntity
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
		return tx.Where("id IN ?", ids).Delete(&PhoneConfirmationEntity{}).Error
	})
}

// PhoneConfirmationEntityActionsSig bundles the actions available for PhoneConfirmationEntity. Extend this (and
// PhoneConfirmationEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type PhoneConfirmationEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *PhoneConfirmationEntity) (*PhoneConfirmationEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input PhoneConfirmationOptionalDto) (*PhoneConfirmationEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*PhoneConfirmationEntity, error)
	Browse             func(tx *gorm.DB, qs PhoneConfirmationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PhoneConfirmationEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*PhoneConfirmationEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var PhoneConfirmationEntityActions PhoneConfirmationEntityActionsSig = PhoneConfirmationEntityActionsSig{
	Create:             PhoneConfirmationEntityCreateFn,
	Update:             PhoneConfirmationEntityUpdateFn,
	Get:                PhoneConfirmationEntityGetFn,
	Browse:             PhoneConfirmationEntityBrowseFn,
	AwareDeletePreview: PhoneConfirmationEntityAwareDeletePreviewFn,
	AwareDelete:        PhoneConfirmationEntityAwareDeleteFn,
}
