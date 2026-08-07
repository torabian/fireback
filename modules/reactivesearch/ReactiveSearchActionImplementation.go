package reactivesearch

import (
	"encoding/json"
	"sync"

	"github.com/torabian/fireback/modules/fireback"
)

// SearchProviderFn is one search source ReactiveSearchModuleConfig.SearchProviders can
// register - it streams its own results (if any) for query.SearchPhrase into
// chanStream, and returns once it's done searching.
type SearchProviderFn = func(query fireback.QueryDSL, chanStream chan *ReactiveSearchResultDto)

// createReactiveSearchHandler builds the reactive handler ModuleSetup wires to the
// ReactiveSearch action - providers is exactly (and only) what the project passed to
// ModuleSetup's config, instead of being read off a FirebackApp field every project
// carried whether or not it used reactive search. Auth resolution (needs a live gin
// context/token) is kept separate from the actual fan-out (runSearchProviders) so the
// fan-out - the part with real logic worth testing - can be tested without a live
// server; see ReactiveSearch_test.go.
func createReactiveSearchHandler(providers []SearchProviderFn) func(
	session ReactiveSearchActionSession,
) (chan []byte, error) {

	return func(
		session ReactiveSearchActionSession,
	) (chan []byte, error) {
		query, err := fireback.ResolveActionContextFromGinContext(session.GinCtx(), &fireback.SecurityModel{
			ResolveStrategy: fireback.ResolveStrategyWorkspace,
		})
		if err != nil {
			return nil, err
		}

		query.RawSocketConnection = session.GetSocket()

		return AdaptResultsToBytes(runSearchProviders(*query, providers)), nil
	}

}

// runSearchProviders fans query out to every provider concurrently, merging whatever
// each one streams back onto a single channel that's closed once they've all
// finished. A provider that finds nothing (or panics/blocks forever) never stops the
// others from being asked or from reporting what they found.
func runSearchProviders(query fireback.QueryDSL, providers []SearchProviderFn) chan *ReactiveSearchResultDto {
	resultChan := make(chan *ReactiveSearchResultDto)

	go func() {
		var wg sync.WaitGroup

		for _, handler := range providers {
			wg.Add(1)

			go func(h SearchProviderFn) {
				defer wg.Done()
				h(query, resultChan)
			}(handler)
		}

		wg.Wait()

		close(resultChan)
	}()

	return resultChan
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
