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

// The base class definition for publicAuthenticationEntity
type PublicAuthenticationEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user which this record belongs to.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// If the application requires totp dual factor upon account creation, we create a secret here and pass the link
	TotpSecret string `json:"totpSecret" yaml:"totpSecret"`
	// The url which will be converted into QR code on client side to scan
	TotpLink string `json:"totpLink" yaml:"totpLink"`
	// The unique-id of the passport this record belongs to.
	PassportId emigo.Nullable[string] `json:"passportId" yaml:"passportId"`
	// This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
	SessionSecret       string               `json:"sessionSecret" yaml:"sessionSecret"`
	PassportValue       string               `json:"passportValue" yaml:"passportValue"`
	IsInCreationProcess emigo.Nullable[bool] `json:"isInCreationProcess" yaml:"isInCreationProcess"`
	Status              string               `json:"status" yaml:"status"`
	BlockedUntil        int64                `json:"blockedUntil" yaml:"blockedUntil"`
	Otp                 string               `json:"otp" yaml:"otp"`
	RecoveryAbsoluteUrl string               `json:"recoveryAbsoluteUrl" sql:"-" yaml:"recoveryAbsoluteUrl"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PublicAuthenticationEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
//
func GetPublicAuthenticationEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Description: "The unique-id of the user which this record belongs to.",
		},
		{
			Name:        prefix + "totp-secret",
			Type:        "string",
			Description: "If the application requires totp dual factor upon account creation, we create a secret here and pass the link",
		},
		{
			Name:        prefix + "totp-link",
			Type:        "string",
			Description: "The url which will be converted into QR code on client side to scan",
		},
		{
			Name:        prefix + "passport-id",
			Type:        "string?",
			Description: "The unique-id of the passport this record belongs to.",
		},
		{
			Name:        prefix + "session-secret",
			Type:        "string",
			Description: "This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account",
		},
		{
			Name: prefix + "passport-value",
			Type: "string",
		},
		{
			Name: prefix + "is-in-creation-process",
			Type: "bool?",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "blocked-until",
			Type: "int64",
		},
		{
			Name: prefix + "otp",
			Type: "string",
		},
		{
			Name: prefix + "recovery-absolute-url",
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
func CastPublicAuthenticationEntityFromCli(c emigo.CliCastable) PublicAuthenticationEntity {
	data := PublicAuthenticationEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("totp-secret") {
		data.TotpSecret = c.String("totp-secret")
	}
	if c.IsSet("totp-link") {
		data.TotpLink = c.String("totp-link")
	}
	if c.IsSet("passport-id") {
		emigo.ParseNullable(c.String("passport-id"), &data.PassportId)
	}
	if c.IsSet("session-secret") {
		data.SessionSecret = c.String("session-secret")
	}
	if c.IsSet("passport-value") {
		data.PassportValue = c.String("passport-value")
	}
	if c.IsSet("is-in-creation-process") {
		emigo.ParseNullable(c.String("is-in-creation-process"), &data.IsInCreationProcess)
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("blocked-until") {
		data.BlockedUntil = int64(c.Int64("blocked-until"))
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	if c.IsSet("recovery-absolute-url") {
		data.RecoveryAbsoluteUrl = c.String("recovery-absolute-url")
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

// PublicAuthenticationEntityCreateFn creates a new PublicAuthenticationEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func PublicAuthenticationEntityCreateFn(tx *gorm.DB, dto *PublicAuthenticationEntity) (*PublicAuthenticationEntity, error) {
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

// PublicAuthenticationEntityUpdateFn applies a partial update to the PublicAuthenticationEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// PublicAuthenticationEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func PublicAuthenticationEntityUpdateFn(tx *gorm.DB, uniqueId string, input PublicAuthenticationOptionalDto) (*PublicAuthenticationEntity, error) {
	var entity PublicAuthenticationEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.TotpSecret.IsSet() {
			changes["TotpSecret"] = input.TotpSecret
		}
		if input.TotpLink.IsSet() {
			changes["TotpLink"] = input.TotpLink
		}
		if input.PassportId.IsSet() {
			changes["PassportId"] = input.PassportId
		}
		if input.SessionSecret.IsSet() {
			changes["SessionSecret"] = input.SessionSecret
		}
		if input.PassportValue.IsSet() {
			changes["PassportValue"] = input.PassportValue
		}
		if input.IsInCreationProcess.IsSet() {
			changes["IsInCreationProcess"] = input.IsInCreationProcess
		}
		if input.Status.IsSet() {
			changes["Status"] = input.Status
		}
		if input.BlockedUntil.IsSet() {
			changes["BlockedUntil"] = input.BlockedUntil
		}
		if input.Otp.IsSet() {
			changes["Otp"] = input.Otp
		}
		if input.RecoveryAbsoluteUrl.IsSet() {
			changes["RecoveryAbsoluteUrl"] = input.RecoveryAbsoluteUrl
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
	var updated PublicAuthenticationEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// PublicAuthenticationEntityGetFn looks up a single PublicAuthenticationEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func PublicAuthenticationEntityGetFn(tx *gorm.DB, uniqueId string) (*PublicAuthenticationEntity, error) {
	var entity PublicAuthenticationEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// PublicAuthenticationEntityBrowseFn returns PublicAuthenticationEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func PublicAuthenticationEntityBrowseFn(tx *gorm.DB, qs PublicAuthenticationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PublicAuthenticationEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&PublicAuthenticationEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*PublicAuthenticationEntity
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

// PublicAuthenticationEntityAwareDeleteAffected reports one relation of PublicAuthenticationEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on PublicAuthenticationEntity itself, so deleting PublicAuthenticationEntity doesn't cascade into them.
type PublicAuthenticationEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// PublicAuthenticationEntityAwareDeletePreview is the result of PublicAuthenticationEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts PublicAuthenticationEntityAwareDeleteFn would delete/clear
// alongside the PublicAuthenticationEntity row(s) themselves.
type PublicAuthenticationEntityAwareDeletePreview struct {
	Message  string                                          `json:"message"`
	Affected []PublicAuthenticationEntityAwareDeleteAffected `json:"affected"`
}

// PublicAuthenticationEntityAwareDeletePreviewFn looks up the PublicAuthenticationEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// PublicAuthenticationEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling PublicAuthenticationEntityAwareDeleteFn.
func PublicAuthenticationEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*PublicAuthenticationEntityAwareDeletePreview, error) {
	var rows []*PublicAuthenticationEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &PublicAuthenticationEntityAwareDeletePreview{Message: "No matching PublicAuthenticationEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []PublicAuthenticationEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d PublicAuthenticationEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &PublicAuthenticationEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// PublicAuthenticationEntityAwareDeleteFn deletes the PublicAuthenticationEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation PublicAuthenticationEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func PublicAuthenticationEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*PublicAuthenticationEntity
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
		return tx.Where("id IN ?", ids).Delete(&PublicAuthenticationEntity{}).Error
	})
}

// PublicAuthenticationEntityActionsSig bundles the actions available for PublicAuthenticationEntity. Extend this (and
// PublicAuthenticationEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type PublicAuthenticationEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *PublicAuthenticationEntity) (*PublicAuthenticationEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input PublicAuthenticationOptionalDto) (*PublicAuthenticationEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*PublicAuthenticationEntity, error)
	Browse             func(tx *gorm.DB, qs PublicAuthenticationBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PublicAuthenticationEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*PublicAuthenticationEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var PublicAuthenticationEntityActions PublicAuthenticationEntityActionsSig = PublicAuthenticationEntityActionsSig{
	Create:             PublicAuthenticationEntityCreateFn,
	Update:             PublicAuthenticationEntityUpdateFn,
	Get:                PublicAuthenticationEntityGetFn,
	Browse:             PublicAuthenticationEntityBrowseFn,
	AwareDeletePreview: PublicAuthenticationEntityAwareDeletePreviewFn,
	AwareDelete:        PublicAuthenticationEntityAwareDeleteFn,
}
