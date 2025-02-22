package workspaces

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func init() {
	// Override the implementation with our actual code.
	ClassicPassportRequestOtpActionImp = ClassicPassportRequestOtpAction
}

func ClassicPassportRequestOtpAction(req *ClassicPassportRequestOtpActionReqDto, q QueryDSL) (*ClassicPassportRequestOtpActionResDto, *IError) {
	if err := ClassicPassportRequestOtpActionReqValidator(req); err != nil {
		return nil, err
	}

	var secondsToUnblock int64 = 120

	var olderEntity *PublicAuthenticationEntity = nil
	GetDbRef().Where(&PublicAuthenticationEntity{PassportValue: req.Value}).Find(&olderEntity)

	if olderEntity != nil && time.Now().UnixNano() < olderEntity.BlockedUntil {
		remaining := (olderEntity.BlockedUntil - time.Now().UnixNano()) / 1000000000
		return &ClassicPassportRequestOtpActionResDto{
			BlockedUntil:     &olderEntity.BlockedUntil,
			SecondsToUnblock: &remaining,
		}, Create401Error(&WorkspacesMessages.OtaRequestBlockedUntil, []string{})
	} else {
		// Let's delete the record, to start the process fresh
		GetDbRef().Where(&PublicAuthenticationEntity{PassportValue: req.Value}).Delete(&PublicAuthenticationEntity{})
	}

	passport, user, err := UnsafeGetUserByPassportValue(*req.Value, q)

	// We only throw error if passport not available, other errors we need to throw
	if err != nil {
		if item := err.Message["$"]; item != "PassportNotAvailable" {
			return nil, err
		}
	}

	uid := UUID()
	otp := GenerateRandomKey(6)
	url := "http://localhost:8888/reset-password?session=" + uid
	secret := UUID_Long() + "." + UUID_Long()

	item := &PublicAuthenticationEntity{
		UniqueId:            uid,
		BlockedUntil:        time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
		Otp:                 &otp,
		RecoveryAbsoluteUrl: &url,
		PassportValue:       req.Value,
		WorkspaceId:         &ROOT_VAR,
		SessionSecret:       &secret,
		IsInCreationProcess: &FALSE,
		Passport:            passport,
		User:                user,
	}

	// add time based dual factor information
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Fireback",
		AccountName: *req.Value,
	})

	totpSecret := key.Secret()
	totpLink := key.URL()

	if totpSecret != "" {
		item.TotpSecret = &totpSecret
	}
	if totpLink != "" {
		item.TotpLink = &totpLink
	}

	// If passport doesn't exists, we assume now user wants to create an account.
	// we will store the entity with details, and after verifying, the account creation process starts
	if passport == nil {

		item.IsInCreationProcess = &TRUE
	}

	if err := GetDbRef().Create(item).Error; err != nil {
		return nil, GormErrorToIError(err)
	}

	_, passportType := validatePassportType(*req.Value)

	if passportType == PASSPORT_METHOD_PHONE {

		result := QuickGetOtpMessage(q, SMS_OTP)
		body, err3 := result.CompileContent(map[string]string{"Otp": otp})
		if err3 != nil {
			return nil, CastToIError(err3)
		}

		if _, err2 := GsmSendSMSUsingNotificationConfig(body, []string{*req.Value}); err2 != nil {
			return nil, GormErrorToIError(err2)
		}

	} else if passportType == PASSPORT_METHOD_EMAIL {
		result := QuickGetOtpMessage(q, EMAIL_OTP)
		var body = ""
		var title = ""
		if body0, err3 := result.CompileContent(map[string]string{"Otp": otp}); err3 != nil {
			return nil, CastToIError(err3)
		} else {
			body = body0
		}

		if title0, err3 := result.CompileTitle(map[string]string{"Otp": otp}); err3 != nil {
			return nil, CastToIError(err3)
		} else {
			title = title0
		}

		msg := EmailMessageContent{
			Subject:   title,
			Content:   body,
			ToEmail:   *req.Value,
			FromName:  "Account Center",
			FromEmail: "accountcenter@gmail.com",
			ToName:    *req.Value,
		}

		if _, err2 := SendEmailUsingNotificationConfig(&msg, GENERAL_SENDER); err2 != nil {
			return nil, GormErrorToIError(err2)
		}

	} else {
		return nil, &IError{Message: WorkspacesMessages.OtpNotAvailableForThisType}
	}

	return &ClassicPassportRequestOtpActionResDto{
		SecondsToUnblock: &secondsToUnblock,
	}, nil

}
