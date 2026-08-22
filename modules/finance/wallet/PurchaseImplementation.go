package wallet

import (
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

// PurchaseAction is the HTTP-callable twin of the in-process Purchase Go function
// (Purchase.go) - both share applyLedgerEntry, so there is exactly one implementation to
// verify. Gated behind PERM_ROOT_WALLET_PURCHASE: meant for trusted internal/service
// callers debiting a wallet on behalf of another module's sale, never for a wallet owner
// directly (owners can't call this - see PERM_ROOT_WALLET_PURCHASE's doc comment).
func PurchaseAction(c walletdefs.PurchaseActionRequest) (*walletdefs.PurchaseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    false,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WALLET_PURCHASE},
	})
	if err != nil {
		return nil, err
	}

	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.PurchaseActionResponse{StatusCode: status, Payload: payload}, nil
	}

	entry, err := Purchase(PurchaseInput{
		WalletUniqueId: c.Body.WalletId,
		Amount:         c.Body.Amount,
		ReferenceType:  c.Body.ReferenceType,
		ReferenceId:    c.Body.ReferenceId,
		IdempotencyKey: c.Body.IdempotencyKey,
		Note:           c.Body.Note.OrDefault(""),
		CreatedBy:      query.UserId,
	})
	if err != nil {
		if ierr, ok := err.(*fireback.IError); ok {
			status, payload := emiIErrorPayload(ierr, query)
			return &walletdefs.PurchaseActionResponse{StatusCode: status, Payload: payload}, nil
		}
		return nil, err
	}
	return &walletdefs.PurchaseActionResponse{Payload: fireback.GResponseSingleItem(walletTransactionDtoFromEntity(entry))}, nil
}
