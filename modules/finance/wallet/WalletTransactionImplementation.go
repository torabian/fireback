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

// walletTransaction is append-only (see its `features` override in Wallet.emi.yml) - only
// Get/Browse exist, for admin/support auditing across every wallet. Owners see their own
// history through walletpublic's walletHistory action instead, which returns the
// narrower walletTransactionView projection.

func walletTransactionDtoFromEntity(e *walletdefs.WalletTransactionEntity) walletdefs.WalletTransactionDto {
	dto := walletdefs.WalletTransactionDto{
		UniqueId:       emigo.NullableOf(e.UniqueId),
		Direction:      e.Direction,
		Amount:         e.Amount,
		BalanceAfter:   e.BalanceAfter,
		Reason:         e.Reason,
		ReferenceType:  e.ReferenceType,
		ReferenceId:    e.ReferenceId,
		IdempotencyKey: e.IdempotencyKey,
		Note:           e.Note,
		CreatedBy:      e.CreatedBy,
		Metadata:       e.Metadata,
		CreatedAt:      e.CreatedAt,
	}
	if e.Wallet != nil {
		dto.Wallet = emigo.NewOneNullable(walletDtoFromEntity(e.Wallet))
	}
	return dto
}

func WalletTransactionGetAction(c walletdefs.WalletTransactionGetActionRequest) (*walletdefs.WalletTransactionGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	}); err != nil {
		return nil, err
	}
	entity, err := walletdefs.WalletTransactionEntityActions.Get(fireback.GetDbRef().Preload("Wallet"), c.Params.UniqueId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &walletdefs.WalletTransactionGetActionResponse{
				StatusCode: http.StatusNotFound,
				Payload:    map[string]string{"error": "wallet transaction not found"},
			}, nil
		}
		return nil, err
	}
	return &walletdefs.WalletTransactionGetActionResponse{Payload: fireback.GResponseSingleItem(walletTransactionDtoFromEntity(entity))}, nil
}

func WalletTransactionBrowseAction(c walletdefs.WalletTransactionBrowseActionRequest) (*walletdefs.WalletTransactionBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_ADMIN_QUERY},
	})
	if err != nil {
		return nil, err
	}
	qs := walletdefs.WalletTransactionBrowseActionQueryFromString(c.QueryParams.Encode())
	items, meta, err2 := walletdefs.WalletTransactionEntityActions.Browse(fireback.GetDbRef().Preload("Wallet"), qs, "")
	if err2 != nil {
		return nil, err2
	}
	dtos := make([]walletdefs.WalletTransactionDto, len(items))
	for i, item := range items {
		dtos[i] = walletTransactionDtoFromEntity(item)
	}
	firebackMeta := &fireback.QueryResultMeta{TotalItems: meta.TotalItems, Cursor: meta.Cursor}
	return &walletdefs.WalletTransactionBrowseActionResponse{
		Payload: fireback.GResponseQuery(dtos, firebackMeta, query),
	}, nil
}
