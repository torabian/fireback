package fireback

import (
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
