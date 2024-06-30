package workspaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	"github.com/gorilla/websocket"
)

type SocketConnection struct {
	UserId     string
	Connection *websocket.Conn
}

var SocketSessionPool = make(map[string]map[string]*SocketConnection)

// var sessionPool []*SocketConnection = []*SocketConnection{}

func SocketConnectEndpoint(ginContext *gin.Context) { //Usually use c *gin.Context

	wsURL := ginContext.Request.URL
	wsURLParam, err3 := url.ParseQuery(wsURL.RawQuery)

	workspaceId := ginContext.Request.Header.Get("Workspace-id")
	token := ginContext.Request.Header.Get("authorization")

	if err3 == nil && wsURLParam["token"] != nil && len(wsURLParam["token"]) == 1 {

		token = wsURLParam["token"][0]
	}

	if err3 == nil && wsURLParam["workspaceId"] != nil && len(wsURLParam["workspaceId"]) == 1 {

		workspaceId = wsURLParam["workspaceId"][0]
	}

	context := &AuthContextDto{
		WorkspaceId:  &workspaceId,
		Token:        &token,
		Capabilities: []PermissionInfo{},
	}

	res, err2 := WithAuthorizationPure(context)
	if err2 != nil {
		f, _ := json.MarshalIndent(gin.H{"error": err2}, "", "  ")
		http.Error(ginContext.Writer, string(f), int(err2.HttpCode))
		return
	}

	wsSession, err := upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	if SocketSessionPool[*res.WorkspaceId] == nil {
		SocketSessionPool[*res.WorkspaceId] = map[string]*SocketConnection{}
	}

	SocketSessionPool[*res.WorkspaceId][*res.UserId] = &SocketConnection{
		Connection: wsSession,
		UserId:     *res.UserId,
	}

	for {
		var msg interface{}
		err := wsSession.ReadJSON(&msg)
		if err != nil {
			log.Println("read:", err)
			// optional: log the error
			break
		}

	}

	delete(SocketSessionPool[*res.WorkspaceId], *res.UserId)
	if len(SocketSessionPool[*res.WorkspaceId]) == 0 {
		delete(SocketSessionPool, *res.WorkspaceId)

	}

	wsSession.Close()
}

/**
*	Routes all internal events of the application over socket.
* 	This is an exteremly dangerous function, has access to every small details of the application
**/
func UNSAFE_AppEventsOverSocketRouter() {
	dbListener1 := event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		if len(SocketSessionPool) > 0 {

			for _, v := range SocketSessionPool {
				for _, v2 := range v {
					v2.Connection.WriteJSON(gin.H{"data": e.Data()})
				}
			}
		}

		return nil
	})

	event.On("*", dbListener1, event.Normal)

}

func HandleSocket(e *gin.Engine) {
	e.Any("ws", SocketConnectEndpoint)
	UNSAFE_AppEventsOverSocketRouter()
}
