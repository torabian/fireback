//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetDisablePassportTotpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "The passport uniqueId to disable TOTP for.",
		},
	}
}
func CastDisablePassportTotpActionReqFromCli(c emigo.CliCastable) DisablePassportTotpActionReq {
	data := DisablePassportTotpActionReq{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	return data
}
