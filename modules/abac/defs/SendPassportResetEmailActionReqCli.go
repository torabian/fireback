//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetSendPassportResetEmailActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "The passport uniqueId to send the reset email/SMS to.",
		},
	}
}
func CastSendPassportResetEmailActionReqFromCli(c emigo.CliCastable) SendPassportResetEmailActionReq {
	data := SendPassportResetEmailActionReq{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	return data
}
