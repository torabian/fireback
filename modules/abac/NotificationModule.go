package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// NotificationModuleSetup wires notificationConfig only now - GsmProvider, EmailProvider
// and EmailSender (CRUD + the SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider
// actions) moved to their own module, see modules/abac/messaging.ModuleSetup, registered
// alongside this one in AbacCompleteModules().
func NotificationModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name: "abac",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *fireback.FirebackApp) error{
			func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
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
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&NotificationConfigEntity{},
		)
	})

	module.ProvideCliHandlers([]*cli.Command{
		{
			Name:        "notification",
			Description: "Manage the notification system, emails, text messages, templates and so on",
			Usage:       "Manage the workspace-level notification config (invite emails, general senders, and so on)",
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
