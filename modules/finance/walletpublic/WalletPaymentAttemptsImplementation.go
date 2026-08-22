package walletpublic

import (
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// WalletPaymentAttemptsAction returns the caller's own wallet's paginated payment
// attempts, scoped the same way GetWalletAction is.
func WalletPaymentAttemptsAction(c walletpublicdefs.WalletPaymentAttemptsActionRequest) (*walletpublicdefs.WalletPaymentAttemptsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	qs := walletpublicdefs.WalletPaymentAttemptsActionQueryFromString(c.QueryParams.Encode())

	tx := fireback.GetDbRef()
	w, ierr := resolveOwnedWallet(tx, qs.WalletId, query.UserId)
	if ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.WalletPaymentAttemptsActionResponse{StatusCode: status, Payload: payload}, nil
	}

	browseQs := walletdefs.WalletPaymentAttemptBrowseActionQueryFromString(
		pageQueryValues(qs.Filter, qs.Sort, qs.StartIndex, qs.ItemsPerPage, qs.Cursor).Encode(),
	)
	items, meta, err := walletdefs.WalletPaymentAttemptEntityActions.Browse(
		tx.Preload("Gateway"), browseQs, "wallet_id = ?", w.Id,
	)
	if err != nil {
		return nil, err
	}
	dtos := make([]walletpublicdefs.WalletPaymentAttemptViewDto, len(items))
	for i, item := range items {
		gatewayCode := ""
		if item.Gateway != nil {
			gatewayCode = item.Gateway.Code
		}
		dtos[i] = walletPaymentAttemptViewDtoFromEntity(item, gatewayCode)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletpublicdefs.WalletPaymentAttemptsActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
