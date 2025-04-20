package abac

import (
	"encoding/json"
	"strings"

	"github.com/pquerna/otp/totp"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	ClassicSignupActionImp = ClassicSignupAction
}

// Responsible for user creation from public flows in the application.
func ClassicSignupAction(dto *ClassicSignupActionReqDto, q fireback.QueryDSL) (*ClassicSignupActionResDto, *fireback.IError) {
	if err := ClassicSignupActionReqValidator(dto); err != nil {
		return nil, err
	}

	ClearPassportValue(&dto.Value)

	// Look for the configuration to check if the session secret is needed
	config, err := WorkspaceConfigActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR, Tx: q.Tx})
	if err != nil {
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	requiresSessionSecret := false
	if config != nil {
		if config.EnableRecaptcha2.Bool && config.Recaptcha2ServerKey != "" && config.Recaptcha2ClientKey != "" {
			requiresSessionSecret = true
		}
		if config.RequireOtpOnSignup.Bool {
			requiresSessionSecret = true
		}
	}

	var publicSession *PublicAuthenticationEntity = nil
	if requiresSessionSecret {
		if strings.TrimSpace(dto.SessionSecret) == "" {
			return nil, fireback.Create401Error(&AbacMessages.SessionSecretIsNeeded, []string{})
		}

		// Here we need to do some comparison to make sure this is the correct session secret
		fireback.GetDbRef().Where(&PublicAuthenticationEntity{SessionSecret: dto.SessionSecret}).Find(&publicSession)

		if strings.TrimSpace(dto.SessionSecret) == "" {
			return nil, fireback.Create401Error(&AbacMessages.SessionSecretIsNotAvailable, []string{})
		}
	}

	return completeClassicSignupProcess(dto, q, publicSession, config, nil)
}

// This function will complete the signup process.
// the reason it's excluded from action body is, it can be reused for the account creation
// via oauth. Just we pass nil for publicSession and config.
// In case we neeed to change the config slightly, you can get the config
// and change it for this specific function
func completeClassicSignupProcess(
	dto *ClassicSignupActionReqDto,
	q fireback.QueryDSL,
	publicSession *PublicAuthenticationEntity,
	config *WorkspaceConfigEntity,
	beforeProcess func(*UserEntity, *RoleEntity, *WorkspaceEntity, *PassportEntity),
) (*ClassicSignupActionResDto, *fireback.IError) {

	user, role, workspace, passport := GetEmailPassportSignupMechanism(dto)

	// A callback to modify things.
	if beforeProcess != nil {
		beforeProcess(user, role, workspace, passport)
	}

	totpLink := ""
	if publicSession != nil && publicSession.TotpLink != "" {
		totpLink = publicSession.TotpLink
	}

	if config != nil {
		if config.EnableTotp.Bool || config.ForceTotp.Bool {
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

	forcedTotp := config != nil && config.ForceTotp.Bool
	// let's check for totp setup, if the session is successful.
	if config != nil && config.ForceTotp.Bool && session != nil {
		return &ClassicSignupActionResDto{
			ContinueToTotp: true,
			TotpUrl:        totpLink,
			ForcedTotp:     forcedTotp,
		}, nil
	}

	// Clear the value so next time user can login directly
	if sessionError == nil && session != nil && passport != nil && passport.Value != "" {
		fireback.GetRef(q).Where(&PublicAuthenticationEntity{PassportValue: passport.Value}).Delete(&PublicAuthenticationEntity{})
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
