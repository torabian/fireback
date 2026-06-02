package fireback

import (
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ReactiveSearchAction
 */
func ReactiveSearchActionMeta() struct {
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
		Name:        "ReactiveSearchAction",
		URL:         "/reactive-search",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Reactive search is a general purpose search mechanism for different modules, and could be used in mobile apps or front-end to quickly search for a entity.",
	}
}

/**
 * Query parameters for ReactiveSearchAction
 */
// Query wrapper with private fields
type ReactiveSearchActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ReactiveSearchActionQueryFromString(rawQuery string) ReactiveSearchActionQuery {
	v := ReactiveSearchActionQuery{}
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
func ReactiveSearchActionQueryFromHttp(r *http.Request) ReactiveSearchActionQuery {
	return ReactiveSearchActionQueryFromString(r.URL.RawQuery)
}
func (q ReactiveSearchActionQuery) Values() url.Values {
	return q.values
}
func (q ReactiveSearchActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ReactiveSearchActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ReactiveSearchActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ReactiveSearchActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
}

// Developer handler type
type ReactiveSearchActionHandler func(msg ReactiveSearchActionMessage) error
type ReactiveSearchActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan ReactiveSearchActionReadChan
	QueryParams ReactiveSearchActionQuery
}
type ReactiveSearchActionHandlerDuplex func(*ReactiveSearchActionSession)
type ReactiveSearchActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// ReactiveSearchActionClientSession is the client-side mirror of
// ReactiveSearchActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ReactiveSearchActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan ReactiveSearchActionReadChan
	Write  chan []byte
}
