//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPendingWorkspaceInviteAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPendingWorkspaceInviteAwareDeleteActionReqFromCli(c emigo.CliCastable) PendingWorkspaceInviteAwareDeleteActionReq {
	data := PendingWorkspaceInviteAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
