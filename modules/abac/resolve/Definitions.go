package resolve

import (
	"encoding/json"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/application"
)

// Used for actions generally
type SecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can
	// write to this entity from Fireback default functionality
	AllowOnRoot bool `json:"allowOnRoot,omitempty" yaml:"allowOnRoot,omitempty"`

	// Set of permissions which are required for this service.
	ActionRequires []application.PermissionInfo `json:"requires,omitempty" yaml:"requires,omitempty"`

	// Resolve strategy is by default on the workspace, you can change it by user
	// also. Be sure of the consequences
	ResolveStrategy string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty"`
}

type UserAccessPerWorkspaceDto map[string]*struct {
	Name string
	// The access which are available to this workspace, not to the specific user.
	// Even a user has access to many things, these accesses need to reduce those
	WorkspacesAccesses []string

	// The permissions which user has access to
	UserRoles map[string]*struct {
		Name     string
		Accesses []string
	}
}

func (x UserAccessPerWorkspaceDto) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

type UserRoleWorkspacePermissionDto struct {
	WorkspaceName string `json:"workspaceName" yaml:"workspaceName"        `
	WorkspaceId   string `json:"workspaceId" yaml:"workspaceId"        `
	RoleName      string `json:"roleName" yaml:"roleName"        `
	UserId        string `json:"userId" yaml:"userId"        `
	RoleId        string `json:"roleId" yaml:"roleId"        `
	CapabilityId  string `json:"capabilityId" yaml:"capabilityId"        `
	Type          string `json:"type" yaml:"type"        `
}
type UserRoleWorkspacePermissionDtoList struct {
	Items []*UserRoleWorkspacePermissionDto
}

type UserAccessLevelDto struct {
	UserAccessPerWorkspace   *UserAccessPerWorkspaceDto `json:"userAccessPerWorkspace" yaml:"userAccessPerWorkspace"    gorm:"foreignKey:UserAccessPerWorkspaceId;references:UniqueId"      `
	UserAccessPerWorkspaceId emigo.Nullable[string]     `json:"userAccessPerWorkspaceId" yaml:"userAccessPerWorkspaceId"`
}
type UserAccessLevelDtoList struct {
	Items []*UserAccessLevelDto
}
