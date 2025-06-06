package abac

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	CheckClassicPassportActionImp = CheckClassicPassportAction
}

func CheckClassicPassportAction(req *CheckClassicPassportActionReqDto, q fireback.QueryDSL) (*CheckClassicPassportActionResDto, *fireback.IError) {

	if err := validateValueFormat(req.Value); err != nil {
		return nil, err
	}

	// get the configuration, also check for the recaptcha or captcha service here.
	config, prepareErr := prepareTheClassicPassport(req, q)
	if prepareErr != nil {
		return nil, prepareErr
	}

	passport := findPassport(req.Value, q)

	// from here we devide the work flow to existing and non exists passport
	if passport == nil {
		return checkStepsForNonExistingAccount(req.Value, config, q)
	} else {
		return checkStepsForExistingAccount(passport, config, q)
	}
}

// in some operations, the only option is otp either on signin or signup.
// so we send the otp anyway, and next step can be immediately signup.
func implicitlyRequestForOtp(passportValue string, q fireback.QueryDSL) (*CheckClassicPassportResDtoOtpInfo, *fireback.IError) {
	otpInfo, otpFailed := ClassicPassportRequestOtpAction(&ClassicPassportRequestOtpActionReqDto{Value: passportValue}, q)

	// No point of continuing if the type doesn't support otp
	if otpFailed != nil {
		if item := otpFailed.Message["$"]; item == string(OtpNotAvailableForThisType) {
			return nil, otpFailed
		}
	}

	if otpInfo != nil {

		// if request is blocked, we actually did not sent the otp.
		if otpFailed != nil {
			if otpFailed.Message["$"] == string(OtaRequestBlockedUntil) {
				return nil, otpFailed
			}
		}

		return &CheckClassicPassportResDtoOtpInfo{
			SuspendUntil:     otpInfo.SuspendUntil,
			ValidUntil:       otpInfo.ValidUntil,
			BlockedUntil:     otpInfo.BlockedUntil,
			SecondsToUnblock: otpInfo.SecondsToUnblock,
		}, nil
	}

	return nil, otpFailed
}

func findPassport(value string, q fireback.QueryDSL) *PassportEntity {
	var passport *PassportEntity
	if err := fireback.GetRef(q).
		Model(&PassportEntity{}).Where(&PassportEntity{Value: value}).
		First(&passport).Error; err == nil && passport.Value != "" {
		if passport.Value == value {
			return passport
		}
	}
	return nil
}

func validateValueFormat(value string) *fireback.IError {
	if validx, typeof := validatePassportType(value); !validx {
		if typeof == PASSPORT_METHOD_EMAIL {
			return fireback.Create401Error(&AbacMessages.EmailIsNotValid, []string{})
		}
		if typeof == PASSPORT_METHOD_PHONE {
			return fireback.Create401Error(&AbacMessages.PhoneNumberIsNotValid, []string{})
		}
	}

	return nil
}

// Some general tests and preparation which doesn't affect logic much
func prepareTheClassicPassport(req *CheckClassicPassportActionReqDto, q fireback.QueryDSL) (*WorkspaceConfigEntity, *fireback.IError) {
	if err := CheckClassicPassportActionReqValidator(req); err != nil {
		return nil, err
	}

	ClearPassportValue(&req.Value)

	config, err := WorkspaceConfigActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})
	if err != nil {
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	// If recaptcha 2 is needed
	if config != nil && config.EnableRecaptcha2.Bool && config.Recaptcha2ServerKey != "" && config.Recaptcha2ClientKey != "" {
		if req.SecurityToken == "" {
			return nil, &fireback.IError{Message: AbacMessages.Recaptcha2Needed}
		}
		if err := validateRecaptcha(req.SecurityToken, config.Recaptcha2ServerKey); err != nil {
			return nil, &fireback.IError{Message: AbacMessages.Recaptcha2Error}
		}
	}

	return config, nil
}

