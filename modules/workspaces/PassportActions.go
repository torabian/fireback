package workspaces

import (
	"fmt"
	reflect "reflect"
	"strings"
	"time"
)

var PASSPORT_METHOD_EMAIL = "email"
var PASSPORT_METHOD_PHONE = "phonenumber"

func PassportActionCreate(
	dto *PassportEntity, query QueryDSL,
) (*PassportEntity, *IError) {
	return PassportActionCreateFn(dto, query)
}

func PassportActionUpdate(
	query QueryDSL,
	fields *PassportEntity,
) (*PassportEntity, *IError) {
	return PassportActionUpdateFn(query, fields)
}

func GetUserByPassport2(value string) (*UserEntity, *PassportEntity, *IError) {
	passport := &PassportEntity{}
	err := GetDbRef().Where(&PassportEntity{Value: &value}).First(passport).Error

	if err != nil {
		return nil, nil, GormErrorToIError(err)
	}

	user := &UserEntity{}
	GetDbRef().Where(&UserEntity{UniqueId: *passport.UserId}).First(user)

	return user, passport, nil
}

func GetUserByPassport(value string) (*UserEntity, *PassportEntity, error) {
	passport := &PassportEntity{}
	GetDbRef().Where(&PassportEntity{Value: &value}).First(passport)

	user := &UserEntity{}
	GetDbRef().Where(&UserEntity{UniqueId: *passport.UserId}).First(user)

	return user, passport, nil
}

/**
*	Logic of forget password is to generate a row, and tell how long is reset valid,
*	And how long user needs to wait for next one. We return this information to client
*	So it can tell user how long left to wait.
**/
func RequestMailPasswordForget(
	dto *OtpAuthenticateDto,
	query QueryDSL,
) (*EmailOtpResponseDto, *IError) {

	var secondsToUnblock int64 = 30
	user, passport, err := GetUserByPassport(*dto.Value)

	if err != nil {
		return nil, GormErrorToIError(err)
	}

	olderEntity := &ForgetPasswordEntity{}
	GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Find(olderEntity)

	if olderEntity.UniqueId != "" {

		fmt.Println(dto.Otp, olderEntity.Otp)
		// If user has provided the otp, create user session in fact.
		if dto.Otp != nil {

			if *dto.Otp == *olderEntity.Otp {
				session, err := PassportActionCreateSessionOtp(*passport.Value)

				if err != nil {
					return nil, GormErrorToIError(err)
				}

				// Delete the session so user cannot login again
				err2 := GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{}).Error

				if err2 != nil {
					return nil, GormErrorToIError(err)
				}

				return &EmailOtpResponseDto{
					UserSession: session,
				}, nil
			} else {
				return &EmailOtpResponseDto{
					Request: olderEntity,
				}, CreateIErrorString(PassportMessageCode.OtpCodeInvalid, []string{}, 403)
			}
		}

		if time.Now().UnixNano() < olderEntity.BlockedUntil {

			// @todo: fix the until
			// olderEntity.SecondsToUnblock = int64(time.Until(olderEntity.BlockedUntil).Seconds())

			return &EmailOtpResponseDto{
					UserSession: nil,
					Request:     olderEntity,
				}, CreateIErrorString(
					PassportMessageCode.OTARequestBlockedUntil, []string{}, 403,
				)
		} else {
			// In this case, user has the chance to re-request. Let's clean the previous items first
			GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{})
		}
	}

	{

		if err != nil {
			return nil, GormErrorToIError(err)
		}

		if passport == nil || user == nil || user.UniqueId == "" {
			return nil, CreateIErrorString(PassportMessageCode.UserDoesNotExist, []string{}, 403)
		}

		uid := UUID()
		otp := GenerateRandomKey(6)
		url := "http://localhost:8888/reset-password?session=" + uid
		item := &ForgetPasswordEntity{
			User:                user,
			Passport:            passport,
			UniqueId:            uid,
			ValidUntil:          time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			BlockedUntil:        time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			SecondsToUnblock:    &secondsToUnblock,
			Otp:                 &otp,
			RecoveryAbsoluteUrl: &url,
		}

		err = GetDbRef().Create(item).Error

		if err != nil {
			return nil, GormErrorToIError(err)
		}

		// if dto.Type == "sms" {
		// 	err = SendSMSFromTemplate("forget-password.html", item)
		// 	if err != nil {
		// 		return nil, GormErrorToIError(err)
		// 	}

		// } else if dto.Type == "email" {
		// 	err = SendEmailFromTemplate("forget-password.html", item)
		// 	if err != nil {
		// 		return nil, GormErrorToIError(err)
		// 	}
		// }

		return &EmailOtpResponseDto{Request: item}, nil
	}
}

