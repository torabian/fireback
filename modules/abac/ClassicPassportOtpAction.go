package abac

import (
	"time"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func ClassicPassportOtpAction(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	res, err2 := classicPassportOtpCore(c.Body, *query)
	if err2 != nil {
		return nil, err2
	}

	return &ClassicPassportOtpActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

// classicPassportOtpCore holds the actual implementation, reusable by callers which
// already have a resolved QueryDSL (such as the cli-only AuthFlow).
func classicPassportOtpCore(req ClassicPassportOtpActionReq, query fireback.QueryDSL) (*ClassicPassportOtpActionRes, *fireback.IError) {
	ClearPassportValue(&req.Value)

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	olderEntity := &PublicAuthenticationEntity{}
	fireback.GetDbRef().Where(&PublicAuthenticationEntity{
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
		return &ClassicPassportOtpActionRes{
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
				session := &UserSessionDto{}

				if token, err := user.AuthorizeWithToken(query); err != nil {
					return nil, fireback.CastToIError(err)
				} else {
					session.Token = token
				}

				// Delete the session so user cannot login again
				err2 := fireback.GetDbRef().Where(
					&PublicAuthenticationEntity{PassportId: emigo.NullableOf(passport.UniqueId), Otp: req.Otp},
				).Delete(&PublicAuthenticationEntity{}).Error

				if err2 != nil {
					return nil, fireback.GormErrorToIError(err2)
				}

				return &ClassicPassportOtpActionRes{
					Session: emigo.NewOneNullable(*session),
				}, nil
			}
		}
	}
	return nil, fireback.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
}
