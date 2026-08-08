package reactivesearch

import (
	"encoding/json"
	"sync"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// SearchProviderFn is one search source ReactiveSearchModuleConfig.SearchProviders can
// register - it streams its own results (if any) for query.SearchPhrase into
// chanStream, and returns once it's done searching.
type SearchProviderFn = func(query fireback.QueryDSL, chanStream chan *ReactiveSearchResultDto)

// AuthorizeReactiveSearchFn resolves/authorizes an incoming reactive search request,
// returning the QueryDSL SearchProviderFn handlers get called with. It's handed
// anything satisfying emigo.EmiRequestContexts (GetGinCtx/GetCliCtx) - the exact
// contract every other emi-generated request type already implements - so it plugs
// straight into fireback.ResolveActionContext the same way a normal action would, or
// into a project's own auth logic instead: reactivesearch itself never has to know
// which. See ReactiveSearchModuleConfig.Authorize and defaultAuthorize (the fallback
// used when it's left nil).
type AuthorizeReactiveSearchFn func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error)

// reactiveSearchRequestContext adapts a ReactiveSearchActionSession into
// emigo.EmiRequestContexts, so AuthorizeReactiveSearchFn implementations can reuse
// fireback.ResolveActionContext (or anything else already written against that
// interface) without reactivesearch needing its own gin-specific plumbing. Reactive
// search only ever runs over a live gin websocket connection, never CLI, so
// GetCliCtx is always nil.
type reactiveSearchRequestContext struct {
	session *ReactiveSearchActionSession
}

func (r reactiveSearchRequestContext) GetGinCtx() interface{} { return r.session.Ctx }
func (r reactiveSearchRequestContext) GetCliCtx() interface{} { return nil }

// defaultAuthorize is used when ReactiveSearchModuleConfig.Authorize is left nil -
// matches reactivesearch's original, pre-callback behavior: resolve via
// fireback.ResolveActionContext with SecurityModel{ResolveStrategy: ResolveStrategyWorkspace},
// which requires a valid token and depends on abac (or whatever else wires up
// fireback.AuthorizeRequest/fireback.MeetsAccessLevel) having already run its own
// module setup. Pass your own AuthorizeReactiveSearchFn instead if you want
// reactivesearch to not depend on any of that.
func defaultAuthorize(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
	query, err := fireback.ResolveActionContext(req, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyWorkspace,
	})
	if err != nil {
		return fireback.QueryDSL{}, err
	}
	return *query, nil
}

// createReactiveSearchHandler builds the reactive handler ModuleSetup wires to the
// ReactiveSearch action. Authorization (needs a live gin context/token, or whatever
// authorize itself needs) is kept separate from the actual fan-out
// (runSearchProviders) so the fan-out - the part with real logic worth testing - can
// be tested without a live server; see ReactiveSearch_test.go.
func createReactiveSearchHandler(providers []SearchProviderFn, authorize AuthorizeReactiveSearchFn) func(
	session ReactiveSearchActionSession,
) (chan []byte, error) {
	if authorize == nil {
		authorize = defaultAuthorize
	}

	return func(
		session ReactiveSearchActionSession,
	) (chan []byte, error) {
		query, err := authorize(reactiveSearchRequestContext{session: &session})
		if err != nil {
			return nil, err
		}

		query.RawSocketConnection = session.GetSocket()

		return AdaptResultsToBytes(runSearchProviders(query, providers)), nil
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
