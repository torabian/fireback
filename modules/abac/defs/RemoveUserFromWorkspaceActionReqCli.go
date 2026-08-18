//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetRemoveUserFromWorkspaceActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UniqueId of the user to remove.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string",
			Description: "UniqueId of the workspace to remove the user from.",
		},
	}
}
func CastRemoveUserFromWorkspaceActionReqFromCli(c emigo.CliCastable) RemoveUserFromWorkspaceActionReq {
	data := RemoveUserFromWorkspaceActionReq{}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	return data
}
