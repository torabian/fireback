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

// Admin/root-managed CRUD for walletGateway - same shape as WalletCurrencyImplementation.go.
// config/supportedCurrencies are passed through as-is; per Wallet.emi.yml's field doc,
// config must only ever hold secret *references* (e.g. a secrets-manager key name), never
// raw API keys - that's a data-entry contract enforced by whoever administers gateways,
// not something this layer can validate generically.

func walletGatewayDtoFromEntity(e *walletdefs.WalletGatewayEntity) walletdefs.WalletGatewayDto {
	return walletdefs.WalletGatewayDto{
		UniqueId:            emigo.NullableOf(e.UniqueId),
		Code:                e.Code,
		Name:                e.Name,
		Kind:                e.Kind,
		IsActive:            e.IsActive,
		Config:              e.Config,
		SupportedCurrencies: e.SupportedCurrencies,
	}
}

func WalletGatewayCreateAction(c walletdefs.WalletGatewayCreateActionRequest) (*walletdefs.WalletGatewayCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_CREATE},
	})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletGatewayCreateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	entity := &walletdefs.WalletGatewayEntity{
		Code:                c.Body.Code,
		Name:                c.Body.Name,
		Kind:                c.Body.Kind,
		IsActive:            c.Body.IsActive,
		Config:              c.Body.Config,
		SupportedCurrencies: c.Body.SupportedCurrencies,
	}
	created, err := walletdefs.WalletGatewayEntityActions.Create(fireback.GetDbRef(), entity)
	if err != nil {
		return nil, err
	}
	return &walletdefs.WalletGatewayCreateActionResponse{
		StatusCode: http.StatusCreated,
		Payload:    fireback.GResponseSingleItem(walletGatewayDtoFromEntity(created)),
	}, nil
}

func WalletGatewayGetAction(c walletdefs.WalletGatewayGetActionRequest) (*walletdefs.WalletGatewayGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletGatewayEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletGatewayGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet gateway not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletGatewayGetActionResponse{Payload: fireback.GResponseSingleItem(walletGatewayDtoFromEntity(entity))}, nil
}

func WalletGatewayUpdateAction(c walletdefs.WalletGatewayUpdateActionRequest) (*walletdefs.WalletGatewayUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_UPDATE},
	})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, true); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletGatewayUpdateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	if _, err := walletdefs.WalletGatewayEntityActions.Update(tx, c.Params.UniqueId, c.Body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletGatewayUpdateActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet gateway not found"},
			}, nil
		}
		return nil, err
	}
	entity, err := walletdefs.WalletGatewayEntityActions.Get(tx, c.Params.UniqueId)
	if err != nil {
		return nil, err
	}
	return &walletdefs.WalletGatewayUpdateActionResponse{Payload: fireback.GResponseSingleItem(walletGatewayDtoFromEntity(entity))}, nil
}

func WalletGatewayBrowseAction(c walletdefs.WalletGatewayBrowseActionRequest) (*walletdefs.WalletGatewayBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletGatewayBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletGatewayEntityActions.Browse(fireback.GetDbRef(), qs, "")
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletGatewayDto, len(items))
	for i, item := range items {
		dtos[i] = walletGatewayDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletGatewayBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}

func WalletGatewayAwareDeletePreviewAction(c walletdefs.WalletGatewayAwareDeletePreviewActionRequest) (*walletdefs.WalletGatewayAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_DELETE},
	}); err != nil {
		return nil, err
	}
	qs := walletdefs.WalletGatewayAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode())
	preview, err := walletdefs.WalletGatewayEntityActions.AwareDeletePreview(fireback.GetDbRef(), qs.UniqueIds)
	if err != nil {
		return nil, err
	}
	affected := make([]walletdefs.WalletGatewayAwareDeletePreviewActionResAffected, len(preview.Affected))
	for i, a := range preview.Affected {
		affected[i] = walletdefs.WalletGatewayAwareDeletePreviewActionResAffected{Relation: a.Relation, Count: a.Count}
	}
	return &walletdefs.WalletGatewayAwareDeletePreviewActionResponse{
		Payload: fireback.GResponseSingleItem(walletdefs.WalletGatewayAwareDeletePreviewActionRes{
			Message:  preview.Message,
			Affected: emigo.ArrayReplace(affected),
		}),
	}, nil
}

func WalletGatewayAwareDeleteAction(c walletdefs.WalletGatewayAwareDeleteActionRequest) (*walletdefs.WalletGatewayAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_GATEWAY_DELETE},
	}); err != nil {
		return nil, err
	}
	if err := walletdefs.WalletGatewayEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err != nil {
		return nil, err
	}
	return &walletdefs.WalletGatewayAwareDeleteActionResponse{
		Payload: map[string]any{"deleted": c.Body.UniqueIds},
	}, nil
}
