//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetUserAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastUserAwareDeleteActionReqFromCli(c emigo.CliCastable) UserAwareDeleteActionReq {
	data := UserAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
