//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetSetPassportPasswordActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "The passport uniqueId whose password will be replaced.",
		},
		{
			Name:        prefix + "password",
			Type:        "string",
			Description: "New password meeting the security requirements.",
		},
	}
}
func CastSetPassportPasswordActionReqFromCli(c emigo.CliCastable) SetPassportPasswordActionReq {
	data := SetPassportPasswordActionReq{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}
