package wallet

import (
	"errors"
	"net/http"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
)

// Root-only, DB-backed CRUD for walletProviderConfig - same shape as
// WalletCurrencyImplementation.go, but gated by SecurityModel.AllowOnRoot (a caller
// either is root or isn't) rather than a grantable permission, matching
// getWalletConfig/updateWalletConfig/adjustBalance's convention: this configures which
// payment provider adapters are live and how, arguably more sensitive than a currency or
// gateway row, so it isn't delegable via the regular permission system at all.
//
// Create/Update both go through emigo's regular validator; the composite unique index on
// (providerType, region) declared in Wallet.emi.yml is the actual "one provider type per
// region" enforcement - a second Create for the same pair fails at the DB level and
// surfaces as a normal validation-shaped error via fireback.GormErrorToIError, not
// something this file needs to check by hand.

func walletProviderConfigDtoFromEntity(e *walletdefs.WalletProviderConfigEntity) walletdefs.WalletProviderConfigDto {
	return walletdefs.WalletProviderConfigDto{
		UniqueId:     emigo.NullableOf(e.UniqueId),
		ProviderType: e.ProviderType,
		Region:       e.Region,
		IsEnabled:    e.IsEnabled,
		Config:       e.Config,
	}
}

func WalletProviderConfigCreateAction(c walletdefs.WalletProviderConfigCreateActionRequest) (*walletdefs.WalletProviderConfigCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletProviderConfigCreateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	entity := &walletdefs.WalletProviderConfigEntity{
		ProviderType: c.Body.ProviderType,
		Region:       c.Body.Region,
		IsEnabled:    c.Body.IsEnabled,
		Config:       c.Body.Config,
	}
	created, err := walletdefs.WalletProviderConfigEntityActions.Create(fireback.GetDbRef(), entity)
	if err != nil {
		if ierr := fireback.GormErrorToIError(err); ierr != nil {
			status, payload := emiIErrorPayload(ierr, query)
			return &walletdefs.WalletProviderConfigCreateActionResponse{StatusCode: status, Payload: payload}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletProviderConfigCreateActionResponse{
		StatusCode: http.StatusCreated,
		Payload:    fireback.GResponseSingleItem(walletProviderConfigDtoFromEntity(created)),
	}, nil
}

func WalletProviderConfigGetAction(c walletdefs.WalletProviderConfigGetActionRequest) (*walletdefs.WalletProviderConfigGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletProviderConfigEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletProviderConfigGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet provider config not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletProviderConfigGetActionResponse{Payload: fireback.GResponseSingleItem(walletProviderConfigDtoFromEntity(entity))}, nil
}

func WalletProviderConfigUpdateAction(c walletdefs.WalletProviderConfigUpdateActionRequest) (*walletdefs.WalletProviderConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, true); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.WalletProviderConfigUpdateActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	if _, err := walletdefs.WalletProviderConfigEntityActions.Update(tx, c.Params.UniqueId, c.Body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletProviderConfigUpdateActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet provider config not found"},
			}, nil
		}
		if ierr := fireback.GormErrorToIError(err); ierr != nil {
			status, payload := emiIErrorPayload(ierr, query)
			return &walletdefs.WalletProviderConfigUpdateActionResponse{StatusCode: status, Payload: payload}, nil
		}
		return nil, err
	}
	entity, err := walletdefs.WalletProviderConfigEntityActions.Get(tx, c.Params.UniqueId)
	if err != nil {
		return nil, err
	}
	return &walletdefs.WalletProviderConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(walletProviderConfigDtoFromEntity(entity))}, nil
}

func WalletProviderConfigBrowseAction(c walletdefs.WalletProviderConfigBrowseActionRequest) (*walletdefs.WalletProviderConfigBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletProviderConfigBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletProviderConfigEntityActions.Browse(fireback.GetDbRef(), qs, "")
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletProviderConfigDto, len(items))
	for i, item := range items {
		dtos[i] = walletProviderConfigDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletProviderConfigBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}

func WalletProviderConfigAwareDeletePreviewAction(c walletdefs.WalletProviderConfigAwareDeletePreviewActionRequest) (*walletdefs.WalletProviderConfigAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true}); err != nil {
		return nil, err
	}
	qs := walletdefs.WalletProviderConfigAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode())
	preview, err := walletdefs.WalletProviderConfigEntityActions.AwareDeletePreview(fireback.GetDbRef(), qs.UniqueIds)
	if err != nil {
		return nil, err
	}
	affected := make([]walletdefs.WalletProviderConfigAwareDeletePreviewActionResAffected, len(preview.Affected))
	for i, a := range preview.Affected {
		affected[i] = walletdefs.WalletProviderConfigAwareDeletePreviewActionResAffected{Relation: a.Relation, Count: a.Count}
	}
	return &walletdefs.WalletProviderConfigAwareDeletePreviewActionResponse{
		Payload: fireback.GResponseSingleItem(walletdefs.WalletProviderConfigAwareDeletePreviewActionRes{
			Message:  preview.Message,
			Affected: emigo.ArrayReplace(affected),
		}),
	}, nil
}

func WalletProviderConfigAwareDeleteAction(c walletdefs.WalletProviderConfigAwareDeleteActionRequest) (*walletdefs.WalletProviderConfigAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err := walletdefs.WalletProviderConfigEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err != nil {
		return nil, err
	}
	return &walletdefs.WalletProviderConfigAwareDeleteActionResponse{
		Payload: map[string]any{"deleted": c.Body.UniqueIds},
	}, nil
}
