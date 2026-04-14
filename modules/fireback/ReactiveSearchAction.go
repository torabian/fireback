package fireback

import (
	"encoding/json"
	"sync"
)

func init() {
	// Override the implementation with our actual code.
	// ReactiveSearchImpl = // Trigger intelisense, it would auto complete.
}

func CreateReactiveSearchHanlder(app *FirebackApp) func(
	session ReactiveSearchActionSession,
) (chan []byte, error) {

	return func(
		session ReactiveSearchActionSession,
	) (chan []byte, error) {
		query := ExtractQueryDslFromGinContext(session.Ctx)
		query.RawSocketConnection = session.Socket
		resultChan := make(chan *ReactiveSearchResultDto)

		go func() {
			var wg sync.WaitGroup

			for _, handler := range app.SearchProviders {
				wg.Add(1)

				go func(h SearchProviderFn) {
					defer wg.Done()
					h(query, resultChan)
				}(handler)
			}

			wg.Wait()

			close(resultChan)
		}()

		return AdaptResultsToBytes(resultChan), nil
	}

}

func AdaptResultsToBytes(input chan *ReactiveSearchResultDto) chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)

		for res := range input {
			b, err := json.Marshal(res)
			if err != nil {
				continue // or log error
			}
			out <- b
		}
	}()

	return out
}
