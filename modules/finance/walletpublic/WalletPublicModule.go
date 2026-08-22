package walletpublic

// See WalletPublic.emi.yml - business logic lives in the various *Implementation.go
// files. No permissions are declared here: every action is gated purely by
// login+ownership (see Scope.go), matching
// modules/entitlement/MySubscriptionsImplementation.go's convention for logged-in-only,
// non-permission-gated actions. Same module layout as modules/category/CategoryModule.go,
// minus EntityBundles - this module owns no entities of its own (see the emi file's
// module doc comment).
import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
	"github.com/urfave/cli/v3"
)

type WalletPublicModuleConfig struct {
	// Add whatever you need to pass to this module for setup.
}

func WalletPublicModuleSetup(cfg *WalletPublicModuleConfig) *application.ModuleProvider {
	module := &application.ModuleProvider{
		Name: "walletpublic",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				walletpublicdefs.CreateWalletActionGin(g, CreateWalletAction)
				walletpublicdefs.MyWalletsActionGin(g, MyWalletsAction)
				walletpublicdefs.MyWorkspaceWalletsActionGin(g, MyWorkspaceWalletsAction)
				walletpublicdefs.GetWalletActionGin(g, GetWalletAction)
				walletpublicdefs.WalletHistoryActionGin(g, WalletHistoryAction)
				walletpublicdefs.WalletPaymentAttemptsActionGin(g, WalletPaymentAttemptsAction)
				walletpublicdefs.TopupActionGin(g, TopupAction)
				walletpublicdefs.CancelPaymentAttemptActionGin(g, CancelPaymentAttemptAction)
				walletpublicdefs.UpdateWalletSettingsActionGin(g, UpdateWalletSettingsAction)
				return nil
			},
		},

		CliHandlers: []*cli.Command{
			{
				Name:     "walletp",
				Usage:    "Owner-facing wallet actions (create/list/topup/history)",
				Commands: RouterCliManifest(),
			},
		},
	}

	return module
}
