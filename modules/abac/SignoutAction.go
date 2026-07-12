package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	SignoutImpl = SignoutAction
}

func SignoutAction(c SignoutActionRequest, query fireback.QueryDSL) (*SignoutActionResponse, error) {

	// Clear secure cookie
	if c.IsGin() {
		c.GinCtx.(*gin.Context).SetCookie("authorization", "", 3600*24, "/", "", true, true)
	}

	return &SignoutActionResponse{
		Payload: SignoutActionRes{
			Okay: true,
		},
	}, nil
}
