package wallet

import (
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"github.com/urfave/cli/v3"
)

// RouterCliManifest bundles every wallet-module action's *cli.Command - same convention
// as modules/category/RouterManifest.go. GatewayWebhookHandler is intentionally absent:
// it's a raw Gin route (see GatewayWebhookImplementation.go), not an emi action, so it has
// no CLI counterpart - a real webhook only ever arrives over HTTP from a gateway.
func RouterCliManifestManual() []*cli.Command {
	return []*cli.Command{
		// wallet (admin/support read-only + delete)
		walletdefs.WalletGetActionCliHandler(WalletGetAction),
		walletdefs.WalletBrowseActionCliHandler(WalletBrowseAction),
		walletdefs.WalletAwareDeletePreviewActionCliHandler(WalletAwareDeletePreviewAction),
		walletdefs.WalletAwareDeleteActionCliHandler(WalletAwareDeleteAction),

		// custom actions
		walletdefs.GetWalletConfigActionCliHandler(GetWalletConfigAction),
		walletdefs.UpdateWalletConfigActionCliHandler(UpdateWalletConfigAction),
		walletdefs.AdminCreateWalletActionCliHandler(AdminCreateWalletAction),
		walletdefs.AdjustBalanceActionCliHandler(AdjustBalanceAction),
		walletdefs.PurchaseActionCliHandler(PurchaseAction),

		{
			Name: "currency",
			Commands: []*cli.Command{
				// walletCurrency (admin CRUD)
				walletdefs.WalletCurrencyCreateActionCliHandler(WalletCurrencyCreateAction),
				walletdefs.WalletCurrencyUpdateActionCliHandler(WalletCurrencyUpdateAction),
				walletdefs.WalletCurrencyGetActionCliHandler(WalletCurrencyGetAction),
				walletdefs.WalletCurrencyBrowseActionCliHandler(WalletCurrencyBrowseAction),
				walletdefs.WalletCurrencyAwareDeletePreviewActionCliHandler(WalletCurrencyAwareDeletePreviewAction),
				walletdefs.WalletCurrencyAwareDeleteActionCliHandler(WalletCurrencyAwareDeleteAction),
			},
		},
		{
			Name: "gateway",
			Commands: []*cli.Command{
				walletdefs.WalletGatewayCreateActionCliHandler(WalletGatewayCreateAction),
				walletdefs.WalletGatewayUpdateActionCliHandler(WalletGatewayUpdateAction),
				walletdefs.WalletGatewayGetActionCliHandler(WalletGatewayGetAction),
				walletdefs.WalletGatewayBrowseActionCliHandler(WalletGatewayBrowseAction),
				walletdefs.WalletGatewayAwareDeletePreviewActionCliHandler(WalletGatewayAwareDeletePreviewAction),
				walletdefs.WalletGatewayAwareDeleteActionCliHandler(WalletGatewayAwareDeleteAction),
			},
		},
		{
			Name: "transaction",
			Commands: []*cli.Command{
				// walletTransaction (append-only ledger, admin read-only)
				walletdefs.WalletTransactionGetActionCliHandler(WalletTransactionGetAction),
				walletdefs.WalletTransactionBrowseActionCliHandler(WalletTransactionBrowseAction),
			},
		},
		{
			Name: "provider",
			Commands: []*cli.Command{

				// walletProviderConfig (root-only CRUD - see WalletProviderConfigImplementation.go)
				walletdefs.WalletProviderConfigCreateActionCliHandler(WalletProviderConfigCreateAction),
				walletdefs.WalletProviderConfigUpdateActionCliHandler(WalletProviderConfigUpdateAction),
				walletdefs.WalletProviderConfigGetActionCliHandler(WalletProviderConfigGetAction),
				walletdefs.WalletProviderConfigBrowseActionCliHandler(WalletProviderConfigBrowseAction),
				walletdefs.WalletProviderConfigAwareDeletePreviewActionCliHandler(WalletProviderConfigAwareDeletePreviewAction),
				walletdefs.WalletProviderConfigAwareDeleteActionCliHandler(WalletProviderConfigAwareDeleteAction),
			},
		},
		{
			Name: "payment-attempt",
			Commands: []*cli.Command{
				// walletPaymentAttempt (admin read-only)
				walletdefs.WalletPaymentAttemptGetActionCliHandler(WalletPaymentAttemptGetAction),
				walletdefs.WalletPaymentAttemptBrowseActionCliHandler(WalletPaymentAttemptBrowseAction),
			},
		},
		{
			Name: "event",
			Commands: []*cli.Command{
				// walletEvent (admin read-only)
				walletdefs.WalletEventGetActionCliHandler(WalletEventGetAction),
				walletdefs.WalletEventBrowseActionCliHandler(WalletEventBrowseAction),
			},
		},
	}
}