func PassportActionAuthorize2(
	dto *OtpAuthenticateDto,
	query QueryDSL,
) (*EmailOtpResponseDto, *IError) {

	var secondsToUnblock int64 = 30
	user, passport, err := GetUserByPassport2(*dto.Value)

	if err != nil {
		return nil, err
	}

	olderEntity := &ForgetPasswordEntity{}
	GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Find(olderEntity)

	if olderEntity.UniqueId != "" {

		fmt.Println(dto.Otp, olderEntity.Otp)
		// If user has provided the otp, create user session in fact.
		if dto.Otp != nil {

			if *dto.Otp == *olderEntity.Otp {
				session, err := PassportActionCreateSessionOtp(*passport.Value)

				if err != nil {
					return nil, GormErrorToIError(err)
				}

				// Delete the session so user cannot login again
				err2 := GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{}).Error

				if err2 != nil {
					return nil, GormErrorToIError(err)
				}

				return &EmailOtpResponseDto{
					UserSession: session,
				}, nil
			} else {
				return &EmailOtpResponseDto{
					Request: olderEntity,
				}, CreateIErrorString(PassportMessageCode.OtpCodeInvalid, []string{}, 403)
			}
		}

		if time.Now().Nanosecond() < int(olderEntity.BlockedUntil) {
			// olderEntity.SecondsToUnblock = int64(time.Until(olderEntity.BlockedUntil).Seconds())
			// @todo fix the seconds

			return &EmailOtpResponseDto{
					UserSession: nil,
					Request:     olderEntity,
				}, CreateIErrorString(
					PassportMessageCode.OTARequestBlockedUntil, []string{}, 403,
				)
		} else {
			// In this case, user has the chance to re-request. Let's clean the previous items first
			GetDbRef().Where(&ForgetPasswordEntity{PassportId: &passport.UniqueId}).Delete(&ForgetPasswordEntity{})
		}
	}

	{

		if err != nil {
			return nil, GormErrorToIError(err)
		}

		if passport == nil || user == nil || user.UniqueId == "" {
			return nil, CreateIErrorString(PassportMessageCode.UserDoesNotExist, []string{}, 403)
		}

		uid := UUID()

		otp := GenerateRandomKey(6)
		url := "http://localhost:8888/reset-password?session=" + uid
		item := &ForgetPasswordEntity{
			User:                user,
			Passport:            passport,
			UniqueId:            uid,
			ValidUntil:          time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			BlockedUntil:        time.Now().Add(time.Second * time.Duration(secondsToUnblock)).UnixNano(),
			SecondsToUnblock:    &secondsToUnblock,
			Otp:                 &otp,
			RecoveryAbsoluteUrl: &url,
		}

		err := GetDbRef().Create(item).Error

		if err != nil {
			return nil, GormErrorToIError(err)
		}

		// if dto.Type == "sms" {
		// 	err = SendSMSFromTemplate("forget-password.html", item)
		// 	if err != nil {
		// 		return nil, GormErrorToIError(err)
		// 	}

		// } else if dto.Type == "email" {
		// 	err = SendEmailFromTemplate("forget-password.html", item)
		// 	if err != nil {
		// 		return nil, GormErrorToIError(err)
		// 	}
		// }

		return &EmailOtpResponseDto{Request: item}, nil
	}
}

/**
*	This is the information we pass to the Email templates generally
 */

func SendSMSFromTemplate(templateName string, dto *ForgetPasswordEntity) error {
	fmt.Println("SMS needs to be implemented")
	return nil
}

// func SendEmailFromTemplate(templateName string, dto *ForgetPasswordEntity) error {
// 	var config = GetAppConfig()

