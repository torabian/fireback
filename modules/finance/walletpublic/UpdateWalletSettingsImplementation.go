package walletpublic

import (
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// UpdateWalletSettingsAction updates a wallet's non-financial fields only
// (label/status/isDefault) - balance is never reachable through this action's dto.
// Scoped to the caller via resolveOwnedWallet.
func UpdateWalletSettingsAction(c walletpublicdefs.UpdateWalletSettingsActionRequest) (*walletpublicdefs.UpdateWalletSettingsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, true); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.UpdateWalletSettingsActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	w, ierr := resolveOwnedWallet(tx, c.Body.WalletId, query.UserId)
	if ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.UpdateWalletSettingsActionResponse{StatusCode: status, Payload: payload}, nil
	}

	updated, err := walletdefs.WalletEntityActions.Update(tx, w.UniqueId, walletdefs.WalletOptionalDto{
		Label:     c.Body.Label,
		Status:    c.Body.Status,
		IsDefault: c.Body.IsDefault,
	})
	if err != nil {
		return nil, err
	}
	return &walletpublicdefs.UpdateWalletSettingsActionResponse{
		Payload: fireback.GResponseSingleItem(walletViewDtoFromEntity(updated)),
	}, nil
}
