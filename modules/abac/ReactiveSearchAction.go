package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ReactiveSearchActionImp = ReactiveSearchAction
}

func ReactiveSearchAction(query fireback.QueryDSL, done chan bool, read chan string) (chan *string, error) {
	// Implement the logic here.

	m := make(chan *string)
	return m, nil
}
