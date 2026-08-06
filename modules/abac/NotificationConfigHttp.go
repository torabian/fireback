package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func HttpSendTestMail(c *gin.Context) {
	fireback.HttpPostEntity(c, NotificationTestMailAction)
}

func HttpGetNotificationWorkspaceConfig(c *gin.Context) {
	fireback.HttpGetEntity(c, NotificationWorkspaecConfigActionGet)
}

func HttpUpdateNotificationWorkspaceConfig(c *gin.Context) {
	fireback.HttpUpdateEntity(c, NotificationWorkspaceConfigActionUpdate)
}

// AppendNotificationConfigRouter wires the custom /notification/testmail and
// /notification/workspace/config routes directly into gin - moved out of the old
// Module3Action + GetNotificationConfigModule3Actions() indirection (which no longer
// exists now that notificationConfig moved to Abac.emi.yml), called directly from
// NotificationModule.go's GinWebServerInitHooks instead.
func AppendNotificationConfigRouter(g *gin.RouterGroup) {
	g.POST("/notification/testmail", HttpSendTestMail)
	g.GET("/notification/workspace/config",
		WithAuthorization(&fireback.SecurityModel{
			ActionRequires: []fireback.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
		}),
		HttpGetNotificationWorkspaceConfig,
	)
	g.PATCH("/notification/workspace/config",
		WithAuthorization(&fireback.SecurityModel{
			ActionRequires: []fireback.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
		}),
		HttpUpdateNotificationWorkspaceConfig,
	)
}
