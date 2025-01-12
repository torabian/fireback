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

	AppendNotificationConfigRouter = func(r *[]Module3Action) {
		*r = append(*r,
			Module3Action{
				Method: "POST",
				Url:    "/notification/testmail",
				Handlers: []gin.HandlerFunc{
					HttpSendTestMail,
				},
				RequestEntity:  &TestMailDto{},
				ResponseEntity: &OkayResponseDto{},
				Out: &Module3ActionBody{
					Dto: "OkayResponseDto",
				},
				In: &Module3ActionBody{
					Dto: "TestMailDto",
				},
			},
			Module3Action{
				Method: "GET",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&SecurityModel{
						ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
					}),
					HttpGetNotificationWorkspaceConfig,
				},
				ResponseEntity: &NotificationConfigEntity{},
				Out: &Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
			Module3Action{
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
				Out: &Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
				In: &Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
		)

	}
}
