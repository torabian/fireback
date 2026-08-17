package abac

import (
	"fmt"
	neturl "net/url"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/messaging"
	"github.com/torabian/fireback/modules/fireback"
)

func ClassicPassportRequestOtpAction(c abacdefs.ClassicPassportRequestOtpActionRequest) (*abacdefs.ClassicPassportRequestOtpActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	res, err2 := classicPassportRequestOtpCore(c.Body, *query)
	if err2 != nil {
		return nil, err2
	}

	return &abacdefs.ClassicPassportRequestOtpActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

// classicPassportRequestOtpCore holds the actual implementation, reusable by callers
// which already have a resolved QueryDSL (such as checkClassicPassportCore, or the cli-only AuthFlow).
func classicPassportRequestOtpCore(req abacdefs.ClassicPassportRequestOtpActionReq, query fireback.QueryDSL) (*abacdefs.ClassicPassportRequestOtpActionRes, *fireback.IError) {

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	// Configurable via Abac.emi.yml's `config:` block (env OTP_LOCKOUT_SECONDS, or
	// "abac config otp-lockout-seconds set") - see modules/abac/defs/Configuration.go.
	var secondsToUnblock int64 = abacdefs.LoadConfiguration().OtpLockoutSeconds

	var olderEntity *abacdefs.PublicAuthenticationEntity = nil
	fireback.GetDbRef().Where(&abacdefs.PublicAuthenticationEntity{PassportValue: req.Value}).Find(&olderEntity)

	if olderEntity != nil && time.Now().UnixNano() < olderEntity.BlockedUntil {
		remaining := (olderEntity.BlockedUntil - time.Now().UnixNano()) / 1000000000
		return &abacdefs.ClassicPassportRequestOtpActionRes{
			BlockedUntil:     olderEntity.BlockedUntil,
			SecondsToUnblock: remaining,
		}, fireback.Create401Error(&AbacMessages.OtaRequestBlockedUntil, []string{})
	} else {
		// Let's delete the record, to start the process fresh
		fireback.GetDbRef().Where(&abacdefs.PublicAuthenticationEntity{PassportValue: req.Value}).Delete(&abacdefs.PublicAuthenticationEntity{})
	}

	passport, user, err := UnsafeGetUserByPassportValue(req.Value, query)

	// We only throw error if passport not available, other errors we need to throw
	if err != nil {
		if item := err.Message["$"]; item != "PassportNotAvailable" {
			return nil, err
		}
	}

	uid := fireback.UUID()
	otp := fireback.GenerateRandomKey(6)
	// Configurable via Abac.emi.yml's `config:` block (env SELF_SERVICE_BASE_URL, or
	// "abac config self-service-base-url set") - points at
	// ui/packages/selfservice's own "reset-password" route (SelfServiceRoutes.tsx),
	// which reads ?value= itself (ResetPassword.presenter.tsx) to prefill the form -
	// not the session uid below, which is only ever read back by this same row,
	// never by the frontend. CompletePassportPasswordResetActionImplementation.go
	// verifies the actual otp against (value, otp), the same pair
	// ClassicPassportOtpAction does.
	url := abacdefs.LoadConfiguration().SelfServiceBaseUrl + "/en/selfservice/reset-password?value=" + neturl.QueryEscape(req.Value)
	secret := fireback.UUID_Long() + "." + fireback.UUID_Long()

	item := &abacdefs.PublicAuthenticationEntity{
		UniqueId:            uid,
		BlockedUntil:        time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
		Otp:                 otp,
		RecoveryAbsoluteUrl: url,
		PassportValue:       req.Value,
		WorkspaceId:         emigo.NullableOf(ROOT_VAR),
		SessionSecret:       secret,
		IsInCreationProcess: emigo.NullableOf(false),
	}
	if passport != nil {
		item.PassportId = emigo.NullableOf(passport.UniqueId)
	}
	if user != nil {
		item.UserId = emigo.NullableOf(user.UniqueId)
	}

	// add time based dual factor information
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Fireback",
		AccountName: req.Value,
	})

	totpSecret := key.Secret()
	totpLink := key.URL()

	if totpSecret != "" {
		item.TotpSecret = totpSecret
	}
	if totpLink != "" {
		item.TotpLink = totpLink
	}

	// If passport doesn't exists, we assume now user wants to create an account.
	// we will store the entity with details, and after verifying, the account creation process starts
	if passport == nil {

		item.IsInCreationProcess = emigo.NullableOf(true)
	}

	if err := fireback.GetDbRef().Create(item).Error; err != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	_, passportType := validatePassportType(req.Value)

	if passportType == PASSPORT_METHOD_PHONE {

		result := QuickGetOtpMessage(query, SMS_OTP)
		body, err3 := CompileContent(result, map[string]string{"Otp": otp})
		if err3 != nil {
			return nil, fireback.CastToIError(err3)
		}

		if _, err2 := GsmSendSMSUsingNotificationConfig(body, []string{req.Value}); err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}

	} else if passportType == PASSPORT_METHOD_EMAIL {
		result := QuickGetOtpMessage(query, EMAIL_OTP)
		var body = ""
		var title = ""
		if body0, err3 := CompileContent(result, map[string]string{"Otp": otp}); err3 != nil {
			return nil, fireback.CastToIError(err3)
		} else {
			body = body0
		}

		if title0, err3 := CompileTitle(result, map[string]string{"Otp": otp}); err3 != nil {
			return nil, fireback.CastToIError(err3)
		} else {
			title = title0
		}

		msg := messaging.EmailMessageContent{
			Subject:   title,
			Content:   body,
			ToEmail:   req.Value,
			FromName:  "Account Center",
			FromEmail: "accountcenter@gmail.com",
			ToName:    req.Value,
		}

		if _, err2 := SendEmailUsingNotificationConfig(&msg, GENERAL_SENDER); err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}

	} else if passportType == ANONYMOUS_AUTHENTICATION && query.C != nil && query.G == nil {
		// If only, in cli mode, we print the otp information on the cli screen
		fmt.Println("Your otp code for anonymous login on cli only: ", otp)
	} else {
		return nil, &fireback.IError{Message: AbacMessages.OtpNotAvailableForThisType}
	}

	return &abacdefs.ClassicPassportRequestOtpActionRes{
		SecondsToUnblock: secondsToUnblock,
	}, nil
}
