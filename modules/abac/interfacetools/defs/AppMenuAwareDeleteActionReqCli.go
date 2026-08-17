//go:build !wasm

package interfacetoolsdefs

import "github.com/torabian/emi/emigo"

func GetAppMenuAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastAppMenuAwareDeleteActionReqFromCli(c emigo.CliCastable) AppMenuAwareDeleteActionReq {
	data := AppMenuAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
