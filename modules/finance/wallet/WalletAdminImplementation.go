package wallet

import (
	"errors"
	"net/http"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
)

// WalletAdminImplementation.go covers the wallet entity's Get/Browse/AwareDelete -
// PERM_ROOT_WALLET_ADMIN_QUERY/PERM_ROOT_WALLET_ADMIN_DELETE-gated admin/support views
// across every owner's wallets. No Create/Update action exists (see the wallet entity's
// `features` override in Wallet.emi.yml) - wallets are only ever created via
// walletpublic's createWallet, and balance never changes except through
// Purchase.go/adjustBalance; only label/status/isDefault can change, via walletpublic's
// updateWalletSettings.

func walletDtoFromEntity(e *walletdefs.WalletEntity) walletdefs.WalletDto {
	return walletdefs.WalletDto{
		UniqueId:    emigo.NullableOf(e.UniqueId),
		OwnerType:   e.OwnerType,
		UserId:      e.UserId,
		WorkspaceId: e.WorkspaceId,
		Currency:    e.Currency,
		Balance:     e.Balance,
		Status:      e.Status,
		Label:       e.Label,
		IsDefault:   e.IsDefault,
		Version:     e.Version,
	}
}

func WalletGetAction(c walletdefs.WalletGetActionRequest) (*walletdefs.WalletGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletGetActionResponse{Payload: fireback.GResponseSingleItem(walletDtoFromEntity(entity))}, nil
}

func WalletBrowseAction(c walletdefs.WalletBrowseActionRequest) (*walletdefs.WalletBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletEntityActions.Browse(fireback.GetDbRef(), qs, "")
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletDto, len(items))
	for i, item := range items {
		dtos[i] = walletDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}

func WalletAwareDeletePreviewAction(c walletdefs.WalletAwareDeletePreviewActionRequest) (*walletdefs.WalletAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_DELETE},
	}); err != nil {
		return nil, err
	}
	qs := walletdefs.WalletAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode())
	preview, err := walletdefs.WalletEntityActions.AwareDeletePreview(fireback.GetDbRef(), qs.UniqueIds)
	if err != nil {
		return nil, err
	}
	affected := make([]walletdefs.WalletAwareDeletePreviewActionResAffected, len(preview.Affected))
	for i, a := range preview.Affected {
		affected[i] = walletdefs.WalletAwareDeletePreviewActionResAffected{Relation: a.Relation, Count: a.Count}
	}
	return &walletdefs.WalletAwareDeletePreviewActionResponse{
		Payload: fireback.GResponseSingleItem(walletdefs.WalletAwareDeletePreviewActionRes{
			Message:  preview.Message,
			Affected: emigo.ArrayReplace(affected),
		}),
	}, nil
}

func WalletAwareDeleteAction(c walletdefs.WalletAwareDeleteActionRequest) (*walletdefs.WalletAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_DELETE},
	}); err != nil {
		return nil, err
	}
	if err := walletdefs.WalletEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err != nil {
		return nil, err
	}
	return &walletdefs.WalletAwareDeleteActionResponse{
		Payload: map[string]any{"deleted": c.Body.UniqueIds},
	}, nil
}
