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
* Action to communicate with the action EventBusSubscription2Action
 */
func EventBusSubscription2ActionMeta() struct {
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
		Name:        "EventBusSubscription2Action",
		URL:         "/ws2",
		Method:      "REACTIVE",
		CliName:     "",
		Description: "Connects a client to all events related to their user profile, or workspace they are in",
	}
}

/**
 * Query parameters for EventBusSubscription2Action
 */
// Query wrapper with private fields
type EventBusSubscription2ActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EventBusSubscription2ActionQueryFromString(rawQuery string) EventBusSubscription2ActionQuery {
	v := EventBusSubscription2ActionQuery{}
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
func EventBusSubscription2ActionQueryFromGin(c *gin.Context) EventBusSubscription2ActionQuery {
	return EventBusSubscription2ActionQueryFromString(c.Request.URL.RawQuery)
}
func EventBusSubscription2ActionQueryFromHttp(r *http.Request) EventBusSubscription2ActionQuery {
	return EventBusSubscription2ActionQueryFromString(r.URL.RawQuery)
}
func (q EventBusSubscription2ActionQuery) Values() url.Values {
	return q.values
}
func (q EventBusSubscription2ActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EventBusSubscription2ActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EventBusSubscription2ActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

// WebSocket upgrader
var upgraderEventBusSubscription2Action = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type EventBusSubscription2ActionMessage struct {
	Raw         []byte
	Conn        *websocket.Conn
	MessageType int
	Error       error
	QueryParams EventBusSubscription2ActionQuery
}

// Developer handler type
type EventBusSubscription2ActionHandler func(msg EventBusSubscription2ActionMessage) error

// Generated handler
func EventBusSubscription2Action(r *gin.Engine, handler EventBusSubscription2ActionHandler) {
	meta := EventBusSubscription2ActionMeta()
	r.GET(meta.URL, func(c *gin.Context) {
		ws, err := upgraderEventBusSubscription2Action.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
			return
		}
		defer ws.Close()
		for {
			mt, raw, err := ws.ReadMessage()
			msg := EventBusSubscription2ActionMessage{
				Conn:        ws,
				Raw:         raw,
				Error:       err,
				MessageType: mt,
			}
			msg.QueryParams = EventBusSubscription2ActionQueryFromGin(c)
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

type EventBusSubscription2ActionSession struct {
	In          <-chan EventBusSubscription2ActionMessage
	Out         chan<- EventBusSubscription2ActionMessage
	Done        <-chan struct{}
	Close       func(err error)
	QueryParams EventBusSubscription2ActionQuery
	G           *gin.Context
	Socket      *websocket.Conn
}
type EventBusSubscription2ActionHandlerDuplex func(*EventBusSubscription2ActionSession)

// EventBusSubscription2ActionDuplex upgrades the HTTP connection to a WebSocket and
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
//	external.EventBusSubscription2ActionDuplex(r, func(ctx *external.EventBusSubscription2ActionSession) {
//		for {
//			select {
//			case msg, ok := <-ctx.In:
//				if !ok {
//					return // client disconnected
//				}
//				ctx.Out <- external.EventBusSubscription2ActionMessage{
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
func EventBusSubscription2ActionDuplex(r *gin.Engine, handler EventBusSubscription2ActionHandlerDuplex) {
	meta := EventBusSubscription2ActionMeta()
	// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.
	r.GET(meta.URL, func(ctx *gin.Context) {
		EventBusSubscription2ActionDuplexGinHandler(ctx, handler)
	})
}
func EventBusSubscription2ActionDuplexGinHandler(c *gin.Context, handler EventBusSubscription2ActionHandlerDuplex) {
	ws, err := upgraderEventBusSubscription2Action.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
	in := make(chan EventBusSubscription2ActionMessage)
	out := make(chan EventBusSubscription2ActionMessage)
	done := make(chan struct{})
	session := &EventBusSubscription2ActionSession{
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
	session.QueryParams = EventBusSubscription2ActionQueryFromGin(c)
	// Read loop
	go func() {
		defer close(in)
		for {
			mt, raw, err := ws.ReadMessage()
			in <- EventBusSubscription2ActionMessage{MessageType: mt, Raw: raw, Error: err}
		}
	}()
	// Write loop
	go func() {
		for msg := range out {
			if err := ws.WriteMessage(msg.MessageType, msg.Raw); err != nil {
				// When message is -1, means it's internal error coming out
				in <- EventBusSubscription2ActionMessage{MessageType: -1, Error: err}
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
