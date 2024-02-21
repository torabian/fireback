package workspaces

import (
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func EmailProviderActionCreate(
	dto *EmailProviderEntity, query QueryDSL,
) (*EmailProviderEntity, *IError) {
	return EmailProviderActionCreateFn(dto, query)
}

func EmailProviderActionUpdate(
	query QueryDSL,
	fields *EmailProviderEntity,
) (*EmailProviderEntity, *IError) {
	return EmailProviderActionUpdateFn(query, fields)
}

type EmailSenderCategory string

const (
	GENERAL_SENDER EmailSenderCategory = "GENERAL_SENDER"
)

func SendEmailUsingNotificationConfig(content *EmailMessageContent, sender EmailSenderCategory) (*SendEmailWithProviderActionResDto, *IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		return nil, err
	}

	provider := config.GeneralEmailProvider

	/// I was working here. Now I need to read config.ForgetPasswordSender
	if provider == nil {
		return nil, CreateIErrorString(WorkspacesMessageCode.EmailConfigurationIsNotAvailable, []string{}, 403)
	} else {

		// @todo: Give the option to set custom senders everywhere
		if config.AccountCenterEmailSender != nil {
			content.FromEmail = *config.AccountCenterEmailSender.FromEmailAddress
			content.FromName = *config.AccountCenterEmailSender.FromName
		}

		if err := SendMail(*content, provider); err != nil {
			return nil, CastToIError(err)
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

func getCurrentNotificationConfiguration(query QueryDSL) (*NotificationConfigEntity, error) {
	query.Deep = true
	items, _, err := NotificationConfigActionQuery(query)

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
	// At this moment we send using the sendgrid
	var config = GetAppConfig()

	fmt.Println("+++ New email")
	fmt.Println(message)
	fmt.Println("+++ End of email")

	if os.Getenv("ENV") != "production" && config.Development.ForwardMails != "" {
		fmt.Println("Emails are not going to be sent, forwarding instead to " + config.Development.ForwardMails)
		message.Content = "This mail was forwarded for development purposes, it had to be sent to " + message.ToEmail + " named " + message.ToName + "<br>------<br>" + message.Content
		message.ToEmail = config.Development.ForwardMails
		message.ToName = config.Development.ForwardMailsName
	}

	if provider == nil {
		return errors.New("GENERAL_EMAIL_PROVIDER_IS_NEEDED")
	}

	if *provider.Type == EmailProviderType.Sendgrid {
		res, err := SendMailViaSendGrid(message, *provider.ApiKey)

		if res != nil {
			fmt.Println(res.Body)
		}

		return err

	} else if *provider.Type == EmailProviderType.Terminal {

		fmt.Println(message.Json())

		return nil

	} else {
		return errors.New("EMAIL_PROVIDER_IS_NOT_AVAILABLE")
	}

}
