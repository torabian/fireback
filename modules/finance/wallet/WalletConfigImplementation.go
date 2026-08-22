package wallet

import (
	"errors"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"gorm.io/gorm"
)

// walletConfig has features.actions:false in Wallet.emi.yml - no generic entity actions
// exist for it at all. It's read/written exclusively through these two actions, both
// root-only (SecurityModel.AllowOnRoot: true, no ActionRequires - a caller either is root
// or isn't, there's no grantable permission for it), matching
// modules/entitlement/GenerateActivationCodeImplementation.go's convention. In practice
// there is exactly one config row, created on first read/write with the defaults declared
// on the entity.

func walletConfigDtoFromEntity(e *walletdefs.WalletConfigEntity) walletdefs.WalletConfigDto {
	return walletdefs.WalletConfigDto{
		UniqueId:                          emigo.NullableOf(e.UniqueId),
		MaxWalletsPerUser:                 e.MaxWalletsPerUser,
		MaxWalletsPerWorkspace:            e.MaxWalletsPerWorkspace,
		MaxWalletsPerUserPerCurrency:      e.MaxWalletsPerUserPerCurrency,
		MaxWalletsPerWorkspacePerCurrency: e.MaxWalletsPerWorkspacePerCurrency,
	}
}

// getOrCreateWalletConfig returns the single well-known walletConfig row, creating it with
// the entity's declared defaults (5/5, no per-currency limit) if it doesn't exist yet.
func getOrCreateWalletConfig(tx *gorm.DB) (*walletdefs.WalletConfigEntity, error) {
	var cfg walletdefs.WalletConfigEntity
	err := tx.First(&cfg).Error
	if err == nil {
		return &cfg, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return walletdefs.WalletConfigEntityActions.Create(tx, &walletdefs.WalletConfigEntity{
		MaxWalletsPerUser:      5,
		MaxWalletsPerWorkspace: 5,
	})
}

// CurrentConfig is the exported form of getOrCreateWalletConfig, for walletpublic's
// createWallet action to enforce the configured per-owner wallet limits without going
// through HTTP.
func CurrentConfig() (*walletdefs.WalletConfigEntity, error) {
	return getOrCreateWalletConfig(fireback.GetDbRef())
}

func GetWalletConfigAction(c walletdefs.GetWalletConfigActionRequest) (*walletdefs.GetWalletConfigActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot: true,
	}); err != nil {
		return nil, err
	}
	cfg, err := getOrCreateWalletConfig(fireback.GetDbRef())
	if err != nil {
		return nil, err
	}
	return &walletdefs.GetWalletConfigActionResponse{Payload: fireback.GResponseSingleItem(walletConfigDtoFromEntity(cfg))}, nil
}

func UpdateWalletConfigAction(c walletdefs.UpdateWalletConfigActionRequest) (*walletdefs.UpdateWalletConfigActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot: true,
	}); err != nil {
		return nil, err
	}

	tx := fireback.GetDbRef()
	cfg, err := getOrCreateWalletConfig(tx)
	if err != nil {
		return nil, err
	}

	updated, err := walletdefs.WalletConfigEntityActions.Update(tx, cfg.UniqueId, walletdefs.WalletConfigOptionalDto{
		MaxWalletsPerUser:                 c.Body.MaxWalletsPerUser,
		MaxWalletsPerWorkspace:            c.Body.MaxWalletsPerWorkspace,
		MaxWalletsPerUserPerCurrency:      c.Body.MaxWalletsPerUserPerCurrency,
		MaxWalletsPerWorkspacePerCurrency: c.Body.MaxWalletsPerWorkspacePerCurrency,
	})
	if err != nil {
		return nil, err
	}
	return &walletdefs.UpdateWalletConfigActionResponse{Payload: fireback.GResponseSingleItem(walletConfigDtoFromEntity(updated))}, nil
}
