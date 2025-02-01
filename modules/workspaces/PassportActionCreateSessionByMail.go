package workspaces

import (
	"fmt"
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

// @unsafe - only internal calls
func SigninUserWithEmail(email string) (*UserEntity, string, error) {

	hash, User := GetUserOnlyByMail(email)

	if hash == "" {
		return &UserEntity{}, "", Create401Error(&WorkspacesMessages.UserDoesNotExist, []string{})
	}

	tokenString := GenerateSecureToken(32)
	// until := time.Now().Add(time.Hour * time.Duration(12)).String()

	GetDbRef().Create(&TokenEntity{
		UniqueId: tokenString,
		UserId:   &User.UniqueId,
		// ValidUntil:  &until,
		WorkspaceId: &ROOT_VAR,
	})

	return User, tokenString, nil
}

// Implementation of generating token for specific user.
// Tokens belong to a user, and if they are removed from that workspace
// Token would not be any useful.
// Before creating each token, we are looking for existing token in the database
// and if it exists and is still valid for that specific user,
// we skip generating new one.
func (x *UserEntity) AuthorizeWithToken(q QueryDSL) (string, error) {

	ref := GetRef(q)
	tokenString := GenerateSecureToken(32)
	q.ResolveStrategy = "user"
	tokens, _, err := TokenActionQuery(q)

	if err != nil {
		return "", err
	}

	for _, token := range tokens {
		if token.ValidUntil == nil {
			continue
		}

		t, err := token.ValidUntil.GetTime()
		if err != nil {
			continue
		}

		if t.After(time.Now()) {
			fmt.Println("Found a token which is still available", token.UniqueId)
			return token.UniqueId, nil
		}
	}

	until := XDateTimeFromTime(time.Now().Add(time.Minute * time.Duration(2)))
	token := &TokenEntity{
		UniqueId:    tokenString,
		UserId:      &x.UniqueId,
		ValidUntil:  until,
		WorkspaceId: &ROOT_VAR,
	}
	if err3 := ref.Create(token).Error; err3 != nil {
		return "", err3
	}

	return tokenString, nil
}
