package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for assignRoleDto
type AssignRoleDto struct {
	RoleId     string `json:"roleId" yaml:"roleId"`
	UserId     string `json:"userId" yaml:"userId"`
	Visibility string `json:"visibility" yaml:"visibility"`
	Updated    int64  `json:"updated" yaml:"updated"`
	Created    int64  `json:"created" yaml:"created"`
}

func (x *AssignRoleDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
