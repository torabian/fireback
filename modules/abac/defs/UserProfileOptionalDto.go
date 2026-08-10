package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for userProfileOptionalDto
type UserProfileOptionalDto struct {
	UniqueId  emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	LastName  emigo.Nullable[string] `json:"lastName" yaml:"lastName"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *UserProfileOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserProfileOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "first-name",
			Type: "string?",
		},
		{
			Name: prefix + "last-name",
			Type: "string?",
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
func CastUserProfileOptionalDtoFromCli(c emigo.CliCastable) UserProfileOptionalDto {
	data := UserProfileOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
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
