//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetWorkspaceRoleAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastWorkspaceRoleAwareDeleteActionReqFromCli(c emigo.CliCastable) WorkspaceRoleAwareDeleteActionReq {
	data := WorkspaceRoleAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
