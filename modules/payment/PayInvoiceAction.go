package payment

import (
	"fmt"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	PayInvoiceActionImp = PayInvoiceAction
}

// Implement this if you want to know about an invoice status.
// Even better is that passing it via Payment Module config
var OnInvoiceStatusChange func(ie *InvoiceEntity) = nil

func payInvoiceHTTP(q fireback.QueryDSL) (string, *fireback.IError) {
	invoiceId := q.G.Param("uniqueId")
	q.UniqueId = invoiceId

	provider, err := resolveProvider(invoiceId, "http")
	if err != nil {
		return "", err
	}

	if q.G.Query("is_callback") == "true" {

		handleErr := handlePaymentResult(provider, q.G.Query("session_id"), q)

		if handleErr != nil {
			return "", handleErr

		} else {

			return "✅ Payment initialized. Awaiting confirmation...", nil
		}
	}

	result, err := initiatePaymentFromInvoice(provider, invoiceId, q)
	if err != nil {
		return "", err
	}

	if result.URL != "" {
		q.G.Redirect(301, result.URL)
	} else if result.Instructions != "" {
		// Maybe an html page here instead
		return result.Instructions, nil
	} else {
		return "✅ Payment initialized. Awaiting confirmation...", nil
	}

	return "", nil
}

func payInvoiceCLI(q fireback.QueryDSL) (string, *fireback.IError) {
	invoiceId := q.C.String("invoice-id")
	q.UniqueId = invoiceId

	provider, err := resolveProvider(invoiceId, "cli")
	if err != nil {
		return "", err
	}

	result, err := initiatePaymentFromInvoice(provider, invoiceId, q)
	if err != nil {
		return "", err
	}

	if result.URL != "" {
		fmt.Println("Opening", result.URL)
		openBrowser(result.URL)
	} else if result.Instructions != "" {
		fmt.Println("🔔 Follow these steps to pay:")
		fmt.Println(result.Instructions)
	} else {
		fmt.Println("✅ Payment initialized. Awaiting confirmation...")
	}

	sessionID := cliTemporaryHttpListener()
	return "", handlePaymentResult(provider, sessionID, q)
}

func PayInvoiceAction(q fireback.QueryDSL) (string, *fireback.IError) {
	if q.G != nil {
		return payInvoiceHTTP(q)
	}
	if q.C != nil {
		return payInvoiceCLI(q)
	}
	return "", nil
}
