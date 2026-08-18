//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetCapabilityAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastCapabilityAwareDeleteActionReqFromCli(c emigo.CliCastable) CapabilityAwareDeleteActionReq {
	data := CapabilityAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
