package workspaces

import (
	"github.com/gin-gonic/gin"
)

func HttpSendTestMail(c *gin.Context) {
	HttpPostEntity(c, NotificationTestMailAction)
}

func HttpGetNotificationWorkspaceConfig(c *gin.Context) {
	HttpGetEntity(c, NotificationWorkspaecConfigActionGet)
}

func HttpUpdateNotificationWorkspaceConfig(c *gin.Context) {
	HttpUpdateEntity(c, NotificationWorkspaceConfigActionUpdate)
}

func init() {

	AppendNotificationConfigRouter = func(r *[]Module2Action) {
		*r = append(*r,
			Module2Action{
				Method: "POST",
				Url:    "/notification/testmail",
				Handlers: []gin.HandlerFunc{
					HttpSendTestMail,
				},
				RequestEntity:  &TestMailDto{},
				ResponseEntity: &OkayResponseDto{},
				Out: Module2ActionBody{
					Dto: "OkayResponseDto",
				},
				In: Module2ActionBody{
					Dto: "TestMailDto",
				},
			},
			Module2Action{
				Method: "GET",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&SecurityModel{
						ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
					}),
					HttpGetNotificationWorkspaceConfig,
				},
				ResponseEntity: &NotificationConfigEntity{},
				Out: Module2ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
			Module2Action{
				Method: "PATCH",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&SecurityModel{
						ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
					}),
					HttpUpdateNotificationWorkspaceConfig,
				},
				RequestEntity:  &NotificationConfigEntity{},
				ResponseEntity: &NotificationConfigEntity{},
				Out: Module2ActionBody{
					Entity: "NotificationConfigEntity",
				},
				In: Module2ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
		)

	}
}
