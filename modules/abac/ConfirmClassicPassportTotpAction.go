package abac

import (
	"github.com/pquerna/otp/totp"
	"github.com/torabian/fireback/modules/workspaces"
)

func init() {
	// Override the implementation with our actual code.
	ConfirmClassicPassportTotpActionImp = ConfirmClassicPassportTotpAction
}

func ConfirmClassicPassportTotpAction(
	req *ConfirmClassicPassportTotpActionReqDto,
	q workspaces.QueryDSL) (*ConfirmClassicPassportTotpActionResDto,
	*workspaces.IError,
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
		return nil, workspaces.Create401Error(&AbacMessages.TotpIsNotAvailableForThisPassport, []string{})
	}

	if !totp.Validate(req.TotpCode, singinResult.Session.Passport.TotpSecret) {
		return nil, workspaces.Create401Error(&AbacMessages.TotpCodeIsNotValid, []string{})
	}

	// Update the passport entity that it's confirmed
	if _, err := PassportActions.Update(workspaces.QueryDSL{
		WorkspaceId: ROOT_VAR,
		UniqueId:    singinResult.Session.Passport.UniqueId,
	}, &PassportEntity{TotpConfirmed: workspaces.NewBool(true), UniqueId: singinResult.Session.Passport.UniqueId}); err != nil {
		return nil, workspaces.Create401Error(&AbacMessages.PassportTotpNotConfirmed, []string{})
	}

	// Implement the logic here.
	return &ConfirmClassicPassportTotpActionResDto{
		Session: singinResult.Session,
	}, nil
}