func validateRecaptcha(token string, RECAPTCHA_SECRET_KEY string) error {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {RECAPTCHA_SECRET_KEY}, "response": {token}})

	if err != nil {
		return errors.New("failed to connect to reCAPTCHA service")
	}
	defer resp.Body.Close()

	var googleResp struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil || !googleResp.Success {
		return errors.New("captcha verification failed")
	}

	return nil
}

// checks if value is email or phone number
func validatePassportType(input string) (bool, string) {
	// Phone: Only numbers and optional leading +
	phoneRegex := regexp.MustCompile(`^\+?[0-9]+$`)

	if phoneRegex.MatchString(input) {
		return len(input) >= 4 && len(input) <= 40, "phone"
	}

	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return len(input) <= 80 && emailRegex.MatchString(input), "email"
}

// used when a passport (email/phone) exists and they want to authenticate
func checkStepsForExistingAccount(passport *PassportEntity, config *WorkspaceConfigEntity, q fireback.QueryDSL) (*CheckClassicPassportActionResDto, *fireback.IError) {
	res := &CheckClassicPassportActionResDto{}

	// if otp is forced, then user can only authenticate using otp.
	// basically password and 2FA for signin will become useless, because the otp
	// will be used to reset user access anyway.
	envForcedOtp := config != nil && config.RequireOtpOnSignin.Bool
	if envForcedOtp {
		res.Next = []string{"otp"}
		res.OtpInfo, _ = implicitlyRequestForOtp(passport.Value, q)
		return res, nil
	}

	// let's check the passport configuration first.
	userHasPassword := passport.Password != ""

	// time based dual factor authentication
	userHasTotp := passport.TotpSecret != ""

	// check if otp is enabled, then we give the user 2 choices, either join with password
	// or join with password.
	envEnabledOtp := config != nil && config.EnableOtp.Bool
	res.Next = []string{}

	if envEnabledOtp {
		res.Next = append(res.Next, "otp")
	}

	if userHasPassword {
		res.Next = append(res.Next, "signin-with-password")

		// not sure if we have to expose this information before user provided the password,
		// but we are providing that user already exists, seems no harm to me expose this info as well
		if userHasTotp {
			res.Flags = append(res.Flags, "totp")
		}
	}

	// if the otp is only option, then we send the otp request implicitly to help the next steps on ui
	// be directly showing form to fullfill the otp and signin.
	if envEnabledOtp && !userHasPassword {
		res.OtpInfo, _ = implicitlyRequestForOtp(passport.Value, q)
	}

	return res, nil
}

func checkStepsForNonExistingAccount(value string, config *WorkspaceConfigEntity, q fireback.QueryDSL) (*CheckClassicPassportActionResDto, *fireback.IError) {
	res := &CheckClassicPassportActionResDto{}

	enableTotp := config != nil && config.EnableTotp.Bool
	forceTotp := config != nil && config.ForceTotp.Bool
	if enableTotp {
		res.Flags = append(res.Flags, "enable-totp")
	}
	if forceTotp {
		res.Flags = append(res.Flags, "force-totp")
	}

	// if environment has forced the otp, then there is no option other than
	// this condition has higher priority and needs to be checked first
	// so it won't expose existing users for setups that they do not want to
	// reveal that.
	envForcedOtp := config != nil && config.RequireOtpOnSignup.Bool
	if envForcedOtp {
		res.Next = []string{"otp"}
		info, errMsg := implicitlyRequestForOtp(value, q)
		res.OtpInfo = info

		// since the otp is only option, and if it has been failed then we should tell client
		if res.OtpInfo == nil {
			log.Default().Println("Failed to send otp:", errMsg)
			return nil, fireback.Create401Error(&AbacMessages.OtpFailed, []string{})
		}

		return res, nil
	}

	// check if otp is enabled, then we give the user 2 choices, either join with password
	// or join with password.
	envEnabledOtp := config != nil && config.EnableOtp.Bool
	if envEnabledOtp {
		res.Next = []string{"otp", "create-with-password"}
		return res, nil
	}

	// if otp is not available at all, the only option is to set password and create account
	res.Next = []string{"create-with-password"}
	return res, nil
}
