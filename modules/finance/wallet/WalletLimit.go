package wallet

import (
	"net/http"

	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
)

// WalletLimit.go holds the walletConfig.maxWalletsPer{User,Workspace}[PerCurrency]
// enforcement shared by every wallet-creation path: walletpublic's owner-facing
// createWallet and this module's own adminCreateWallet (AdminCreateWalletImplementation.go).
// Exported from here (rather than duplicated per caller) since walletpublic already
// depends on this package for everything else - keeping the one real implementation
// here, not the reverse, avoids a wallet -> walletpublic import.

// CheckWalletLimit enforces walletConfig's maxWalletsPer{User,Workspace}[PerCurrency]
// limits (root-configured, see getWalletConfig/updateWalletConfig). 0 means unlimited
// for the two blanket limits; the per-currency limits only apply when set at all
// (emigo.Nullable unset means "no extra limit").
func CheckWalletLimit(tx *gorm.DB, cfg *walletdefs.WalletConfigEntity, ownerType, userId, workspaceId, currency string) *fireback.IError {
	var ownerFilter string
	var ownerArg string
	var total, perCurrencyLimit int64
	var hasPerCurrencyLimit bool

	switch ownerType {
	case "user":
		ownerFilter, ownerArg = "owner_type = ? AND user_id = ?", userId
		total = cfg.MaxWalletsPerUser
		if v, ok := cfg.MaxWalletsPerUserPerCurrency.Get(); ok {
			perCurrencyLimit, hasPerCurrencyLimit = *v, true
		}
	case "workspace":
		ownerFilter, ownerArg = "owner_type = ? AND workspace_id = ?", workspaceId
		total = cfg.MaxWalletsPerWorkspace
		if v, ok := cfg.MaxWalletsPerWorkspacePerCurrency.Get(); ok {
			perCurrencyLimit, hasPerCurrencyLimit = *v, true
		}
	}

	if total > 0 {
		var count int64
		if err := tx.Model(&walletdefs.WalletEntity{}).Where(ownerFilter, ownerType, ownerArg).Count(&count).Error; err != nil {
			return fireback.GormErrorToIError(err)
		}
		if count >= total {
			return WalletLimitExceededError()
		}
	}
	if hasPerCurrencyLimit && perCurrencyLimit > 0 {
		var count int64
		if err := tx.Model(&walletdefs.WalletEntity{}).
			Where(ownerFilter+" AND currency = ?", ownerType, ownerArg, currency).
			Count(&count).Error; err != nil {
			return fireback.GormErrorToIError(err)
		}
		if count >= perCurrencyLimit {
			return WalletLimitExceededError()
		}
	}
	return nil
}

func WalletLimitExceededError() *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusConflict,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.limitExceeded",
			"en": "You have reached the maximum number of wallets allowed.",
			"fa": "شما به حداکثر تعداد کیف پول مجاز رسیده‌اید.",
			"ru": "Вы достигли максимально допустимого количества кошельков.",
			"pl": "Osiągnięto maksymalną dozwoloną liczbę portfeli.",
		},
	}
}

func InactiveCurrencyError(code string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusBadRequest,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.inactiveCurrency",
			"en": "This currency is not available for new wallets.",
			"fa": "این ارز برای ایجاد کیف پول جدید در دسترس نیست.",
			"ru": "Эта валюта недоступна для новых кошельков.",
			"pl": "Ta waluta nie jest dostępna dla nowych portfeli.",
		},
		MessageParams: map[string]any{"currency": code},
	}
}
