package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/security"
)

// SetPassportPasswordAction lets a root administrator directly set a new password on
// any passport - unlike ChangePasswordAction (ChangePasswordActionImplementation.go),
// which only ever touches the calling user's own passport (it resolves the target
// passport from the session, ignoring any uniqueId in its body). This one takes the
// target passport's uniqueId from the body and requires PERM_ROOT_PASSPORT_UPDATE, for
// an admin resetting a locked-out user's password on their behalf - see
// SendPassportResetEmailActionImplementation.go for the alternative that lets the user
// pick their own new password instead.
func SetPassportPasswordAction(c abacdefs.SetPassportPasswordActionRequest) (*abacdefs.SetPassportPasswordActionResponse, error) {
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

	// Same minimum-length rule ChangePasswordAction enforces for a user's own
	// password - an admin setting one on someone's behalf shouldn't be able to
	// set a weaker one than the user could themselves.
	if len(c.Body.Password) < 6 {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	q := *query
	q.UniqueId = c.Body.UniqueId
	existing, err3 := PassportActions.GetOne(q)
	if err3 != nil {
		return nil, err3
	}

	hashed, err1 := security.HashPassword(c.Body.Password)
	if err1 != nil {
		return nil, fireback.CastToIError(err1)
	}

	updated, err4 := PassportActions.Update(q, &abacdefs.PassportEntity{
		UniqueId: existing.UniqueId,
		Password: hashed,
	})
	if err4 != nil {
		return nil, err4
	}

	return &abacdefs.SetPassportPasswordActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.SetPassportPasswordActionRes{Changed: updated.Password != existing.Password}),
	}, nil
}
