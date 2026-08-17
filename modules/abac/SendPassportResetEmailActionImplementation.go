package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// SendPassportResetEmailAction lets a root administrator trigger the same OTP/recovery
// email or SMS ClassicPassportRequestOtpAction sends for self-service login
// (ClassicPassportRequestOtpActionImplementation.go's classicPassportRequestOtpCore,
// reused here as-is) - but addressed by passport uniqueId, rather than requiring the
// admin to already know the passport's raw value, and gated behind
// PERM_ROOT_PASSPORT_UPDATE instead of being publicly callable. Lets an admin hand a
// locked-out user a way to reset their own password (via
// CompletePassportPasswordResetActionImplementation.go) rather than the admin setting
// it directly (see SetPassportPasswordActionImplementation.go for that alternative).
func SendPassportResetEmailAction(c abacdefs.SendPassportResetEmailActionRequest) (*abacdefs.SendPassportResetEmailActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_UPDATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}

	q := *query
	q.UniqueId = c.Body.UniqueId
	passport, err3 := PassportActions.GetOne(q)
	if err3 != nil {
		return nil, err3
	}

	if _, err4 := classicPassportRequestOtpCore(abacdefs.ClassicPassportRequestOtpActionReq{Value: passport.Value}, *query); err4 != nil {
		return nil, err4
	}

	return &abacdefs.SendPassportResetEmailActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.SendPassportResetEmailActionRes{Sent: true}),
	}, nil
}
