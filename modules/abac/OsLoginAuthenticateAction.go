package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	OsLoginAuthenticateImpl = OsLoginAuthenticateAction
}

func OsLoginAuthenticateAction(
	c OsLoginAuthenticateActionRequest, query fireback.QueryDSL,
) (*OsLoginAuthenticateActionResponse, error) {

	res, err := SigninWithOsUser2(query)

	return &OsLoginAuthenticateActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, err
}
