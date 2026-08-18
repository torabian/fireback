//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPublicAuthenticationAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPublicAuthenticationAwareDeleteActionReqFromCli(c emigo.CliCastable) PublicAuthenticationAwareDeleteActionReq {
	data := PublicAuthenticationAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
