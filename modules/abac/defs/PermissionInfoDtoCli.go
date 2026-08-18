//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPermissionInfoDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "complete-key",
			Type: "string",
		},
	}
}
func CastPermissionInfoDtoFromCli(c emigo.CliCastable) PermissionInfoDto {
	data := PermissionInfoDto{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("complete-key") {
		data.CompleteKey = c.String("complete-key")
	}
	return data
}