// 	cfg := config.MailTemplates.ForgetPasswordRequest

// 	if !cfg.Enabled {
// 		fmt.Println("Email server is not available, so this won't work. config.MailTemplates.ForgetPasswordRequest.Enabled = false")
// 		return errors.New(PassportMessageCode.EmailIsNotConfigured)
// 	}

// 	t, err := template.ParseFS(templates.PassportEmailTemplates, templateName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var tpl bytes.Buffer
// 	err = t.Execute(&tpl, dto)
// 	if err != nil {
// 		panic(err)
// 	}

// 	result := tpl.String()

// 	return notification.SendMail(
// 		notification.EmailMessageContent{
// 			ToName:  dto.User.FirstName + " " + dto.User.LastName,
// 			ToEmail: dto.Passport.Value,
// 			Content: result,
// 			Subject: cfg.Subject,
// 		},
// 	)

// }

func GetResetPasswordInfo(requestUniqueId string) ForgetPasswordEntity {
	var entity ForgetPasswordEntity
	GetDbRef().Where(RealEscape("unique_id = ? and status != ?", requestUniqueId, "used")).First(&entity)

	return entity
}

func ResetUserPasswordWithRequest(requestUniqueId string, newPassword string) (bool, string) {

	var entity ForgetPasswordEntity
	GetDbRef().Preload("User").Preload("Passport").Where(RealEscape("unique_id = ? and status != ?", requestUniqueId, "used")).First(&entity)

	if entity.UniqueId == "" {
		return false, ""
	}

	passwordHashed, _ := HashPassword(newPassword)

	GetDbRef().Model(&PassportEntity{}).Where(RealEscape("unique_id = ?", entity.Passport.UniqueId)).Update("password", passwordHashed)

	GetDbRef().Model(&ForgetPasswordEntity{}).Where(RealEscape("unique_id = ?", requestUniqueId)).Update("status", "used")

	return true, *entity.Passport.Value
}

// func SigninUserWithGoogleClaim(user UserEntity, claim GoogleAuthClaim) (string, error) {

// 	fmt.Println(claim)
// 	tokenString := GenerateSecureToken(32)

// 	GetDbRef().Create(&Token{
// 		Hash:       tokenString,
// 		UserId:     user.UniqueId,
// 		UniqueId:   UUID(),
// 		ValidUntil: time.Now().Add(time.Hour * time.Duration(12)),
// 	})

// 	return tokenString, nil

// }

func ConfirmEmailAddress(confirmUniqueId string) bool {
	return GetDbRef().
		Model(&EmailConfirmationEntity{}).
		Where(RealEscape("unique_id = ? and status != ?", confirmUniqueId, "confirmed")).
		Update("status", "confirmed").RowsAffected > 0
}

func GetUserOnlyByMail(email string) (string, *UserEntity) {
	email = strings.ToLower(email)
	var passport PassportEntity
	var user UserEntity

	GetDbRef().Where(RealEscape("value = ?", email)).First(&passport)

	if passport.UniqueId == "" {
		return "", nil
	}

	GetDbRef().Where(RealEscape("unique_id = ?", *passport.UserId)).First(&user)

	if user.UniqueId == "" {
		return "", nil
	}

	return *passport.Password, &user
}

func GetUserPhoneNumber(phoneNumber string) (PassportEntity, UserEntity) {

	var passport PassportEntity
	var user UserEntity
	GetDbRef().Where(RealEscape("value = ?", phoneNumber)).First(&passport)
	GetDbRef().Where(RealEscape("unique_id = ?", *passport.UserId)).First(&user)

	return passport, user
}

func SendUserMailConfirmation(email string, user UserEntity) bool {

	var item EmailConfirmationEntity = EmailConfirmationEntity{
		User:     &user,
		UniqueId: UUID(),
	}

	err := GetDbRef().Create(&item)

	// SendUserAccountConfirmationMail(email, email, item.UniqueId)

	return err == nil
}

func PublicJoinKeyActionGetOnePublic(query QueryDSL) (*PublicJoinKeyEntity, *IError) {
	refl := reflect.ValueOf(&PublicJoinKeyEntity{})
	return GetOneEntity[PublicJoinKeyEntity](query, refl)
}
