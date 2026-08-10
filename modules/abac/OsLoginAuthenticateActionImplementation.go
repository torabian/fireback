package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func OsLoginAuthenticateAction(c abacdefs.OsLoginAuthenticateActionRequest) (*abacdefs.OsLoginAuthenticateActionResponse, error) {
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

	return &abacdefs.OsLoginAuthenticateActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}
