package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func NotificationModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name: "abac",

		// gsmProvider moved from AbacModule3.yml's old entities: section to
		// Abac.emi.yml, so it's wired directly here now, the same way
		// FirebackModuleSetup wires Capability* actions.
		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *fireback.FirebackApp) error{
			func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
				GsmProviderBrowseActionGin(g, GsmProviderBrowseAction)
				GsmProviderGetActionGin(g, GsmProviderGetAction)
				GsmProviderCreateActionGin(g, GsmProviderCreateAction)
				GsmProviderUpdateActionGin(g, GsmProviderUpdateAction)
				GsmProviderAwareDeletePreviewActionGin(g, GsmProviderAwareDeletePreviewAction)
				GsmProviderAwareDeleteActionGin(g, GsmProviderAwareDeleteAction)

				EmailProviderBrowseActionGin(g, EmailProviderBrowseAction)
				EmailProviderGetActionGin(g, EmailProviderGetAction)
				EmailProviderCreateActionGin(g, EmailProviderCreateAction)
				EmailProviderUpdateActionGin(g, EmailProviderUpdateAction)
				EmailProviderAwareDeletePreviewActionGin(g, EmailProviderAwareDeletePreviewAction)
				EmailProviderAwareDeleteActionGin(g, EmailProviderAwareDeleteAction)

				EmailSenderBrowseActionGin(g, EmailSenderBrowseAction)
				EmailSenderGetActionGin(g, EmailSenderGetAction)
				EmailSenderCreateActionGin(g, EmailSenderCreateAction)
				EmailSenderUpdateActionGin(g, EmailSenderUpdateAction)
				EmailSenderAwareDeletePreviewActionGin(g, EmailSenderAwareDeletePreviewAction)
				EmailSenderAwareDeleteActionGin(g, EmailSenderAwareDeleteAction)

				NotificationConfigBrowseActionGin(g, NotificationConfigBrowseAction)
				NotificationConfigGetActionGin(g, NotificationConfigGetAction)
				NotificationConfigCreateActionGin(g, NotificationConfigCreateAction)
				NotificationConfigUpdateActionGin(g, NotificationConfigUpdateAction)
				NotificationConfigAwareDeletePreviewActionGin(g, NotificationConfigAwareDeletePreviewAction)
				NotificationConfigAwareDeleteActionGin(g, NotificationConfigAwareDeleteAction)

				AppendNotificationConfigRouter(g)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&EmailProviderEntity{},
			&EmailSenderEntity{},
			&NotificationConfigEntity{},
		)
	})

	module.ProvideCliHandlers([]*cli.Command{
		{
			Name:        "notification",
			Description: "Manage the notification system, emails, text messages, templates and so on",
			Usage:       "Manage email accounts, templates, email providers and so on",
			Aliases:     []string{"nt"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "language",
					Value: "en",
				},
			},
			Commands: []*cli.Command{
				&NotificationModuleAuditCmd,
				&EmailProviderTestCmd,
				EmailProviderBrowseActionCliHandler(EmailProviderBrowseAction),
				EmailProviderGetActionCliHandler(EmailProviderGetAction),
				EmailProviderCreateActionCliHandler(EmailProviderCreateAction),
				EmailProviderUpdateActionCliHandler(EmailProviderUpdateAction),
				EmailProviderAwareDeletePreviewActionCliHandler(EmailProviderAwareDeletePreviewAction),
				EmailProviderAwareDeleteActionCliHandler(EmailProviderAwareDeleteAction),
				EmailSenderBrowseActionCliHandler(EmailSenderBrowseAction),
				EmailSenderGetActionCliHandler(EmailSenderGetAction),
				EmailSenderCreateActionCliHandler(EmailSenderCreateAction),
				EmailSenderUpdateActionCliHandler(EmailSenderUpdateAction),
				EmailSenderAwareDeletePreviewActionCliHandler(EmailSenderAwareDeletePreviewAction),
				EmailSenderAwareDeleteActionCliHandler(EmailSenderAwareDeleteAction),
				&GsmProviderTestCmd,
				GsmProviderBrowseActionCliHandler(GsmProviderBrowseAction),
				GsmProviderGetActionCliHandler(GsmProviderGetAction),
				GsmProviderCreateActionCliHandler(GsmProviderCreateAction),
				GsmProviderUpdateActionCliHandler(GsmProviderUpdateAction),
				GsmProviderAwareDeletePreviewActionCliHandler(GsmProviderAwareDeletePreviewAction),
				GsmProviderAwareDeleteActionCliHandler(GsmProviderAwareDeleteAction),
				&NotificationConfigTestCmd,
				NotificationConfigBrowseActionCliHandler(NotificationConfigBrowseAction),
				NotificationConfigGetActionCliHandler(NotificationConfigGetAction),
				NotificationConfigCreateActionCliHandler(NotificationConfigCreateAction),
				NotificationConfigUpdateActionCliHandler(NotificationConfigUpdateAction),
				NotificationConfigAwareDeletePreviewActionCliHandler(NotificationConfigAwareDeletePreviewAction),
				NotificationConfigAwareDeleteActionCliHandler(NotificationConfigAwareDeleteAction),
			},
		},
	})

	return module
}
