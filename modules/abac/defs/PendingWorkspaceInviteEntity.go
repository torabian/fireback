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

// The base class definition for pendingWorkspaceInviteEntity
type PendingWorkspaceInviteEntity struct {
	Id            int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId      string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Value         string `json:"value" yaml:"value"`
	Type          string `json:"type" yaml:"type"`
	CoverLetter   string `json:"coverLetter" yaml:"coverLetter"`
	WorkspaceName string `json:"workspaceName" yaml:"workspaceName"`
	// The unique-id of the role which invitee will get if they accept the request.
	RoleId emigo.Nullable[string] `json:"roleId" yaml:"roleId"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PendingWorkspaceInviteEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPendingWorkspaceInviteEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string",
		},
		{
			Name: prefix + "workspace-name",
			Type: "string",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string?",
			Description: "The unique-id of the role which invitee will get if they accept the request.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
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
func CastPendingWorkspaceInviteEntityFromCli(c emigo.CliCastable) PendingWorkspaceInviteEntity {
	data := PendingWorkspaceInviteEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	if c.IsSet("workspace-name") {
		data.WorkspaceName = c.String("workspace-name")
	}
	if c.IsSet("role-id") {
		emigo.ParseNullable(c.String("role-id"), &data.RoleId)
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
// PendingWorkspaceInviteEntityCreateFn creates a new PendingWorkspaceInviteEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func PendingWorkspaceInviteEntityCreateFn(tx *gorm.DB, dto *PendingWorkspaceInviteEntity) (*PendingWorkspaceInviteEntity, error) {
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

// PendingWorkspaceInviteEntityUpdateFn applies a partial update to the PendingWorkspaceInviteEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// PendingWorkspaceInviteEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func PendingWorkspaceInviteEntityUpdateFn(tx *gorm.DB, uniqueId string, input PendingWorkspaceInviteOptionalDto) (*PendingWorkspaceInviteEntity, error) {
	var entity PendingWorkspaceInviteEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Value.IsSet() {
			changes["Value"] = input.Value
		}
		if input.Type.IsSet() {
			changes["Type"] = input.Type
		}
		if input.CoverLetter.IsSet() {
			changes["CoverLetter"] = input.CoverLetter
		}
		if input.WorkspaceName.IsSet() {
			changes["WorkspaceName"] = input.WorkspaceName
		}
		if input.RoleId.IsSet() {
			changes["RoleId"] = input.RoleId
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
	var updated PendingWorkspaceInviteEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// PendingWorkspaceInviteEntityGetFn looks up a single PendingWorkspaceInviteEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func PendingWorkspaceInviteEntityGetFn(tx *gorm.DB, uniqueId string) (*PendingWorkspaceInviteEntity, error) {
	var entity PendingWorkspaceInviteEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// PendingWorkspaceInviteEntityBrowseFn returns PendingWorkspaceInviteEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func PendingWorkspaceInviteEntityBrowseFn(tx *gorm.DB, qs PendingWorkspaceInviteBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PendingWorkspaceInviteEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&PendingWorkspaceInviteEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*PendingWorkspaceInviteEntity
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

// PendingWorkspaceInviteEntityAwareDeleteAffected reports one relation of PendingWorkspaceInviteEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on PendingWorkspaceInviteEntity itself, so deleting PendingWorkspaceInviteEntity doesn't cascade into them.
type PendingWorkspaceInviteEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// PendingWorkspaceInviteEntityAwareDeletePreview is the result of PendingWorkspaceInviteEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts PendingWorkspaceInviteEntityAwareDeleteFn would delete/clear
// alongside the PendingWorkspaceInviteEntity row(s) themselves.
type PendingWorkspaceInviteEntityAwareDeletePreview struct {
	Message  string                                            `json:"message"`
	Affected []PendingWorkspaceInviteEntityAwareDeleteAffected `json:"affected"`
}

// PendingWorkspaceInviteEntityAwareDeletePreviewFn looks up the PendingWorkspaceInviteEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// PendingWorkspaceInviteEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling PendingWorkspaceInviteEntityAwareDeleteFn.
func PendingWorkspaceInviteEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*PendingWorkspaceInviteEntityAwareDeletePreview, error) {
	var rows []*PendingWorkspaceInviteEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &PendingWorkspaceInviteEntityAwareDeletePreview{Message: "No matching PendingWorkspaceInviteEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []PendingWorkspaceInviteEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d PendingWorkspaceInviteEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &PendingWorkspaceInviteEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// PendingWorkspaceInviteEntityAwareDeleteFn deletes the PendingWorkspaceInviteEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation PendingWorkspaceInviteEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func PendingWorkspaceInviteEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*PendingWorkspaceInviteEntity
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
		return tx.Where("id IN ?", ids).Delete(&PendingWorkspaceInviteEntity{}).Error
	})
}

// PendingWorkspaceInviteEntityActionsSig bundles the actions available for PendingWorkspaceInviteEntity. Extend this (and
// PendingWorkspaceInviteEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type PendingWorkspaceInviteEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *PendingWorkspaceInviteEntity) (*PendingWorkspaceInviteEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input PendingWorkspaceInviteOptionalDto) (*PendingWorkspaceInviteEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*PendingWorkspaceInviteEntity, error)
	Browse             func(tx *gorm.DB, qs PendingWorkspaceInviteBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*PendingWorkspaceInviteEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*PendingWorkspaceInviteEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var PendingWorkspaceInviteEntityActions PendingWorkspaceInviteEntityActionsSig = PendingWorkspaceInviteEntityActionsSig{
	Create:             PendingWorkspaceInviteEntityCreateFn,
	Update:             PendingWorkspaceInviteEntityUpdateFn,
	Get:                PendingWorkspaceInviteEntityGetFn,
	Browse:             PendingWorkspaceInviteEntityBrowseFn,
	AwareDeletePreview: PendingWorkspaceInviteEntityAwareDeletePreviewFn,
	AwareDelete:        PendingWorkspaceInviteEntityAwareDeleteFn,
}
