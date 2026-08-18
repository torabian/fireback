//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetPendingWorkspaceInviteDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string",
		},
		{
			Name: prefix + "workspace-name",
			Type: "string",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string?",
			Description: "The unique-id of the role which invitee will get if they accept the request.",
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
func CastPendingWorkspaceInviteDtoFromCli(c emigo.CliCastable) PendingWorkspaceInviteDto {
	data := PendingWorkspaceInviteDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	if c.IsSet("workspace-name") {
		data.WorkspaceName = c.String("workspace-name")
	}
	if c.IsSet("role-id") {
		emigo.ParseNullable(c.String("role-id"), &data.RoleId)
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
