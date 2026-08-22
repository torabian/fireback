package walletpublic

import (
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
	"github.com/urfave/cli/v3"
)

// RouterCliManifest bundles every walletpublic action's *cli.Command - same convention as
// modules/category/RouterManifest.go.
func RouterCliManifest() []*cli.Command {
	return []*cli.Command{
		walletpublicdefs.CreateWalletActionCliHandler(CreateWalletAction),
		walletpublicdefs.MyWalletsActionCliHandler(MyWalletsAction),
		walletpublicdefs.MyWorkspaceWalletsActionCliHandler(MyWorkspaceWalletsAction),
		walletpublicdefs.GetWalletActionCliHandler(GetWalletAction),
		walletpublicdefs.WalletHistoryActionCliHandler(WalletHistoryAction),
		walletpublicdefs.WalletPaymentAttemptsActionCliHandler(WalletPaymentAttemptsAction),
		walletpublicdefs.TopupActionCliHandler(TopupAction),
		walletpublicdefs.CancelPaymentAttemptActionCliHandler(CancelPaymentAttemptAction),
		walletpublicdefs.UpdateWalletSettingsActionCliHandler(UpdateWalletSettingsAction),
	}
}
