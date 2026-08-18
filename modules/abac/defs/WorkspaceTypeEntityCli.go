//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetWorkspaceTypeEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name:        prefix + "slug",
			Type:        "string",
			Description: "Unique, URL-safe identifier for this workspace type. Must start with \"/\", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:\"unique\" here is the DB-level backstop for the uniqueness half of that.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string",
			Description: "The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.",
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
func CastWorkspaceTypeEntityFromCli(c emigo.CliCastable) WorkspaceTypeEntity {
	data := WorkspaceTypeEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("slug") {
		data.Slug = c.String("slug")
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
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
