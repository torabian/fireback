package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func OsLoginAuthenticateAction(c OsLoginAuthenticateActionRequest) (*OsLoginAuthenticateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	res, err2 := SigninWithOsUser2(*query)

	return &OsLoginAuthenticateActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, err2
}
