//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetClassicAuthDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
		{
			Name: prefix + "invite-id",
			Type: "string",
		},
		{
			Name: prefix + "public-join-key-id",
			Type: "string",
		},
		{
			Name: prefix + "workspace-type-id",
			Type: "string",
		},
	}
}
func CastClassicAuthDtoFromCli(c emigo.CliCastable) ClassicAuthDto {
	data := ClassicAuthDto{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("invite-id") {
		data.InviteId = c.String("invite-id")
	}
	if c.IsSet("public-join-key-id") {
		data.PublicJoinKeyId = c.String("public-join-key-id")
	}
	if c.IsSet("workspace-type-id") {
		data.WorkspaceTypeId = c.String("workspace-type-id")
	}
	return data
}
