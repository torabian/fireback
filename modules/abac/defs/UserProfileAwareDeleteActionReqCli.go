//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetUserProfileAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastUserProfileAwareDeleteActionReqFromCli(c emigo.CliCastable) UserProfileAwareDeleteActionReq {
	data := UserProfileAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
