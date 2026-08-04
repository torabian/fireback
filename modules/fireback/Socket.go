package fireback

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SocketConnection struct {
	UserId     string                    `json:"userId"`
	Connection *websocket.Conn           `json:"-"`
	URW        UserAccessPerWorkspaceDto `json:"urw"`
	UniqueId   string                    `json:"uniqueId"`
}

var (
	SocketSessionPool = make(map[string]map[string][]*SocketConnection) // workspaceId -> userId -> []*SocketConnection
	socketMutex       sync.Mutex
)
