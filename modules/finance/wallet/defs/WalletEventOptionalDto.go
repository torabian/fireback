package walletdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for walletEventOptionalDto
type WalletEventOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The gateway this event came from.
	Gateway emigo.OneNullable[WalletGatewayDto] `json:"gateway" validate:"required" yaml:"gateway"`
	// Gateway-specific event type string, e.g. "payment_intent.succeeded".
	EventType emigo.Nullable[string] `json:"eventType" validate:"required" yaml:"eventType"`
	// The gateway's own id for this event, when it provides one - used to deduplicate webhook retries.
	ExternalEventId emigo.Nullable[string] `json:"externalEventId" yaml:"externalEventId"`
	// The full raw event payload as received from the gateway.
	Payload complexes.JSON `json:"payload" yaml:"payload"`
	// Whether this event was successfully applied (e.g. wallet credited).
	Processed emigo.Nullable[bool] `json:"processed" yaml:"processed"`
	// Error message from the last failed processing attempt, if any.
	ProcessingError emigo.Nullable[string] `json:"processingError" yaml:"processingError"`
	// The payment attempt this event relates to, if identifiable.
	PaymentAttempt emigo.OneNullable[WalletPaymentAttemptDto] `json:"paymentAttempt" yaml:"paymentAttempt"`
	// When this event was received.
	ReceivedAt complexes.XDate `json:"receivedAt" yaml:"receivedAt"`
}

func (x *WalletEventOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletEventOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "gateway",
			Type:        "one?",
			Description: "The gateway this event came from.",
		},
		{
			Name:        prefix + "event-type",
			Type:        "string?",
			Description: "Gateway-specific event type string, e.g. \"payment_intent.succeeded\".",
		},
		{
			Name:        prefix + "external-event-id",
			Type:        "string?",
			Description: "The gateway's own id for this event, when it provides one - used to deduplicate webhook retries.",
		},
		{
			Name:        prefix + "payload",
			Type:        "complex",
			Description: "The full raw event payload as received from the gateway.",
		},
		{
			Name:        prefix + "processed",
			Type:        "bool?",
			Description: "Whether this event was successfully applied (e.g. wallet credited).",
		},
		{
			Name:        prefix + "processing-error",
			Type:        "string?",
			Description: "Error message from the last failed processing attempt, if any.",
		},
		{
			Name:        prefix + "payment-attempt",
			Type:        "one?",
			Description: "The payment attempt this event relates to, if identifiable.",
		},
		{
			Name:        prefix + "received-at",
			Type:        "complex",
			Description: "When this event was received.",
		},
	}
}
func CastWalletEventOptionalDtoFromCli(c emigo.CliCastable) WalletEventOptionalDto {
	data := WalletEventOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("gateway") {
		data.Gateway = emigo.CapturePossibleOneNullable(CastWalletGatewayDtoFromCli, "gateway", c)
	}
	if c.IsSet("event-type") {
		emigo.ParseNullable(c.String("event-type"), &data.EventType)
	}
	if c.IsSet("external-event-id") {
		emigo.ParseNullable(c.String("external-event-id"), &data.ExternalEventId)
	}
	if c.IsSet("payload") {
		if u, ok := any(&data.Payload).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("payload")))
		}
	}
	if c.IsSet("processed") {
		emigo.ParseNullable(c.String("processed"), &data.Processed)
	}
	if c.IsSet("processing-error") {
		emigo.ParseNullable(c.String("processing-error"), &data.ProcessingError)
	}
	if c.IsSet("payment-attempt") {
		data.PaymentAttempt = emigo.CapturePossibleOneNullable(CastWalletPaymentAttemptDtoFromCli, "payment-attempt", c)
	}
	if c.IsSet("received-at") {
		if u, ok := any(&data.ReceivedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("received-at")))
		}
	}
	return data
}
