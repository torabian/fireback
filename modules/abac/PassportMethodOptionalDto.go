package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for passportMethodOptionalDto
type PassportMethodOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Type     emigo.Nullable[string] `json:"type" validate:"oneof=email phone google facebook,required" yaml:"type"`
	// The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
	Region emigo.Nullable[string] `json:"region" validate:"required,oneof=global" yaml:"region"`
	// Client key for those methods such as 'google' which require oauth client key
	ClientKey emigo.Nullable[string] `json:"clientKey" yaml:"clientKey"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PassportMethodOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPassportMethodOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "type",
			Type: "enum?",
		},
		{
			Name:        prefix + "region",
			Type:        "enum?",
			Description: "The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.",
		},
		{
			Name:        prefix + "client-key",
			Type:        "string?",
			Description: "Client key for those methods such as 'google' which require oauth client key",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which created/owns the record.",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
	}
}
func CastPassportMethodOptionalDtoFromCli(c emigo.CliCastable) PassportMethodOptionalDto {
	data := PassportMethodOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("type") {
		emigo.ParseNullable(c.String("type"), &data.Type)
	}
	if c.IsSet("region") {
		emigo.ParseNullable(c.String("region"), &data.Region)
	}
	if c.IsSet("client-key") {
		emigo.ParseNullable(c.String("client-key"), &data.ClientKey)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("updated-at") {
		if u, ok := any(&data.UpdatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("updated-at")))
		}
	}
	return data
}
