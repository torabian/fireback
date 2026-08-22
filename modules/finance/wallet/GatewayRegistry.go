package wallet

// gatewayRegistry maps a walletGateway.code to the GatewayAdapter implementing it.
// Populated once at startup from WalletModuleConfig.Gateways (see WalletModule.go) - not
// meant to change at runtime, so no locking.
var gatewayRegistry = map[string]GatewayAdapter{}

func registerGatewayAdapters(adapters []GatewayAdapter) {
	for _, a := range adapters {
		gatewayRegistry[a.Code()] = a
	}
}

// GetGatewayAdapter looks up the adapter registered for a walletGateway code. Used by
// walletpublic's topup action and by GatewayWebhookHandler.
func GetGatewayAdapter(code string) (GatewayAdapter, bool) {
	a, ok := gatewayRegistry[code]
	return a, ok
}
