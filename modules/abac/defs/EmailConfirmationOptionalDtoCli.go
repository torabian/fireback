//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetEmailConfirmationOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user this confirmation belongs to.",
		},
		{
			Name: prefix + "status",
			Type: "string?",
		},
		{
			Name: prefix + "email",
			Type: "string?",
		},
		{
			Name: prefix + "key",
			Type: "string?",
		},
		{
			Name: prefix + "expires-at",
			Type: "string?",
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
func CastEmailConfirmationOptionalDtoFromCli(c emigo.CliCastable) EmailConfirmationOptionalDto {
	data := EmailConfirmationOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("status") {
		emigo.ParseNullable(c.String("status"), &data.Status)
	}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("key") {
		emigo.ParseNullable(c.String("key"), &data.Key)
	}
	if c.IsSet("expires-at") {
		emigo.ParseNullable(c.String("expires-at"), &data.ExpiresAt)
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
