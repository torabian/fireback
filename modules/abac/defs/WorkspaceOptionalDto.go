package abacdefs

import (
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
