package messaging

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// ModuleSetup registers messaging (email/gsm providers, email senders, and the
// SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider actions) as its own
// fireback.ModuleProvider - split out of abac's NotificationModuleSetup, which now only
// wires notificationConfig itself. GsmSendSms (the "use the workspace's default
// provider" variant) stays wired in abac's WorkspaceModuleSetup, since it depends on
// NotificationConfigEntity, which lives in abac.
func ModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name: "abac",

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

				SendEmailActionGin(g, SendEmailAction)
				SendEmailWithProviderActionGin(g, SendEmailWithProviderAction)
				GsmSendSmsWithProviderActionGin(g, GsmSendSmsWithProviderAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_GSM_PROVIDER_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&EmailProviderEntity{},
			&EmailSenderEntity{},
			&GsmProviderEntity{},
		)
	})

	module.ProvideCliHandlers([]*cli.Command{
		{
			Name:        "messaging",
			Description: "Manage email/gsm providers and named email senders",
			Usage:       "Manage email accounts, email providers, gsm providers and so on",
			Aliases:     []string{"msg"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "language",
					Value: "en",
				},
			},
			Commands: []*cli.Command{
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
				SendEmailActionCliHandler(SendEmailAction),
				SendEmailWithProviderActionCliHandler(SendEmailWithProviderAction),
				GsmSendSmsWithProviderActionCliHandler(GsmSendSmsWithProviderAction),
			},
		},
	})

	return module
}
