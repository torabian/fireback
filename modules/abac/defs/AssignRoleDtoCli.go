//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetAssignRoleDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "role-id",
			Type: "string",
		},
		{
			Name: prefix + "user-id",
			Type: "string",
		},
		{
			Name: prefix + "visibility",
			Type: "string",
		},
		{
			Name: prefix + "updated",
			Type: "int64",
		},
		{
			Name: prefix + "created",
			Type: "int64",
		},
	}
}
func CastAssignRoleDtoFromCli(c emigo.CliCastable) AssignRoleDto {
	data := AssignRoleDto{}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("visibility") {
		data.Visibility = c.String("visibility")
	}
	if c.IsSet("updated") {
		data.Updated = int64(c.Int64("updated"))
	}
	if c.IsSet("created") {
		data.Created = int64(c.Int64("created"))
	}
	return data
}
