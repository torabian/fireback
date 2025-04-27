package fireback

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func init() {
	// Override the implementation with our actual code.
	EventBusSubscriptionActionImp = EventBusSubscriptionAction
}

func cleanUserFromSocketPool(query QueryDSL) {
	workspaceId := query.WorkspaceId
	userId := query.UserId

	// Automatically happens, perhaps should not be called anymore.
	query.RawSocketConnection.Close()

	socketMutex.Lock()
	conns := SocketSessionPool[workspaceId][userId]
	for i, conn := range conns {
		if conn.Connection == query.RawSocketConnection {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(conns) == 0 {
		delete(SocketSessionPool[workspaceId], userId)
	} else {
		SocketSessionPool[workspaceId][userId] = conns
	}
	if len(SocketSessionPool[workspaceId]) == 0 {
		delete(SocketSessionPool, workspaceId)
	}
	socketMutex.Unlock()

	GetEventBusInstance().RemoveUser(SERVER_INSTANCE, userId)
}

func addUserToEventBus(query QueryDSL) {
	workspaceId := query.WorkspaceId
	userId := query.UserId

	GetEventBusInstance().AddUser(SERVER_INSTANCE, userId)

	socket := &SocketConnection{
		UserId:     userId,
		Connection: query.RawSocketConnection,
		URW:        *query.UserAccessPerWorkspace,
	}

	socketMutex.Lock()
	if SocketSessionPool[workspaceId] == nil {
		SocketSessionPool[workspaceId] = make(map[string][]*SocketConnection)
	}
	SocketSessionPool[workspaceId][userId] = append(SocketSessionPool[workspaceId][userId], socket)
	socketMutex.Unlock()

}

func EventBusSubscriptionAction(query QueryDSL, done chan bool, read chan SocketReadChan) (chan []byte, error) {

	addUserToEventBus(query)

	out := make(chan []byte)

	go func() {
		defer close(out)
		defer cleanUserFromSocketPool(query)

		for {
			select {
			case msg, ok := <-read:
				if !ok {
					return
				}

				if websocket.IsCloseError(msg.Error, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					LOG.Debug("WebSocket closed normally", zap.String("user", query.UserId))

					// Only break in this case
					break
				} else {
					// We need to handle this kinda.
					LOG.Debug("WebSocket read error", zap.Error(msg.Error))
					out <- []byte("Socket interaction with webserver only supports json at this moment.")
				}

			case <-done:
				return
			}
		}
	}()

	return out, nil
}
