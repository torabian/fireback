package abac

import (
	"strings"

	"github.com/pquerna/otp/totp"
	"github.com/torabian/fireback/modules/workspaces"
)

func init() {
	ClassicSigninActionImp = ClassicSigninAction
}

func ClassicSigninAction(req *ClassicSigninActionReqDto, q workspaces.QueryDSL) (*ClassicSigninActionResDto, *workspaces.IError) {

	if err := ClassicSigninActionReqValidator(req); err != nil {
		return nil, err
	}

	config, err2 := WorkspaceConfigActions.GetByWorkspace(workspaces.QueryDSL{WorkspaceId: ROOT_VAR, Tx: q.Tx})
	if err2 != nil {
		if err2.HttpCode != 404 {
			return nil, err2
		}
	}

	requiresSessionSecret := false
	if config != nil {

		if config.EnableRecaptcha2.Bool && config.Recaptcha2ServerKey != "" && config.Recaptcha2ClientKey != "" {
			requiresSessionSecret = true
		}
		if config.RequireOtpOnSignin.Bool {
			requiresSessionSecret = true
		}
	}

	if requiresSessionSecret {
		if strings.TrimSpace(req.SessionSecret) == "" {
			return nil, workspaces.Create401Error(&AbacMessages.SessionSecretIsNeeded, []string{})
		}

		// Here we need to do some comparison to make sure this is the correct session secret
		var publicSession *PublicAuthenticationEntity = nil
		workspaces.GetDbRef().Where(&PublicAuthenticationEntity{SessionSecret: req.SessionSecret}).Find(&publicSession)

		if strings.TrimSpace(req.SessionSecret) == "" {
			return nil, workspaces.Create401Error(&AbacMessages.SessionSecretIsNotAvailable, []string{})
		}
	}

	session := &UserSessionDto{}

	if err := fetchPureUserAndPassToSession(req.Value, req.Password, session, q); err != nil {
		return nil, err
	}

	// if user doesn't have totp setup, then move him
	if config != nil && config.ForceTotp.Bool {
		if session.Passport.TotpSecret == "" ||
			!session.Passport.TotpConfirmed.Bool {

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

	if session.Passport.TotpSecret != "" && config != nil && config.EnableTotp.Bool {
		// Assume this is first time, so do not fail the response and allow user to go there.
		if req.TotpCode == "" {
			return &ClassicSigninActionResDto{
				Next: []string{"enter-totp"},
			}, nil
		}

		if !totp.Validate(req.TotpCode, session.Passport.TotpSecret) {
			return nil, workspaces.Create401Error(&AbacMessages.TotpCodeIsNotValid, []string{})
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
func classicSinginInternalUnsafe(req *ClassicSigninActionReqDto, q workspaces.QueryDSL) (*ClassicSigninActionResDto, *workspaces.IError) {

	session := &UserSessionDto{}

	fetchPureUserAndPassToSession(req.Value, req.Password, session, q)
	applyUserTokenAndWorkspacesToToken(session, q)

	return &ClassicSigninActionResDto{
		Session: session,
	}, nil
}

func applyUserTokenAndWorkspacesToToken(session *UserSessionDto, q workspaces.QueryDSL) *workspaces.IError {
	// Get the user workspaces as well
	q.UserId = session.User.UniqueId
	q.ResolveStrategy = "user"
	workspacesItems, _, err := UserWorkspaceActions.Query(q)
	if err != nil {
		return workspaces.GormErrorToIError(err)
	}
	session.UserWorkspaces = workspacesItems

	// Authorize the session, put the token
	if token, err := session.User.AuthorizeWithToken(q); err != nil {
		return workspaces.CastToIError(err)
	} else {
		session.Token = token
	}

	return nil
}

// Gets the user via the passport value.
// This is an unsafe function and should not be exposed to outside.
// If the password is nil, it means it would work without a password.
// So make sure you have
func fetchPureUserAndPassToSession(value string, password string, session *UserSessionDto, q workspaces.QueryDSL) *workspaces.IError {

	passportPassword, err := fetchUserAndPassToSession(value, session, q)

	if err != nil {
		return err
	}

	if !workspaces.CheckPasswordHash(password, *passportPassword) {
		return workspaces.Create401Error(&AbacMessages.PassportNotAvailable, []string{})
	}

	return nil
}

// Unsafe function, which reads a user passport and finds him and assigns to the
// session. Just use in password less scenarios, such as oauth.
func fetchUserAndPassToSession(value string, session *UserSessionDto, q workspaces.QueryDSL) (*string, *workspaces.IError) {
	ClearPassportValue(&value)

	var passportPassword = ""
	if passport, user, err := UnsafeGetUserByPassportValue(value, q); err != nil {
		return nil, err
	} else {
		session.User = user
		session.Passport = passport
		passportPassword = passport.Password
	}

	if session.User == nil {
		return nil, workspaces.Create401Error(&AbacMessages.PassportNotAvailable, []string{})
	}

	return &passportPassword, nil
}

func UnsafeGetUserByPassportValue(value string, q workspaces.QueryDSL) (*PassportEntity, *UserEntity, *workspaces.IError) {

	// Check the passport if exists
	var item PassportEntity
	if err := workspaces.GetRef(q).Model(&PassportEntity{}).Where(&PassportEntity{Value: value}).First(&item).Error; err != nil || item.Value == "" {

		return nil, nil, workspaces.Create401Error(&AbacMessages.PassportNotAvailable, []string{})
	}

	var user UserEntity
	if err := workspaces.GetRef(q).Model(&UserEntity{}).Where(&UserEntity{UniqueId: item.UserId.String}).First(&user).Error; err != nil {
		return nil, nil, workspaces.Create401Error(&AbacMessages.PassportNotAvailable, []string{})
	}

	return &item, &user, nil
}

// Delete the spaces in the email and make it lower case
// before any operation
func ClearPassportValue(str *string) {
	v := strings.ToLower(strings.TrimSpace(*str))
	*str = v
}
