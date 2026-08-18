//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetWorkspaceAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastWorkspaceAwareDeleteActionReqFromCli(c emigo.CliCastable) WorkspaceAwareDeleteActionReq {
	data := WorkspaceAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
