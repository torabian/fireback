package internalstatsdefs

import (
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action InternalStatsStreamAction
 */
func InternalStatsStreamActionMeta() struct {
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
		Name:        "InternalStatsStreamAction",
		URL:         "/internal-stats/stream",
		Method:      "REACTIVE",
		CliName:     "stream",
		Description: "Reactive (websocket) live feed of the same snapshot InternalStatsSnapshot returns, pushed on an interval (see InternalStatsModuleConfig.Interval) until the client disconnects. Each frame is the JSON-encoded snapshot object - identical shape to InternalStatsSnapshot's out fields. Root workspace token required by default - see InternalStatsModuleConfig.Authorize.",
	}
}

/**
 * Query parameters for InternalStatsStreamAction
 */
// Query wrapper with private fields
type InternalStatsStreamActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func InternalStatsStreamActionQueryFromString(rawQuery string) InternalStatsStreamActionQuery {
	v := InternalStatsStreamActionQuery{}
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
func InternalStatsStreamActionQueryFromHttp(r *http.Request) InternalStatsStreamActionQuery {
	return InternalStatsStreamActionQueryFromString(r.URL.RawQuery)
}
func (q InternalStatsStreamActionQuery) Values() url.Values {
	return q.values
}
func (q InternalStatsStreamActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *InternalStatsStreamActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *InternalStatsStreamActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type InternalStatsStreamActionMessage struct {
	Raw []byte
	// Conn *websocket.Conn
	Conn        interface{}
	MessageType int
	Error       error
}

// Developer handler type
type InternalStatsStreamActionHandler func(msg InternalStatsStreamActionMessage) error
type InternalStatsStreamActionSession struct {
	// Ctx    *gin.Context
	// Socket *websocket.Conn
	Ctx         interface{}
	Socket      interface{}
	Done        chan bool
	Read        chan InternalStatsStreamActionReadChan
	QueryParams InternalStatsStreamActionQuery
}
type InternalStatsStreamActionHandlerDuplex func(*InternalStatsStreamActionSession)
type InternalStatsStreamActionReadChan struct {
	Data        []byte
	Error       error
	MessageType int
}

// InternalStatsStreamActionClientSession is the client-side mirror of
// InternalStatsStreamActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type InternalStatsStreamActionClientSession struct {
	// Socket *websocket.Conn
	Socket interface{}
	Done   chan bool
	Read   chan InternalStatsStreamActionReadChan
	Write  chan []byte
}
