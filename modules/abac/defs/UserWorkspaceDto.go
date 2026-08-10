package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for userWorkspaceDto
type UserWorkspaceDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this record belongs to.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId          emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserPermissions      []string                `json:"userPermissions" sql:"-" yaml:"userPermissions"`
	RolePermission       []interface{}           `json:"rolePermission" sql:"-" yaml:"rolePermission"`
	WorkspacePermissions []string                `json:"workspacePermissions" sql:"-" yaml:"workspacePermissions"`
	CreatedAt            abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt            abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *UserWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserWorkspaceDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user this record belongs to.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name: prefix + "user-permissions",
			Type: "slice",
		},
		{
			Name: prefix + "role-permission",
			Type: "slice",
		},
		{
			Name: prefix + "workspace-permissions",
			Type: "slice",
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
func CastUserWorkspaceDtoFromCli(c emigo.CliCastable) UserWorkspaceDto {
	data := UserWorkspaceDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-permissions") {
		emigo.InflatePossibleSlice(c.String("user-permissions"), &data.UserPermissions)
	}
	if c.IsSet("role-permission") {
		emigo.InflatePossibleSlice(c.String("role-permission"), &data.RolePermission)
	}
	if c.IsSet("workspace-permissions") {
		emigo.InflatePossibleSlice(c.String("workspace-permissions"), &data.WorkspacePermissions)
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
