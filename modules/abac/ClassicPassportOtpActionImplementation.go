package abac

import (
	"time"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func ClassicPassportOtpAction(c abacdefs.ClassicPassportOtpActionRequest) (*abacdefs.ClassicPassportOtpActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	res, err2 := classicPassportOtpCore(c.Body, *query)
	if err2 != nil {
		return nil, err2
	}

	return &abacdefs.ClassicPassportOtpActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

// classicPassportOtpCore holds the actual implementation, reusable by callers which
// already have a resolved QueryDSL (such as the cli-only AuthFlow).
func classicPassportOtpCore(req abacdefs.ClassicPassportOtpActionReq, query fireback.QueryDSL) (*abacdefs.ClassicPassportOtpActionRes, *fireback.IError) {
	ClearPassportValue(&req.Value)

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	olderEntity := &abacdefs.PublicAuthenticationEntity{}
	fireback.GetDbRef().Where(&abacdefs.PublicAuthenticationEntity{
		PassportValue: req.Value,
		Otp:           req.Otp,
	}).Order("id DESC").Find(olderEntity)

	if olderEntity == nil || time.Now().UnixNano() >= olderEntity.BlockedUntil {
		return nil, fireback.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
	}

	if olderEntity.IsInCreationProcess.OrDefault(false) {
		// in some cases, the otp alone should be enough and can complete signup process.
		// for example, phone number often is enough for authroization of sms or phone call
		// has been through
		return &abacdefs.ClassicPassportOtpActionRes{
			ContinueWithCreation: true,
			SessionSecret:        olderEntity.SessionSecret,
			TotpUrl:              olderEntity.TotpLink,
		}, nil
	}

	passport, user, err := UnsafeGetUserByPassportValue(req.Value, query)
	if err != nil {
		return nil, err
	}

	if olderEntity.UniqueId != "" {
		if req.Otp != "" {

			if req.Otp == olderEntity.Otp {
				session := &abacdefs.UserSessionDto{}

				if token, err := AuthorizeWithToken(user, query); err != nil {
					return nil, fireback.CastToIError(err)
				} else {
					session.Token = token
				}

				// Delete the session so user cannot login again
				err2 := fireback.GetDbRef().Where(
					&abacdefs.PublicAuthenticationEntity{PassportId: emigo.NullableOf(passport.UniqueId), Otp: req.Otp},
				).Delete(&abacdefs.PublicAuthenticationEntity{}).Error

				if err2 != nil {
					return nil, fireback.GormErrorToIError(err2)
				}

				return &abacdefs.ClassicPassportOtpActionRes{
					Session: emigo.NewOneNullable(*session),
				}, nil
			}
		}
	}
	return nil, fireback.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
}
