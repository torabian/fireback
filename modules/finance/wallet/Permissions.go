package wallet

import "github.com/torabian/fireback/modules/fireback/application"

// The newer emi entity compiler doesn't auto-generate PermissionInfo constants (unlike
// the old Module3 compiler), so these are declared by hand here - same convention as
// modules/category/Permissions.go.
//
// walletCurrency and walletGateway get full per-verb permissions (they're independently
// admin-managed resources). wallet/walletTransaction/walletEvent/walletPaymentAttempt
// only ever expose Get/Browse publicly to an admin/support caller (their Create/Update/
// Delete either don't exist - see the "actions" overrides in Wallet.emi.yml - or only run
// through the internal wallet.Purchase engine), so they share one read-only permission
// instead of one per entity.
var PERM_ROOT_WALLET = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.*",
	Name:        "Entire wallet module (*)",
	Description: "",
}

var PERM_ROOT_WALLET_CURRENCY_CREATE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.currency.create",
	Name:        "Create wallet currency",
	Description: "",
}
var PERM_ROOT_WALLET_CURRENCY_UPDATE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.currency.update",
	Name:        "Update wallet currency",
	Description: "",
}
var PERM_ROOT_WALLET_CURRENCY_QUERY = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.currency.query",
	Name:        "Query wallet currencies",
	Description: "",
}
var PERM_ROOT_WALLET_CURRENCY_DELETE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.currency.delete",
	Name:        "Delete wallet currency",
	Description: "",
}

var PERM_ROOT_WALLET_GATEWAY_CREATE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.gateway.create",
	Name:        "Register wallet gateway",
	Description: "",
}
var PERM_ROOT_WALLET_GATEWAY_UPDATE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.gateway.update",
	Name:        "Update wallet gateway",
	Description: "",
}
var PERM_ROOT_WALLET_GATEWAY_QUERY = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.gateway.query",
	Name:        "Query wallet gateways",
	Description: "",
}
var PERM_ROOT_WALLET_GATEWAY_DELETE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.gateway.delete",
	Name:        "Delete wallet gateway",
	Description: "",
}

// PERM_ROOT_WALLET_ADMIN_QUERY guards every admin/support read-only view over
// wallet/walletTransaction/walletEvent/walletPaymentAttempt (Get/Browse only - see the
// entity feature overrides in Wallet.emi.yml for why these have no public Create/Update).
var PERM_ROOT_WALLET_ADMIN_QUERY = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.admin.query",
	Name:        "View all wallets, ledger entries, payment attempts and gateway events",
	Description: "",
}

// PERM_ROOT_WALLET_ADMIN_DELETE guards AwareDelete/AwareDeletePreview on the wallet
// entity itself - deleting a wallet (as opposed to closing it) is a rare, destructive,
// support-only operation.
var PERM_ROOT_WALLET_ADMIN_DELETE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.admin.delete",
	Name:        "Delete a wallet",
	Description: "",
}

// PERM_ROOT_WALLET_PURCHASE is required by the purchase HTTP action - it's meant for
// trusted internal/service callers (other modules debiting a wallet for a sale), never for
// a wallet owner directly. Owners can only spend indirectly, through whatever feature
// calls wallet.Purchase on their behalf.
var PERM_ROOT_WALLET_PURCHASE = application.PermissionInfo{
	CompleteKey: "root.modules.wallet.purchase",
	Name:        "Debit a wallet for an internal purchase",
	Description: "",
}

// ALL_WALLET_PERMISSIONS is registered with the module (see WalletModule.go's
// ProvidePermissionHandler call) so these show up in the capability tree roles can be
// granted from. adjustBalance and the walletConfig get/update actions are intentionally
// absent here - they're gated by SecurityModel.AllowOnRoot (true root-only), not by a
// grantable permission, matching modules/entitlement/GenerateActivationCodeImplementation.go's
// convention.
var ALL_WALLET_PERMISSIONS = []application.PermissionInfo{
	PERM_ROOT_WALLET,
	PERM_ROOT_WALLET_CURRENCY_CREATE,
	PERM_ROOT_WALLET_CURRENCY_UPDATE,
	PERM_ROOT_WALLET_CURRENCY_QUERY,
	PERM_ROOT_WALLET_CURRENCY_DELETE,
	PERM_ROOT_WALLET_GATEWAY_CREATE,
	PERM_ROOT_WALLET_GATEWAY_UPDATE,
	PERM_ROOT_WALLET_GATEWAY_QUERY,
	PERM_ROOT_WALLET_GATEWAY_DELETE,
	PERM_ROOT_WALLET_ADMIN_QUERY,
	PERM_ROOT_WALLET_ADMIN_DELETE,
	PERM_ROOT_WALLET_PURCHASE,
}
