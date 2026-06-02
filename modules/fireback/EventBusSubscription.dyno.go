package fireback

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"unicode/utf8"
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
	Ctx         *gin.Context
	Socket      *websocket.Conn
	Done        chan bool
	Read        chan EventBusSubscriptionActionReadChan
	QueryParams EventBusSubscriptionActionQuery
}
type EventBusSubscriptionActionHandlerDuplex func(*EventBusSubscriptionActionSession)
type EventBusSubscriptionActionReadChan struct {
	Data  []byte
	Error error
}

func EventBusSubscriptionActionReactiveHandler(factory func(
	session EventBusSubscriptionActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan EventBusSubscriptionActionReadChan)
		done := make(chan bool)
		c, err := upgraderEventBusSubscriptionAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := EventBusSubscriptionActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = EventBusSubscriptionActionQueryFromGin(ctx)
		write, err := factory(session)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- EventBusSubscriptionActionReadChan{
					Data:  data,
					Error: err,
				}
				if err != nil {
					return
				}
			}
		}()
		go func() {
			for {
				select {
				case msg, ok := <-write:
					if !ok {
						// Channel closed; shutdown
						c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
						done <- true
						return
					}
					msgType := websocket.TextMessage
					if !utf8.Valid(msg) {
						msgType = websocket.BinaryMessage
					}
					err := c.WriteMessage(msgType, msg)
					if err != nil {
						// Optionally log the error or send to a logger
						done <- true
						return
					}
				case <-done:
					c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					return
				}
			}
		}()
	}
}

// EventBusSubscriptionActionClientSession is the client-side mirror of
// EventBusSubscriptionActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type EventBusSubscriptionActionClientSession struct {
	Socket *websocket.Conn
	Done   chan bool
	Read   chan EventBusSubscriptionActionReadChan
	Write  chan []byte
}

// EventBusSubscriptionActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type EventBusSubscriptionActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// EventBusSubscriptionActionClient dials the EventBusSubscriptionAction endpoint at baseURL
// (e.g. "ws://localhost:8080" or "https://host" — http/https are auto-rewritten
// to ws/wss) and returns a session whose channels behave like the server's.
//
// Connect, then read and write concurrently. The Read goroutine bails on
// msg.Error (server closed, network drop, etc.); the Done channel fires once
// for any disconnect — use it to unblock main or trigger reconnect logic.
//
// For mTLS, build a *tls.Config with the client keypair + server CA and pass
// it via opts.TLSConfig. The handshake completes before the websocket upgrade
// is sent, so a bad cert fails this call rather than later on Read/Write.
//
//	sess, err := EventBusSubscriptionActionClient(
//	    "wss://hub.example.com",
//	    EventBusSubscriptionActionClientOptions{TLSConfig: tlsCfg},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Reader: pull frames off sess.Read until error.
//	go func() {
//	    for {
//	        msg := <-sess.Read
//	        if msg.Error != nil {
//	            log.Println("read error:", msg.Error)
//	            return
//	        }
//	        log.Printf("server sent %d bytes: %s", len(msg.Data), msg.Data)
//	    }
//	}()
//
//	// Writer: send frames whenever you have something to say. Bytes that
//	// aren't valid UTF-8 are sent as binary frames automatically.
//	sess.Write <- []byte("hello server")
//
//	// Block until the connection drops, then exit. Alternatively, close
//	// sess.Write to initiate a clean shutdown from the client side:
//	//   close(sess.Write)
//	<-sess.Done
func EventBusSubscriptionActionClient(baseURL string, opts *EventBusSubscriptionActionClientOptions) (*EventBusSubscriptionActionClientSession, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}
	u.Path = EventBusSubscriptionActionMeta().URL
	if opts != nil && opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}
	var headers http.Header
	if opts != nil {
		headers = opts.Headers
	}
	dialer := &websocket.Dialer{}
	if opts != nil {
		dialer.TLSClientConfig = opts.TLSConfig
	}
	c, _, err := dialer.Dial(u.String(), headers)
	if err != nil {
		return nil, err
	}
	session := &EventBusSubscriptionActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan EventBusSubscriptionActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- EventBusSubscriptionActionReadChan{
				Data:  data,
				Error: err,
			}
			if err != nil {
				select {
				case session.Done <- true:
				default:
				}
				return
			}
		}
	}()
	// Writer goroutine: drains Write to the socket. Closing Write triggers a
	// clean close handshake; an error or Done signal closes the socket.
	go func() {
		defer c.Close()
		for {
			select {
			case msg, ok := <-session.Write:
				if !ok {
					c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					select {
					case session.Done <- true:
					default:
					}
					return
				}
				msgType := websocket.TextMessage
				if !utf8.Valid(msg) {
					msgType = websocket.BinaryMessage
				}
				if err := c.WriteMessage(msgType, msg); err != nil {
					select {
					case session.Done <- true:
					default:
					}
					return
				}
			case <-session.Done:
				c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				return
			}
		}
	}()
	return session, nil
}
