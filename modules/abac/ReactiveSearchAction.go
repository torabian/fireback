package workspaces

func init() {
	// Override the implementation with our actual code.
	ReactiveSearchActionImp = ReactiveSearchAction
}

func ReactiveSearchAction(query QueryDSL, done chan bool, read chan string) (chan *string, error) {
	// Implement the logic here.

	m := make(chan *string)
	return m, nil
}
