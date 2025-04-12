package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
)

func init() {

	AppendPassportRouter = func(r *[]workspaces.Module3Action) {

		*r = append(*r,

			workspaces.Module3Action{
				Method: "POST",
				Url:    ("/passport/authorizeOs"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						workspaces.HttpPostEntity(c, PassportActionAuthorizeOs2)
					},
				},
				RequestEntity:  &EmailAccountSigninDto{},
				ResponseEntity: &UserSessionDto{},
				In: &workspaces.Module3ActionBody{
					Dto: "EmailAccountSigninDto",
				},
				Out: &workspaces.Module3ActionBody{
					Dto: "UserSessionDto",
				},
				Action: PassportActionAuthorizeOs2,
				Format: "POST_ONE",
			},
		)

	}
}
