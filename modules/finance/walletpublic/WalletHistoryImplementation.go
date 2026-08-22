package walletpublic

import (
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// WalletHistoryAction returns the caller's own wallet's paginated balance-history ledger,
// scoped the same way GetWalletAction is.
func WalletHistoryAction(c walletpublicdefs.WalletHistoryActionRequest) (*walletpublicdefs.WalletHistoryActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	qs := walletpublicdefs.WalletHistoryActionQueryFromString(c.QueryParams.Encode())

	tx := fireback.GetDbRef()
	w, ierr := resolveOwnedWallet(tx, qs.WalletId, query.UserId)
	if ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.WalletHistoryActionResponse{StatusCode: status, Payload: payload}, nil
	}

	browseQs := walletdefs.WalletTransactionBrowseActionQueryFromString(
		pageQueryValues(qs.Filter, qs.Sort, qs.StartIndex, qs.ItemsPerPage, qs.Cursor).Encode(),
	)
	items, meta, err := walletdefs.WalletTransactionEntityActions.Browse(tx, browseQs, "wallet_id = ?", w.Id)
	if err != nil {
		return nil, err
	}
	dtos := make([]walletpublicdefs.WalletTransactionViewDto, len(items))
	for i, item := range items {
		dtos[i] = walletTransactionViewDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletpublicdefs.WalletHistoryActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
