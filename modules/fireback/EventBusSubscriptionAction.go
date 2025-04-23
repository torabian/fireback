package fireback

import "fmt"

func init() {
	// Override the implementation with our actual code.
	EventBusSubscriptionActionImp = EventBusSubscriptionAction
}

func EventBusSubscriptionAction(query QueryDSL, done chan bool, read chan []byte) (chan []byte, error) {

	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-read:
				if !ok {
					return
				}

				fmt.Println("Message came:", msg, query.UserId, query.WorkspaceId)
			case <-done:
				return
			}
		}
	}()

	return out, nil
}
