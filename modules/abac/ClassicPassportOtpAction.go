package abac

import (
	"time"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	ClassicPassportOtpImpl = ClassicPassportOtpAction
}

func ClassicPassportOtpAction(c ClassicPassportOtpActionRequest, query fireback.QueryDSL) (*ClassicPassportOtpActionResponse, error) {
	req := c.Body
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

	// if olderEntity.IsInCreationProcess.Bool {
	// 	completeClassicSignupProcess(
	// 		ClassicSignupActionReq{},
	// 		query,
	// 		olderEntity,
	// 	)
	// }

	if olderEntity.IsInCreationProcess.Bool {
		// in some cases, the otp alone should be enough and can complete signup process.
		// for example, phone number often is enough for authroization of sms or phone call
		// has been through

		// Not possible, because user needs to choose workspace type id
		// ALLOW_PHONE_PASS := true
		// if ok, ptype := validatePassportType(*req.Value); ok && ptype == PASSPORT_METHOD_PHONE && ALLOW_PHONE_PASS {

		// 	user, role, workspace, passport := getPhoneQuickMechanism(*req.Value,)
		// 	session, sessionError := UnsafeGenerateUser(&GenerateUserDto{

		// 		createUser:      true,
		// 		createWorkspace: true,
		// 		createRole:      true,
		// 		createPassport:  true,

		// 		user:      user,
		// 		role:      role,
		// 		workspace: workspace,
		// 		passport:  passport,

		// 		// We want always to be able to login regardless
		// 		restricted: true,
		// 	}, q)

		// 	if sessionError != nil {
		// 		return nil, sessionError
		// 	} else {
		// 		return &ClassicPassportOtpActionResDto{
		// 			Session: session,
		// 		}, nil
		// 	}
		// }

		return &ClassicPassportOtpActionResponse{
			Payload: fireback.GResponseSingleItem(ClassicPassportOtpActionRes{
				ContinueWithCreation: true,
				SessionSecret:        olderEntity.SessionSecret,
				TotpUrl:              olderEntity.TotpLink,
			}),
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

				if err != nil {
					return nil, fireback.GormErrorToIError(err)
				}

				// Delete the session so user cannot login again
				err2 := fireback.GetDbRef().Where(
					&PublicAuthenticationEntity{PassportId: fireback.NewString(passport.UniqueId), Otp: req.Otp},
				).Delete(&PublicAuthenticationEntity{}).Error

				if err2 != nil {
					return nil, fireback.GormErrorToIError(err)
				}

				return &ClassicPassportOtpActionResponse{
					Payload: fireback.GResponseSingleItem(ClassicPassportOtpActionRes{
						Session: emigo.NullableOf(*session),
					}),
				}, nil
			}
		}
	}
	return nil, fireback.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
}
