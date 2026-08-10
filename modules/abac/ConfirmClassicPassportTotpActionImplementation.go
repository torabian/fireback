package abac

import (
	"github.com/pquerna/otp/totp"
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func ConfirmClassicPassportTotpAction(c abacdefs.ConfirmClassicPassportTotpActionRequest) (*abacdefs.ConfirmClassicPassportTotpActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	res, err2 := confirmClassicPassportTotpCore(c.Body, *query)
	if err2 != nil {
		return nil, err2
	}

	return &abacdefs.ConfirmClassicPassportTotpActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

// confirmClassicPassportTotpCore holds the actual implementation, reusable by callers
// which already have a resolved QueryDSL (such as the cli-only AuthFlow).
func confirmClassicPassportTotpCore(req abacdefs.ConfirmClassicPassportTotpActionReq, q fireback.QueryDSL) (*abacdefs.ConfirmClassicPassportTotpActionRes, *fireback.IError) {
	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	singinResult, signinError := classicSinginInternalUnsafe(&abacdefs.ClassicSigninActionReq{
		Value:    req.Value,
		Password: req.Password,
	}, q)

	if signinError != nil {
		return nil, signinError
	}

	passport, _ := singinResult.Session.Item.Passport.Get()
	if passport.Item.TotpSecret == "" {
		return nil, fireback.Create401Error(&AbacMessages.TotpIsNotAvailableForThisPassport, []string{})
	}

	if !totp.Validate(req.TotpCode, passport.Item.TotpSecret) {
		return nil, fireback.Create401Error(&AbacMessages.TotpCodeIsNotValid, []string{})
	}

	// Update the passport entity that it's confirmed
	if _, err := PassportActions.Update(fireback.QueryDSL{
		WorkspaceId: ROOT_VAR,
		UniqueId:    passport.Item.UniqueId,
	}, &abacdefs.PassportEntity{TotpConfirmed: emigo.NullableOf(true), UniqueId: passport.Item.UniqueId}); err != nil {
		return nil, fireback.Create401Error(&AbacMessages.PassportTotpNotConfirmed, []string{})
	}

	return &abacdefs.ConfirmClassicPassportTotpActionRes{
		Session: singinResult.Session,
	}, nil
}
