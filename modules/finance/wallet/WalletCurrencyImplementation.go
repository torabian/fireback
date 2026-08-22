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

// Admin/root-managed CRUD for walletCurrency - same shape as
// modules/category/CategoryImplementation.go. Currencies aren't workspace-scoped (they're
// installation-wide), so unlike Category's Browse there is no workspace_id filter here.

func walletCurrencyDtoFromEntity(e *walletdefs.WalletCurrencyEntity) walletdefs.WalletCurrencyDto {
	return walletdefs.WalletCurrencyDto{
		UniqueId: emigo.NullableOf(e.UniqueId),
		Code:     e.Code,
		Name:     e.Name,
		Kind:     e.Kind,
		Decimals: e.Decimals,
		Symbol:   e.Symbol,
		IsActive: e.IsActive,
	}
}

func WalletCurrencyCreateAction(c walletdefs.WalletCurrencyCreateActionRequest) (*walletdefs.WalletCurrencyCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_CREATE},
	})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletCurrencyCreateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	entity := &walletdefs.WalletCurrencyEntity{
		Code:     c.Body.Code,
		Name:     c.Body.Name,
		Kind:     c.Body.Kind,
		Decimals: c.Body.Decimals,
		Symbol:   c.Body.Symbol,
		IsActive: c.Body.IsActive,
	}
	created, err := walletdefs.WalletCurrencyEntityActions.Create(fireback.GetDbRef(), entity)
	if err != nil {
		return nil, err
	}
	return &walletdefs.WalletCurrencyCreateActionResponse{
		StatusCode: http.StatusCreated,
		Payload:    fireback.GResponseSingleItem(walletCurrencyDtoFromEntity(created)),
	}, nil
}

func WalletCurrencyGetAction(c walletdefs.WalletCurrencyGetActionRequest) (*walletdefs.WalletCurrencyGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletCurrencyEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletCurrencyGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet currency not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletCurrencyGetActionResponse{Payload: fireback.GResponseSingleItem(walletCurrencyDtoFromEntity(entity))}, nil
}

func WalletCurrencyUpdateAction(c walletdefs.WalletCurrencyUpdateActionRequest) (*walletdefs.WalletCurrencyUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_UPDATE},
	})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, true); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletCurrencyUpdateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	if _, err := walletdefs.WalletCurrencyEntityActions.Update(tx, c.Params.UniqueId, c.Body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletCurrencyUpdateActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet currency not found"},
			}, nil
		}
		return nil, err
	}
	entity, err := walletdefs.WalletCurrencyEntityActions.Get(tx, c.Params.UniqueId)
	if err != nil {
		return nil, err
	}
	return &walletdefs.WalletCurrencyUpdateActionResponse{Payload: fireback.GResponseSingleItem(walletCurrencyDtoFromEntity(entity))}, nil
}

func WalletCurrencyBrowseAction(c walletdefs.WalletCurrencyBrowseActionRequest) (*walletdefs.WalletCurrencyBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletCurrencyBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletCurrencyEntityActions.Browse(fireback.GetDbRef(), qs, "")
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletCurrencyDto, len(items))
	for i, item := range items {
		dtos[i] = walletCurrencyDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletCurrencyBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}

func WalletCurrencyAwareDeletePreviewAction(c walletdefs.WalletCurrencyAwareDeletePreviewActionRequest) (*walletdefs.WalletCurrencyAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_DELETE},
	}); err != nil {
		return nil, err
	}
	qs := walletdefs.WalletCurrencyAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode())
	preview, err := walletdefs.WalletCurrencyEntityActions.AwareDeletePreview(fireback.GetDbRef(), qs.UniqueIds)
	if err != nil {
		return nil, err
	}
	affected := make([]walletdefs.WalletCurrencyAwareDeletePreviewActionResAffected, len(preview.Affected))
	for i, a := range preview.Affected {
		affected[i] = walletdefs.WalletCurrencyAwareDeletePreviewActionResAffected{Relation: a.Relation, Count: a.Count}
	}
	return &walletdefs.WalletCurrencyAwareDeletePreviewActionResponse{
		Payload: fireback.GResponseSingleItem(walletdefs.WalletCurrencyAwareDeletePreviewActionRes{
			Message:  preview.Message,
			Affected: emigo.ArrayReplace(affected),
		}),
	}, nil
}

func WalletCurrencyAwareDeleteAction(c walletdefs.WalletCurrencyAwareDeleteActionRequest) (*walletdefs.WalletCurrencyAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_CURRENCY_DELETE},
	}); err != nil {
		return nil, err
	}
	if err := walletdefs.WalletCurrencyEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err != nil {
		return nil, err
	}
	return &walletdefs.WalletCurrencyAwareDeleteActionResponse{
		Payload: map[string]any{"deleted": c.Body.UniqueIds},
	}, nil
}
