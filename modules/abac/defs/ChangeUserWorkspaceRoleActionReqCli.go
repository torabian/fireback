//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetChangeUserWorkspaceRoleActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string",
			Description: "UniqueId of the workspace the membership belongs to.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string",
			Description: "UniqueId of the new role (must belong to workspaceId) to assign to the user.",
		},
	}
}
func CastChangeUserWorkspaceRoleActionReqFromCli(c emigo.CliCastable) ChangeUserWorkspaceRoleActionReq {
	data := ChangeUserWorkspaceRoleActionReq{}
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
