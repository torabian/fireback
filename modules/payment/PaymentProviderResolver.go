package payment

import (
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

// Extend this function to detect the different payment methods enabled in the project
// for now it's only stripe
func resolveProvider(invoiceId string, mode string) (PaymentProvider, *fireback.IError) {

	if mode == "http" {

		config, err := PaymentConfigActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: "root"})
		if err != nil || config == nil {
			return nil, fireback.Create401Error(&PaymentMessages.PaymentProviderIsNotAvailable, []string{})
		}

		provider := &StripeProvider{
			CompleteRedirectUrl: strings.ReplaceAll(config.StripeCallbackUrl, "{invoiceId}", invoiceId),
			ApiKey:              config.StripeSecretKey,
		}

		return provider, nil
	}

	provider := &StripeProvider{
		CompleteRedirectUrl: "http://localhost:45678/payment-callback",
	}

	return provider, nil
}
