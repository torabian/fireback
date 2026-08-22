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

// walletPaymentAttempt has no public Create/Update/Delete (see its `features` override in
// Wallet.emi.yml) - it's only created/updated by walletpublic's topup/gatewayWebhook/
// cancelPaymentAttempt logic. Get/Browse here are the admin/support view across every
// wallet's attempts; owners see their own through walletpublic's walletPaymentAttempts
// action instead.

func walletPaymentAttemptDtoFromEntity(e *walletdefs.WalletPaymentAttemptEntity) walletdefs.WalletPaymentAttemptDto {
	dto := walletdefs.WalletPaymentAttemptDto{
		UniqueId:         emigo.NullableOf(e.UniqueId),
		Purpose:          e.Purpose,
		Amount:           e.Amount,
		Currency:         e.Currency,
		Status:           e.Status,
		GatewayReference: e.GatewayReference,
		IdempotencyKey:   e.IdempotencyKey,
		FailureReason:    e.FailureReason,
		RawRequest:       e.RawRequest,
		RawResponse:      e.RawResponse,
		CreatedAt:        e.CreatedAt,
		ExpiresAt:        e.ExpiresAt,
		CompletedAt:      e.CompletedAt,
	}
	if e.Wallet != nil {
		dto.Wallet = emigo.NewOneNullable(walletDtoFromEntity(e.Wallet))
	}
	if e.Gateway != nil {
		dto.Gateway = emigo.NewOneNullable(walletGatewayDtoFromEntity(e.Gateway))
	}
	if e.WalletTransactionId != 0 {
		dto.WalletTransaction = emigo.NewOneNullable(walletTransactionDtoFromEntity(&e.WalletTransaction))
	}
	return dto
}

func WalletPaymentAttemptGetAction(c walletdefs.WalletPaymentAttemptGetActionRequest) (*walletdefs.WalletPaymentAttemptGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletPaymentAttemptEntityActions.Get(
		fireback.GetDbRef().Preload("Wallet").Preload("Gateway").Preload("WalletTransaction"),
		c.Params.UniqueId,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletPaymentAttemptGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet payment attempt not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletPaymentAttemptGetActionResponse{Payload: fireback.GResponseSingleItem(walletPaymentAttemptDtoFromEntity(entity))}, nil
}

func WalletPaymentAttemptBrowseAction(c walletdefs.WalletPaymentAttemptBrowseActionRequest) (*walletdefs.WalletPaymentAttemptBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletPaymentAttemptBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletPaymentAttemptEntityActions.Browse(
		fireback.GetDbRef().Preload("Wallet").Preload("Gateway").Preload("WalletTransaction"),
		qs, "",
	)
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletPaymentAttemptDto, len(items))
	for i, item := range items {
		dtos[i] = walletPaymentAttemptDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletPaymentAttemptBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
