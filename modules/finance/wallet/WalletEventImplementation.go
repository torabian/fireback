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

// walletEvent is append-only, written only by the gatewayWebhook handler (see
// WalletModule.go's plain Gin route) - Get/Browse here are the admin/support audit view.

func walletEventDtoFromEntity(e *walletdefs.WalletEventEntity) walletdefs.WalletEventDto {
	dto := walletdefs.WalletEventDto{
		UniqueId:        emigo.NullableOf(e.UniqueId),
		EventType:       e.EventType,
		ExternalEventId: e.ExternalEventId,
		Payload:         e.Payload,
		Processed:       e.Processed,
		ProcessingError: e.ProcessingError,
		ReceivedAt:      e.ReceivedAt,
	}
	if e.Gateway != nil {
		dto.Gateway = emigo.NewOneNullable(walletGatewayDtoFromEntity(e.Gateway))
	}
	if e.PaymentAttemptId != 0 {
		dto.PaymentAttempt = emigo.NewOneNullable(walletPaymentAttemptDtoFromEntity(&e.PaymentAttempt))
	}
	return dto
}

func WalletEventGetAction(c walletdefs.WalletEventGetActionRequest) (*walletdefs.WalletEventGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletEventEntityActions.Get(
		fireback.GetDbRef().Preload("Gateway").Preload("PaymentAttempt"),
		c.Params.UniqueId,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletEventGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet event not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletEventGetActionResponse{Payload: fireback.GResponseSingleItem(walletEventDtoFromEntity(entity))}, nil
}

func WalletEventBrowseAction(c walletdefs.WalletEventBrowseActionRequest) (*walletdefs.WalletEventBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletEventBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletEventEntityActions.Browse(
		fireback.GetDbRef().Preload("Gateway").Preload("PaymentAttempt"),
		qs, "",
	)
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletEventDto, len(items))
	for i, item := range items {
		dtos[i] = walletEventDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletEventBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
