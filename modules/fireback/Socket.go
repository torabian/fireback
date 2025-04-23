package fireback

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

// General purpose reader from socket and writer into a io.Writer.
// could be used for sending audio, video, or actually any other file over web socket.
// output channel is not used, to make it similar to other reactive functionality we
// need to return that empty output channel and make sure closing it, maybe later we need it.
func StreamToWriter(done chan bool, read chan []byte, writer io.Writer) (chan []byte, error) {
	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-read:
				if !ok {
					return
				}
				writer.Write(msg) // you could add error check if needed
			case <-done:
				return
			}
		}
	}()

	return out, nil
}

// Stream incoming socket into a file. Here is a front-end sample:
//
//	async function startStreaming() {
//		const ws = new WebSocket("ws://localhost:4500/audiostream");
//		const stream = await navigator.mediaDevices.getUserMedia({ audio: true, video: true });
//		const recorder = new MediaRecorder(stream, { mimeType: 'video/webm; codecs=vp8,opus' });
//			recorder.ondataavailable = (e) => {
//				if (e.data.size > 0) {
//					e.data.arrayBuffer().then(buf => ws.send(buf)); // no prefix needed
//				}
//			};
//			recorder.start(250);
//		   document.getElementById("preview").srcObject = stream;
//		}
func SocketStreamToFile(done chan bool, read chan []byte, fileAddressOnDisk string) (chan []byte, error) {
	file, err := os.Create(fileAddressOnDisk)
	if err != nil {
		return nil, err
	}

	return StreamToWriter(done, read, file)
}
