package workspaces

import (
	"strings"

	"github.com/pquerna/otp/totp"
)

func init() {
	ClassicSigninActionImp = ClassicSigninAction
}

func ClassicSigninAction(req *ClassicSigninActionReqDto, q QueryDSL) (*ClassicSigninActionResDto, *IError) {

	if err := ClassicSigninActionReqValidator(req); err != nil {
		return nil, err
	}

	config, err2 := WorkspaceConfigActions.GetByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR, Tx: q.Tx})
	if err2 != nil {
		if err2.HttpCode != 404 {
			return nil, err2
		}
	}

	requiresSessionSecret := false
	if config != nil {

		if config.EnableRecaptcha2 {
			requiresSessionSecret = true
		}
		if config.RequireOtpOnSignin {
			requiresSessionSecret = true
		}
	}

	if requiresSessionSecret {
		if strings.TrimSpace(req.SessionSecret) == "" {
			return nil, Create401Error(&WorkspacesMessages.SessionSecretIsNeeded, []string{})
		}

		// Here we need to do some comparison to make sure this is the correct session secret
		var publicSession *PublicAuthenticationEntity = nil
		GetDbRef().Where(&PublicAuthenticationEntity{SessionSecret: req.SessionSecret}).Find(&publicSession)

		if strings.TrimSpace(req.SessionSecret) == "" {
			return nil, Create401Error(&WorkspacesMessages.SessionSecretIsNotAvailable, []string{})
		}
	}

	session := &UserSessionDto{}

	if err := fetchPureUserAndPassToSession(req.Value, req.Password, session, q); err != nil {
		return nil, err
	}

	// if user doesn't have totp setup, then move him
	if config != nil && config.ForceTotp {
		if session.Passport.TotpSecret == "" ||
			!session.Passport.TotpConfirmed {

			// Let's create and assign to passport
			key, _ := totp.Generate(totp.GenerateOpts{
				Issuer:      "Fireback",
				AccountName: req.Value,
			})

			totpSecret := key.Secret()
			totpLink := key.URL()

			if _, err := PassportActions.Update(q, &PassportEntity{
				UniqueId:   session.Passport.UniqueId,
				TotpSecret: totpSecret,
			}); err != nil {
				return nil, err
			}

			return &ClassicSigninActionResDto{
				TotpUrl: totpLink,
				Next:    []string{"setup-totp"},
			}, nil
		}
	}

	if session.Passport.TotpSecret != "" {
		// Assume this is first time, so do not fail the response and allow user to go there.
		if req.TotpCode == "" {
			return &ClassicSigninActionResDto{
				Next: []string{"enter-totp"},
			}, nil
		}

		if !totp.Validate(req.TotpCode, session.Passport.TotpSecret) {
			return nil, Create401Error(&WorkspacesMessages.TotpCodeIsNotValid, []string{})
		}
	}

	if err := applyUserTokenAndWorkspacesToToken(session, q); err != nil {
		return nil, err
	}

	return &ClassicSigninActionResDto{
		Session: session,
	}, nil
}

// Can be used to authenticate only using value and passport.
// Do not expose this publicly, by passes recaptcha and all other securities.
func classicSinginInternalUnsafe(req *ClassicSigninActionReqDto, q QueryDSL) (*ClassicSigninActionResDto, *IError) {

	session := &UserSessionDto{}

	fetchPureUserAndPassToSession(req.Value, req.Password, session, q)
	applyUserTokenAndWorkspacesToToken(session, q)

	return &ClassicSigninActionResDto{
		Session: session,
	}, nil
}

func applyUserTokenAndWorkspacesToToken(session *UserSessionDto, q QueryDSL) *IError {
	// Get the user workspaces as well
	q.UserId = session.User.UniqueId
	q.ResolveStrategy = "user"
	workspaces, _, err := UserWorkspaceActions.Query(q)
	if err != nil {
		return GormErrorToIError(err)
	}
	session.UserWorkspaces = workspaces

	// Authorize the session, put the token
	if token, err := session.User.AuthorizeWithToken(q); err != nil {
		return CastToIError(err)
	} else {
		session.Token = token
	}

	return nil
}

func fetchPureUserAndPassToSession(value string, password string, session *UserSessionDto, q QueryDSL) *IError {

	ClearShot(&value)

	var passportPassword = ""
	if passport, user, err := UnsafeGetUserByPassportValue(value, q); err != nil {
		return err
	} else {
		session.User = user
		session.Passport = passport
		passportPassword = passport.Password
	}

	if !CheckPasswordHash(password, passportPassword) {
		return Create401Error(&WorkspacesMessages.PassportNotAvailable, []string{})
	}

	if session.User == nil {
		return Create401Error(&WorkspacesMessages.PassportNotAvailable, []string{})
	}

	return nil
}
