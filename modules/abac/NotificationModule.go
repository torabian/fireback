//go:build !wasm

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

				// NotificationEntity (an actual notification delivered to a user) is a
				// separate entity from NotificationConfig (per-workspace email/sms settings)
				// above - see NotificationActions.go. Create/Update aren't wired here: the
				// only sanctioned way to create one is SendNotificationAction, which calls
				// abacdefs.NotificationEntityActions.Create directly rather than going through
				// a public "notification create" endpoint.
				abacdefs.NotificationBrowseActionGin(g, NotificationBrowseAction)
				abacdefs.NotificationGetActionGin(g, NotificationGetAction)
				abacdefs.NotificationAwareDeletePreviewActionGin(g, NotificationAwareDeletePreviewAction)
				abacdefs.NotificationAwareDeleteActionGin(g, NotificationAwareDeleteAction)
				abacdefs.SendNotificationActionGin(g, SendNotificationAction)

				// Self-service side (any authenticated user, not root) - see
				// NotificationSelfServiceActionImplementation.go.
				abacdefs.MyNotificationsActionGin(g, MyNotificationsAction)
				abacdefs.MarkNotificationReadActionGin(g, MarkNotificationReadAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
		ALL_NOTIFICATION_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&abacdefs.NotificationConfigEntity{},
			&abacdefs.NotificationEntity{},
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

				// The actual "send a notification to users" entry point - `notification send
				// --user-ids u1,u2 --title ... --body ...` (see SendNotificationAction).
				abacdefs.SendNotificationActionCliHandler(SendNotificationAction),

				// Self-service side - runs as whichever user the CLI's own session belongs
				// to, not root (see NotificationSelfServiceActionImplementation.go).
				abacdefs.MyNotificationsActionCliHandler(MyNotificationsAction),
				abacdefs.MarkNotificationReadActionCliHandler(MarkNotificationReadAction),
				{
					Name:        "list",
					Description: "Manage delivered notifications (root only) - browse/get/delete. See 'notification send' to create one.",
					Commands: []*cli.Command{
						abacdefs.NotificationBrowseActionCliHandler(NotificationBrowseAction),
						abacdefs.NotificationGetActionCliHandler(NotificationGetAction),
						abacdefs.NotificationAwareDeletePreviewActionCliHandler(NotificationAwareDeletePreviewAction),
						abacdefs.NotificationAwareDeleteActionCliHandler(NotificationAwareDeleteAction),
					},
				},
			},
		},
	})

	return module
}
