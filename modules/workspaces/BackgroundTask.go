package workspaces

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type BackgroundReactiveProcess struct {
	Done      chan bool
	Read      chan string
	Listeners []func(*string)
	Group     string
}

var ProcessPool map[string]*BackgroundReactiveProcess = map[string]*BackgroundReactiveProcess{}

func (x *BackgroundReactiveProcess) Terminate() {
	close(x.Done)
}

func (x *BackgroundReactiveProcess) AttachListener(listener func(*string)) {
	x.Listeners = append(x.Listeners, listener)
}

func (x *BackgroundReactiveProcess) Send(v string) {
	x.Read <- v
}

// If the operation exists in the pool, it will return that instead of creating new one
func BeginOrAttachOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	if ProcessPool[query.UniqueId] != nil {
		return ProcessPool[query.UniqueId], nil
	}

	return BeginOperation(query, fn)
}

type BackgroundOptFn func(query QueryDSL, done chan bool, read chan string) (chan *string, error)

func BeginOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	done := make(chan bool)
	read := make(chan string)
	ref := query.UniqueId

	act, err := fn(query, done, read)

	if err != nil {
		return nil, err
	}
	ProcessPool[ref] = &BackgroundReactiveProcess{
		Done:      done,
		Read:      read,
		Group:     "ControlSheet",
		Listeners: []func(*string){},
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

func ReactiveSocketHandler(factory func(
	query QueryDSL, done chan bool,
	read chan string,
) (chan *string, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		HttpSocketRequest(ctx, func(query QueryDSL, write func(string)) {
			opt, err := BeginOrAttachOperation(query, factory)
			fmt.Println("Err:", err)
			opt.AttachListener(func(s *string) {
				write(*s)
			})

		}, func(query QueryDSL, i interface{}) {
			opt, err := BeginOrAttachOperation(query, factory)
			fmt.Println("Err:", err)
			var kv string = i.(string)
			opt.Send(kv)
		})

	}
}

func sendStringWithInterval(ctx context.Context, interval time.Duration, out chan *string) {

	for {
		select {
		case <-ctx.Done():
			return
		default:

			js := "Hello :)"
			out <- &js
			time.Sleep(time.Millisecond * 1000)
		}
	}

}

// This is a sample which will be used as default action for reactive items
func DefaultEmptyReactiveAction(
	query QueryDSL,
	done chan bool,
	read chan string,
) (chan *string, error) {

	stream := make(chan *string)

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
	read := make(chan string)

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
		text := scanner.Text()
		read <- text
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}
