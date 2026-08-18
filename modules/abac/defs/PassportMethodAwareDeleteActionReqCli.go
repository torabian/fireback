//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPassportMethodAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPassportMethodAwareDeleteActionReqFromCli(c emigo.CliCastable) PassportMethodAwareDeleteActionReq {
	data := PassportMethodAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
