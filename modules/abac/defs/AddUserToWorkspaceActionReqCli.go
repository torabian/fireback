//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetAddUserToWorkspaceActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UniqueId of the existing user to add to the workspace.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string",
			Description: "UniqueId of the workspace to add the user to.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string",
			Description: "UniqueId of the role (must belong to workspaceId) to assign to the user.",
		},
	}
}
func CastAddUserToWorkspaceActionReqFromCli(c emigo.CliCastable) AddUserToWorkspaceActionReq {
	data := AddUserToWorkspaceActionReq{}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	return data
}
