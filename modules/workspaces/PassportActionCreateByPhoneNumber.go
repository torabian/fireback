package workspaces

import (
	"crypto/rand"
	"io"
	"time"
)

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func UserActionCreateByPhoneNumber(dto PhoneNumberAccountCreationDto, query QueryDSL) (string, *IError) {
	if iError := UserWithPhoneValidator(dto, false); iError != nil {
		return "", iError
	}

	u := UUID()
	code := EncodeToString(6)

	expire := time.Now().Local().Add(time.Minute * 3).String()
	prevalidate := "Prevalidate"
	var item PhoneConfirmationEntity = PhoneConfirmationEntity{
		ExpiresAt:   &expire,
		PhoneNumber: dto.PhoneNumber,
		Status:      &prevalidate,
		Key:         &code,
		UniqueId:    u,
	}

	err := GetDbRef().Create(&item).Error

	if err != nil {
		return "", GormErrorToIError(err)
	}

	return code, nil
}

func UserActionCreateByPhoneNumberConfirm(dto PhoneNumberAccountCreationDto, query QueryDSL) (UserSessionDto, *IError) {
	// var config = utils.GetAppConfig()

	session := UserSessionDto{}
	if iError := UserWithPhoneValidator(dto, false); iError != nil {
		return session, iError
	}

	user, passport, token, err := CreateUserWithPhoneNumber(*dto.PhoneNumber)
	// workspacesList, _ := GetUserWorkspaces(query)

	if err != nil {
		e := GormErrorToIError(err)
		return session, e
	}

	// if config.MailTemplates.ConfirmMail.Enabled {
	// 	SendUserSMSConfirmation(dto.PhoneNumber, user)
	// }

	session.Passport = &passport
	session.User = &user
	session.Token = &token
	// session.UserWorkspaces = workspacesList
	ek := PutTokenInExchangePool(token)
	session.ExchangeKey = &ek

	return session, nil
}

func CreateUserWithPhoneNumber(phoneNumber string) (UserEntity, PassportEntity, string, error) {
	// IMPORTANT: THIS IS BROKEN, DOES NOT HAVE QUERYDSL
	userUniqueId := UUID()
	passportUniqueId := UUID()

	var user UserEntity = UserEntity{
		UniqueId: userUniqueId,
	}

	err := GetDbRef().Create(&user).Error

	if err != nil {
		return UserEntity{}, PassportEntity{}, "", err
	}

	var passport = &PassportEntity{Value: &phoneNumber, Type: &PassportTypes.PhoneNumber, UserId: &userUniqueId, UniqueId: passportUniqueId}
	err = GetDbRef().Create(&passport).Error

	if err != nil {
		return UserEntity{}, PassportEntity{}, "", err
	}

	// 3. Create the workspace for him
	// 4. Assign his owner role
	// _, err = WorkspaceActionCreate()

	if err != nil {
		return UserEntity{}, PassportEntity{}, "", err
	}

	// 5. Log him in
	user, token, _ := SigninUserWithPhoneNumber(phoneNumber)

	return user, *passport, token, nil
}
