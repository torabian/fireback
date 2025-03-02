package workspaces

import (
	"encoding/json"
	"strings"

	"github.com/pquerna/otp/totp"
)

func init() {
	ClassicSignupActionImp = ClassicSignupAction
}

// Responsible for user creation from public flows in the application.
func ClassicSignupAction(dto *ClassicSignupActionReqDto, q QueryDSL) (*ClassicSignupActionResDto, *IError) {
	if err := ClassicSignupActionReqValidator(dto); err != nil {
		return nil, err
	}

	ClearShot(&dto.Value)

	// Look for the configuration to check if the session secret is needed
	config, err := WorkspaceConfigActions.GetByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR, Tx: q.Tx})
	if err != nil {
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	requiresSessionSecret := false
	if config != nil {
		if config.EnableRecaptcha2 {
			requiresSessionSecret = true
		}
		if config.RequireOtpOnSignup {
			requiresSessionSecret = true
		}
	}

	var publicSession *PublicAuthenticationEntity = nil
	if requiresSessionSecret {
		if strings.TrimSpace(dto.SessionSecret) == "" {
			return nil, Create401Error(&WorkspacesMessages.SessionSecretIsNeeded, []string{})
		}

		// Here we need to do some comparison to make sure this is the correct session secret
		GetDbRef().Where(&PublicAuthenticationEntity{SessionSecret: dto.SessionSecret}).Find(&publicSession)

		if strings.TrimSpace(dto.SessionSecret) == "" {
			return nil, Create401Error(&WorkspacesMessages.SessionSecretIsNotAvailable, []string{})
		}
	}

	user, role, workspace, passport := GetEmailPassportSignupMechanism(dto)

	totpLink := ""
	if publicSession != nil && publicSession.TotpLink != "" {
		totpLink = publicSession.TotpLink
	}

	if config != nil {
		if config.EnableTotp || config.ForceTotp {
			// add time based dual factor information
			key, _ := totp.Generate(totp.GenerateOpts{
				Issuer:      "Fireback",
				AccountName: passport.Value,
			})
			secret := key.Secret()
			link := key.URL()
			passport.TotpSecret = secret
			totpLink = link
		}
	}

	session, sessionError := UnsafeGenerateUser(&GenerateUserDto{

		createUser:      true,
		createWorkspace: true,
		createRole:      true,
		createPassport:  true,

		user:      user,
		role:      role,
		workspace: workspace,
		passport:  passport,

		// We want always to be able to login regardless
		restricted: true,
	}, q)

	if sessionError != nil {
		return nil, sessionError
	}

	forcedTotp := config != nil && config.ForceTotp
	// let's check for totp setup, if the session is successful.
	if config != nil && config.ForceTotp && session != nil {
		return &ClassicSignupActionResDto{
			ContinueToTotp: true,
			TotpUrl:        totpLink,
			ForcedTotp:     forcedTotp,
		}, nil
	}

	// Clear the value so next time user can login directly
	if sessionError == nil && session != nil && passport != nil && passport.Value != "" {
		GetRef(q).Where(&PublicAuthenticationEntity{PassportValue: passport.Value}).Delete(&PublicAuthenticationEntity{})
	}

	return &ClassicSignupActionResDto{
		Session:        session,
		ContinueToTotp: false,
		ForcedTotp:     forcedTotp,
	}, sessionError
}

func (x *ClassicSignupActionResDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
