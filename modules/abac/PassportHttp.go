package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {

	AppendPassportRouter = func(r *[]fireback.Module3Action) {

		*r = append(*r,

			fireback.Module3Action{
				Method: "POST",
				Url:    ("/passport/authorizeOs"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						fireback.HttpPostEntity(c, PassportActionAuthorizeOs2)
					},
				},
				RequestEntity:  &EmailAccountSigninDto{},
				ResponseEntity: &UserSessionDto{},
				In: &fireback.Module3ActionBody{
					Dto: "EmailAccountSigninDto",
				},
				Out: &fireback.Module3ActionBody{
					Dto: "UserSessionDto",
				},
				Action: PassportActionAuthorizeOs2,
				Format: "POST_ONE",
			},
		)

	}
}
