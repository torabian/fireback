package abac

import (
	"github.com/pquerna/otp/totp"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	ConfirmClassicPassportTotpActionImp = ConfirmClassicPassportTotpAction
}

func ConfirmClassicPassportTotpAction(
	req *ConfirmClassicPassportTotpActionReqDto,
	q fireback.QueryDSL) (*ConfirmClassicPassportTotpActionResDto,
	*fireback.IError,
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
		return nil, fireback.Create401Error(&AbacMessages.TotpIsNotAvailableForThisPassport, []string{})
	}

	if !totp.Validate(req.TotpCode, singinResult.Session.Passport.TotpSecret) {
		return nil, fireback.Create401Error(&AbacMessages.TotpCodeIsNotValid, []string{})
	}

	// Update the passport entity that it's confirmed
	if _, err := PassportActions.Update(fireback.QueryDSL{
		WorkspaceId: ROOT_VAR,
		UniqueId:    singinResult.Session.Passport.UniqueId,
	}, &PassportEntity{TotpConfirmed: fireback.NewBool(true), UniqueId: singinResult.Session.Passport.UniqueId}); err != nil {
		return nil, fireback.Create401Error(&AbacMessages.PassportTotpNotConfirmed, []string{})
	}

	// Implement the logic here.
	return &ConfirmClassicPassportTotpActionResDto{
		Session: singinResult.Session,
	}, nil
}
