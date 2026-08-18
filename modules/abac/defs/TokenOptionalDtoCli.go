//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetTokenOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user this token belongs to.",
		},
		{
			Name: prefix + "token",
			Type: "string?",
		},
		{
			Name: prefix + "valid-until",
			Type: "complex",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
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
func CastTokenOptionalDtoFromCli(c emigo.CliCastable) TokenOptionalDto {
	data := TokenOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("token") {
		emigo.ParseNullable(c.String("token"), &data.Token)
	}
	if c.IsSet("valid-until") {
		if u, ok := any(&data.ValidUntil).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("valid-until")))
		}
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
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
