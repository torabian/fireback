//go:build !wasm

package abac

import "github.com/torabian/emi/emigo"

func GetSignoutActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "clear",
			Type: "bool?",
		},
	}
}
func CastSignoutActionReqFromCli(c emigo.CliCastable) SignoutActionReq {
	data := SignoutActionReq{}
	if c.IsSet("clear") {
		emigo.ParseNullable(c.String("clear"), &data.Clear)
	}
	return data
}
