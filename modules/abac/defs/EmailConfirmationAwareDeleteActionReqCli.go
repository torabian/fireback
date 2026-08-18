//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetEmailConfirmationAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastEmailConfirmationAwareDeleteActionReqFromCli(c emigo.CliCastable) EmailConfirmationAwareDeleteActionReq {
	data := EmailConfirmationAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
