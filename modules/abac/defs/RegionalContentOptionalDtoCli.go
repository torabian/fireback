//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetRegionalContentOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "content",
			Type:        "string?",
			Description: "The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.",
		},
		{
			Name:        prefix + "region",
			Type:        "string?",
			Description: "Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.",
		},
		{
			Name:        prefix + "title",
			Type:        "string?",
			Description: "Optional subject line - only used for email-type content.",
		},
		{
			Name:        prefix + "language-id",
			Type:        "string?",
			Description: "Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.",
		},
		{
			Name:        prefix + "key-group",
			Type:        "enum?",
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
