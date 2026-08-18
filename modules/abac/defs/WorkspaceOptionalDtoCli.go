//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetWorkspaceOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "description",
			Type: "string?",
		},
		{
			Name: prefix + "name",
			Type: "string?",
		},
		{
			Name:        prefix + "type-id",
			Type:        "string?",
			Description: "The unique-id of the workspace type which defines this workspace's role.",
		},
		{
			Name:        prefix + "parent-id",
			Type:        "string?",
			Description: "The unique-id of the parent workspace, for nested/tree workspaces.",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
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
func CastWorkspaceOptionalDtoFromCli(c emigo.CliCastable) WorkspaceOptionalDto {
	data := WorkspaceOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("description") {
		emigo.ParseNullable(c.String("description"), &data.Description)
	}
	if c.IsSet("name") {
		emigo.ParseNullable(c.String("name"), &data.Name)
	}
	if c.IsSet("type-id") {
		emigo.ParseNullable(c.String("type-id"), &data.TypeId)
	}
	if c.IsSet("parent-id") {
		emigo.ParseNullable(c.String("parent-id"), &data.ParentId)
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
