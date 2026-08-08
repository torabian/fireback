package eventbus

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/torabian/fireback/modules/fireback"
)

type SocketConnection struct {
	UserId     string                             `json:"userId"`
	Connection *websocket.Conn                    `json:"-"`
	URW        fireback.UserAccessPerWorkspaceDto `json:"urw"`
	UniqueId   string                             `json:"uniqueId"`
}

var (
	SocketSessionPool = make(map[string]map[string][]*SocketConnection) // workspaceId -> userId -> []*SocketConnection
	socketMutex       sync.Mutex
)
