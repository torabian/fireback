//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetRoleAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastRoleAwareDeleteActionReqFromCli(c emigo.CliCastable) RoleAwareDeleteActionReq {
	data := RoleAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
