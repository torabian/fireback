package messagingdefs

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

// The base class definition for messagingConfigEntity
type MessagingConfigEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the email provider service, which will be used to send the messages using it's service.
	GeneralEmailProviderId emigo.Nullable[string] `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
	// The unique-id of the general service which would be used to send text messages (sms).
	GeneralGsmProviderId emigo.Nullable[string] `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
	// The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
	InviteToWorkspaceContentId emigo.Nullable[string] `json:"inviteToWorkspaceContentId" yaml:"inviteToWorkspaceContentId"`
	// The unique-id of the template used to fill the message for email one-time-password requests.
	EmailOtpContentId emigo.Nullable[string] `json:"emailOtpContentId" yaml:"emailOtpContentId"`
	// The unique-id of the template used for OTP text messages, including the one time password code.
	SmsOtpContentId emigo.Nullable[string]  `json:"smsOtpContentId" yaml:"smsOtpContentId"`
	CreatedAt       abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt       abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
	WorkspaceId     string                  `json:"workspaceId" yaml:"workspaceId"`
}

func (x *MessagingConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetMessagingConfigEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "general-email-provider-id",
			Type:        "string?",
			Description: "The unique-id of the email provider service, which will be used to send the messages using it's service.",
		},
		{
			Name:        prefix + "general-gsm-provider-id",
			Type:        "string?",
			Description: "The unique-id of the general service which would be used to send text messages (sms).",
		},
		{
			Name:        prefix + "invite-to-workspace-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used as default when a user is inviting a third-party into their own workspace.",
		},
		{
			Name:        prefix + "email-otp-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used to fill the message for email one-time-password requests.",
		},
		{
			Name:        prefix + "sms-otp-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used for OTP text messages, including the one time password code.",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string",
		},
	}
}
func CastMessagingConfigEntityFromCli(c emigo.CliCastable) MessagingConfigEntity {
	data := MessagingConfigEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("general-email-provider-id") {
		emigo.ParseNullable(c.String("general-email-provider-id"), &data.GeneralEmailProviderId)
	}
	if c.IsSet("general-gsm-provider-id") {
		emigo.ParseNullable(c.String("general-gsm-provider-id"), &data.GeneralGsmProviderId)
	}
	if c.IsSet("invite-to-workspace-content-id") {
		emigo.ParseNullable(c.String("invite-to-workspace-content-id"), &data.InviteToWorkspaceContentId)
	}
	if c.IsSet("email-otp-content-id") {
		emigo.ParseNullable(c.String("email-otp-content-id"), &data.EmailOtpContentId)
	}
	if c.IsSet("sms-otp-content-id") {
		emigo.ParseNullable(c.String("sms-otp-content-id"), &data.SmsOtpContentId)
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
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// MessagingConfigEntityCreateFn creates a new MessagingConfigEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func MessagingConfigEntityCreateFn(tx *gorm.DB, dto *MessagingConfigEntity) (*MessagingConfigEntity, error) {
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

// MessagingConfigEntityUpdateFn applies a partial update to the MessagingConfigEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// MessagingConfigEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func MessagingConfigEntityUpdateFn(tx *gorm.DB, uniqueId string, input MessagingConfigOptionalDto) (*MessagingConfigEntity, error) {
	var entity MessagingConfigEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.GeneralEmailProviderId.IsSet() {
			changes["GeneralEmailProviderId"] = input.GeneralEmailProviderId
		}
		if input.GeneralGsmProviderId.IsSet() {
			changes["GeneralGsmProviderId"] = input.GeneralGsmProviderId
		}
		if input.InviteToWorkspaceContentId.IsSet() {
			changes["InviteToWorkspaceContentId"] = input.InviteToWorkspaceContentId
		}
		if input.EmailOtpContentId.IsSet() {
			changes["EmailOtpContentId"] = input.EmailOtpContentId
		}
		if input.SmsOtpContentId.IsSet() {
			changes["SmsOtpContentId"] = input.SmsOtpContentId
		}
		changes["CreatedAt"] = input.CreatedAt
		changes["UpdatedAt"] = input.UpdatedAt
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
		}
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
	var updated MessagingConfigEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// MessagingConfigEntityGetFn looks up a single MessagingConfigEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func MessagingConfigEntityGetFn(tx *gorm.DB, uniqueId string) (*MessagingConfigEntity, error) {
	var entity MessagingConfigEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// MessagingConfigEntityBrowseFn returns MessagingConfigEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func MessagingConfigEntityBrowseFn(tx *gorm.DB, qs MessagingConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*MessagingConfigEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&MessagingConfigEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*MessagingConfigEntity
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

/**
 * Query parameters for MessagingConfigBrowseAction
 */
// Query wrapper with private fields
type MessagingConfigBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func MessagingConfigBrowseActionQueryFromString(rawQuery string) MessagingConfigBrowseActionQuery {
	v := MessagingConfigBrowseActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func MessagingConfigBrowseActionQueryFromHttp(r *http.Request) MessagingConfigBrowseActionQuery {
	return MessagingConfigBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q MessagingConfigBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q MessagingConfigBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *MessagingConfigBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *MessagingConfigBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

// MessagingConfigEntityAwareDeleteAffected reports one relation of MessagingConfigEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on MessagingConfigEntity itself, so deleting MessagingConfigEntity doesn't cascade into them.
type MessagingConfigEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// MessagingConfigEntityAwareDeletePreview is the result of MessagingConfigEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts MessagingConfigEntityAwareDeleteFn would delete/clear
// alongside the MessagingConfigEntity row(s) themselves.
type MessagingConfigEntityAwareDeletePreview struct {
	Message  string                                     `json:"message"`
	Affected []MessagingConfigEntityAwareDeleteAffected `json:"affected"`
}

// MessagingConfigEntityAwareDeletePreviewFn looks up the MessagingConfigEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// MessagingConfigEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling MessagingConfigEntityAwareDeleteFn.
func MessagingConfigEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*MessagingConfigEntityAwareDeletePreview, error) {
	var rows []*MessagingConfigEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &MessagingConfigEntityAwareDeletePreview{Message: "No matching MessagingConfigEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []MessagingConfigEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d MessagingConfigEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &MessagingConfigEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// MessagingConfigEntityAwareDeleteFn deletes the MessagingConfigEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation MessagingConfigEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func MessagingConfigEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*MessagingConfigEntity
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
		return tx.Where("id IN ?", ids).Delete(&MessagingConfigEntity{}).Error
	})
}

// MessagingConfigEntityActionsSig bundles the actions available for MessagingConfigEntity. Extend this (and
// MessagingConfigEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type MessagingConfigEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *MessagingConfigEntity) (*MessagingConfigEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input MessagingConfigOptionalDto) (*MessagingConfigEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*MessagingConfigEntity, error)
	Browse             func(tx *gorm.DB, qs MessagingConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*MessagingConfigEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*MessagingConfigEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var MessagingConfigEntityActions MessagingConfigEntityActionsSig = MessagingConfigEntityActionsSig{
	Create:             MessagingConfigEntityCreateFn,
	Update:             MessagingConfigEntityUpdateFn,
	Get:                MessagingConfigEntityGetFn,
	Browse:             MessagingConfigEntityBrowseFn,
	AwareDeletePreview: MessagingConfigEntityAwareDeletePreviewFn,
	AwareDelete:        MessagingConfigEntityAwareDeleteFn,
}
