package fireback

import (
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action EventBusSubscriptionAction
 */
func EventBusSubscriptionActionMeta() struct {
	Name        string
	URL         string
	Method      string
	CliName     string
	Description string
} {
	return struct {
		Name        string
		URL         string
		Method      string
		CliName     string
		Description string
	}{
		Name:        "EventBusSubscriptionAction",
		URL:         "/ws",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Connects a client to all events related to their user profile, or workspace they are in",
	}
}

/**
 * Query parameters for EventBusSubscriptionAction
 */
// Query wrapper with private fields
type EventBusSubscriptionActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EventBusSubscriptionActionQueryFromString(rawQuery string) EventBusSubscriptionActionQuery {
	v := EventBusSubscriptionActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func EventBusSubscriptionActionQueryFromHttp(r *http.Request) EventBusSubscriptionActionQuery {
	return EventBusSubscriptionActionQueryFromString(r.URL.RawQuery)
}
func (q EventBusSubscriptionActionQuery) Values() url.Values {
	return q.values
}
func (q EventBusSubscriptionActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EventBusSubscriptionActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EventBusSubscriptionActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EventBusSubscriptionActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
}

// Developer handler type
type EventBusSubscriptionActionHandler func(msg EventBusSubscriptionActionMessage) error
type EventBusSubscriptionActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan EventBusSubscriptionActionReadChan
	QueryParams EventBusSubscriptionActionQuery
}
type EventBusSubscriptionActionHandlerDuplex func(*EventBusSubscriptionActionSession)
type EventBusSubscriptionActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// EventBusSubscriptionActionClientSession is the client-side mirror of
// EventBusSubscriptionActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type EventBusSubscriptionActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan EventBusSubscriptionActionReadChan
	Write  chan []byte
}
