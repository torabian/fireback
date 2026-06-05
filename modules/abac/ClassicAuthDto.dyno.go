package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for classicAuthDto
type ClassicAuthDto struct {
	Value           string `json:"value" validate:"required" yaml:"value"`
	Password        string `json:"password" validate:"required" yaml:"password"`
	FirstName       string `json:"firstName" validate:"required" yaml:"firstName"`
	LastName        string `json:"lastName" validate:"required" yaml:"lastName"`
	InviteId        string `json:"inviteId" yaml:"inviteId"`
	PublicJoinKeyId string `json:"publicJoinKeyId" yaml:"publicJoinKeyId"`
	WorkspaceTypeId string `json:"workspaceTypeId" yaml:"workspaceTypeId"`
}

func (x *ClassicAuthDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
