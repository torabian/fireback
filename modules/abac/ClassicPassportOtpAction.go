package abac

import (
	"time"

	"github.com/torabian/fireback/modules/workspaces"
)

func init() {
	// Override the implementation with our actual code.
	ClassicPassportOtpActionImp = ClassicPassportOtpAction
}

func ClassicPassportOtpAction(req *ClassicPassportOtpActionReqDto, q workspaces.QueryDSL) (
	*ClassicPassportOtpActionResDto, *workspaces.IError,
) {

	ClearPassportValue(&req.Value)
	if err := ClassicPassportOtpActionReqValidator(req); err != nil {
		return nil, err
	}

	olderEntity := &PublicAuthenticationEntity{}
	workspaces.GetDbRef().Where(&PublicAuthenticationEntity{
		PassportValue: req.Value,
		Otp:           req.Otp,
	}).Order("id DESC").Find(olderEntity)

	if olderEntity == nil || time.Now().UnixNano() >= olderEntity.BlockedUntil {
		return nil, workspaces.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
	}

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

		return &ClassicPassportOtpActionResDto{
			ContinueWithCreation: true,
			SessionSecret:        olderEntity.SessionSecret,
			TotpUrl:              olderEntity.TotpLink,
		}, nil
	}

	passport, user, err := UnsafeGetUserByPassportValue(req.Value, q)
	if err != nil {
		return nil, err
	}

	if olderEntity.UniqueId != "" {
		if req.Otp != "" {

			if req.Otp == olderEntity.Otp {
				session := &UserSessionDto{}

				if token, err := user.AuthorizeWithToken(q); err != nil {
					return nil, workspaces.CastToIError(err)
				} else {
					session.Token = token
				}

				if err != nil {
					return nil, workspaces.GormErrorToIError(err)
				}

				// Delete the session so user cannot login again
				err2 := workspaces.GetDbRef().Where(
					&PublicAuthenticationEntity{PassportId: workspaces.NewString(passport.UniqueId), Otp: req.Otp},
				).Delete(&PublicAuthenticationEntity{}).Error

				if err2 != nil {
					return nil, workspaces.GormErrorToIError(err)
				}

				return &ClassicPassportOtpActionResDto{
					Session: session,
				}, nil
			}
		}
	}
	return nil, workspaces.Create401Error(&AbacMessages.OtpCodeInvalid, []string{})
}
