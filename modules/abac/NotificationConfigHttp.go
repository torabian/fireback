package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
)

func HttpSendTestMail(c *gin.Context) {
	workspaces.HttpPostEntity(c, NotificationTestMailAction)
}

func HttpGetNotificationWorkspaceConfig(c *gin.Context) {
	workspaces.HttpGetEntity(c, NotificationWorkspaecConfigActionGet)
}

func HttpUpdateNotificationWorkspaceConfig(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, NotificationWorkspaceConfigActionUpdate)
}

func init() {

	AppendNotificationConfigRouter = func(r *[]workspaces.Module3Action) {
		*r = append(*r,
			workspaces.Module3Action{
				Method: "POST",
				Url:    "/notification/testmail",
				Handlers: []gin.HandlerFunc{
					HttpSendTestMail,
				},
				RequestEntity:  &TestMailDto{},
				ResponseEntity: &OkayResponseDto{},
				Out: &workspaces.Module3ActionBody{
					Dto: "OkayResponseDto",
				},
				In: &workspaces.Module3ActionBody{
					Dto: "TestMailDto",
				},
			},
			workspaces.Module3Action{
				Method: "GET",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&workspaces.SecurityModel{
						ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
					}),
					HttpGetNotificationWorkspaceConfig,
				},
				ResponseEntity: &NotificationConfigEntity{},
				Out: &workspaces.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
			workspaces.Module3Action{
				Method: "PATCH",
				Url:    "/notification/workspace/config",
				Handlers: []gin.HandlerFunc{
					WithAuthorization(&workspaces.SecurityModel{
						ActionRequires: []workspaces.PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
					}),
					HttpUpdateNotificationWorkspaceConfig,
				},
				RequestEntity:  &NotificationConfigEntity{},
				ResponseEntity: &NotificationConfigEntity{},
				Out: &workspaces.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
				In: &workspaces.Module3ActionBody{
					Entity: "NotificationConfigEntity",
				},
			},
		)

	}
}
