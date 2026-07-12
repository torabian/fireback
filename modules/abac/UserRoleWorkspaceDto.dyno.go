package abac

import "encoding/json"

// The base class definition for userRoleWorkspaceDto
type UserRoleWorkspaceDto struct {
	RoleId       string   `json:"roleId" yaml:"roleId"`
	Capabilities []string `json:"capabilities" yaml:"capabilities"`
}

func (x *UserRoleWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
