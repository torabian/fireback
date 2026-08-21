package abac

import (
	"errors"
	"log"

	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/messaging"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// EmailProvider/EmailSender CRUD, and the raw email-sending mechanics (SendMail et al.),
// moved to modules/abac/messaging - see messaging.EmailProviderActions/EmailSenderActions
// and messaging.SendMail. What's left here is the one orchestration function that ties
// NotificationConfigEntity (which stays here, in abac) to a specific provider from
// messaging.

type EmailSenderCategory string

const (
	GENERAL_SENDER EmailSenderCategory = "GENERAL_SENDER"
)

func SendEmailUsingNotificationConfig(content *messaging.EmailMessageContent, sender EmailSenderCategory) (*messagingdefs.SendEmailWithProviderActionRes, *fireback.IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		// If there are no configuration, skip returning error, we use some terminal stuff for development.
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	// config is nil whenever NotificationConfigActionGetOneByWorkspace 404'd above (no
	// row configured yet, e.g. a fresh install nobody has set email sending up on) -
	// config.GeneralEmailProviderId.Get() below would panic with a nil pointer
	// dereference on a nil *NotificationConfigEntity receiver otherwise, which crashed
	// the whole request (OTP emails, workspace invites, ...) on every send attempt
	// until any admin configured a provider - same fix as
	// GsmProviderActions.go's GsmSendSMSUsingNotificationConfig.
	printToTerminal := func() (*messagingdefs.SendEmailWithProviderActionRes, *fireback.IError) {
		log.Default().Println("There are no email providers configured, we are printing the email into the console assuming this is development.")
		log.Default().Println(content.Json())
		return &messagingdefs.SendEmailWithProviderActionRes{QueueId: "printed-to-terminal"}, nil
	}

	if config == nil {
		return printToTerminal()
	}

	generalEmailProviderId, hasProvider := config.GeneralEmailProviderId.Get()
	if !hasProvider || *generalEmailProviderId == "" {
		return printToTerminal()
	} else {

		provider, providerErr := messaging.EmailProviderActions.GetOne(fireback.QueryDSL{UniqueId: *generalEmailProviderId})
		if providerErr != nil {
			return nil, providerErr
		}

		// @todo: Give the option to set custom senders everywhere
		if senderId, ok := config.AccountCenterEmailSenderId.Get(); ok && *senderId != "" {
			sender, senderErr := messaging.EmailSenderActions.GetOne(fireback.QueryDSL{UniqueId: *senderId})
			if senderErr == nil && sender != nil {
				content.FromEmail = sender.FromEmailAddress
				content.FromName = sender.FromName
			}
		}

		if err := messaging.SendMail(*content, provider); err != nil {
			return nil, fireback.CastToIError(err)
		} else {
			return &messagingdefs.SendEmailWithProviderActionRes{}, nil
		}
	}
}

func getCurrentNotificationConfiguration(query fireback.QueryDSL) (*abacdefs.NotificationConfigEntity, error) {
	query.Deep = true
	items, _, err := NotificationConfigActions.Query(query)

	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("NO_NOTIFICATION_CONFIGURATION_AVAILABLE")
	}

	// @todo handle the region information based on the query

	return items[0], nil
}
