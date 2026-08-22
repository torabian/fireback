package wallet

// See Wallet.emi.yml - business logic lives in the various *Implementation.go files,
// permissions are declared by hand in Permissions.go. Same layout as
// modules/category/CategoryModule.go.
import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	"github.com/urfave/cli/v3"
)

type WalletModuleConfig struct {
	// Gateways are the concrete GatewayAdapter implementations available to this
	// installation, registered under their own Code(). Always include MockGatewayAdapter
	// (and usually ManualGatewayAdapter) for CLI/dev use even alongside real providers -
	// see GatewayAdapter.go. A walletGateway row with no matching adapter here fails
	// topup/webhook calls with a clear "no adapter registered" error rather than a panic.
	Gateways []GatewayAdapter
}

func WalletModuleSetup(cfg *WalletModuleConfig) *application.ModuleProvider {
	if cfg == nil {
		cfg = &WalletModuleConfig{}
	}
	if len(cfg.Gateways) == 0 {
		cfg.Gateways = []GatewayAdapter{MockGatewayAdapter{}, ManualGatewayAdapter{}}
	}
	registerGatewayAdapters(cfg.Gateways)

	module := &application.ModuleProvider{
		Name: "wallet",

		EntityBundles: []application.EntityBundle{
			{
				AutoMigrationEntities: []interface{}{
					&walletdefs.WalletEntity{},
					&walletdefs.WalletCurrencyEntity{},
					&walletdefs.WalletTransactionEntity{},
					&walletdefs.WalletGatewayEntity{},
					&walletdefs.WalletPaymentAttemptEntity{},
					&walletdefs.WalletEventEntity{},
					&walletdefs.WalletConfigEntity{},
					&walletdefs.WalletProviderConfigEntity{},
				},
			},
		},

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				walletdefs.WalletGetActionGin(g, WalletGetAction)
				walletdefs.WalletBrowseActionGin(g, WalletBrowseAction)
				walletdefs.WalletAwareDeletePreviewActionGin(g, WalletAwareDeletePreviewAction)
				walletdefs.WalletAwareDeleteActionGin(g, WalletAwareDeleteAction)

				walletdefs.WalletCurrencyCreateActionGin(g, WalletCurrencyCreateAction)
				walletdefs.WalletCurrencyUpdateActionGin(g, WalletCurrencyUpdateAction)
				walletdefs.WalletCurrencyGetActionGin(g, WalletCurrencyGetAction)
				walletdefs.WalletCurrencyBrowseActionGin(g, WalletCurrencyBrowseAction)
				walletdefs.WalletCurrencyAwareDeletePreviewActionGin(g, WalletCurrencyAwareDeletePreviewAction)
				walletdefs.WalletCurrencyAwareDeleteActionGin(g, WalletCurrencyAwareDeleteAction)

				walletdefs.WalletGatewayCreateActionGin(g, WalletGatewayCreateAction)
				walletdefs.WalletGatewayUpdateActionGin(g, WalletGatewayUpdateAction)
				walletdefs.WalletGatewayGetActionGin(g, WalletGatewayGetAction)
				walletdefs.WalletGatewayBrowseActionGin(g, WalletGatewayBrowseAction)
				walletdefs.WalletGatewayAwareDeletePreviewActionGin(g, WalletGatewayAwareDeletePreviewAction)
				walletdefs.WalletGatewayAwareDeleteActionGin(g, WalletGatewayAwareDeleteAction)

				walletdefs.WalletTransactionGetActionGin(g, WalletTransactionGetAction)
				walletdefs.WalletTransactionBrowseActionGin(g, WalletTransactionBrowseAction)

				walletdefs.WalletPaymentAttemptGetActionGin(g, WalletPaymentAttemptGetAction)
				walletdefs.WalletPaymentAttemptBrowseActionGin(g, WalletPaymentAttemptBrowseAction)

				walletdefs.WalletEventGetActionGin(g, WalletEventGetAction)
				walletdefs.WalletEventBrowseActionGin(g, WalletEventBrowseAction)

				walletdefs.WalletProviderConfigCreateActionGin(g, WalletProviderConfigCreateAction)
				walletdefs.WalletProviderConfigUpdateActionGin(g, WalletProviderConfigUpdateAction)
				walletdefs.WalletProviderConfigGetActionGin(g, WalletProviderConfigGetAction)
				walletdefs.WalletProviderConfigBrowseActionGin(g, WalletProviderConfigBrowseAction)
				walletdefs.WalletProviderConfigAwareDeletePreviewActionGin(g, WalletProviderConfigAwareDeletePreviewAction)
				walletdefs.WalletProviderConfigAwareDeleteActionGin(g, WalletProviderConfigAwareDeleteAction)

				walletdefs.GetWalletConfigActionGin(g, GetWalletConfigAction)
				walletdefs.UpdateWalletConfigActionGin(g, UpdateWalletConfigAction)
				walletdefs.AdminCreateWalletActionGin(g, AdminCreateWalletAction)
				walletdefs.AdjustBalanceActionGin(g, AdjustBalanceAction)
				walletdefs.PurchaseActionGin(g, PurchaseAction)

				// Raw route, not emi-generated - see GatewayWebhookImplementation.go.
				// Any, not POST: most gateways call back via a signed POST, but
				// ZarinPal has no server-to-server webhook at all - it redirects the
				// browser via GET (?Authority=...&Status=...). See
				// GatewayWebhookHandler for how both shapes get normalized into the
				// same rawBody the GatewayAdapter interface expects.
				g.Any("/wallet/gateway/:code/webhook", GatewayWebhookHandler)

				return nil
			},
		},

		CliHandlers: []*cli.Command{
			{
				Name:     "wallet",
				Usage:    "Manage wallets, currencies, gateways, and the balance-history ledger",
				Commands: RouterCliManifestManual(),
			},
		},
	}

	module.ProvidePermissionHandler(ALL_WALLET_PERMISSIONS)

	return module
}
