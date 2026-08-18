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

// The base class definition for workspaceInviteEntity
type WorkspaceInviteEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
	PublicKey string `json:"publicKey" yaml:"publicKey"`
	// The content that user will receive to understand the reason of the letter.
	CoverLetter string `json:"coverLetter" yaml:"coverLetter"`
	// If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
	TargetUserLocale string `json:"targetUserLocale" yaml:"targetUserLocale"`
	// The email address of the person which is invited.
	Email string `json:"email" yaml:"email"`
	// The phone number of the person which is invited.
	Phonenumber string `json:"phonenumber" yaml:"phonenumber"`
	// The unique-id of the workspace which user is being invited to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// First name of the person which is invited
	FirstName string `json:"firstName" validate:"required" yaml:"firstName"`
	// Last name of the person which is invited.
	LastName string `json:"lastName" validate:"required" yaml:"lastName"`
	// If forced, the email address cannot be changed by the user which has been invited.
	ForceEmailAddress emigo.Nullable[bool] `json:"forceEmailAddress" yaml:"forceEmailAddress"`
	// If forced, user cannot change the phone number and needs to complete signup.
	ForcePhoneNumber emigo.Nullable[bool] `json:"forcePhoneNumber" yaml:"forcePhoneNumber"`
	// The role which invitee get if they accept the request.
	RoleId emigo.Nullable[string] `json:"roleId" validate:"required" yaml:"roleId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceInviteEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetWorkspaceInviteEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "public-key",
			Type:        "string",
			Description: "A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.",
		},
		{
			Name:        prefix + "cover-letter",
			Type:        "string",
			Description: "The content that user will receive to understand the reason of the letter.",
		},
		{
			Name:        prefix + "target-user-locale",
			Type:        "string",
			Description: "If the invited person has a different language, then you can define that so the interface for him will be automatically translated.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "The email address of the person which is invited.",
		},
		{
			Name:        prefix + "phonenumber",
			Type:        "string",
			Description: "The phone number of the person which is invited.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which user is being invited to.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string",
			Description: "First name of the person which is invited",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string",
			Description: "Last name of the person which is invited.",
		},
		{
			Name:        prefix + "force-email-address",
			Type:        "bool?",
			Description: "If forced, the email address cannot be changed by the user which has been invited.",
		},
		{
			Name:        prefix + "force-phone-number",
			Type:        "bool?",
			Description: "If forced, user cannot change the phone number and needs to complete signup.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string?",
			Description: "The role which invitee get if they accept the request.",
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
func CastWorkspaceInviteEntityFromCli(c emigo.CliCastable) WorkspaceInviteEntity {
	data := WorkspaceInviteEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("public-key") {
		data.PublicKey = c.String("public-key")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	if c.IsSet("target-user-locale") {
		data.TargetUserLocale = c.String("target-user-locale")
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("phonenumber") {
		data.Phonenumber = c.String("phonenumber")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("force-email-address") {
		emigo.ParseNullable(c.String("force-email-address"), &data.ForceEmailAddress)
	}
	if c.IsSet("force-phone-number") {
		emigo.ParseNullable(c.String("force-phone-number"), &data.ForcePhoneNumber)
	}
	if c.IsSet("role-id") {
		emigo.ParseNullable(c.String("role-id"), &data.RoleId)
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

// WorkspaceInviteEntityCreateFn creates a new WorkspaceInviteEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WorkspaceInviteEntityCreateFn(tx *gorm.DB, dto *WorkspaceInviteEntity) (*WorkspaceInviteEntity, error) {
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

// WorkspaceInviteEntityUpdateFn applies a partial update to the WorkspaceInviteEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WorkspaceInviteEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WorkspaceInviteEntityUpdateFn(tx *gorm.DB, uniqueId string, input WorkspaceInviteOptionalDto) (*WorkspaceInviteEntity, error) {
	var entity WorkspaceInviteEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.PublicKey.IsSet() {
			changes["PublicKey"] = input.PublicKey
		}
		if input.CoverLetter.IsSet() {
			changes["CoverLetter"] = input.CoverLetter
		}
		if input.TargetUserLocale.IsSet() {
			changes["TargetUserLocale"] = input.TargetUserLocale
		}
		if input.Email.IsSet() {
			changes["Email"] = input.Email
		}
		if input.Phonenumber.IsSet() {
			changes["Phonenumber"] = input.Phonenumber
		}
		if input.WorkspaceId.IsSet() {
			changes["WorkspaceId"] = input.WorkspaceId
		}
		if input.FirstName.IsSet() {
			changes["FirstName"] = input.FirstName
		}
		if input.LastName.IsSet() {
			changes["LastName"] = input.LastName
		}
		if input.ForceEmailAddress.IsSet() {
			changes["ForceEmailAddress"] = input.ForceEmailAddress
		}
		if input.ForcePhoneNumber.IsSet() {
			changes["ForcePhoneNumber"] = input.ForcePhoneNumber
		}
		if input.RoleId.IsSet() {
			changes["RoleId"] = input.RoleId
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
	var updated WorkspaceInviteEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WorkspaceInviteEntityGetFn looks up a single WorkspaceInviteEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WorkspaceInviteEntityGetFn(tx *gorm.DB, uniqueId string) (*WorkspaceInviteEntity, error) {
	var entity WorkspaceInviteEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WorkspaceInviteEntityBrowseFn returns WorkspaceInviteEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WorkspaceInviteEntityBrowseFn(tx *gorm.DB, qs WorkspaceInviteBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceInviteEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WorkspaceInviteEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WorkspaceInviteEntity
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

// WorkspaceInviteEntityAwareDeleteAffected reports one relation of WorkspaceInviteEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WorkspaceInviteEntity itself, so deleting WorkspaceInviteEntity doesn't cascade into them.
type WorkspaceInviteEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WorkspaceInviteEntityAwareDeletePreview is the result of WorkspaceInviteEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WorkspaceInviteEntityAwareDeleteFn would delete/clear
// alongside the WorkspaceInviteEntity row(s) themselves.
type WorkspaceInviteEntityAwareDeletePreview struct {
	Message  string                                     `json:"message"`
	Affected []WorkspaceInviteEntityAwareDeleteAffected `json:"affected"`
}

// WorkspaceInviteEntityAwareDeletePreviewFn looks up the WorkspaceInviteEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WorkspaceInviteEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WorkspaceInviteEntityAwareDeleteFn.
func WorkspaceInviteEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WorkspaceInviteEntityAwareDeletePreview, error) {
	var rows []*WorkspaceInviteEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WorkspaceInviteEntityAwareDeletePreview{Message: "No matching WorkspaceInviteEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WorkspaceInviteEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WorkspaceInviteEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WorkspaceInviteEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WorkspaceInviteEntityAwareDeleteFn deletes the WorkspaceInviteEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WorkspaceInviteEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WorkspaceInviteEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WorkspaceInviteEntity
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
		return tx.Where("id IN ?", ids).Delete(&WorkspaceInviteEntity{}).Error
	})
}

// WorkspaceInviteEntityActionsSig bundles the actions available for WorkspaceInviteEntity. Extend this (and
// WorkspaceInviteEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WorkspaceInviteEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WorkspaceInviteEntity) (*WorkspaceInviteEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WorkspaceInviteOptionalDto) (*WorkspaceInviteEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WorkspaceInviteEntity, error)
	Browse             func(tx *gorm.DB, qs WorkspaceInviteBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WorkspaceInviteEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WorkspaceInviteEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WorkspaceInviteEntityActions WorkspaceInviteEntityActionsSig = WorkspaceInviteEntityActionsSig{
	Create:             WorkspaceInviteEntityCreateFn,
	Update:             WorkspaceInviteEntityUpdateFn,
	Get:                WorkspaceInviteEntityGetFn,
	Browse:             WorkspaceInviteEntityBrowseFn,
	AwareDeletePreview: WorkspaceInviteEntityAwareDeletePreviewFn,
	AwareDelete:        WorkspaceInviteEntityAwareDeleteFn,
}
