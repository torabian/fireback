package workspaces

import (
	"errors"
	"time"
)

// @unsafe - only internal call
func PassportActionCreateSessionOtp(email string) (*UserSessionDto, *IError) {

	session := &UserSessionDto{}

	user, token, err := SigninUserWithEmail(email)

	if err != nil {
		e := GormErrorToIError(err)
		return session, e
	}

	// workspacesList := GetUserWorkspaces(user.UniqueId)
	session.User = user
	session.Token = &token
	// session.UserRoleWorkspaces = workspacesList
	ev := PutTokenInExchangePool(token)
	session.ExchangeKey = &ev

	return session, nil
}

/**
*	Does the authorization with the current logged in user on host
*	Usually, this makes sense for desktop/mobile apps which are having a light or complete
*	Version of the backend installed on them. Make sure this is not accessbile on the web version or cloud version.
**/
func PassportActionAuthorizeOs2(dto *EmptyRequest, query QueryDSL) (*UserSessionDto, *IError) {
	return SigninWithOsUser2(query)
}

func PassportActionEmailSignin(dto *EmailAccountSigninDto, query QueryDSL) (*UserSessionDto, *IError) {

	session := &UserSessionDto{}
	if iError := UserSigninEmailAndPasswordValidator(dto, false); iError != nil {
		return session, iError
	}

	user, token, err := SigninUserWithEmailAndPassword(*dto.Email, *dto.Password)

	if err != nil {
		e := GormErrorToIError(err)
		return session, e
	}

	// workspacesList := GetUserWorkspaces(user.UniqueId)
	session.User = user
	session.Token = &token
	// session.UserRoleWorkspaces = workspacesList
	ev := PutTokenInExchangePool(token)
	session.ExchangeKey = &ev

	return session, nil
}

func SigninUserWithEmailAndPassword(email string, password string) (*UserEntity, string, error) {

	hash, User := GetUserOnlyByMail(email)

	if User == nil || hash == "" {
		return nil, "", errors.New(PassportMessageCode.UserDoesNotExist)
	}

	if CheckPasswordHash(password, hash) {
		tokenString := GenerateSecureToken(32)

		until := time.Now().Add(time.Hour * time.Duration(12)).String()
		GetDbRef().Create(&TokenEntity{
			UniqueId:   tokenString,
			UserId:     &User.UniqueId,
			ValidUntil: &until,
		})

		return User, tokenString, nil
	}

	return User, "", errors.New(PassportMessageCode.UserDoesNotExist)
}

// @unsafe - only internal calls
func SigninUserWithEmail(email string) (*UserEntity, string, error) {

	hash, User := GetUserOnlyByMail(email)

	if hash == "" {
		return &UserEntity{}, "", errors.New(PassportMessageCode.UserDoesNotExist)
	}

	tokenString := GenerateSecureToken(32)
	until := time.Now().Add(time.Hour * time.Duration(12)).String()
	GetDbRef().Create(&TokenEntity{
		UniqueId:   tokenString,
		UserId:     &User.UniqueId,
		ValidUntil: &until,
	})

	return User, tokenString, nil
}
