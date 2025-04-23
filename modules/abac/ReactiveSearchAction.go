package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ReactiveSearchActionImp = func(query fireback.QueryDSL, done chan bool, read chan []byte) (chan []byte, error) {
		return fireback.DefaultEmptyReactiveAction(query, done, read)
	}
}
