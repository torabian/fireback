package abacdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"gorm.io/gorm"
)

// The base class definition for notificationConfigEntity
type NotificationConfigEntity struct {
	Id                                     int64                   `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId                               string                  `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	CascadeToSubWorkspaces                 bool                    `json:"cascadeToSubWorkspaces" yaml:"cascadeToSubWorkspaces"`
	ForcedCascadeEmailProvider             bool                    `json:"forcedCascadeEmailProvider" yaml:"forcedCascadeEmailProvider"`
	GeneralEmailProviderId                 emigo.Nullable[string]  `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
	GeneralGsmProviderId                   emigo.Nullable[string]  `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
	InviteToWorkspaceContent               string                  `gorm:"text" json:"inviteToWorkspaceContent" yaml:"inviteToWorkspaceContent"`
	InviteToWorkspaceContentExcerpt        string                  `gorm:"text" json:"inviteToWorkspaceContentExcerpt" yaml:"inviteToWorkspaceContentExcerpt"`
	InviteToWorkspaceContentDefault        string                  `gorm:"text" json:"inviteToWorkspaceContentDefault" sql:"-" yaml:"inviteToWorkspaceContentDefault"`
	InviteToWorkspaceContentDefaultExcerpt string                  `gorm:"text" json:"inviteToWorkspaceContentDefaultExcerpt" sql:"-" yaml:"inviteToWorkspaceContentDefaultExcerpt"`
	InviteToWorkspaceTitle                 string                  `json:"inviteToWorkspaceTitle" yaml:"inviteToWorkspaceTitle"`
	InviteToWorkspaceTitleDefault          string                  `json:"inviteToWorkspaceTitleDefault" sql:"-" yaml:"inviteToWorkspaceTitleDefault"`
	InviteToWorkspaceSenderId              emigo.Nullable[string]  `json:"inviteToWorkspaceSenderId" yaml:"inviteToWorkspaceSenderId"`
	AccountCenterEmailSenderId             emigo.Nullable[string]  `json:"accountCenterEmailSenderId" yaml:"accountCenterEmailSenderId"`
	ForgetPasswordContent                  string                  `gorm:"text" json:"forgetPasswordContent" yaml:"forgetPasswordContent"`
	ForgetPasswordContentExcerpt           string                  `gorm:"text" json:"forgetPasswordContentExcerpt" yaml:"forgetPasswordContentExcerpt"`
	ForgetPasswordContentDefault           string                  `gorm:"text" json:"forgetPasswordContentDefault" sql:"-" yaml:"forgetPasswordContentDefault"`
	ForgetPasswordContentDefaultExcerpt    string                  `gorm:"text" json:"forgetPasswordContentDefaultExcerpt" sql:"-" yaml:"forgetPasswordContentDefaultExcerpt"`
	ForgetPasswordTitle                    string                  `gorm:"text" json:"forgetPasswordTitle" yaml:"forgetPasswordTitle"`
	ForgetPasswordTitleDefault             string                  `gorm:"text" json:"forgetPasswordTitleDefault" sql:"-" yaml:"forgetPasswordTitleDefault"`
	ForgetPasswordSenderId                 emigo.Nullable[string]  `json:"forgetPasswordSenderId" yaml:"forgetPasswordSenderId"`
	AcceptLanguage                         string                  `json:"acceptLanguage" yaml:"acceptLanguage"`
	ConfirmEmailSenderId                   emigo.Nullable[string]  `json:"confirmEmailSenderId" yaml:"confirmEmailSenderId"`
	ConfirmEmailContent                    string                  `gorm:"text" json:"confirmEmailContent" yaml:"confirmEmailContent"`
	ConfirmEmailContentExcerpt             string                  `gorm:"text" json:"confirmEmailContentExcerpt" yaml:"confirmEmailContentExcerpt"`
	ConfirmEmailContentDefault             string                  `gorm:"text" json:"confirmEmailContentDefault" sql:"-" yaml:"confirmEmailContentDefault"`
	ConfirmEmailContentDefaultExcerpt      string                  `gorm:"text" json:"confirmEmailContentDefaultExcerpt" sql:"-" yaml:"confirmEmailContentDefaultExcerpt"`
	ConfirmEmailTitle                      string                  `json:"confirmEmailTitle" yaml:"confirmEmailTitle"`
	ConfirmEmailTitleDefault               string                  `json:"confirmEmailTitleDefault" sql:"-" yaml:"confirmEmailTitleDefault"`
	WorkspaceId                            emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId                                 emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt                              abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt                              abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *NotificationConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// NotificationConfigEntityCreateFn creates a new NotificationConfigEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func NotificationConfigEntityCreateFn(tx *gorm.DB, dto *NotificationConfigEntity) (*NotificationConfigEntity, error) {
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

// NotificationConfigEntityUpdateFn applies a partial update to the NotificationConfigEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// NotificationConfigEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func NotificationConfigEntityUpdateFn(tx *gorm.DB, uniqueId string, input NotificationConfigOptionalDto) (*NotificationConfigEntity, error) {
	var entity NotificationConfigEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.CascadeToSubWorkspaces.IsSet() {
			changes["CascadeToSubWorkspaces"] = input.CascadeToSubWorkspaces
		}
		if input.ForcedCascadeEmailProvider.IsSet() {
			changes["ForcedCascadeEmailProvider"] = input.ForcedCascadeEmailProvider
		}
		if input.GeneralEmailProviderId.IsSet() {
			changes["GeneralEmailProviderId"] = input.GeneralEmailProviderId
		}
		if input.GeneralGsmProviderId.IsSet() {
			changes["GeneralGsmProviderId"] = input.GeneralGsmProviderId
		}
		if input.InviteToWorkspaceContent.IsSet() {
			changes["InviteToWorkspaceContent"] = input.InviteToWorkspaceContent
		}
		if input.InviteToWorkspaceContentExcerpt.IsSet() {
			changes["InviteToWorkspaceContentExcerpt"] = input.InviteToWorkspaceContentExcerpt
		}
		if input.InviteToWorkspaceContentDefault.IsSet() {
			changes["InviteToWorkspaceContentDefault"] = input.InviteToWorkspaceContentDefault
		}
		if input.InviteToWorkspaceContentDefaultExcerpt.IsSet() {
			changes["InviteToWorkspaceContentDefaultExcerpt"] = input.InviteToWorkspaceContentDefaultExcerpt
		}
		if input.InviteToWorkspaceTitle.IsSet() {
			changes["InviteToWorkspaceTitle"] = input.InviteToWorkspaceTitle
		}
		if input.InviteToWorkspaceTitleDefault.IsSet() {
			changes["InviteToWorkspaceTitleDefault"] = input.InviteToWorkspaceTitleDefault
		}
		if input.InviteToWorkspaceSenderId.IsSet() {
			changes["InviteToWorkspaceSenderId"] = input.InviteToWorkspaceSenderId
		}
		if input.AccountCenterEmailSenderId.IsSet() {
			changes["AccountCenterEmailSenderId"] = input.AccountCenterEmailSenderId
		}
		if input.ForgetPasswordContent.IsSet() {
			changes["ForgetPasswordContent"] = input.ForgetPasswordContent
		}
		if input.ForgetPasswordContentExcerpt.IsSet() {
			changes["ForgetPasswordContentExcerpt"] = input.ForgetPasswordContentExcerpt
		}
		if input.ForgetPasswordContentDefault.IsSet() {
			changes["ForgetPasswordContentDefault"] = input.ForgetPasswordContentDefault
		}
		if input.ForgetPasswordContentDefaultExcerpt.IsSet() {
			changes["ForgetPasswordContentDefaultExcerpt"] = input.ForgetPasswordContentDefaultExcerpt
		}
		if input.ForgetPasswordTitle.IsSet() {
			changes["ForgetPasswordTitle"] = input.ForgetPasswordTitle
		}
		if input.ForgetPasswordTitleDefault.IsSet() {
			changes["ForgetPasswordTitleDefault"] = input.ForgetPasswordTitleDefault
		}
		if input.ForgetPasswordSenderId.IsSet() {
			changes["ForgetPasswordSenderId"] = input.ForgetPasswordSenderId
		}
		if input.AcceptLanguage.IsSet() {
			changes["AcceptLanguage"] = input.AcceptLanguage
		}
		if input.ConfirmEmailSenderId.IsSet() {
			changes["ConfirmEmailSenderId"] = input.ConfirmEmailSenderId
		}
		if input.ConfirmEmailContent.IsSet() {
			changes["ConfirmEmailContent"] = input.ConfirmEmailContent
		}
		if input.ConfirmEmailContentExcerpt.IsSet() {
			changes["ConfirmEmailContentExcerpt"] = input.ConfirmEmailContentExcerpt
		}
		if input.ConfirmEmailContentDefault.IsSet() {
			changes["ConfirmEmailContentDefault"] = input.ConfirmEmailContentDefault
		}
		if input.ConfirmEmailContentDefaultExcerpt.IsSet() {
			changes["ConfirmEmailContentDefaultExcerpt"] = input.ConfirmEmailContentDefaultExcerpt
		}
		if input.ConfirmEmailTitle.IsSet() {
			changes["ConfirmEmailTitle"] = input.ConfirmEmailTitle
		}
		if input.ConfirmEmailTitleDefault.IsSet() {
			changes["ConfirmEmailTitleDefault"] = input.ConfirmEmailTitleDefault
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
	var updated NotificationConfigEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// NotificationConfigEntityGetFn looks up a single NotificationConfigEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func NotificationConfigEntityGetFn(tx *gorm.DB, uniqueId string) (*NotificationConfigEntity, error) {
	var entity NotificationConfigEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// NotificationConfigEntityBrowseFn returns NotificationConfigEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func NotificationConfigEntityBrowseFn(tx *gorm.DB, qs NotificationConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*NotificationConfigEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&NotificationConfigEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*NotificationConfigEntity
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

// NotificationConfigEntityAwareDeleteAffected reports one relation of NotificationConfigEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on NotificationConfigEntity itself, so deleting NotificationConfigEntity doesn't cascade into them.
type NotificationConfigEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// NotificationConfigEntityAwareDeletePreview is the result of NotificationConfigEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts NotificationConfigEntityAwareDeleteFn would delete/clear
// alongside the NotificationConfigEntity row(s) themselves.
type NotificationConfigEntityAwareDeletePreview struct {
	Message  string                                        `json:"message"`
	Affected []NotificationConfigEntityAwareDeleteAffected `json:"affected"`
}

// NotificationConfigEntityAwareDeletePreviewFn looks up the NotificationConfigEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// NotificationConfigEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling NotificationConfigEntityAwareDeleteFn.
func NotificationConfigEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*NotificationConfigEntityAwareDeletePreview, error) {
	var rows []*NotificationConfigEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &NotificationConfigEntityAwareDeletePreview{Message: "No matching NotificationConfigEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []NotificationConfigEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d NotificationConfigEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &NotificationConfigEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// NotificationConfigEntityAwareDeleteFn deletes the NotificationConfigEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation NotificationConfigEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func NotificationConfigEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*NotificationConfigEntity
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
		return tx.Where("id IN ?", ids).Delete(&NotificationConfigEntity{}).Error
	})
}

// NotificationConfigEntityActionsSig bundles the actions available for NotificationConfigEntity. Extend this (and
// NotificationConfigEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type NotificationConfigEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *NotificationConfigEntity) (*NotificationConfigEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input NotificationConfigOptionalDto) (*NotificationConfigEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*NotificationConfigEntity, error)
	Browse             func(tx *gorm.DB, qs NotificationConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*NotificationConfigEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*NotificationConfigEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var NotificationConfigEntityActions NotificationConfigEntityActionsSig = NotificationConfigEntityActionsSig{
	Create:             NotificationConfigEntityCreateFn,
	Update:             NotificationConfigEntityUpdateFn,
	Get:                NotificationConfigEntityGetFn,
	Browse:             NotificationConfigEntityBrowseFn,
	AwareDeletePreview: NotificationConfigEntityAwareDeletePreviewFn,
	AwareDelete:        NotificationConfigEntityAwareDeleteFn,
}
