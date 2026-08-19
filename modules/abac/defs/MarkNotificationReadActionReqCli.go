package abacdefs

import "github.com/torabian/emi/emigo"

func GetMarkNotificationReadActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "The notification uniqueId to mark as read.",
		},
	}
}
func CastMarkNotificationReadActionReqFromCli(c emigo.CliCastable) MarkNotificationReadActionReq {
	data := MarkNotificationReadActionReq{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	return data
}
