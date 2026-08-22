package walletpublic

import (
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// MyWalletsAction lists the caller's own user-owned wallets. Logged-in-only, no extra
// permission - same convention as modules/entitlement/MySubscriptionsImplementation.go.
func MyWalletsAction(c walletpublicdefs.MyWalletsActionRequest) (*walletpublicdefs.MyWalletsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}

	qs := walletpublicdefs.MyWalletsActionQueryFromString(c.QueryParams.Encode())
	browseQs := walletdefs.WalletBrowseActionQueryFromString(
		pageQueryValues(qs.Filter, qs.Sort, qs.StartIndex, qs.ItemsPerPage, qs.Cursor).Encode(),
	)
	items, meta, err := walletdefs.WalletEntityActions.Browse(
		fireback.GetDbRef(), browseQs, "owner_type = ? AND user_id = ?", "user", query.UserId,
	)
	if err != nil {
		return nil, err
	}
	dtos := make([]walletpublicdefs.WalletViewDto, len(items))
	for i, item := range items {
		dtos[i] = walletViewDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletpublicdefs.MyWalletsActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
