package abac

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/torabian/fireback/modules/fireback"
)

/**
*	Does the authorization with the current logged in user on host
*	Usually, this makes sense for desktop/mobile apps which are having a light or complete
*	Version of the backend installed on them. Make sure this is not accessbile on the web version or cloud version.
**/
func PassportActionAuthorizeOs2(dto *fireback.EmptyRequest, query fireback.QueryDSL) (*UserSessionDto, *fireback.IError) {
	return SigninWithOsUser2(query)
}

// Implementation of generating token for specific user.
// Tokens belong to a user, and if they are removed from that workspace
// Token would not be any useful.
// Before creating each token, we are looking for existing token in the database
// and if it exists and is still valid for that specific user,
// we skip generating new one.
func (x *UserEntity) AuthorizeWithToken(q fireback.QueryDSL) (string, error) {
	appConfig := fireback.GetConfig()

	ref := fireback.GetRef(q)

	// generating token based on random hash, or jwt here can be decided.
	var tokenString string

	if appConfig.TokenGenerationStrategy == "jwt" {
		claims := jwt.MapClaims{

			"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if jwttoken, err := token.SignedString([]byte(appConfig.JwtSecretKey)); err != nil {
			tokenString = fireback.GenerateSecureToken(32)
		} else {
			tokenString = jwttoken
		}
	} else {
		tokenString = fireback.GenerateSecureToken(32)

	}

	q.ResolveStrategy = "user"
	tokens, _, err := TokenActions.Query(q)

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

		if t.After(time.Now()) && token.Token != "" {
			return token.Token, nil
		}
	}

	until := fireback.XDateTimeFromTime(time.Now().Add(time.Minute * time.Duration(2)))
	token := &TokenEntity{
		UniqueId:    fireback.UUID(),
		UserId:      fireback.NewString(x.UniqueId),
		Token:       tokenString,
		ValidUntil:  until,
		WorkspaceId: fireback.NewString(ROOT_VAR),
	}
	if err3 := ref.Create(token).Error; err3 != nil {
		return "", err3
	}

	return tokenString, nil
}
