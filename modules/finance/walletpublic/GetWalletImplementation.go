package walletpublic

import (
	"github.com/torabian/fireback/modules/fireback"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// GetWalletAction reads one wallet by id, scoped to the caller (owning user, or a member
// of the owning workspace) via resolveOwnedWallet.
func GetWalletAction(c walletpublicdefs.GetWalletActionRequest) (*walletpublicdefs.GetWalletActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	qs := walletpublicdefs.GetWalletActionQueryFromString(c.QueryParams.Encode())

	w, ierr := resolveOwnedWallet(fireback.GetDbRef(), qs.WalletId, query.UserId)
	if ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.GetWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}
	return &walletpublicdefs.GetWalletActionResponse{
		Payload: fireback.GResponseSingleItem(walletViewDtoFromEntity(w)),
	}, nil
}
