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
	// err2 is a concrete *fireback.IError; returning it directly as the built-in error
	// interface would wrap a nil pointer in a non-nil interface value on success (the
	// classic Go typed-nil-in-interface gotcha - see
	// CheckClassicPassportActionImplementation.go's wrapCheckPassportResult for the
	// same issue spelled out in full), which crashes the generated Gin handler with a
	// nil-pointer panic on *every* call, since it always calls err.Error() /
	// ToPublicJSON() whenever the interface itself is non-nil.
	if err2 != nil {
		return nil, err2
	}

	return &OsLoginAuthenticateActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}
