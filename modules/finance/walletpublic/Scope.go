package walletpublic

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
	"gorm.io/gorm"
)

// optionalString wraps a possibly-empty string into emigo.Nullable[string], unset when
// empty - used for GatewayInitResult fields (RedirectUrl/ClientSecret) that most gateways
// leave blank.
func optionalString(s string) emigo.Nullable[string] {
	if s == "" {
		return emigo.Nullable[string]{}
	}
	return emigo.NullableOf(s)
}

// Scope.go centralizes ownership-scoping and the internal-entity -> public-view-dto
// mapping every walletpublic action needs, so "does this caller own this wallet" and
// "what does a wallet look like from the outside" are each defined exactly once.

// walletViewDtoFromEntity projects a walletdefs.WalletEntity onto the public
// WalletViewDto - deliberately omitting nothing sensitive (unlike walletTransaction/
// walletPaymentAttempt, wallet has no internal-only fields), but kept as its own named
// mapper for symmetry and so a future internal-only field doesn't leak by accident.
func walletViewDtoFromEntity(e *walletdefs.WalletEntity) walletpublicdefs.WalletViewDto {
	return walletpublicdefs.WalletViewDto{
		UniqueId:    e.UniqueId,
		OwnerType:   e.OwnerType,
		WorkspaceId: e.WorkspaceId,
		Currency:    e.Currency,
		Balance:     e.Balance,
		Status:      e.Status,
		Label:       e.Label,
		IsDefault:   e.IsDefault,
	}
}

func walletTransactionViewDtoFromEntity(e *walletdefs.WalletTransactionEntity) walletpublicdefs.WalletTransactionViewDto {
	return walletpublicdefs.WalletTransactionViewDto{
		UniqueId:      e.UniqueId,
		Direction:     e.Direction,
		Amount:        e.Amount,
		BalanceAfter:  e.BalanceAfter,
		Reason:        e.Reason,
		ReferenceType: e.ReferenceType,
		ReferenceId:   e.ReferenceId,
		Note:          e.Note,
		CreatedAt:     e.CreatedAt,
	}
}

func walletPaymentAttemptViewDtoFromEntity(e *walletdefs.WalletPaymentAttemptEntity, gatewayCode string) walletpublicdefs.WalletPaymentAttemptViewDto {
	return walletpublicdefs.WalletPaymentAttemptViewDto{
		UniqueId:      e.UniqueId,
		Purpose:       e.Purpose,
		Amount:        e.Amount,
		Currency:      e.Currency,
		Status:        e.Status,
		GatewayCode:   gatewayCode,
		FailureReason: e.FailureReason,
		CreatedAt:     e.CreatedAt,
	}
}

// resolveOwnedWallet looks up walletId and checks the caller is allowed to see/act on it:
// the wallet's own user (ownerType "user") or a member of its owning workspace (ownerType
// "workspace"). Returns a not-found IError for both a genuinely missing wallet and one
// that exists but isn't the caller's - deliberately not distinguishing the two, so a
// caller can't probe for other people's wallet ids by the error they get back.
func resolveOwnedWallet(tx *gorm.DB, walletId string, userId string) (*walletdefs.WalletEntity, *fireback.IError) {
	var w walletdefs.WalletEntity
	if err := tx.First(&w, "unique_id = ?", walletId).Error; err != nil {
		return nil, walletNotFoundOrNotOwnedError()
	}
	if !ownsWallet(&w, userId) {
		return nil, walletNotFoundOrNotOwnedError()
	}
	return &w, nil
}

func ownsWallet(w *walletdefs.WalletEntity, userId string) bool {
	switch w.OwnerType {
	case "user":
		return w.UserId.OrDefault("") == userId
	case "workspace":
		return isMemberOfWorkspace(userId, w.WorkspaceId.OrDefault(""))
	default:
		return false
	}
}

// isMemberOfWorkspace reuses abac's own user<->workspace join table (the same one every
// other workspace-scoped feature in this codebase relies on) rather than re-implementing
// membership.
func isMemberOfWorkspace(userId, workspaceId string) bool {
	if userId == "" || workspaceId == "" {
		return false
	}
	var row abacdefs.UserWorkspaceEntity
	err := fireback.GetDbRef().
		Where("user_id = ? AND workspace_id = ?", userId, workspaceId).
		First(&row).Error
	return err == nil
}

func walletNotFoundOrNotOwnedError() *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusNotFound,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.notFound",
			"en": "Wallet not found.",
			"fa": "کیف پول یافت نشد.",
			"ru": "Кошелёк не найден.",
			"pl": "Nie znaleziono portfela.",
		},
	}
}

func notWorkspaceMemberError() *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusForbidden,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.notWorkspaceMember",
			"en": "You are not a member of this workspace.",
			"fa": "شما عضو این فضای کاری نیستید.",
			"ru": "Вы не являетесь участником этого рабочего пространства.",
			"pl": "Nie jesteś członkiem tej przestrzeni roboczej.",
		},
	}
}

// pageQueryValues builds a url.Values from the plain filter/sort/pagination fields every
// list-style walletpublic action's `qs:` shares, for feeding into the wallet module's own
// generated {X}BrowseActionQueryFromString builders (so browsing here honors the exact
// same filter/sort/cursor semantics as the admin Browse actions).
func pageQueryValues(filter, sort string, startIndex, itemsPerPage int, cursor string) url.Values {
	v := url.Values{}
	if filter != "" {
		v.Set("filter", filter)
	}
	if sort != "" {
		v.Set("sort", sort)
	}
	if startIndex != 0 {
		v.Set("startIndex", strconv.Itoa(startIndex))
	}
	if itemsPerPage != 0 {
		v.Set("itemsPerPage", strconv.Itoa(itemsPerPage))
	}
	if cursor != "" {
		v.Set("cursor", cursor)
	}
	return v
}
