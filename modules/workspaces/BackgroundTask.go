package workspaces

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BackgroundReactiveProcess struct {
	Done      chan bool
	Read      chan map[string]interface{}
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

func (x *BackgroundReactiveProcess) Send(v map[string]interface{}) {
	x.Read <- v
}

// If the operation exists in the pool, it will return that instead of creating new one
func BeginOrAttachOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	if ProcessPool[query.UniqueId] != nil {
		return ProcessPool[query.UniqueId], nil
	}

	return BeginOperation(query, fn)
}

type BackgroundOptFn func(query QueryDSL, done chan bool, read chan map[string]interface{}) (chan *string, error)

func BeginOperation(query QueryDSL, fn BackgroundOptFn) (*BackgroundReactiveProcess, error) {
	done := make(chan bool)
	read := make(chan map[string]interface{})
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
	read chan map[string]interface{},
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
			var kv map[string]interface{} = i.(map[string]interface{})
			opt.Send(kv)
		})

	}
}
