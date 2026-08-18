//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPreferenceAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPreferenceAwareDeleteActionReqFromCli(c emigo.CliCastable) PreferenceAwareDeleteActionReq {
	data := PreferenceAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
