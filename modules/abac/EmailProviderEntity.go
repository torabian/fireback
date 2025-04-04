package abac

import (
	"errors"
	"fmt"
	"log"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/torabian/fireback/modules/workspaces"
)

type EmailSenderCategory string

const (
	GENERAL_SENDER EmailSenderCategory = "GENERAL_SENDER"
)

func SendEmailUsingNotificationConfig(content *EmailMessageContent, sender EmailSenderCategory) (*SendEmailWithProviderActionResDto, *workspaces.IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(workspaces.QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		// If there are no configuration, skip returning error, we use some terminal stuff for development.
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	/// I was working here. Now I need to read config.ForgetPasswordSender
	if config.GeneralEmailProvider == nil {
		log.Default().Println("There are no email providers configured, we are printing the email into the console assuming this is development.")
		log.Default().Println(content.Json())

		QueueId := "printed-to-terminal"
		return &SendEmailWithProviderActionResDto{QueueId: QueueId}, nil
	} else {

		// @todo: Give the option to set custom senders everywhere
		if config.AccountCenterEmailSender != nil {
			content.FromEmail = config.AccountCenterEmailSender.FromEmailAddress
			content.FromName = config.AccountCenterEmailSender.FromName
		}

		if err := SendMail(*content, config.GeneralEmailProvider); err != nil {
			return nil, workspaces.CastToIError(err)
		} else {
			return &SendEmailWithProviderActionResDto{}, nil
		}
	}
}

func SendMailViaSendGrid(message EmailMessageContent, apiKey string) (*rest.Response, error) {

	from := mail.NewEmail(message.FromName, message.FromEmail)
	to := mail.NewEmail(message.ToName, message.ToEmail)

	message2 := mail.NewSingleEmail(from, message.Subject, to, message.Content, message.Content)

	client := sendgrid.NewSendClient(apiKey)
	res, err := client.Send(message2)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func getCurrentNotificationConfiguration(query workspaces.QueryDSL) (*NotificationConfigEntity, error) {
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

func SendMail(message EmailMessageContent, provider *EmailProviderEntity) error {

	if provider == nil {
		return errors.New("GENERAL_EMAIL_PROVIDER_IS_NEEDED")
	}

	if provider.Type == EmailProviderType.Sendgrid {
		res, err := SendMailViaSendGrid(message, provider.ApiKey)

		if res != nil {
			fmt.Println(res.Body)
		}

		return err

	} else if provider.Type == EmailProviderType.Terminal {

		log.Default().Println(message.Json())

		return nil

	} else {
		return errors.New("EMAIL_PROVIDER_IS_NOT_AVAILABLE")
	}

}
