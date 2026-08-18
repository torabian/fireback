//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetWorkspaceConfigAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastWorkspaceConfigAwareDeleteActionReqFromCli(c emigo.CliCastable) WorkspaceConfigAwareDeleteActionReq {
	data := WorkspaceConfigAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
