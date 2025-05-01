package fireback

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

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

// General purpose reader from socket and writer into a io.Writer.
// could be used for sending audio, video, or actually any other file over web socket.
// output channel is not used, to make it similar to other reactive functionality we
// need to return that empty output channel and make sure closing it, maybe later we need it.
func StreamToWriter(done chan bool, read chan SocketReadChan, writer io.Writer) (chan []byte, error) {
	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-read:
				if !ok {
					return
				}
				writer.Write(msg.Data) // you could add error check if needed
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
func SocketStreamToFile(done chan bool, read chan SocketReadChan, fileAddressOnDisk string) (chan []byte, error) {
	file, err := os.Create(fileAddressOnDisk)
	if err != nil {
		return nil, err
	}

	return StreamToWriter(done, read, file)
}

type SocketReadChan struct {
	Data  []byte
	Error error
}

func StreamSlowMp3OverSocket(
	filePath string,

	// Use 2048 for a simulation
	chunkSize int,
	query QueryDSL,
	done chan bool,
	read chan SocketReadChan,
) (chan []byte, error) {
	stream := make(chan []byte)
	go func() {
		defer close(stream)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		buf := make([]byte, 2048) // ~2KB chunks
		for {
			n, err := file.Read(buf)
			if n > 0 {
				chunk := make([]byte, n)
				copy(chunk, buf[:n])
				stream <- chunk
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			time.Sleep(40 * time.Millisecond) // mimic ~25 FPS (can be adjusted)
		}
	}()

	return stream, nil
}
