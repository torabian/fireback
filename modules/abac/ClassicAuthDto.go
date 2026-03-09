package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetClassicAuthDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
		{
			Name: prefix + "invite-id",
			Type: "string",
		},
		{
			Name: prefix + "public-join-key-id",
			Type: "string",
		},
		{
			Name: prefix + "workspace-type-id",
			Type: "string",
		},
	}
}
func CastClassicAuthDtoFromCli(c emigo.CliCastable) ClassicAuthDto {
	data := ClassicAuthDto{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("invite-id") {
		data.InviteId = c.String("invite-id")
	}
	if c.IsSet("public-join-key-id") {
		data.PublicJoinKeyId = c.String("public-join-key-id")
	}
	if c.IsSet("workspace-type-id") {
		data.WorkspaceTypeId = c.String("workspace-type-id")
	}
	return data
}

// The base class definition for classicAuthDto
type ClassicAuthDto struct {
	Value           string `json:"value" yaml:"value" validate:"required"`
	Password        string `yaml:"password" validate:"required" json:"password"`
	FirstName       string `validate:"required" json:"firstName" yaml:"firstName"`
	LastName        string `validate:"required" json:"lastName" yaml:"lastName"`
	InviteId        string `json:"inviteId" yaml:"inviteId"`
	PublicJoinKeyId string `json:"publicJoinKeyId" yaml:"publicJoinKeyId"`
	WorkspaceTypeId string `yaml:"workspaceTypeId" json:"workspaceTypeId"`
}

func (x *ClassicAuthDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
