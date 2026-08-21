package abac

import (
	"log"

	"github.com/torabian/fireback/modules/abac/messaging"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// GsmProvider CRUD, and the raw sms-sending mechanics (GsmSendSMS/SendSms et al.), moved
// to modules/abac/messaging - see messagingdefs.GsmProviderActions/GsmSendSMS. What's left
// here is the one orchestration function that ties NotificationConfigEntity (which stays
// here, in abac) to a specific provider from messagingdefs.

func GsmSendSMSUsingNotificationConfig(message string, recp []string) (*messagingdefs.GsmSendSmsWithProviderActionRes, *fireback.IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})
	if err != nil {
		// If there are no configuration, skip returning error, we use some terminal stuff for development.
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	// config is nil whenever NotificationConfigActionGetOneByWorkspace 404'd above (no
	// row configured yet, e.g. a fresh install nobody has set GSM sending up on) -
	// config.GeneralGsmProviderId.Get() below would panic with a nil pointer
	// dereference on a nil *NotificationConfigEntity receiver if checked first, which
	// crashed the whole request on every send attempt until any admin configured a
	// provider.
	if config == nil {
		log.Default().Println("There is no gsm configuration unfortunately. We are printing the sms to the terminal for the sake of development.")
		log.Default().Println(message, recp)

		return &messagingdefs.GsmSendSmsWithProviderActionRes{QueueId: "print-to-terminal"}, nil
	}

	generalGsmProviderId, hasProvider := config.GeneralGsmProviderId.Get()
	if !hasProvider || *generalGsmProviderId == "" {
		log.Default().Println("There is no gsm configuration unfortunately. We are printing the sms to the terminal for the sake of development.")
		log.Default().Println(message, recp)

		terminalQueue := "print-to-terminal"
		return &messagingdefs.GsmSendSmsWithProviderActionRes{QueueId: terminalQueue}, nil
	}

	return messaging.GsmSendSMS(*generalGsmProviderId, message, recp)
}
