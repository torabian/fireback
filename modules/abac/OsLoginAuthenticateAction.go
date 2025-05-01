package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	OsLoginAuthenticateActionImp = OsLoginAuthenticateAction
}
func OsLoginAuthenticateAction(
	q fireback.QueryDSL) (*UserSessionDto,
	*fireback.IError,
) {
	return SigninWithOsUser2(q)

}
