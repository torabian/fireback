package walletpublic

import (
	"net/http"

	"github.com/torabian/fireback/modules/fireback"
	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// CreateWalletAction creates a wallet for the caller (ownerType "user") or for a
// workspace the caller belongs to (ownerType "workspace"). No permission is required
// beyond being logged in (matching modules/entitlement/MySubscriptionsImplementation.go's
// convention) - what a caller is allowed to create for is entirely determined by
// ownership/membership, checked here, not by a grantable permission.
func CreateWalletAction(c walletpublicdefs.CreateWalletActionRequest) (*walletpublicdefs.CreateWalletActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.CreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()

	entity := &walletdefs.WalletEntity{
		OwnerType: c.Body.OwnerType,
		Currency:  c.Body.Currency,
		Balance:   "0",
		Status:    "active",
		Label:     c.Body.Label,
	}

	switch c.Body.OwnerType {
	case "user":
		entity.UserId.Set(&query.UserId)
	case "workspace":
		workspaceId := c.Body.WorkspaceId.OrDefault("")
		if !isMemberOfWorkspace(query.UserId, workspaceId) {
			status, payload := emiIErrorPayload(notWorkspaceMemberError(), query)
			return &walletpublicdefs.CreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
		}
		entity.WorkspaceId.Set(&workspaceId)
	}

	var currency walletdefs.WalletCurrencyEntity
	if err := tx.First(&currency, "code = ? AND is_active = ?", c.Body.Currency, true).Error; err != nil {
		status, payload := emiIErrorPayload(wallet.InactiveCurrencyError(c.Body.Currency), query)
		return &walletpublicdefs.CreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}

	cfg, err := wallet.CurrentConfig()
	if err != nil {
		return nil, err
	}
	if ierr := wallet.CheckWalletLimit(tx, cfg, c.Body.OwnerType, query.UserId, entity.WorkspaceId.OrDefault(""), c.Body.Currency); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.CreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}

	created, err := walletdefs.WalletEntityActions.Create(tx, entity)
	if err != nil {
		return nil, err
	}
	return &walletpublicdefs.CreateWalletActionResponse{
		StatusCode: http.StatusCreated,
		Payload:    fireback.GResponseSingleItem(walletViewDtoFromEntity(created)),
	}, nil
}

// The maxWalletsPer{User,Workspace}[PerCurrency] check and its errors moved to
// wallet.CheckWalletLimit (modules/wallet/WalletLimit.go) - shared with this module's
// own admin-facing counterpart (wallet.AdminCreateWalletAction) instead of duplicated.
