package fireback

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type BackgroundReactiveProcess struct {
	Done      chan bool
	Read      chan SocketReadChan
	Listeners []func([]byte)
	Group     string
}

var ProcessPool map[string]*BackgroundReactiveProcess = map[string]*BackgroundReactiveProcess{}

func (x *BackgroundReactiveProcess) Terminate() {
	close(x.Done)
}

func (x *BackgroundReactiveProcess) AttachListener(listener func([]byte)) {
	x.Listeners = append(x.Listeners, listener)
}

func (x *BackgroundReactiveProcess) Send(v SocketReadChan) {
	x.Read <- v
}

// If the operation exists in the pool, it will return that instead of creating new one
func BeginOrAttachOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	if ProcessPool[query.UniqueId] != nil {
		return ProcessPool[query.UniqueId], nil
	}

	return BeginOperation(query, fn)
}

type BackgroundOptFn func(query QueryDSL, done chan bool, read chan SocketReadChan) (chan []byte, error)

func BeginOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	done := make(chan bool)
	read := make(chan SocketReadChan)
	ref := query.UniqueId

	act, err := fn(query, done, read)

	if err != nil {
		return nil, err
	}
	ProcessPool[ref] = &BackgroundReactiveProcess{
		Done:      done,
		Read:      read,
		Group:     "ControlSheet",
		Listeners: []func([]byte){},
	}

	go func() {

		for {
			select {
			case <-done:
				return
			case row, more := <-act:
				if ProcessPool[ref] != nil && len(ProcessPool[ref].Listeners) > 0 {

					for _, fnx := range ProcessPool[ref].Listeners {
						fnx(row)

					}
				}

				if !more {
					return
				}

			}
		}
	}()

	return ProcessPool[ref], nil

}

type ReactiveFactory = func(
	query QueryDSL, done chan bool,
	read chan SocketReadChan,
) (chan []byte, error)

func ReactiveSocketHandler(factory ReactiveFactory) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		read := make(chan SocketReadChan)
		done := make(chan bool)
		f := ExtractQueryDslFromGinContext(ctx)

		c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteJSON(GormErrorToIError(err))

			c.Close()
			return
		}

		f.RawSocketConnection = c

		write, err := factory(f, done, read)
		if err != nil {
			c.WriteJSON(GormErrorToIError(err))
		}

		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- SocketReadChan{
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

func sendStringWithInterval(ctx context.Context, interval time.Duration, out chan []byte) {

	for {
		select {
		case <-ctx.Done():
			return
		default:

			js := "Hello :)"
			out <- []byte(js)
			time.Sleep(time.Millisecond * 1000)
		}
	}

}

type WebRtcSignal struct{}

func WebRtcAction(
	query QueryDSL,
) (*WebRtcSignal, error) {
	return nil, nil
}

// This is a sample which will be used as default action for reactive items
func DefaultEmptyReactiveAction(
	query QueryDSL,
	done chan bool,
	read chan SocketReadChan,
) (chan []byte, error) {

	stream := make(chan []byte)

	go func() {
		var ctx context.Context = nil
		var cancel context.CancelFunc = nil

		for {
			select {
			case <-done:
				fmt.Println("Completed actually")
				return

			case row, more := <-read:

				// Do somehting with the row actuall
				fmt.Println("Row:", row)
				if cancel != nil {
					cancel()
				}

				ctx, cancel = context.WithCancel(context.Background())
				defer cancel()

				go sendStringWithInterval(ctx, 1000, stream)

				if !more {
					return
				}
			}
		}
	}()

	return stream, nil
}

func CliReactivePipeHandler(query QueryDSL, fn BackgroundOptFn) {
	done := make(chan bool)
	read := make(chan SocketReadChan)

	out, err := fn(query, done, read)

	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for {
			select {
			case <-done:
				return

			case data, more := <-out:
				fmt.Print(data)
				if !more {
					return
				}
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Bytes()
		read <- SocketReadChan{
			Data: text,
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}
