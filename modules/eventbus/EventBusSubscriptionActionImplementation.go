package eventbus

import (
	"github.com/gorilla/websocket"
	eventbusdefs "github.com/torabian/fireback/modules/eventbus/defs"
	"github.com/torabian/fireback/modules/fireback"
	"go.uber.org/zap"
)

func cleanUserFromSocketPool(query fireback.QueryDSL) {
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

	GetEventBusInstance().RemoveUser(fireback.SERVER_INSTANCE, userId)
}

func addUserToEventBus(query fireback.QueryDSL) {

	workspaceId := query.WorkspaceId
	userId := query.UserId

	GetEventBusInstance().AddUser(fireback.SERVER_INSTANCE, userId)

	socket := &SocketConnection{
		UserId:     userId,
		Connection: query.RawSocketConnection,
		UniqueId:   fireback.UUID_SHORT(),
	}

	if query.UserAccessPerWorkspace != nil {
		socket.URW = *query.UserAccessPerWorkspace
	}

	socketMutex.Lock()
	if SocketSessionPool[workspaceId] == nil {
		SocketSessionPool[workspaceId] = make(map[string][]*SocketConnection)
	}
	SocketSessionPool[workspaceId][userId] = append(SocketSessionPool[workspaceId][userId], socket)
	socketMutex.Unlock()

}

func EventBusSubscriptionActionSig(session eventbusdefs.EventBusSubscriptionActionSession) (chan []byte, error) {
	query, err := fireback.ResolveActionContextFromGinContext(session.GinCtx(), &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyWorkspace,
	})
	if err != nil {
		return nil, err
	}

	query.RawSocketConnection = session.GetSocket()

	fireback.LOG.Debug(
		"Event bus subscription has been started",
		zap.String("workspace-id", query.WorkspaceId),
		zap.String("user", query.UserId),
	)

	addUserToEventBus(*query)

	out := make(chan []byte)

	go func() {
		defer close(out)
		defer cleanUserFromSocketPool(*query)

		for {
			select {
			case msg, ok := <-session.Read:
				if !ok {
					return
				}

				if websocket.IsCloseError(
					msg.Error,
					websocket.CloseNormalClosure,
					websocket.CloseAbnormalClosure,
					websocket.CloseGoingAway,
					websocket.CloseInternalServerErr,
					websocket.CloseInvalidFramePayloadData,
					websocket.CloseTLSHandshake,
				) {
					fireback.LOG.Debug("WebSocket closed", zap.String("user", query.UserId), zap.String("ws", query.WorkspaceId))
					cleanUserFromSocketPool(*query)
					break
				} else {
					// We need to handle this kinda.
					fireback.LOG.Debug("WebSocket read error", zap.Error(msg.Error))
					out <- []byte("Socket interaction with webserver only supports json at this moment.")
				}

			case <-session.Done:
				return
			}
		}
	}()

	return out, nil
}
