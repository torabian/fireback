//go:build !wasm

package interfacetoolsdefs

import "github.com/torabian/emi/emigo"

func GetTableViewSizingAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastTableViewSizingAwareDeleteActionReqFromCli(c emigo.CliCastable) TableViewSizingAwareDeleteActionReq {
	data := TableViewSizingAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
