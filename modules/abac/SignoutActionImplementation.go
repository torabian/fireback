package abac

import (
	"github.com/gin-gonic/gin"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func SignoutAction(c abacdefs.SignoutActionRequest) (*abacdefs.SignoutActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	// Clear secure cookie
	if c.IsGin() {
		c.GinCtx.(*gin.Context).SetCookie("authorization", "", 3600*24, "/", "", true, true)
	}

	return &abacdefs.SignoutActionResponse{
		Payload: abacdefs.SignoutActionRes{
			Okay: true,
		},
	}, nil
}
