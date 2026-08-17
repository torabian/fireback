//go:build !wasm

package interfacetoolsdefs

import "github.com/torabian/emi/emigo"

func GetTimezoneGroupAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastTimezoneGroupAwareDeleteActionReqFromCli(c emigo.CliCastable) TimezoneGroupAwareDeleteActionReq {
	data := TimezoneGroupAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
