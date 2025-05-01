package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ReactiveSearchActionImp = func(query fireback.QueryDSL, done chan bool, read chan fireback.SocketReadChan) (chan []byte, error) {
		return fireback.DefaultEmptyReactiveAction(query, done, read)
	}
}
