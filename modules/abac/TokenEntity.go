package abac

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

// The base class definition for tokenEntity
type TokenEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this token belongs to.
	UserId     emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Token      string                 `json:"token" yaml:"token"`
	ValidUntil complexes.XDateTime    `json:"validUntil" yaml:"validUntil"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *TokenEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTokenEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Description: "The unique-id of the user this token belongs to.",
		},
		{
			Name: prefix + "token",
			Type: "string",
		},
		{
			Name: prefix + "valid-until",
			Type: "complex",
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
func CastTokenEntityFromCli(c emigo.CliCastable) TokenEntity {
	data := TokenEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("token") {
		data.Token = c.String("token")
	}
	if c.IsSet("valid-until") {
		if u, ok := any(&data.ValidUntil).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("valid-until")))
		}
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
// TokenEntityCreateFn creates a new TokenEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func TokenEntityCreateFn(tx *gorm.DB, dto *TokenEntity) (*TokenEntity, error) {
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

// TokenEntityUpdateFn applies a partial update to the TokenEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// TokenEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func TokenEntityUpdateFn(tx *gorm.DB, uniqueId string, input TokenOptionalDto) (*TokenEntity, error) {
	var entity TokenEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.UserId.IsSet() {
			changes["UserId"] = input.UserId
		}
		if input.Token.IsSet() {
			changes["Token"] = input.Token
		}
		changes["ValidUntil"] = input.ValidUntil
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
	var updated TokenEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// TokenEntityGetFn looks up a single TokenEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func TokenEntityGetFn(tx *gorm.DB, uniqueId string) (*TokenEntity, error) {
	var entity TokenEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// TokenEntityBrowseFn returns TokenEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func TokenEntityBrowseFn(tx *gorm.DB, qs TokenBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TokenEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&TokenEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*TokenEntity
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

// TokenEntityAwareDeleteAffected reports one relation of TokenEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on TokenEntity itself, so deleting TokenEntity doesn't cascade into them.
type TokenEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// TokenEntityAwareDeletePreview is the result of TokenEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts TokenEntityAwareDeleteFn would delete/clear
// alongside the TokenEntity row(s) themselves.
type TokenEntityAwareDeletePreview struct {
	Message  string                           `json:"message"`
	Affected []TokenEntityAwareDeleteAffected `json:"affected"`
}

// TokenEntityAwareDeletePreviewFn looks up the TokenEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// TokenEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling TokenEntityAwareDeleteFn.
func TokenEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*TokenEntityAwareDeletePreview, error) {
	var rows []*TokenEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &TokenEntityAwareDeletePreview{Message: "No matching TokenEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []TokenEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d TokenEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &TokenEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// TokenEntityAwareDeleteFn deletes the TokenEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation TokenEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func TokenEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*TokenEntity
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
		return tx.Where("id IN ?", ids).Delete(&TokenEntity{}).Error
	})
}

// TokenEntityActionsSig bundles the actions available for TokenEntity. Extend this (and
// TokenEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type TokenEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *TokenEntity) (*TokenEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input TokenOptionalDto) (*TokenEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*TokenEntity, error)
	Browse             func(tx *gorm.DB, qs TokenBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*TokenEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*TokenEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var TokenEntityActions TokenEntityActionsSig = TokenEntityActionsSig{
	Create:             TokenEntityCreateFn,
	Update:             TokenEntityUpdateFn,
	Get:                TokenEntityGetFn,
	Browse:             TokenEntityBrowseFn,
	AwareDeletePreview: TokenEntityAwareDeletePreviewFn,
	AwareDelete:        TokenEntityAwareDeleteFn,
}
