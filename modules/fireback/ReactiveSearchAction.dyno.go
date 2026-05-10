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
func ReactiveSearchActionQueryFromGin(c *gin.Context) ReactiveSearchActionQuery {
	return ReactiveSearchActionQueryFromString(c.Request.URL.RawQuery)
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

// WebSocket upgrader
var upgraderReactiveSearchAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type ReactiveSearchActionMessage struct {
	Raw         []byte
	Conn        *websocket.Conn
	MessageType int
	Error       error
}

// Developer handler type
type ReactiveSearchActionHandler func(msg ReactiveSearchActionMessage) error

// Generated handler
func ReactiveSearchAction(r *gin.Engine, handler ReactiveSearchActionHandler) {
	meta := ReactiveSearchActionMeta()
	r.GET(meta.URL, func(c *gin.Context) {
		ws, err := upgraderReactiveSearchAction.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
			return
		}
		defer ws.Close()
		for {
			mt, raw, err := ws.ReadMessage()
			msg := ReactiveSearchActionMessage{
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

type ReactiveSearchActionSession struct {
	Ctx         *gin.Context
	Socket      *websocket.Conn
	Done        chan bool
	Read        chan ReactiveSearchActionReadChan
	QueryParams ReactiveSearchActionQuery
}
type ReactiveSearchActionHandlerDuplex func(*ReactiveSearchActionSession)
type ReactiveSearchActionReadChan struct {
	Data  []byte
	Error error
}

func ReactiveSearchActionReactiveHandler(factory func(
	session ReactiveSearchActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan ReactiveSearchActionReadChan)
		done := make(chan bool)
		c, err := upgraderReactiveSearchAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := ReactiveSearchActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = ReactiveSearchActionQueryFromGin(ctx)
		write, err := factory(session)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- ReactiveSearchActionReadChan{
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

// ReactiveSearchActionClientSession is the client-side mirror of
// ReactiveSearchActionSession. Receive frames on Read, send frames on Write,
// and close Write (or send on Done) to tear the connection down. Done also
// fires when the server closes or the socket errors, so the caller can use it
// as a single disconnect signal.
type ReactiveSearchActionClientSession struct {
	Socket *websocket.Conn
	Done   chan bool
	Read   chan ReactiveSearchActionReadChan
	Write  chan []byte
}

// ReactiveSearchActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type ReactiveSearchActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// ReactiveSearchActionClient dials the ReactiveSearchAction endpoint at baseURL
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
//	sess, err := ReactiveSearchActionClient(
//	    "wss://hub.example.com",
//	    ReactiveSearchActionClientOptions{TLSConfig: tlsCfg},
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
func ReactiveSearchActionClient(baseURL string, opts *ReactiveSearchActionClientOptions) (*ReactiveSearchActionClientSession, error) {
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
	u.Path = ReactiveSearchActionMeta().URL
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
	session := &ReactiveSearchActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan ReactiveSearchActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- ReactiveSearchActionReadChan{
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
