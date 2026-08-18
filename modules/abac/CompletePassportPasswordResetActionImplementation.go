package abac

import (
	"time"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/security"
)

// CompletePassportPasswordResetAction is the public counterpart to
// ClassicPassportRequestOtpAction/SendPassportResetEmailAction: given the passport
// value and the OTP code that was emailed/texted to it, it sets a new password and
// signs the user in - the same OTP row lookup ClassicPassportOtpActionImplementation.go
// uses to sign a user in without a password reset, just followed by actually setting
// one instead of only issuing a session.
func CompletePassportPasswordResetAction(c abacdefs.CompletePassportPasswordResetActionRequest) (*abacdefs.CompletePassportPasswordResetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}

	if len(c.Body.Password) < 6 {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	// Same lookup ClassicPassportOtpActionImplementation.go's classicPassportOtpCore
	// does: the most recent PublicAuthenticationEntity row matching both value and otp,
	// still within its BlockedUntil window - a zero-value (not-found) row's
	// BlockedUntil is 0, which is always in the past, so it's rejected the same way an
	// expired one is.
	var otpEntity abacdefs.PublicAuthenticationEntity
	fireback.GetDbRef().Where(&abacdefs.PublicAuthenticationEntity{
		PassportValue: c.Body.Value,
		Otp:           c.Body.Otp,
	}).Order("id DESC").Find(&otpEntity)

	if otpEntity.UniqueId == "" || time.Now().UnixNano() >= otpEntity.BlockedUntil {
		return nil, fireback.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
	}

	passport, user, err3 := UnsafeGetUserByPassportValue(c.Body.Value, *query)
	if err3 != nil {
		return nil, err3
	}

	hashed, err4 := security.HashPassword(c.Body.Password)
	if err4 != nil {
		return nil, fireback.CastToIError(err4)
	}

	if _, err5 := PassportActions.Update(*query, &abacdefs.PassportEntity{
		UniqueId: passport.UniqueId,
		Password: hashed,
	}); err5 != nil {
		return nil, err5
	}

	// Consume the OTP so it can't be replayed to reset the password again - same
	// cleanup classicPassportOtpCore does on a successful login.
	if err6 := fireback.GetDbRef().Where(
		&abacdefs.PublicAuthenticationEntity{PassportId: emigo.NullableOf(passport.UniqueId), Otp: c.Body.Otp},
	).Delete(&abacdefs.PublicAuthenticationEntity{}).Error; err6 != nil {
		return nil, fireback.GormErrorToIError(err6)
	}

	session := &abacdefs.UserSessionDto{}
	if token, err7 := AuthorizeWithToken(user, *query); err7 != nil {
		return nil, fireback.CastToIError(err7)
	} else {
		session.Token = token
	}

	return &abacdefs.CompletePassportPasswordResetActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.CompletePassportPasswordResetActionRes{
			Reset:   true,
			Session: emigo.NewOneNullable(*session),
		}),
	}, nil
}
