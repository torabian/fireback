package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for workspaceOptionalDto
type WorkspaceOptionalDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Description emigo.Nullable[string] `json:"description" yaml:"description"`
	Name        emigo.Nullable[string] `json:"name" validate:"required" yaml:"name"`
	// The unique-id of the workspace type which defines this workspace's role.
	TypeId emigo.Nullable[string] `json:"typeId" validate:"required" yaml:"typeId"`
	// The unique-id of the parent workspace, for nested/tree workspaces.
	ParentId    emigo.Nullable[string]  `json:"parentId" yaml:"parentId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "description",
			Type: "string?",
		},
		{
			Name: prefix + "name",
			Type: "string?",
		},
		{
			Name:        prefix + "type-id",
			Type:        "string?",
			Description: "The unique-id of the workspace type which defines this workspace's role.",
		},
		{
			Name:        prefix + "parent-id",
			Type:        "string?",
			Description: "The unique-id of the parent workspace, for nested/tree workspaces.",
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
func CastWorkspaceOptionalDtoFromCli(c emigo.CliCastable) WorkspaceOptionalDto {
	data := WorkspaceOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("description") {
		emigo.ParseNullable(c.String("description"), &data.Description)
	}
	if c.IsSet("name") {
		emigo.ParseNullable(c.String("name"), &data.Name)
	}
	if c.IsSet("type-id") {
		emigo.ParseNullable(c.String("type-id"), &data.TypeId)
	}
	if c.IsSet("parent-id") {
		emigo.ParseNullable(c.String("parent-id"), &data.ParentId)
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
