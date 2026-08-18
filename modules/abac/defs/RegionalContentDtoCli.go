//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetRegionalContentDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "content",
			Type:        "string",
			Description: "The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.",
		},
		{
			Name:        prefix + "region",
			Type:        "string",
			Description: "Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.",
		},
		{
			Name:        prefix + "title",
			Type:        "string",
			Description: "Optional subject line - only used for email-type content.",
		},
		{
			Name:        prefix + "language-id",
			Type:        "string",
			Description: "Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.",
		},
		{
			Name:        prefix + "key-group",
			Type:        "enum",
			Description: "Which kind of message this content is used for.",
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
