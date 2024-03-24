package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func NotificationModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name: "workspaces",
	}

	module.ProvidePermissionHandler(
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
	)

	module.Actions = [][]Module2Action{
		GetEmailProviderModule2Actions(),
		GetEmailSenderModule2Actions(),
		GetNotificationConfigModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(&EmailProviderEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&EmailSenderEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&NotificationConfigEntity{}); err != nil {
			fmt.Println(err.Error())
		}
	})

	module.ProvideCliHandlers([]cli.Command{
		{
			Name:        "notification",
			Description: "Manage the notification system, emails, text messages, templates and so on",
			Usage:       "Manage email accounts, templates, email providers and so on",
			ShortName:   "nt",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "language",
					Value: "en",
				},
			},
			Subcommands: cli.Commands{
				NotificationModuleAuditCmd,
				EmailProviderTestCmd,
				EmailProviderCliFn(),
				EmailSenderCliFn(),
				GsmProviderCliFn(),
				NotificationConfigCliFn(),
			},
		},
	})

	return module
}
