package messaging

import (
	"github.com/gin-gonic/gin"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// ModuleSetup registers messaging (email/gsm providers, email senders, the
// SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider actions, and webPushConfig -
// moved here from modules/fireback, where its CRUD had never actually been implemented
// or wired) as its own application.ModuleProvider - split out of abac's
// NotificationModuleSetup, which now only wires notificationConfig itself. GsmSendSms
// (the "use the workspace's default provider" variant) stays wired in abac's
// WorkspaceModuleSetup, since it depends on NotificationConfigEntity, which lives in
// abac.
func ModuleSetup() *application.ModuleProvider {
	module := &application.ModuleProvider{
		Name: "abac",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				messagingdefs.GsmProviderBrowseActionGin(g, GsmProviderBrowseAction)
				messagingdefs.GsmProviderGetActionGin(g, GsmProviderGetAction)
				messagingdefs.GsmProviderCreateActionGin(g, GsmProviderCreateAction)
				messagingdefs.GsmProviderUpdateActionGin(g, GsmProviderUpdateAction)
				messagingdefs.GsmProviderAwareDeletePreviewActionGin(g, GsmProviderAwareDeletePreviewAction)
				messagingdefs.GsmProviderAwareDeleteActionGin(g, GsmProviderAwareDeleteAction)

				messagingdefs.EmailProviderBrowseActionGin(g, EmailProviderBrowseAction)
				messagingdefs.EmailProviderGetActionGin(g, EmailProviderGetAction)
				messagingdefs.EmailProviderCreateActionGin(g, EmailProviderCreateAction)
				messagingdefs.EmailProviderUpdateActionGin(g, EmailProviderUpdateAction)
				messagingdefs.EmailProviderAwareDeletePreviewActionGin(g, EmailProviderAwareDeletePreviewAction)
				messagingdefs.EmailProviderAwareDeleteActionGin(g, EmailProviderAwareDeleteAction)

				messagingdefs.EmailSenderBrowseActionGin(g, EmailSenderBrowseAction)
				messagingdefs.EmailSenderGetActionGin(g, EmailSenderGetAction)
				messagingdefs.EmailSenderCreateActionGin(g, EmailSenderCreateAction)
				messagingdefs.EmailSenderUpdateActionGin(g, EmailSenderUpdateAction)
				messagingdefs.EmailSenderAwareDeletePreviewActionGin(g, EmailSenderAwareDeletePreviewAction)
				messagingdefs.EmailSenderAwareDeleteActionGin(g, EmailSenderAwareDeleteAction)

				messagingdefs.WebPushConfigBrowseActionGin(g, WebPushConfigBrowseAction)
				messagingdefs.WebPushConfigGetActionGin(g, WebPushConfigGetAction)
				messagingdefs.WebPushConfigCreateActionGin(g, WebPushConfigCreateAction)
				messagingdefs.WebPushConfigUpdateActionGin(g, WebPushConfigUpdateAction)
				messagingdefs.WebPushConfigAwareDeletePreviewActionGin(g, WebPushConfigAwareDeletePreviewAction)
				messagingdefs.WebPushConfigAwareDeleteActionGin(g, WebPushConfigAwareDeleteAction)

				messagingdefs.SendEmailActionGin(g, SendEmailAction)
				messagingdefs.SendEmailWithProviderActionGin(g, SendEmailWithProviderAction)
				messagingdefs.GsmSendSmsWithProviderActionGin(g, GsmSendSmsWithProviderAction)

				messagingdefs.MessagingConfigGetActionGin(g, MessagingConfigGetAction)
				messagingdefs.MessagingConfigUpdateActionGin(g, MessagingConfigUpdateAction)

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
			&messagingdefs.EmailProviderEntity{},
			&messagingdefs.EmailSenderEntity{},
			&messagingdefs.GsmProviderEntity{},
			&messagingdefs.WebPushConfigEntity{},
			&messagingdefs.MessagingConfigEntity{},
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
				&cli.Command{
					Name:        "email",
					Description: `Email providers (sendgrid/mailgun/postmark/resend/smtp/terminal) and named "from" email senders.`,
					Usage:       `Email providers (sendgrid/mailgun/postmark/resend/smtp/terminal) and named "from" email senders.`,
					Commands: []*cli.Command{
						&cli.Command{
							Name:        "provider",
							Aliases:     []string{"ep"},
							Description: `Email server/thirdparty service configuration used to actually send emails.`,
							Usage:       `Email server/thirdparty service configuration used to actually send emails.`,
							Commands: []*cli.Command{
								messagingdefs.EmailProviderBrowseActionCliHandler(EmailProviderBrowseAction),
								messagingdefs.EmailProviderGetActionCliHandler(EmailProviderGetAction),
								messagingdefs.EmailProviderCreateActionCliHandler(EmailProviderCreateAction),
								messagingdefs.EmailProviderUpdateActionCliHandler(EmailProviderUpdateAction),
								messagingdefs.EmailProviderAwareDeletePreviewActionCliHandler(EmailProviderAwareDeletePreviewAction),
								messagingdefs.EmailProviderAwareDeleteActionCliHandler(EmailProviderAwareDeleteAction),
							},
						},
						&cli.Command{
							Name:        "sender",
							Aliases:     []string{"es"},
							Description: `Named "from" identities (name/address, reply-to) an email is sent as.`,
							Usage:       `Named "from" identities (name/address, reply-to) an email is sent as.`,
							Commands: []*cli.Command{
								messagingdefs.EmailSenderBrowseActionCliHandler(EmailSenderBrowseAction),
								messagingdefs.EmailSenderGetActionCliHandler(EmailSenderGetAction),
								messagingdefs.EmailSenderCreateActionCliHandler(EmailSenderCreateAction),
								messagingdefs.EmailSenderUpdateActionCliHandler(EmailSenderUpdateAction),
								messagingdefs.EmailSenderAwareDeletePreviewActionCliHandler(EmailSenderAwareDeletePreviewAction),
								messagingdefs.EmailSenderAwareDeleteActionCliHandler(EmailSenderAwareDeleteAction),
							},
						},
						messagingdefs.SendEmailActionCliHandler(SendEmailAction),
						messagingdefs.SendEmailWithProviderActionCliHandler(SendEmailWithProviderAction),
					},
				},
				&cli.Command{
					Name:        "gsm",
					Aliases:     []string{"sms"},
					Description: `SMS/GSM gateway configuration (url/terminal/mediana) and sending.`,
					Usage:       `SMS/GSM gateway configuration (url/terminal/mediana) and sending.`,
					Commands: []*cli.Command{
						&cli.Command{
							Name:        "provider",
							Aliases:     []string{"gp"},
							Description: `SMS/GSM gateway configuration used to actually send text messages.`,
							Usage:       `SMS/GSM gateway configuration used to actually send text messages.`,
							Commands: []*cli.Command{
								messagingdefs.GsmProviderBrowseActionCliHandler(GsmProviderBrowseAction),
								messagingdefs.GsmProviderGetActionCliHandler(GsmProviderGetAction),
								messagingdefs.GsmProviderCreateActionCliHandler(GsmProviderCreateAction),
								messagingdefs.GsmProviderUpdateActionCliHandler(GsmProviderUpdateAction),
								messagingdefs.GsmProviderAwareDeletePreviewActionCliHandler(GsmProviderAwareDeletePreviewAction),
								messagingdefs.GsmProviderAwareDeleteActionCliHandler(GsmProviderAwareDeleteAction),
								&GsmProviderTestCmd,
							},
						},
						messagingdefs.GsmSendSmsWithProviderActionCliHandler(GsmSendSmsWithProviderAction),
					},
				},
				&cli.Command{
					Name:        "webpush",
					Description: `Per-user browser web-push subscriptions.`,
					Usage:       `Per-user browser web-push subscriptions.`,
					Commands: []*cli.Command{
						messagingdefs.WebPushConfigBrowseActionCliHandler(WebPushConfigBrowseAction),
						messagingdefs.WebPushConfigGetActionCliHandler(WebPushConfigGetAction),
						messagingdefs.WebPushConfigCreateActionCliHandler(WebPushConfigCreateAction),
						messagingdefs.WebPushConfigUpdateActionCliHandler(WebPushConfigUpdateAction),
						messagingdefs.WebPushConfigAwareDeletePreviewActionCliHandler(WebPushConfigAwareDeletePreviewAction),
						messagingdefs.WebPushConfigAwareDeleteActionCliHandler(WebPushConfigAwareDeleteAction),
					},
				},
				&cli.Command{
					Name:        "config",
					Description: `The single, global messaging config (general email/gsm provider, invite/otp template ids).`,
					Usage:       `The single, global messaging config (general email/gsm provider, invite/otp template ids).`,
					Commands: []*cli.Command{
						messagingdefs.MessagingConfigGetActionCliHandler(MessagingConfigGetAction),
						messagingdefs.MessagingConfigUpdateActionCliHandler(MessagingConfigUpdateAction),
					},
				},
			},
		},
	})

	return module
}
