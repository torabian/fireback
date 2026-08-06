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

// The base class definition for passportEntity
type PassportEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
	ThirdPartyVerifier string                 `json:"thirdPartyVerifier" yaml:"thirdPartyVerifier"`
	Type               string                 `json:"type" validate:"required" yaml:"type"`
	UserId             emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Value              string                 `gorm:"unique" json:"value" validate:"required" yaml:"value"`
	// Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
	TotpSecret string `json:"totpSecret" yaml:"totpSecret"`
	// Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
	TotpConfirmed emigo.Nullable[bool]    `json:"totpConfirmed" yaml:"totpConfirmed"`
	Password      string                  `json:"-" yaml:"-"`
	Confirmed     emigo.Nullable[bool]    `json:"confirmed" yaml:"confirmed"`
	AccessToken   string                  `json:"accessToken" yaml:"accessToken"`
	WorkspaceId   emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt     abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt     abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PassportEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPassportEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "third-party-verifier",
			Type:        "string",
			Description: "When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
		},
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name:        prefix + "totp-secret",
			Type:        "string",
			Description: "Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.",
		},
		{
			Name:        prefix + "totp-confirmed",
			Type:        "bool?",
			Description: "Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "confirmed",
			Type: "bool?",
		},
		{
			Name: prefix + "access-token",
			Type: "string",
		},
		{
			Name: prefix + "workspace-id",
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
func CastPassportEntityFromCli(c emigo.CliCastable) PassportEntity {
	data := PassportEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("third-party-verifier") {
		data.ThirdPartyVerifier = c.String("third-party-verifier")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("totp-secret") {
		data.TotpSecret = c.String("totp-secret")
	}
	if c.IsSet("totp-confirmed") {
		emigo.ParseNullable(c.String("totp-confirmed"), &data.TotpConfirmed)
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("confirmed") {
		emigo.ParseNullable(c.String("confirmed"), &data.Confirmed)
	}
	if c.IsSet("access-token") {
		data.AccessToken = c.String("access-token")
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
// PassportEntityCreateFn creates a new PassportEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func PassportEntityCreateFn(tx *gorm.DB, dto *PassportEntity) (*PassportEntity, error) {
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

// PassportEntityUpdateFn applies a partial update to the PassportEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// PassportEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func PassportEntityUpdateFn(tx *gorm.DB, uniqueId string, input PassportOptionalDto) (*PassportEntity, error) {
	var entity PassportEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.ThirdPartyVerifier.IsSet() {
			changes["ThirdPartyVerifier"] = input.ThirdPartyVerifier
		}
		if input.Type.IsSet() {
			changes["Type"] = input.Type
		}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.Value.IsSet() {
			changes["Value"] = input.Value
		}
		if input.TotpSecret.IsSet() {
			changes["TotpSecret"] = input.TotpSecret
		}
		if input.TotpConfirmed.IsSet() {
			changes["TotpConfirmed"] = input.TotpConfirmed
		}
		if input.Password.IsSet() {
			changes["Password"] = input.Password
		}
		if input.Confirmed.IsSet() {
			changes["Confirmed"] = input.Confirmed
		}
		if input.AccessToken.IsSet() {
			changes["AccessToken"] = input.AccessToken
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
	var updated PassportEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// PassportEntityGetFn looks up a single PassportEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func PassportEntityGetFn(tx *gorm.DB, uniqueId string) (*PassportEntity, error) {
	var entity PassportEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// PassportEntityBrowseFn returns PassportEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func PassportEntityBrowseFn(tx *gorm.DB, qs PassportBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PassportEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&PassportEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*PassportEntity
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

// PassportEntityAwareDeleteAffected reports one relation of PassportEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on PassportEntity itself, so deleting PassportEntity doesn't cascade into them.
type PassportEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// PassportEntityAwareDeletePreview is the result of PassportEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts PassportEntityAwareDeleteFn would delete/clear
// alongside the PassportEntity row(s) themselves.
type PassportEntityAwareDeletePreview struct {
	Message  string                              `json:"message"`
	Affected []PassportEntityAwareDeleteAffected `json:"affected"`
}

// PassportEntityAwareDeletePreviewFn looks up the PassportEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// PassportEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling PassportEntityAwareDeleteFn.
func PassportEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*PassportEntityAwareDeletePreview, error) {
	var rows []*PassportEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &PassportEntityAwareDeletePreview{Message: "No matching PassportEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []PassportEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d PassportEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &PassportEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// PassportEntityAwareDeleteFn deletes the PassportEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation PassportEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func PassportEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*PassportEntity
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
		return tx.Where("id IN ?", ids).Delete(&PassportEntity{}).Error
	})
}

// PassportEntityActionsSig bundles the actions available for PassportEntity. Extend this (and
// PassportEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type PassportEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *PassportEntity) (*PassportEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input PassportOptionalDto) (*PassportEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*PassportEntity, error)
	Browse             func(tx *gorm.DB, qs PassportBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PassportEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*PassportEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var PassportEntityActions PassportEntityActionsSig = PassportEntityActionsSig{
	Create:             PassportEntityCreateFn,
	Update:             PassportEntityUpdateFn,
	Get:                PassportEntityGetFn,
	Browse:             PassportEntityBrowseFn,
	AwareDeletePreview: PassportEntityAwareDeletePreviewFn,
	AwareDelete:        PassportEntityAwareDeleteFn,
}
