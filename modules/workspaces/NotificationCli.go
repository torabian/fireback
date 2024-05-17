package workspaces

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"
)

func GetWorkspaceNotificationConfig(workspaceId string) (*NotificationConfigEntity, *IError) {

	var item NotificationConfigEntity

	err := GetDbRef().Where(RealEscape("workspace_id = ?", workspaceId)).First(&item).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}

	return &item, nil
}

var NotificationModuleAuditCmd cli.Command = cli.Command{

	Name:  "audit",
	Usage: "Runs several tests, and checks if the notification module has been configurated correctly",
	Action: func(c *cli.Context) error {

		query := CommonCliQueryDSLBuilder(c)
		fmt.Println("Workspace:", query.WorkspaceId)

		fmt.Println("1. Check if there is a configuration for email senders, and main email templates")

		config, err := GetWorkspaceNotificationConfig(query.WorkspaceId)
		if err != nil {
			if err.HttpCode == 404 {
				log.Fatalln("You do not have any configuration, create the configuration first, set the mail server, and senders, and get back")
			}
			log.Fatalln(err)
		}

		if config.GeneralEmailProvider == nil {
			log.Fatalln("You need to specify a general email provider. Email provider, is a service, such as sendgrid, smtp, which lets you send emails. Software depends on sending emails for user interactions, its important to configurate it and make sure the emails arrive.")
		}

		return nil
	},
}

func GetEmailSenderAsStringList(items []*EmailSenderEntity) ([]string, error) {

	result := []string{}
	for _, entity := range items {
		result = append(result, entity.UniqueId+" >>> "+*entity.FromEmailAddress+" - "+*entity.FromName)
	}
	return result, nil
}

var EmailProviderTestCmd cli.Command = cli.Command{

	Name:  "test-mail",
	Usage: "Sends a test mail to verify the mail server is working correctly",

	Action: func(c *cli.Context) error {

		query := CommonCliQueryDSLBuilder(c)
		items, count, err := EmailSenderActionQuery(QueryDSL{ItemsPerPage: 20})

		if err != nil {
			log.Fatalln(err.Error())
		}

		senders, err := GetEmailSenderAsStringList(items)

		if err != nil {
			log.Fatalln(err.Error())
		}

		senderId := ""
		if count.TotalItems <= 20 {
			senderId = AskForSelect("Select the sender, which test mail will be sent on their behalf", senders)
			index := strings.Index(senderId, ">>>")
			senderId = strings.Trim(senderId[0:index], " ")
		} else {
			senderId = AskForInput("Too many workspaces, enter the unique id", "")
		}

		if senderId == "" {
			log.Fatalln("A valid sender is required first, create at least a no-reply email address")
		}

		var toName string = AskForInput("Reciepent name: (eg. Ali Torabi)", "Ali Torabi")
		var toEmail string = AskForInput("Reciepent email address:", "ali-torabian@outlook.com")
		var subject string = AskForInput("Subject", "Testing mail server")
		var content string = AskForInput("Content", "This is a test, to see if our mail server is actually working")

		_, err = NotificationTestMailAction(&TestMailDto{
			SenderId: &senderId,
			ToName:   &toName,
			ToEmail:  &toEmail,
			Subject:  &subject,
			Content:  &content,
		}, query)

		return err
	},
}
