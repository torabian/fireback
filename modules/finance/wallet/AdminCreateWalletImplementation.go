package wallet

import (
	"net/http"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

// AdminCreateWalletAction is root-only: the admin/support counterpart to walletpublic's
// createWallet, which can only ever create a wallet for the caller themselves. This can
// target any user or workspace, still going through the same
// wallet.CheckWalletLimit/currency-active checks as the owner-facing path (see
// WalletLimit.go) - an admin bypasses ownership, not the configured limits.
func AdminCreateWalletAction(c walletdefs.AdminCreateWalletActionRequest) (*walletdefs.AdminCreateWalletActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.AdminCreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
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
		userId := c.Body.UserId.OrDefault("")
		var user abacdefs.UserEntity
		if err := tx.First(&user, "unique_id = ?", userId).Error; err != nil {
			status, payload := emiIErrorPayload(targetNotFoundError("user", userId), query)
			return &walletdefs.AdminCreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
		}
		entity.UserId = emigo.NullableOf(userId)
	case "workspace":
		workspaceId := c.Body.WorkspaceId.OrDefault("")
		var workspace abacdefs.WorkspaceEntity
		if err := tx.First(&workspace, "unique_id = ?", workspaceId).Error; err != nil {
			status, payload := emiIErrorPayload(targetNotFoundError("workspace", workspaceId), query)
			return &walletdefs.AdminCreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
		}
		entity.WorkspaceId = emigo.NullableOf(workspaceId)
	}

	var currency walletdefs.WalletCurrencyEntity
	if err := tx.First(&currency, "code = ? AND is_active = ?", c.Body.Currency, true).Error; err != nil {
		status, payload := emiIErrorPayload(InactiveCurrencyError(c.Body.Currency), query)
		return &walletdefs.AdminCreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}

	cfg, err := CurrentConfig()
	if err != nil {
		return nil, err
	}
	if ierr := CheckWalletLimit(tx, cfg, c.Body.OwnerType, entity.UserId.OrDefault(""), entity.WorkspaceId.OrDefault(""), c.Body.Currency); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletdefs.AdminCreateWalletActionResponse{StatusCode: status, Payload: payload}, nil
	}

	created, err := walletdefs.WalletEntityActions.Create(tx, entity)
	if err != nil {
		return nil, err
	}
	return &walletdefs.AdminCreateWalletActionResponse{
		StatusCode: http.StatusCreated,
		Payload:    fireback.GResponseSingleItem(walletDtoFromEntity(created)),
	}, nil
}

func targetNotFoundError(kind, id string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusBadRequest,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.targetNotFound",
			"en": "The target " + kind + " does not exist.",
			"fa": "این " + kind + " وجود ندارد.",
			"ru": "Указанный " + kind + " не существует.",
			"pl": "Wskazany " + kind + " nie istnieje.",
		},
		MessageParams: map[string]any{"kind": kind, "id": id},
	}
}
