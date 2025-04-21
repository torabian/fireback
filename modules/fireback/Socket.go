package fireback

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type SocketConnection struct {
	UserId     string                    `json:"userId"`
	Connection *websocket.Conn           `json:"-"`
	URW        UserAccessPerWorkspaceDto `json:"urw"`
}

var (
	SocketSessionPool = make(map[string]map[string][]*SocketConnection) // workspaceId -> userId -> []*SocketConnection
	socketMutex       sync.Mutex
)

func SocketConnectEndpoint(c *gin.Context) {
	wsURLParam, _ := url.ParseQuery(c.Request.URL.RawQuery)

	workspaceId := c.Request.Header.Get("Workspace-id")
	token := c.Request.Header.Get("authorization")

	if val, ok := wsURLParam["token"]; ok && len(val) == 1 {
		token = val[0]
	}
	if val, ok := wsURLParam["workspaceId"]; ok && len(val) == 1 {
		workspaceId = val[0]
	}

	context := &AuthContextDto{
		WorkspaceId:  workspaceId,
		Token:        token,
		Capabilities: []PermissionInfo{},
	}

	res, err := WithAuthorizationPure(context)
	if err != nil {
		resp, _ := json.MarshalIndent(gin.H{"error": err}, "", "  ")
		http.Error(c.Writer, string(resp), int(err.HttpCode))
		return
	}

	ws, err2 := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err2 != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	GetEventBusInstance().AddUser(SERVER_INSTANCE, res.UserId.String)

	socket := &SocketConnection{
		UserId:     res.UserId.String,
		Connection: ws,
		URW:        *res.UserAccessPerWorkspace,
	}

	socketMutex.Lock()
	if SocketSessionPool[workspaceId] == nil {
		SocketSessionPool[workspaceId] = make(map[string][]*SocketConnection)
	}
	SocketSessionPool[workspaceId][res.UserId.String] = append(SocketSessionPool[workspaceId][res.UserId.String], socket)
	socketMutex.Unlock()

	for {
		var msg interface{}
		if err := ws.ReadJSON(&msg); err != nil {
			// WebSocket close error?
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				LOG.Debug("WebSocket closed normally", zap.String("user", res.UserId.String))

				// Only break in this case
				break
			} else {
				// We need to handle this kinda.
				LOG.Debug("WebSocket read error", zap.Error(err))
				ws.WriteJSON("Socket interaction with webserver only supports json at this moment.")
			}

		}
		LOG.Debug("Message from socket connection:", zap.Any("message", msg))
	}

	// Cleanup
	ws.Close()
	socketMutex.Lock()
	conns := SocketSessionPool[workspaceId][res.UserId.String]
	for i, conn := range conns {
		if conn.Connection == ws {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(conns) == 0 {
		delete(SocketSessionPool[workspaceId], res.UserId.String)
	} else {
		SocketSessionPool[workspaceId][res.UserId.String] = conns
	}
	if len(SocketSessionPool[workspaceId]) == 0 {
		delete(SocketSessionPool, workspaceId)
	}
	socketMutex.Unlock()

	GetEventBusInstance().RemoveUser(SERVER_INSTANCE, res.UserId.String)
}

func HandleSocket(e *gin.Engine) {
	e.Any("ws", SocketConnectEndpoint)
}
