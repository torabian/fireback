package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func SignoutAction(c SignoutActionRequest) (*SignoutActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

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
