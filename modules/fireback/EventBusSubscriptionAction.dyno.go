package fireback

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
func EventBusSubscriptionActionQueryFromGin(c *gin.Context) EventBusSubscriptionActionQuery {
	return EventBusSubscriptionActionQueryFromString(c.Request.URL.RawQuery)
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

// WebSocket upgrader
var upgraderEventBusSubscriptionAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type EventBusSubscriptionActionMessage struct {
	Raw         []byte
	Conn        *websocket.Conn
	MessageType int
	Error       error
	QueryParams EventBusSubscriptionActionQuery
}

// Developer handler type
type EventBusSubscriptionActionHandler func(msg EventBusSubscriptionActionMessage) error

// Generated handler
func EventBusSubscriptionAction(r *gin.Engine, handler EventBusSubscriptionActionHandler) {
	meta := EventBusSubscriptionActionMeta()
	r.GET(meta.URL, func(c *gin.Context) {
		ws, err := upgraderEventBusSubscriptionAction.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
			return
		}
		defer ws.Close()
		for {
			mt, raw, err := ws.ReadMessage()
			msg := EventBusSubscriptionActionMessage{
				Conn:        ws,
				Raw:         raw,
				Error:       err,
				MessageType: mt,
			}
			msg.QueryParams = EventBusSubscriptionActionQueryFromGin(c)
			// Provide raw message to developer handler
			if err := handler(msg); err != nil {
				errMsg := fmt.Sprintf("handler error: %v", err)
				if writeErr := ws.WriteMessage(mt, []byte(errMsg)); writeErr != nil {
					break
				}
			}
		}
	})
}

type EventBusSubscriptionActionSession struct {
	In          <-chan EventBusSubscriptionActionMessage
	Out         chan<- EventBusSubscriptionActionMessage
	Done        <-chan struct{}
	Close       func(err error)
	QueryParams EventBusSubscriptionActionQuery
	G           *gin.Context
	Socket      *websocket.Conn
}
type EventBusSubscriptionActionHandlerDuplex func(*EventBusSubscriptionActionSession)

// EventBusSubscriptionActionDuplex upgrades the HTTP connection to a WebSocket and
// exposes it as a full-duplex, blocking session.
//
// The provided handler owns the lifetime of the connection.
// The WebSocket remains open as long as the handler is running.
// Returning from the handler will close the connection.
//
// Session channels:
//   - ctx.In   : incoming messages from the client (closed on disconnect)
//   - ctx.Out  : outgoing messages to the client (blocking send)
//   - ctx.Done : closed when the server terminates the session
//
// Usage pattern:
//
//	external.EventBusSubscriptionActionDuplex(r, func(ctx *external.EventBusSubscriptionActionSession) {
//		for {
//			select {
//			case msg, ok := <-ctx.In:
//				if !ok {
//					return // client disconnected
//				}
//				ctx.Out <- external.EventBusSubscriptionActionMessage{
//					MessageType: websocket.TextMessage,
//					Raw:         msg.Raw,
//				}
//
//			case <-ctx.Done:
//				return // server-side close
//			}
//		}
//	})
//
// Important:
//   - Always read the generated code, don't use blindly.
//   - If there is an error on write, you'll get a message back, with message type -1 (instead of default websocket message type int.)
//   - The handler MUST block (typically via a loop).
//   - Returning from the handler closes the WebSocket.
//   - Do not treat this as a per-message callback.
func EventBusSubscriptionActionDuplex(r *gin.Engine, handler EventBusSubscriptionActionHandlerDuplex) {
	meta := EventBusSubscriptionActionMeta()
	// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.
	r.GET(meta.URL, func(ctx *gin.Context) {
		EventBusSubscriptionActionDuplexGinHandler(ctx, handler)
	})
}
func EventBusSubscriptionActionDuplexGinHandler(c *gin.Context, handler EventBusSubscriptionActionHandlerDuplex) {
	ws, err := upgraderEventBusSubscriptionAction.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
	in := make(chan EventBusSubscriptionActionMessage)
	out := make(chan EventBusSubscriptionActionMessage)
	done := make(chan struct{})
	session := &EventBusSubscriptionActionSession{
		In:     in,
		Out:    out,
		Done:   done,
		Socket: ws,
		G:      c,
		Close: func(err error) {
			close(done)
			ws.Close()
		},
	}
	session.QueryParams = EventBusSubscriptionActionQueryFromGin(c)
	// Read loop
	go func() {
		defer close(in)
		for {
			mt, raw, err := ws.ReadMessage()
			in <- EventBusSubscriptionActionMessage{MessageType: mt, Raw: raw, Error: err}
		}
	}()
	// Write loop
	go func() {
		for msg := range out {
			if err := ws.WriteMessage(msg.MessageType, msg.Raw); err != nil {
				// When message is -1, means it's internal error coming out
				in <- EventBusSubscriptionActionMessage{MessageType: -1, Error: err}
				return
			}
		}
	}()
	// Run developer code (blocking)
	handler(session)
	// Cleanup
	close(out)
	ws.Close()
}
