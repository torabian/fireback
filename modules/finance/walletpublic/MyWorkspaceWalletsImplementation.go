package walletpublic

import (
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// MyWorkspaceWalletsAction lists the wallets owned by a workspace the caller belongs to.
func MyWorkspaceWalletsAction(c walletpublicdefs.MyWorkspaceWalletsActionRequest) (*walletpublicdefs.MyWorkspaceWalletsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}

	qs := walletpublicdefs.MyWorkspaceWalletsActionQueryFromString(c.QueryParams.Encode())
	if !isMemberOfWorkspace(query.UserId, qs.WorkspaceId) {
		status, payload := emiIErrorPayload(notWorkspaceMemberError(), query)
		return &walletpublicdefs.MyWorkspaceWalletsActionResponse{StatusCode: status, Payload: payload}, nil
	}

	browseQs := walletdefs.WalletBrowseActionQueryFromString(
		pageQueryValues(qs.Filter, qs.Sort, qs.StartIndex, qs.ItemsPerPage, qs.Cursor).Encode(),
	)
	items, meta, err := walletdefs.WalletEntityActions.Browse(
		fireback.GetDbRef(), browseQs, "owner_type = ? AND workspace_id = ?", "workspace", qs.WorkspaceId,
	)
	if err != nil {
		return nil, err
	}
	dtos := make([]walletpublicdefs.WalletViewDto, len(items))
	for i, item := range items {
		dtos[i] = walletViewDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletpublicdefs.MyWorkspaceWalletsActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
