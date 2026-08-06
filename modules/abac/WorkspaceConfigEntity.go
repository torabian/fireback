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

// The base class definition for workspaceConfigEntity
type WorkspaceConfigEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// Enables the recaptcha2 for authentication flow.
	EnableRecaptcha2 emigo.Nullable[bool] `json:"enableRecaptcha2" yaml:"enableRecaptcha2"`
	// Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
	EnableOtp emigo.Nullable[bool] `json:"enableOtp" yaml:"enableOtp"`
	// Forces the user to have otp verification before can create an account. They can define their password still.
	RequireOtpOnSignup emigo.Nullable[bool] `json:"requireOtpOnSignup" yaml:"requireOtpOnSignup"`
	// Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
	RequireOtpOnSignin emigo.Nullable[bool] `json:"requireOtpOnSignin" yaml:"requireOtpOnSignin"`
	// Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
	Recaptcha2ServerKey string `json:"recaptcha2ServerKey" yaml:"recaptcha2ServerKey"`
	// Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
	Recaptcha2ClientKey string `json:"recaptcha2ClientKey" yaml:"recaptcha2ClientKey"`
	// Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
	EnableTotp emigo.Nullable[bool] `json:"enableTotp" yaml:"enableTotp"`
	// Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
	ForceTotp emigo.Nullable[bool] `json:"forceTotp" yaml:"forceTotp"`
	// Forces users who want to create account using phone number to also set a password on their account
	ForcePasswordOnPhone emigo.Nullable[bool] `json:"forcePasswordOnPhone" yaml:"forcePasswordOnPhone"`
	// Forces the creation of account using phone number to ask for user first name and last name
	ForcePersonNameOnPhone emigo.Nullable[bool] `json:"forcePersonNameOnPhone" yaml:"forcePersonNameOnPhone"`
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
	WorkspaceId     emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId          emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt       abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt       abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceConfigEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "enable-recaptcha2",
			Type:        "bool?",
			Description: "Enables the recaptcha2 for authentication flow.",
		},
		{
			Name:        prefix + "enable-otp",
			Type:        "bool?",
			Description: "Enables the otp option. It's not forcing it, so user can choose if they want otp or password.",
		},
		{
			Name:        prefix + "require-otp-on-signup",
			Type:        "bool?",
			Description: "Forces the user to have otp verification before can create an account. They can define their password still.",
		},
		{
			Name:        prefix + "require-otp-on-signin",
			Type:        "bool?",
			Description: "Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.",
		},
		{
			Name:        prefix + "recaptcha2-server-key",
			Type:        "string",
			Description: "Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.",
		},
		{
			Name:        prefix + "recaptcha2-client-key",
			Type:        "string",
			Description: "Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.",
		},
		{
			Name:        prefix + "enable-totp",
			Type:        "bool?",
			Description: "Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.",
		},
		{
			Name:        prefix + "force-totp",
			Type:        "bool?",
			Description: "Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.",
		},
		{
			Name:        prefix + "force-password-on-phone",
			Type:        "bool?",
			Description: "Forces users who want to create account using phone number to also set a password on their account",
		},
		{
			Name:        prefix + "force-person-name-on-phone",
			Type:        "bool?",
			Description: "Forces the creation of account using phone number to ask for user first name and last name",
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
func CastWorkspaceConfigEntityFromCli(c emigo.CliCastable) WorkspaceConfigEntity {
	data := WorkspaceConfigEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("enable-recaptcha2") {
		emigo.ParseNullable(c.String("enable-recaptcha2"), &data.EnableRecaptcha2)
	}
	if c.IsSet("enable-otp") {
		emigo.ParseNullable(c.String("enable-otp"), &data.EnableOtp)
	}
	if c.IsSet("require-otp-on-signup") {
		emigo.ParseNullable(c.String("require-otp-on-signup"), &data.RequireOtpOnSignup)
	}
	if c.IsSet("require-otp-on-signin") {
		emigo.ParseNullable(c.String("require-otp-on-signin"), &data.RequireOtpOnSignin)
	}
	if c.IsSet("recaptcha2-server-key") {
		data.Recaptcha2ServerKey = c.String("recaptcha2-server-key")
	}
	if c.IsSet("recaptcha2-client-key") {
		data.Recaptcha2ClientKey = c.String("recaptcha2-client-key")
	}
	if c.IsSet("enable-totp") {
		emigo.ParseNullable(c.String("enable-totp"), &data.EnableTotp)
	}
	if c.IsSet("force-totp") {
		emigo.ParseNullable(c.String("force-totp"), &data.ForceTotp)
	}
	if c.IsSet("force-password-on-phone") {
		emigo.ParseNullable(c.String("force-password-on-phone"), &data.ForcePasswordOnPhone)
	}
	if c.IsSet("force-person-name-on-phone") {
		emigo.ParseNullable(c.String("force-person-name-on-phone"), &data.ForcePersonNameOnPhone)
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
// WorkspaceConfigEntityCreateFn creates a new WorkspaceConfigEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WorkspaceConfigEntityCreateFn(tx *gorm.DB, dto *WorkspaceConfigEntity) (*WorkspaceConfigEntity, error) {
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

// WorkspaceConfigEntityUpdateFn applies a partial update to the WorkspaceConfigEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WorkspaceConfigEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WorkspaceConfigEntityUpdateFn(tx *gorm.DB, uniqueId string, input WorkspaceConfigOptionalDto) (*WorkspaceConfigEntity, error) {
	var entity WorkspaceConfigEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.EnableRecaptcha2.IsSet() {
			changes["EnableRecaptcha2"] = input.EnableRecaptcha2
		}
		if input.EnableOtp.IsSet() {
			changes["EnableOtp"] = input.EnableOtp
		}
		if input.RequireOtpOnSignup.IsSet() {
			changes["RequireOtpOnSignup"] = input.RequireOtpOnSignup
		}
		if input.RequireOtpOnSignin.IsSet() {
			changes["RequireOtpOnSignin"] = input.RequireOtpOnSignin
		}
		if input.Recaptcha2ServerKey.IsSet() {
			changes["Recaptcha2ServerKey"] = input.Recaptcha2ServerKey
		}
		if input.Recaptcha2ClientKey.IsSet() {
			changes["Recaptcha2ClientKey"] = input.Recaptcha2ClientKey
		}
		if input.EnableTotp.IsSet() {
			changes["EnableTotp"] = input.EnableTotp
		}
		if input.ForceTotp.IsSet() {
			changes["ForceTotp"] = input.ForceTotp
		}
		if input.ForcePasswordOnPhone.IsSet() {
			changes["ForcePasswordOnPhone"] = input.ForcePasswordOnPhone
		}
		if input.ForcePersonNameOnPhone.IsSet() {
			changes["ForcePersonNameOnPhone"] = input.ForcePersonNameOnPhone
		}
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
	var updated WorkspaceConfigEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WorkspaceConfigEntityGetFn looks up a single WorkspaceConfigEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WorkspaceConfigEntityGetFn(tx *gorm.DB, uniqueId string) (*WorkspaceConfigEntity, error) {
	var entity WorkspaceConfigEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WorkspaceConfigEntityBrowseFn returns WorkspaceConfigEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WorkspaceConfigEntityBrowseFn(tx *gorm.DB, qs WorkspaceConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceConfigEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WorkspaceConfigEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WorkspaceConfigEntity
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

// WorkspaceConfigEntityAwareDeleteAffected reports one relation of WorkspaceConfigEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WorkspaceConfigEntity itself, so deleting WorkspaceConfigEntity doesn't cascade into them.
type WorkspaceConfigEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WorkspaceConfigEntityAwareDeletePreview is the result of WorkspaceConfigEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WorkspaceConfigEntityAwareDeleteFn would delete/clear
// alongside the WorkspaceConfigEntity row(s) themselves.
type WorkspaceConfigEntityAwareDeletePreview struct {
	Message  string                                     `json:"message"`
	Affected []WorkspaceConfigEntityAwareDeleteAffected `json:"affected"`
}

// WorkspaceConfigEntityAwareDeletePreviewFn looks up the WorkspaceConfigEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WorkspaceConfigEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WorkspaceConfigEntityAwareDeleteFn.
func WorkspaceConfigEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WorkspaceConfigEntityAwareDeletePreview, error) {
	var rows []*WorkspaceConfigEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WorkspaceConfigEntityAwareDeletePreview{Message: "No matching WorkspaceConfigEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WorkspaceConfigEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WorkspaceConfigEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WorkspaceConfigEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WorkspaceConfigEntityAwareDeleteFn deletes the WorkspaceConfigEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WorkspaceConfigEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WorkspaceConfigEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WorkspaceConfigEntity
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
		return tx.Where("id IN ?", ids).Delete(&WorkspaceConfigEntity{}).Error
	})
}

// WorkspaceConfigEntityActionsSig bundles the actions available for WorkspaceConfigEntity. Extend this (and
// WorkspaceConfigEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WorkspaceConfigEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WorkspaceConfigEntity) (*WorkspaceConfigEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WorkspaceConfigOptionalDto) (*WorkspaceConfigEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WorkspaceConfigEntity, error)
	Browse             func(tx *gorm.DB, qs WorkspaceConfigBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceConfigEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WorkspaceConfigEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WorkspaceConfigEntityActions WorkspaceConfigEntityActionsSig = WorkspaceConfigEntityActionsSig{
	Create:             WorkspaceConfigEntityCreateFn,
	Update:             WorkspaceConfigEntityUpdateFn,
	Get:                WorkspaceConfigEntityGetFn,
	Browse:             WorkspaceConfigEntityBrowseFn,
	AwareDeletePreview: WorkspaceConfigEntityAwareDeletePreviewFn,
	AwareDelete:        WorkspaceConfigEntityAwareDeleteFn,
}
