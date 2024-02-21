package workspaces

import (
	"time"
)

func PassportActionCreateSessionByPhoneNumber(dto PhoneNumberAccountCreationDto, query QueryDSL) (UserSessionDto, *IError) {

	session := UserSessionDto{}

	user, token, err := SigninUserWithPhoneNumber(*dto.PhoneNumber)

	if err != nil {
		e := GormErrorToIError(err)
		return session, e
	}
	// workspacesList := GetUserWorkspaces(user.UniqueId)

	session.User = &user
	session.Token = &token
	// session.UserRoleWorkspaces = workspacesList
	ev := PutTokenInExchangePool(token)
	session.ExchangeKey = &ev

	return session, nil
}

func SigninUserWithPhoneNumber(phoneNumber string) (UserEntity, string, error) {

	_, User := GetUserPhoneNumber(phoneNumber)

	tokenString := GenerateSecureToken(32)

	until := time.Now().Add(time.Hour * time.Duration(12)).String()
	GetDbRef().Create(&TokenEntity{
		UserId:     &User.UniqueId,
		UniqueId:   tokenString,
		ValidUntil: &until,
	})

	return User, tokenString, nil

}
