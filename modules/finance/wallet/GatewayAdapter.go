package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
)

// GatewayAdapter is the abstraction point every concrete payment provider (Stripe, a
// manual bank-transfer flow, an on-chain crypto watcher, ...) implements. A
// walletGateway row (see Wallet.emi.yml) only stores a code/name/config - the actual
// behavior lives in whichever GatewayAdapter is registered under that code via
// WalletModuleConfig.Gateways (see WalletModule.go). New providers plug in without ever
// touching the emi schema.
type GatewayAdapter interface {
	// Code must match the walletGateway.code row this adapter implements.
	Code() string

	// InitiatePayment starts a topup/purchase/withdrawal for the given (already
	// persisted, status "pending") attempt and returns whatever the caller needs to
	// complete it. Implementations should not mutate attempt themselves - the caller
	// (TopupImplementation.go) persists whatever GatewayInitResult reports.
	InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*GatewayInitResult, error)

	// VerifyWebhook authenticates and decodes an inbound gateway callback (raw body +
	// headers, since gateways don't sign requests the way our own generated SDK client
	// does) into a normalized GatewayEvent. Implementations must reject anything that
	// doesn't verify (bad signature, unknown event) with an error rather than guessing.
	VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*GatewayEvent, error)
}

// GatewayInitResult is what a wallet owner's topup call gets back - enough to complete
// the payment client-side, gateway-dependent (a redirect-based gateway sets RedirectUrl,
// a client-side-confirmed one sets ClientSecret; either, both, or neither may be set).
type GatewayInitResult struct {
	GatewayReference string
	RedirectUrl      string
	ClientSecret     string
	RawRequest       any
	RawResponse      any
}

// GatewayEvent is the normalized shape VerifyWebhook decodes a raw callback into.
// ExternalEventId, when the gateway provides one, is used to deduplicate webhook
// retries (see walletEvent.externalEventId).
type GatewayEvent struct {
	ExternalEventId  string
	EventType        string
	GatewayReference string
	Succeeded        bool
	FailureReason    string
	Payload          any
}

// MockGatewayAdapter is a synchronous, always-succeeds adapter for CLI/dev testing and
// for exercising the full topup->webhook->credit path without a real provider. Code
// "mock" must have a matching walletGateway row before it can be used.
type MockGatewayAdapter struct{}

func (MockGatewayAdapter) Code() string { return "mock" }

func (MockGatewayAdapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*GatewayInitResult, error) {
	return &GatewayInitResult{
		GatewayReference: "mock-" + attempt.UniqueId,
	}, nil
}

func (MockGatewayAdapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*GatewayEvent, error) {
	// The mock gateway has no real webhook - GatewayWebhookAction accepts a
	// {"gatewayReference": "...", "succeeded": true} JSON body directly for CLI/dev use.
	return decodeMockWebhookBody(rawBody)
}

// ManualGatewayAdapter backs bank-transfer-style topups: InitiatePayment returns nothing
// but instructions live in the gateway's own config/description, and an admin marks the
// attempt succeeded by hand (see AdjustBalance/admin tooling - a manual gateway attempt is
// completed the same way any other webhook-driven one is, just triggered by an admin
// action instead of an inbound HTTP callback).
type ManualGatewayAdapter struct{}

func (ManualGatewayAdapter) Code() string { return "manual" }

func (ManualGatewayAdapter) InitiatePayment(ctx context.Context, attempt *walletdefs.WalletPaymentAttemptEntity, gateway *walletdefs.WalletGatewayEntity) (*GatewayInitResult, error) {
	return &GatewayInitResult{GatewayReference: attempt.UniqueId}, nil
}

func (ManualGatewayAdapter) VerifyWebhook(ctx context.Context, rawBody []byte, headers http.Header) (*GatewayEvent, error) {
	return decodeMockWebhookBody(rawBody)
}

// mockWebhookBody is the trivial JSON shape the mock/manual adapters' "webhook" accepts -
// no signature, since neither is a real external provider.
type mockWebhookBody struct {
	GatewayReference string `json:"gatewayReference"`
	EventType        string `json:"eventType"`
	Succeeded        bool   `json:"succeeded"`
	FailureReason    string `json:"failureReason"`
}

func decodeMockWebhookBody(rawBody []byte) (*GatewayEvent, error) {
	var body mockWebhookBody
	if err := json.Unmarshal(rawBody, &body); err != nil {
		return nil, fmt.Errorf("invalid mock/manual webhook body: %w", err)
	}
	if body.GatewayReference == "" {
		return nil, fmt.Errorf("gatewayReference is required")
	}
	eventType := body.EventType
	if eventType == "" {
		if body.Succeeded {
			eventType = "payment.succeeded"
		} else {
			eventType = "payment.failed"
		}
	}
	return &GatewayEvent{
		GatewayReference: body.GatewayReference,
		EventType:        eventType,
		Succeeded:        body.Succeeded,
		FailureReason:    body.FailureReason,
		Payload:          body,
	}, nil
}
