//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetQueryWorkspaceRolesActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "workspace-id",
			Type:        "string",
			Description: "UniqueId of the workspace whose roles to list.",
		},
	}
}
func CastQueryWorkspaceRolesActionReqFromCli(c emigo.CliCastable) QueryWorkspaceRolesActionReq {
	data := QueryWorkspaceRolesActionReq{}
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	return data
}
