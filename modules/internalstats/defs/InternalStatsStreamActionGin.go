package internalstatsdefs

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"unicode/utf8"
)

/**
* Action to communicate with the action InternalStatsStreamAction
 */
// WebSocket upgrader
var upgraderInternalStatsStreamAction = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (x *InternalStatsStreamActionMessage) Connection() *websocket.Conn {
	return x.Conn.(*websocket.Conn)
}
func (x *InternalStatsStreamActionClientSession) Connection() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}
func (x *InternalStatsStreamActionSession) GinCtx() *gin.Context {
	return x.Ctx.(*gin.Context)
}
func (x *InternalStatsStreamActionSession) GetSocket() *websocket.Conn {
	return x.Socket.(*websocket.Conn)
}

// Generated handler
func InternalStatsStreamActionGin(r *gin.Engine, factory func(
	session InternalStatsStreamActionSession,
) (chan []byte, error)) {
	meta := InternalStatsStreamActionMeta()
	r.GET(meta.URL, InternalStatsStreamActionReactiveHandler(factory))
}
func InternalStatsStreamActionReactiveHandler(factory func(
	session InternalStatsStreamActionSession,
) (chan []byte, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		read := make(chan InternalStatsStreamActionReadChan)
		done := make(chan bool)
		c, err := upgraderInternalStatsStreamAction.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			c.Close()
			return
		}
		session := InternalStatsStreamActionSession{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = InternalStatsStreamActionQueryFromHttp(ctx.Request)
		write, err := factory(session)
		if err != nil {
			// factory rejected the session (e.g. authorization failed) -
			// report it over the real websocket connection (safe: this is
			// gorilla's own Conn, not gin's response writer, which is no
			// longer usable once Upgrade succeeded above) and tear the
			// connection down instead of falling through to the read/write
			// loops below with a nil write channel, which would otherwise
			// leave a half-open, non-functional connection and a leaked
			// reader goroutine.
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, err.Error()))
			c.Close()
			return
		}
		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- InternalStatsStreamActionReadChan{
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

// InternalStatsStreamActionClientOptions configures a client dial. All fields are
// optional — pass a zero value for a plaintext, header-less ws:// connection.
//
// TLSConfig governs the TLS handshake when dialing wss://. Set Certificates
// and RootCAs (and optionally ServerName) for mTLS; leave nil for ws://.
type InternalStatsStreamActionClientOptions struct {
	Query     url.Values
	TLSConfig *tls.Config
	Headers   http.Header
}

// InternalStatsStreamActionClient dials the InternalStatsStreamAction endpoint at baseURL
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
//	sess, err := InternalStatsStreamActionClient(
//	    "wss://hub.example.com",
//	    InternalStatsStreamActionClientOptions{TLSConfig: tlsCfg},
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
func InternalStatsStreamActionClient(baseURL string, opts *InternalStatsStreamActionClientOptions) (*InternalStatsStreamActionClientSession, error) {
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
	u.Path = InternalStatsStreamActionMeta().URL
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
	session := &InternalStatsStreamActionClientSession{
		Socket: c,
		Done:   make(chan bool, 1),
		Read:   make(chan InternalStatsStreamActionReadChan),
		Write:  make(chan []byte, 16),
	}
	// Reader goroutine: pumps frames from the socket into Read. On error it
	// forwards the error frame, signals Done, and exits.
	go func() {
		for {
			_, data, err := c.ReadMessage()
			session.Read <- InternalStatsStreamActionReadChan{
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
