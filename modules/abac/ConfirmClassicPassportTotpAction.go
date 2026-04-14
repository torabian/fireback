package abac

import (
	"github.com/pquerna/otp/totp"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	ConfirmClassicPassportTotpImpl = ConfirmClassicPassportTotpAction
}

func ConfirmClassicPassportTotpAction(c ConfirmClassicPassportTotpActionRequest, q fireback.QueryDSL) (*ConfirmClassicPassportTotpActionResponse, error) {
	req := c.Body
	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	singinResult, signinError := classicSinginInternalUnsafe(&ClassicSigninActionReq{
		Value:    req.Value,
		Password: req.Password,
	}, q)

	if signinError != nil {
		return nil, signinError
	}

	passport, _ := singinResult.Session.Passport.Get()
	if passport.TotpSecret == "" {
		return nil, fireback.Create401Error(&AbacMessages.TotpIsNotAvailableForThisPassport, []string{})
	}

	if !totp.Validate(req.TotpCode, passport.TotpSecret) {
		return nil, fireback.Create401Error(&AbacMessages.TotpCodeIsNotValid, []string{})
	}

	// Update the passport entity that it's confirmed
	if _, err := PassportActions.Update(fireback.QueryDSL{
		WorkspaceId: ROOT_VAR,
		UniqueId:    passport.UniqueId,
	}, &PassportEntity{TotpConfirmed: fireback.NewBool(true), UniqueId: passport.UniqueId}); err != nil {
		return nil, fireback.Create401Error(&AbacMessages.PassportTotpNotConfirmed, []string{})
	}

	// Implement the logic here.
	return &ConfirmClassicPassportTotpActionResponse{
		Payload: ConfirmClassicPassportTotpActionRes{
			Session: singinResult.Session,
		},
	}, nil
}
