package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for regionalContentOptionalDto
type RegionalContentOptionalDto struct {
	UniqueId   emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Content    emigo.Nullable[string] `json:"content" validate:"required" yaml:"content"`
	Region     emigo.Nullable[string] `json:"region" validate:"required" yaml:"region"`
	Title      emigo.Nullable[string] `json:"title" yaml:"title"`
	LanguageId emigo.Nullable[string] `json:"languageId" validate:"required" yaml:"languageId"`
	KeyGroup   emigo.Nullable[string] `json:"keyGroup" validate:"required" yaml:"keyGroup"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RegionalContentOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetRegionalContentOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "content",
			Type: "string?",
		},
		{
			Name: prefix + "region",
			Type: "string?",
		},
		{
			Name: prefix + "title",
			Type: "string?",
		},
		{
			Name: prefix + "language-id",
			Type: "string?",
		},
		{
			Name: prefix + "key-group",
			Type: "enum?",
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
func CastRegionalContentOptionalDtoFromCli(c emigo.CliCastable) RegionalContentOptionalDto {
	data := RegionalContentOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("content") {
		emigo.ParseNullable(c.String("content"), &data.Content)
	}
	if c.IsSet("region") {
		emigo.ParseNullable(c.String("region"), &data.Region)
	}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("language-id") {
		emigo.ParseNullable(c.String("language-id"), &data.LanguageId)
	}
	if c.IsSet("key-group") {
		emigo.ParseNullable(c.String("key-group"), &data.KeyGroup)
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
