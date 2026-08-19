package abacdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"gorm.io/gorm"
)

// The base class definition for notificationEntity
type NotificationEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// UniqueId of the user this notification was sent to.
	UserId string `json:"userId" validate:"required" yaml:"userId"`
	// UniqueId of the (root) user who sent this notification.
	SenderId emigo.Nullable[string] `json:"senderId" yaml:"senderId"`
	// Short notification title.
	Title string `json:"title" validate:"required" yaml:"title"`
	// Notification message body.
	Body string `json:"body" validate:"required" yaml:"body"`
	// Whether the recipient has read this notification yet.
	IsRead bool `json:"isRead" yaml:"isRead"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *NotificationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// NotificationEntityCreateFn creates a new NotificationEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func NotificationEntityCreateFn(tx *gorm.DB, dto *NotificationEntity) (*NotificationEntity, error) {
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

// NotificationEntityUpdateFn applies a partial update to the NotificationEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// NotificationEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func NotificationEntityUpdateFn(tx *gorm.DB, uniqueId string, input NotificationOptionalDto) (*NotificationEntity, error) {
	var entity NotificationEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.SenderId.IsSet() {
			changes["SenderId"] = input.SenderId
		}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		if input.Body.IsSet() {
			changes["Body"] = input.Body
		}
		if input.IsRead.IsSet() {
			changes["IsRead"] = input.IsRead
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
	var updated NotificationEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// NotificationEntityGetFn looks up a single NotificationEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func NotificationEntityGetFn(tx *gorm.DB, uniqueId string) (*NotificationEntity, error) {
	var entity NotificationEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// NotificationEntityBrowseFn returns NotificationEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func NotificationEntityBrowseFn(tx *gorm.DB, qs NotificationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*NotificationEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&NotificationEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*NotificationEntity
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

// NotificationEntityAwareDeleteAffected reports one relation of NotificationEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on NotificationEntity itself, so deleting NotificationEntity doesn't cascade into them.
type NotificationEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// NotificationEntityAwareDeletePreview is the result of NotificationEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts NotificationEntityAwareDeleteFn would delete/clear
// alongside the NotificationEntity row(s) themselves.
type NotificationEntityAwareDeletePreview struct {
	Message  string                                  `json:"message"`
	Affected []NotificationEntityAwareDeleteAffected `json:"affected"`
}

// NotificationEntityAwareDeletePreviewFn looks up the NotificationEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// NotificationEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling NotificationEntityAwareDeleteFn.
func NotificationEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*NotificationEntityAwareDeletePreview, error) {
	var rows []*NotificationEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &NotificationEntityAwareDeletePreview{Message: "No matching NotificationEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []NotificationEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d NotificationEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &NotificationEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// NotificationEntityAwareDeleteFn deletes the NotificationEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation NotificationEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func NotificationEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*NotificationEntity
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
		return tx.Where("id IN ?", ids).Delete(&NotificationEntity{}).Error
	})
}

// NotificationEntityActionsSig bundles the actions available for NotificationEntity. Extend this (and
// NotificationEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type NotificationEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *NotificationEntity) (*NotificationEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input NotificationOptionalDto) (*NotificationEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*NotificationEntity, error)
	Browse             func(tx *gorm.DB, qs NotificationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*NotificationEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*NotificationEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var NotificationEntityActions NotificationEntityActionsSig = NotificationEntityActionsSig{
	Create:             NotificationEntityCreateFn,
	Update:             NotificationEntityUpdateFn,
	Get:                NotificationEntityGetFn,
	Browse:             NotificationEntityBrowseFn,
	AwareDeletePreview: NotificationEntityAwareDeletePreviewFn,
	AwareDelete:        NotificationEntityAwareDeleteFn,
}
