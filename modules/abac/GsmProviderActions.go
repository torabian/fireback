package abac

import (
	"log"

	"github.com/torabian/fireback/modules/abac/messaging"
	"github.com/torabian/fireback/modules/fireback"
)

// GsmProvider CRUD, and the raw sms-sending mechanics (GsmSendSMS/SendSms et al.), moved
// to modules/abac/messaging - see messaging.GsmProviderActions/GsmSendSMS. What's left
// here is the one orchestration function that ties NotificationConfigEntity (which stays
// here, in abac) to a specific provider from messaging.

func GsmSendSMSUsingNotificationConfig(message string, recp []string) (*messaging.GsmSendSmsWithProviderActionRes, *fireback.IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})
	if err != nil {
		// If there are no configuration, skip returning error, we use some terminal stuff for development.
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	generalGsmProviderId, hasProvider := config.GeneralGsmProviderId.Get()
	if config == nil || !hasProvider || *generalGsmProviderId == "" {
		log.Default().Println("There is no gsm configuration unfortunately. We are printing the sms to the terminal for the sake of development.")
		log.Default().Println(message, recp)

		terminalQueue := "print-to-terminal"
		return &messaging.GsmSendSmsWithProviderActionRes{QueueId: terminalQueue}, nil
	}

	return messaging.GsmSendSMS(*generalGsmProviderId, message, recp)
}
