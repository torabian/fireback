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
		c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
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
