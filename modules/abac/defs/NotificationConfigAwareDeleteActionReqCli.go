//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetNotificationConfigAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastNotificationConfigAwareDeleteActionReqFromCli(c emigo.CliCastable) NotificationConfigAwareDeleteActionReq {
	data := NotificationConfigAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
