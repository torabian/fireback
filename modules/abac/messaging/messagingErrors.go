package messaging

import "github.com/torabian/fireback/modules/fireback"

// SmsNotSent/InvokeUrlMissing mirror the two abac error messages
// (AbacMessages.SmsNotSent/InvokeUrlMissing) the old GsmProviderActions.go raised - copied
// here (not imported from abac) since messaging can't import abac without creating an
// import cycle (see GenericEntityActions.go's comment).
var messagingMessages = struct {
	SmsNotSent      fireback.ErrorItem
	InvokeUrlMissing fireback.ErrorItem
}{
	SmsNotSent: fireback.ErrorItem{
		"$":  "SmsNotSent",
		"en": "Sending text message has failed.",
	},
	InvokeUrlMissing: fireback.ErrorItem{
		"$":  "InvokeUrlMissing",
		"en": "Invoking url is missing",
	},
}
