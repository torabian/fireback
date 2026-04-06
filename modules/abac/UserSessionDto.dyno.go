package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetUserSessionDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "passport",
			Type: "one?",
		},
		{
			Name: prefix + "token",
			Type: "string",
		},
		{
			Name: prefix + "exchange-key",
			Type: "string",
		},
		{
			Name: prefix + "user-workspaces",
			Type: "collection",
		},
		{
			Name: prefix + "user",
			Type: "one?",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
		},
	}
}
func CastUserSessionDtoFromCli(c emigo.CliCastable) UserSessionDto {
	data := UserSessionDto{}
	if c.IsSet("passport") {
		emigo.ParseNullable(c.String("passport"), &data.Passport)
	}
	if c.IsSet("token") {
		data.Token = c.String("token")
	}
	if c.IsSet("exchange-key") {
		data.ExchangeKey = c.String("exchange-key")
	}
	if c.IsSet("user") {
		emigo.ParseNullable(c.String("user"), &data.User)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	return data
}

// The base class definition for userSessionDto
type UserSessionDto struct {
	Passport       emigo.Nullable[PassportEntity] `json:"passport" yaml:"passport"`
	Token          string                         `json:"token" yaml:"token"`
	ExchangeKey    string                         `json:"exchangeKey" yaml:"exchangeKey"`
	UserWorkspaces []UserWorkspaceEntity          `json:"userWorkspaces" yaml:"userWorkspaces"`
	User           emigo.Nullable[UserEntity]     `json:"user" yaml:"user"`
	UserId         emigo.Nullable[string]         `json:"userId" yaml:"userId"`
}

func (x *UserSessionDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
