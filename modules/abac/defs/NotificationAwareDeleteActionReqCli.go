package abacdefs

import "github.com/torabian/emi/emigo"

func GetNotificationAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastNotificationAwareDeleteActionReqFromCli(c emigo.CliCastable) NotificationAwareDeleteActionReq {
	data := NotificationAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}
