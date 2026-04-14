package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SignoutImpl = SignoutAction
}

func SignoutAction(c SignoutActionRequest, query fireback.QueryDSL) (*SignoutActionResponse, error) {

	// Clear secure cookie
	c.GinCtx.SetCookie("authorization", "", 3600*24, "/", "", true, true)

	return &SignoutActionResponse{
		Payload: SignoutActionRes{
			Okay: true,
		},
	}, nil
}
