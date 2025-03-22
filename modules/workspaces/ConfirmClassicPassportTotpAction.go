package workspaces

import "github.com/pquerna/otp/totp"

func init() {
	// Override the implementation with our actual code.
	ConfirmClassicPassportTotpActionImp = ConfirmClassicPassportTotpAction
}

func ConfirmClassicPassportTotpAction(
	req *ConfirmClassicPassportTotpActionReqDto,
	q QueryDSL) (*ConfirmClassicPassportTotpActionResDto,
	*IError,
) {
	if err := ConfirmClassicPassportTotpActionReqValidator(req); err != nil {
		return nil, err
	}

	singinResult, signinError := classicSinginInternalUnsafe(&ClassicSigninActionReqDto{
		Value:    req.Value,
		Password: req.Password,
	}, q)

	if signinError != nil {
		return nil, signinError
	}

	if singinResult.Session.Passport.TotpSecret == "" {
		return nil, Create401Error(&WorkspacesMessages.TotpIsNotAvailableForThisPassport, []string{})
	}

	if !totp.Validate(req.TotpCode, singinResult.Session.Passport.TotpSecret) {
		return nil, Create401Error(&WorkspacesMessages.TotpCodeIsNotValid, []string{})
	}

	// Update the passport entity that it's confirmed
	if _, err := PassportActions.Update(QueryDSL{
		WorkspaceId: ROOT_VAR,
		UniqueId:    singinResult.Session.Passport.UniqueId,
	}, &PassportEntity{TotpConfirmed: NewBool(true), UniqueId: singinResult.Session.Passport.UniqueId}); err != nil {
		return nil, Create401Error(&WorkspacesMessages.PassportTotpNotConfirmed, []string{})
	}

	// Implement the logic here.
	return &ConfirmClassicPassportTotpActionResDto{
		Session: singinResult.Session,
	}, nil
}
