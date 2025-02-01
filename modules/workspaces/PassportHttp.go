package workspaces

import (
	"github.com/gin-gonic/gin"
)

func init() {

	AppendPassportRouter = func(r *[]Module3Action) {

		*r = append(*r,

			Module3Action{
				Method: "POST",
				Url:    ("/passport/authorizeOs"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionAuthorizeOs2)
					},
				},
				RequestEntity:  &EmailAccountSigninDto{},
				ResponseEntity: &UserSessionDto{},
				In: &Module3ActionBody{
					Dto: "EmailAccountSigninDto",
				},
				Out: &Module3ActionBody{
					Dto: "UserSessionDto",
				},
				Action: PassportActionAuthorizeOs2,
				Format: "POST_ONE",
			},
			Module3Action{
				Method: "POST",
				Url:    ("/passport/request-reset-mail-password"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionAuthorize2)
					},
				},
				RequestEntity:  &OtpAuthenticateDto{},
				ResponseEntity: &EmailOtpResponseDto{},
				In: &Module3ActionBody{
					Dto: "OtpAuthenticateDto",
				},
				Out: &Module3ActionBody{
					Dto: "EmailOtpResponseDto",
				},
			},
		)

	}
}
