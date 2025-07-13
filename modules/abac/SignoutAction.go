package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SignoutActionImp = SignoutAction
}
func SignoutAction(
	q fireback.QueryDSL) (string,
	*fireback.IError,
) {

	// Clear secure cookie
	q.G.SetCookie("authorization", "", 3600*24, "/", "", true, true)

	return "OKAY", nil
}
