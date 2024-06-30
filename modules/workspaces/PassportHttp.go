package workspaces

import (
	"github.com/gin-gonic/gin"
)

func init() {

	AppendPassportRouter = func(r *[]Module2Action) {

		*r = append(*r,
			Module2Action{
				Method: "POST",
				Url:    ("/passport/signup/email"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionEmailSignup)
					},
				},
				RequestEntity:  &ClassicAuthDto{},
				ResponseEntity: &UserSessionDto{},
				In: &Module2ActionBody{
					Dto: "ClassicAuthDto",
				},
				Out: &Module2ActionBody{
					Dto: "UserSessionDto",
				},
			},
			Module2Action{
				Method: "POST",
				Url:    ("/passport/signin/email"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionEmailSignin)
					},
				},
				RequestEntity:  &EmailAccountSigninDto{},
				ResponseEntity: &UserSessionDto{},
				In: &Module2ActionBody{
					Dto: "EmailAccountSigninDto",
				},
				Out: &Module2ActionBody{
					Dto: "UserSessionDto",
				},
			},
			Module2Action{
				Method: "POST",
				Url:    ("/passport/authorizeOs"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionAuthorizeOs2)
					},
				},
				RequestEntity:  &EmailAccountSigninDto{},
				ResponseEntity: &UserSessionDto{},
				In: &Module2ActionBody{
					Dto: "EmailAccountSigninDto",
				},
				Out: &Module2ActionBody{
					Dto: "UserSessionDto",
				},
				Action: PassportActionAuthorizeOs2,
				Format: "POST_ONE",
			},
			Module2Action{
				Method: "POST",
				Url:    ("/passport/request-reset-mail-password"),
				Handlers: []gin.HandlerFunc{
					func(c *gin.Context) {
						HttpPostEntity(c, PassportActionAuthorize2)
					},
				},
				RequestEntity:  &OtpAuthenticateDto{},
				ResponseEntity: &EmailOtpResponseDto{},
				In: &Module2ActionBody{
					Dto: "OtpAuthenticateDto",
				},
				Out: &Module2ActionBody{
					Dto: "EmailOtpResponseDto",
				},
			},
		)

	}
}
