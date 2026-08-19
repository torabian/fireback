package abacdefs

import "github.com/torabian/emi/emigo"

func GetSendNotificationActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-ids",
			Type:        "slice",
			Description: "UniqueIds of the existing users to notify.",
		},
		{
			Name:        prefix + "title",
			Type:        "string",
			Description: "Short notification title.",
		},
		{
			Name:        prefix + "body",
			Type:        "string",
			Description: "Notification message body.",
		},
	}
}
func CastSendNotificationActionReqFromCli(c emigo.CliCastable) SendNotificationActionReq {
	data := SendNotificationActionReq{}
	if c.IsSet("user-ids") {
		emigo.InflatePossibleSlice(c.String("user-ids"), &data.UserIds)
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}
