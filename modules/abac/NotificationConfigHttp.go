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

func init() {

	AppendNotificationConfigRouter = func(r *[]fireback.Module3Action) {
		*r = append(*r,
			fireback.Module3Action{
				Method: "POST",
				Url:    "/notification/testmail",
				Handlers: []gin.HandlerFunc{
					HttpSendTestMail,
				},
				RequestEntity:  &TestMailDto{},
				ResponseEntity: &OkayResponseDto{},
				Out: &fireback.Module3ActionBody{
					Dto: "OkayResponseDto",
				},
				In: &fireback.Module3ActionBody{
					Dto: "TestMailDto",
				},
			},
			fireback.Module3Action{
				Method: "GET",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&fireback.SecurityModel{
						ActionRequires: []fireback.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
					}),
					HttpGetNotificationWorkspaceConfig,
				},
				ResponseEntity: &NotificationConfigEntity{},
				Out: &fireback.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
			fireback.Module3Action{
				Method: "PATCH",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&fireback.SecurityModel{
						ActionRequires: []fireback.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
					}),
					HttpUpdateNotificationWorkspaceConfig,
				},
				RequestEntity:  &NotificationConfigEntity{},
				ResponseEntity: &NotificationConfigEntity{},
				Out: &fireback.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
				In: &fireback.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
		)

	}
}
