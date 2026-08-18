//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetRoleEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:        prefix + "capabilities-list-id",
			Type:        "complex",
			Description: "The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).",
		},
		{
			Name: prefix + "is-deletable",
			Type: "bool?",
		},
		{
			Name: prefix + "is-updatable",
			Type: "bool?",
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
func CastRoleEntityFromCli(c emigo.CliCastable) RoleEntity {
	data := RoleEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("capabilities-list-id") {
		if u, ok := any(&data.CapabilitiesListId).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("capabilities-list-id")))
		}
	}
	if c.IsSet("is-deletable") {
		emigo.ParseNullable(c.String("is-deletable"), &data.IsDeletable)
	}
	if c.IsSet("is-updatable") {
		emigo.ParseNullable(c.String("is-updatable"), &data.IsUpdatable)
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
