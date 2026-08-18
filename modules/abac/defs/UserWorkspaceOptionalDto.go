package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for userWorkspaceOptionalDto
type UserWorkspaceOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this record belongs to.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId          emigo.Nullable[string]        `json:"workspaceId" yaml:"workspaceId"`
	UserPermissions      emigo.Nullable[[]string]      `json:"userPermissions" sql:"-" yaml:"userPermissions"`
	RolePermission       emigo.Nullable[[]interface{}] `json:"rolePermission" sql:"-" yaml:"rolePermission"`
	WorkspacePermissions emigo.Nullable[[]string]      `json:"workspacePermissions" sql:"-" yaml:"workspacePermissions"`
	CreatedAt            abaccomplexes.PlainTime       `json:"createdAt" yaml:"createdAt"`
	UpdatedAt            abaccomplexes.PlainTime       `json:"updatedAt" yaml:"updatedAt"`
}

func (x *UserWorkspaceOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
