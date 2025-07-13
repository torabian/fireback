package payment

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/torabian/fireback/modules/fireback"
)

type PaymentProvider interface {
	CreateSession(invoice InvoiceEntity, options CreateSessionOptions) (*StartPaymentResult, error)
	VerifySession(sessionID string, options CreateSessionOptions) (paid bool, err error)
	Name() string

	GetOptions(invoice InvoiceEntity) CreateSessionOptions
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println("Open this URL manually:", url)
	}
}

func handlePaymentResult(provider PaymentProvider, sessionIDFromCallback string, q fireback.QueryDSL) *fireback.IError {

	var transaction *InvoiceTransactionEntity = nil
	err := fireback.
		GetRef(q).
		Model(InvoiceTransactionEntity{}).
		Preload("Invoice").
		Where(InvoiceTransactionEntity{ExternalSessionId: sessionIDFromCallback}).
		Find(&transaction).Error

	if err != nil {
		return fireback.CastToIError(err)
	}

	if transaction == nil {
		return fireback.Create401Error(&PaymentMessages.TransactionNotFound, []string{})
	}

	paid, err3 := provider.VerifySession(sessionIDFromCallback, provider.GetOptions(*transaction.Invoice))
	if err3 != nil || !paid {
		return fireback.CastToIError(err3)
	}

	transaction.PaidAt = fireback.XDateTimeFromTime(time.Now())
	InvoiceTransactionActions.Update(q, transaction)

	transaction.Invoice.FinalStatus = "payed"
	InvoiceActions.Update(q, transaction.Invoice)

	if OnInvoiceStatusChange != nil {
		OnInvoiceStatusChange(transaction.Invoice)
	}

	if q.G != nil && transaction.Invoice.RedirectAfterSuccess != "" {
		q.G.Redirect(301, transaction.Invoice.RedirectAfterSuccess)
	}

	return nil
}

type StartPaymentResult struct {
	URL          string                                 // optional
	SessionID    string                                 // optional
	Instructions string                                 // optional (for manual or USSD)
	Next         func() (finalStatus string, err error) // optional async poller
	Transaction  *InvoiceTransactionEntity
}

func initiatePaymentFromInvoice(
	provider PaymentProvider,
	invoiceId string,
	q fireback.QueryDSL,
) (*StartPaymentResult, *fireback.IError) {
	q.UniqueId = invoiceId
	invoice, err := InvoiceActions.GetOne(q)

	if err != nil {
		return nil, err
	}

	result, err2 := provider.CreateSession(*invoice, provider.GetOptions(*invoice))
	if err2 != nil {
		return nil, fireback.CastToIError(err2)
	}

	transaction, err4 := InvoiceTransactionActions.Create(&InvoiceTransactionEntity{
		Provider:          provider.Name(),
		InvoiceId:         fireback.NewString(invoice.UniqueId),
		Amount:            invoice.Amount.Amount.Int64,
		ExternalSessionId: result.SessionID,
	}, q)
	if err4 != nil {
		return nil, err4
	}

	result.Transaction = transaction

	return result, nil
}
