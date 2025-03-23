package workspaces

import (
	"encoding/json"
)

type UserAccessPerWorkspaceDto map[string]*struct {

	// The access which are available to this workspace, not to the specific user.
	// Even a user has access to many things, these accesses need to reduce those
	WorkspacesAccesses []string

	// The permissions which user has access to
	UserRoles map[string]*struct {
		Name     string
		Accesses []string
	}
}

func (x *UserAccessPerWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
