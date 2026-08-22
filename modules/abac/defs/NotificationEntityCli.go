//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetNotificationEntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UniqueId of the user this notification was sent to.",
		},
		{
			Name:        prefix + "sender-id",
			Type:        "string?",
			Description: "UniqueId of the (root) user who sent this notification.",
		},
		{
			Name:        prefix + "title",
			Type:        "string",
			Description: "Short notification title.",
		},
		{
			Name:        prefix + "body",
			Type:        "string",
			Description: "Notification message body.",
		},
		{
			Name:        prefix + "is-read",
			Type:        "bool",
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
func CastNotificationEntityFromCli(c emigo.CliCastable) NotificationEntity {
	data := NotificationEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("sender-id") {
		emigo.ParseNullable(c.String("sender-id"), &data.SenderId)
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	if c.IsSet("is-read") {
		data.IsRead = bool(c.Bool("is-read"))
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
