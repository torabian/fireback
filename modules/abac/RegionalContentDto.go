package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for regionalContentDto
type RegionalContentDto struct {
	UniqueId   emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Content    string                 `json:"content" validate:"required" yaml:"content"`
	Region     string                 `json:"region" validate:"required" yaml:"region"`
	Title      string                 `json:"title" yaml:"title"`
	LanguageId string                 `json:"languageId" validate:"required" yaml:"languageId"`
	KeyGroup   string                 `json:"keyGroup" validate:"required" yaml:"keyGroup"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RegionalContentDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetRegionalContentDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "content",
			Type: "string",
		},
		{
			Name: prefix + "region",
			Type: "string",
		},
		{
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "language-id",
			Type: "string",
		},
		{
			Name: prefix + "key-group",
			Type: "enum",
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
func CastRegionalContentDtoFromCli(c emigo.CliCastable) RegionalContentDto {
	data := RegionalContentDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("content") {
		data.Content = c.String("content")
	}
	if c.IsSet("region") {
		data.Region = c.String("region")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("language-id") {
		data.LanguageId = c.String("language-id")
	}
	if c.IsSet("key-group") {
		data.KeyGroup = c.String("key-group")
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
