package abac

import (
	"github.com/gin-gonic/gin"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// NotificationModuleSetup wires notificationConfig only now - GsmProvider, EmailProvider
// and EmailSender (CRUD + the SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider
// actions) moved to their own module, see modules/abac/messaging.ModuleSetup, registered
// alongside this one in AbacCompleteModules().
func NotificationModuleSetup() *application.ModuleProvider {
	module := &application.ModuleProvider{
		Name: "abac",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				abacdefs.NotificationConfigBrowseActionGin(g, NotificationConfigBrowseAction)
				abacdefs.NotificationConfigGetActionGin(g, NotificationConfigGetAction)
				abacdefs.NotificationConfigCreateActionGin(g, NotificationConfigCreateAction)
				abacdefs.NotificationConfigUpdateActionGin(g, NotificationConfigUpdateAction)
				abacdefs.NotificationConfigAwareDeletePreviewActionGin(g, NotificationConfigAwareDeletePreviewAction)
				abacdefs.NotificationConfigAwareDeleteActionGin(g, NotificationConfigAwareDeleteAction)

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
			&abacdefs.NotificationConfigEntity{},
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
				abacdefs.NotificationConfigBrowseActionCliHandler(NotificationConfigBrowseAction),
				abacdefs.NotificationConfigGetActionCliHandler(NotificationConfigGetAction),
				abacdefs.NotificationConfigCreateActionCliHandler(NotificationConfigCreateAction),
				abacdefs.NotificationConfigUpdateActionCliHandler(NotificationConfigUpdateAction),
				abacdefs.NotificationConfigAwareDeletePreviewActionCliHandler(NotificationConfigAwareDeletePreviewAction),
				abacdefs.NotificationConfigAwareDeleteActionCliHandler(NotificationConfigAwareDeleteAction),
			},
		},
	})

	return module
}
