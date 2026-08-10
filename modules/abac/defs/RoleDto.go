package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for roleDto
type RoleDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name     string                 `json:"name" validate:"required,omitempty,min=1,max=200" yaml:"name"`
	// The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
	CapabilitiesListId complexes.JSON          `json:"capabilitiesListId" yaml:"capabilitiesListId"`
	IsDeletable        emigo.Nullable[bool]    `json:"isDeletable" yaml:"isDeletable"`
	IsUpdatable        emigo.Nullable[bool]    `json:"isUpdatable" yaml:"isUpdatable"`
	WorkspaceId        emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId             emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt          abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt          abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RoleDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetRoleDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:        prefix + "capabilities-list-id",
			Type:        "complex",
			Description: "The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).",
		},
		{
			Name: prefix + "is-deletable",
			Type: "bool?",
		},
		{
			Name: prefix + "is-updatable",
			Type: "bool?",
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
func CastRoleDtoFromCli(c emigo.CliCastable) RoleDto {
	data := RoleDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("capabilities-list-id") {
		if u, ok := any(&data.CapabilitiesListId).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("capabilities-list-id")))
		}
	}
	if c.IsSet("is-deletable") {
		emigo.ParseNullable(c.String("is-deletable"), &data.IsDeletable)
	}
	if c.IsSet("is-updatable") {
		emigo.ParseNullable(c.String("is-updatable"), &data.IsUpdatable)
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
