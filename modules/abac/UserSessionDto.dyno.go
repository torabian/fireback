package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

// The base class definition for userSessionDto
type UserSessionDto struct {
	Passport       emigo.OneNullable[PassportEntity]     `json:"passport" yaml:"passport"`
	Token          string                                `json:"token" yaml:"token"`
	ExchangeKey    string                                `json:"exchangeKey" yaml:"exchangeKey"`
	UserWorkspaces emigo.Collection[UserWorkspaceEntity] `json:"userWorkspaces" yaml:"userWorkspaces"`
	User           emigo.OneNullable[UserEntity]         `json:"user" yaml:"user"`
	UserId         emigo.Nullable[string]                `json:"userId" yaml:"userId"`
}

func (x *UserSessionDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
