package abac

import (
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func NotificationModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name: "abac",
	}

	module.ProvidePermissionHandler(
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
	)

	module.Actions = [][]fireback.Module3Action{
		GetEmailProviderModule3Actions(),
		GetEmailSenderModule3Actions(),
		GetNotificationConfigModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&EmailProviderEntity{},
			&EmailSenderEntity{},
			&NotificationConfigEntity{},
		)
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
