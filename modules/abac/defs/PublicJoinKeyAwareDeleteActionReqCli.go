//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPublicJoinKeyAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPublicJoinKeyAwareDeleteActionReqFromCli(c emigo.CliCastable) PublicJoinKeyAwareDeleteActionReq {
	data := PublicJoinKeyAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
