//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetNotificationOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "UniqueId of the user this notification was sent to.",
		},
		{
			Name:        prefix + "sender-id",
			Type:        "string?",
			Description: "UniqueId of the (root) user who sent this notification.",
		},
		{
			Name:        prefix + "title",
			Type:        "string?",
			Description: "Short notification title.",
		},
		{
			Name:        prefix + "body",
			Type:        "string?",
			Description: "Notification message body.",
		},
		{
			Name:        prefix + "is-read",
			Type:        "bool?",
			Description: "Whether the recipient has read this notification yet.",
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
func CastNotificationOptionalDtoFromCli(c emigo.CliCastable) NotificationOptionalDto {
	data := NotificationOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("sender-id") {
		emigo.ParseNullable(c.String("sender-id"), &data.SenderId)
	}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("body") {
		emigo.ParseNullable(c.String("body"), &data.Body)
	}
	if c.IsSet("is-read") {
		emigo.ParseNullable(c.String("is-read"), &data.IsRead)
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
