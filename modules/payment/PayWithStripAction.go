package payment

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)

func init() {

}

type StripeProvider struct {
	CompleteRedirectUrl string
	ApiKey              string
}

func (s *StripeProvider) Name() string {
	return "stripe"
}

func (s *StripeProvider) GetOptions(invoice InvoiceEntity) CreateSessionOptions {
	return CreateSessionOptions{
		RedirectUrl: s.CompleteRedirectUrl,
	}
}

type CreateSessionOptions struct {
	RedirectUrl string
}

func (s *StripeProvider) CreateSession(invoice InvoiceEntity, options CreateSessionOptions) (*StartPaymentResult, error) {
	stripe.Key = s.ApiKey

	domain := options.RedirectUrl

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(invoice.Amount.Currency.String),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(invoice.Title),
					},
					UnitAmount: stripe.Int64(invoice.Amount.Amount.Int64), // in cents (or smallest currency unit)
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?is_callback=true&type=success&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(domain + "?is_callback=true&type=cancel&session_id={CHECKOUT_SESSION_ID}"),
	}

	sess, err := session.New(params)
	if err != nil {
		return nil, err
	}
	return &StartPaymentResult{
		URL:       sess.URL,
		SessionID: sess.ID,
	}, nil
}

func (s *StripeProvider) VerifySession(sessionID string, options CreateSessionOptions) (bool, error) {
	stripe.Key = s.ApiKey
	s2, err := session.Get(sessionID, nil)
	if err != nil {
		return false, err
	}
	return s2.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid, nil
}

func cliTemporaryHttpListener() string {
	resultCh := make(chan string)

	http.HandleFunc("/payment-callback", func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.URL.Query().Get("session_id")
		typeof := r.URL.Query().Get("type")

		if typeof == "success" {

			fmt.Fprintln(w, "✅ Payment complete. You can close this tab. ")

			// Send the ID to the result channel
			resultCh <- sessionID
		} else {
			fmt.Fprintln(w, "❌ Payment canceled. You can close this tab.")
			resultCh <- "❌ Payment canceled."
		}
	})

	server := &http.Server{Addr: ":45678"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// Wait for the callback
	result := <-resultCh

	// Cleanly shutdown server after receiving result
	_ = server.Close()
	return result
}
