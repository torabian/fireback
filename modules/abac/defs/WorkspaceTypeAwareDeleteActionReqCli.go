//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetWorkspaceTypeAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastWorkspaceTypeAwareDeleteActionReqFromCli(c emigo.CliCastable) WorkspaceTypeAwareDeleteActionReq {
	data := WorkspaceTypeAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
