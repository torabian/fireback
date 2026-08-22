package wallet

import (
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

// AdjustBalanceAction is root-only (support/ops manual correction) and goes through the
// same applyLedgerEntry path Purchase does - see Purchase.go for the concurrency/
// idempotency guarantees. reason is always "adjustment" and note is required for audit
// (enforced by the emi-generated `validate:required` tag on AdjustBalanceActionReq.Note).
func AdjustBalanceAction(c walletdefs.AdjustBalanceActionRequest) (*walletdefs.AdjustBalanceActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot: true,
	})
	if err != nil {
		return nil, err
	}

	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.AdjustBalanceActionResponse{StatusCode: status, Payload: payload}, nil
	}

	entry, err := applyLedgerEntry(fireback.GetDbRef(), ledgerChange{
		WalletUniqueId: c.Body.WalletId,
		Direction:      c.Body.Direction,
		Amount:         c.Body.Amount,
		Reason:         "adjustment",
		IdempotencyKey: c.Body.IdempotencyKey,
		Note:           c.Body.Note,
		CreatedBy:      query.UserId,
	})
	if err != nil {
		if ierr, ok := err.(*fireback.IError); ok {
			status, payload := emiIErrorPayload(ierr, query)
			return &walletdefs.AdjustBalanceActionResponse{StatusCode: status, Payload: payload}, nil
		}
		return nil, err
	}
	return &walletdefs.AdjustBalanceActionResponse{Payload: fireback.GResponseSingleItem(walletTransactionDtoFromEntity(entry))}, nil
}
