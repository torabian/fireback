//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetUserRoleWorkspaceDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "role-id",
			Type: "string",
		},
		{
			Name: prefix + "capabilities",
			Type: "slice",
		},
	}
}
func CastUserRoleWorkspaceDtoFromCli(c emigo.CliCastable) UserRoleWorkspaceDto {
	data := UserRoleWorkspaceDto{}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	return data
}
