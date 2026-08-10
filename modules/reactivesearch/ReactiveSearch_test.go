package reactivesearch

import (
	"encoding/json"
	"errors"
	"sort"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	reactivesearchdefs "github.com/torabian/fireback/modules/reactivesearch/defs"
)

// collect drains a chan *reactivesearchdefs.ReactiveSearchResultDto until it's closed or the deadline
// passes, returning whatever arrived.
func collect(t *testing.T, ch chan *reactivesearchdefs.ReactiveSearchResultDto) []*reactivesearchdefs.ReactiveSearchResultDto {
	t.Helper()

	var got []*reactivesearchdefs.ReactiveSearchResultDto
	deadline := time.After(2 * time.Second)
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, r)
		case <-deadline:
			t.Fatal("timed out waiting for the result channel to close")
			return got
		}
	}
}

// TestRunSearchProviders_MergesResultsFromAllProviders verifies every registered
// provider gets called and everything each one streams back ends up on the merged
// channel, regardless of how many results each one sends or how long it takes.
func TestRunSearchProviders_MergesResultsFromAllProviders(t *testing.T) {
	menus := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		out <- &reactivesearchdefs.ReactiveSearchResultDto{UniqueId: "menu-1", Group: "menus"}
		out <- &reactivesearchdefs.ReactiveSearchResultDto{UniqueId: "menu-2", Group: "menus"}
	}
	roles := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		time.Sleep(10 * time.Millisecond) // slower provider must not block/drop the others
		out <- &reactivesearchdefs.ReactiveSearchResultDto{UniqueId: "role-1", Group: "roles"}
	}
	empty := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		// Finds nothing - must not stop the channel from eventually closing.
	}

	got := collect(t, runSearchProviders(fireback.QueryDSL{}, []SearchProviderFn{menus, roles, empty}))

	ids := make([]string, len(got))
	for i, r := range got {
		ids[i] = r.UniqueId
	}
	sort.Strings(ids)

	want := []string{"menu-1", "menu-2", "role-1"}
	if len(ids) != len(want) {
		t.Fatalf("expected %v, got %v", want, ids)
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, ids)
		}
	}
}

// TestRunSearchProviders_ClosesWithNoProviders verifies the channel closes right away
// (rather than hanging) when there's nothing registered to search.
func TestRunSearchProviders_ClosesWithNoProviders(t *testing.T) {
	got := collect(t, runSearchProviders(fireback.QueryDSL{}, nil))
	if len(got) != 0 {
		t.Fatalf("expected no results, got %v", got)
	}
}

// TestRunSearchProviders_PassesQueryThrough verifies each provider receives exactly
// the query runSearchProviders was given (e.g. SearchPhrase), not a zero value.
func TestRunSearchProviders_PassesQueryThrough(t *testing.T) {
	seen := make(chan string, 1)
	provider := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		seen <- query.SearchPhrase
	}

	collect(t, runSearchProviders(fireback.QueryDSL{SearchPhrase: "menu"}, []SearchProviderFn{provider}))

	select {
	case phrase := <-seen:
		if phrase != "menu" {
			t.Fatalf("expected provider to see SearchPhrase %q, got %q", "menu", phrase)
		}
	default:
		t.Fatal("provider was never called")
	}
}

// TestAdaptResultsToBytes_EncodesEachResult verifies each result gets JSON-encoded
// individually onto the byte channel, preserving a single provider's own order.
func TestAdaptResultsToBytes_EncodesEachResult(t *testing.T) {
	in := make(chan *reactivesearchdefs.ReactiveSearchResultDto, 2)
	in <- &reactivesearchdefs.ReactiveSearchResultDto{UniqueId: "1", Phrase: "first"}
	in <- &reactivesearchdefs.ReactiveSearchResultDto{UniqueId: "2", Phrase: "second"}
	close(in)

	out := AdaptResultsToBytes(in)

	var got []reactivesearchdefs.ReactiveSearchResultDto
	deadline := time.After(2 * time.Second)
	for i := 0; i < 2; i++ {
		select {
		case b := <-out:
			var r reactivesearchdefs.ReactiveSearchResultDto
			if err := json.Unmarshal(b, &r); err != nil {
				t.Fatalf("decoding %q: %v", b, err)
			}
			got = append(got, r)
		case <-deadline:
			t.Fatal("timed out waiting for encoded results")
		}
	}

	if got[0].UniqueId != "1" || got[1].UniqueId != "2" {
		t.Fatalf("expected results in order [1 2], got %+v", got)
	}

	if _, ok := <-out; ok {
		t.Fatal("expected the output channel to be closed once the input was drained")
	}
}

// TestCreateReactiveSearchHandler_UsesProvidedAuthorize verifies a custom Authorize
// callback is what actually resolves the query providers get called with - proving
// createReactiveSearchHandler doesn't need fireback.ResolveActionContext (and
// therefore abac) at all once Authorize is supplied.
func TestCreateReactiveSearchHandler_UsesProvidedAuthorize(t *testing.T) {
	var gotReq emigo.EmiRequestContexts
	authorize := func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		gotReq = req
		return fireback.QueryDSL{SearchPhrase: "hello"}, nil
	}

	seenPhrase := make(chan string, 1)
	provider := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		seenPhrase <- query.SearchPhrase
	}

	handler := createReactiveSearchHandler([]SearchProviderFn{provider}, authorize)

	// A pure-function Authorize like this one never needs to touch the socket - only
	// the underlying Socket type has to satisfy GetSocket's assertion for the
	// success path to reach runSearchProviders.
	out, err := handler(reactivesearchdefs.ReactiveSearchActionSession{Socket: (*websocket.Conn)(nil)})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for range out {
		// Drain so runSearchProviders' goroutine finishes and the deferred close
		// on the byte channel runs.
	}

	select {
	case phrase := <-seenPhrase:
		if phrase != "hello" {
			t.Fatalf("expected the provider to see the query Authorize returned, got %q", phrase)
		}
	default:
		t.Fatal("provider was never called")
	}

	if gotReq == nil {
		t.Fatal("expected Authorize to be called with a request context")
	}
}

// TestCreateReactiveSearchHandler_PropagatesAuthorizeError verifies a rejected/failed
// Authorize call short-circuits before ever touching the socket or providers.
func TestCreateReactiveSearchHandler_PropagatesAuthorizeError(t *testing.T) {
	wantErr := errors.New("nope")
	called := false
	provider := func(query fireback.QueryDSL, out chan *reactivesearchdefs.ReactiveSearchResultDto) {
		called = true
	}

	handler := createReactiveSearchHandler([]SearchProviderFn{provider}, func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		return fireback.QueryDSL{}, wantErr
	})

	// Socket/Ctx are deliberately left unset (nil interfaces) - if the handler
	// tried to use them before returning the error, this would panic.
	_, err := handler(reactivesearchdefs.ReactiveSearchActionSession{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected Authorize's error to propagate, got %v", err)
	}
	if called {
		t.Fatal("expected providers to never run once Authorize failed")
	}
}

// TestReactiveSearchRequestContext_AdaptsSession verifies the emigo.EmiRequestContexts
// adapter exposes the session's gin context and always reports no CLI context (reactive
// search never runs over CLI).
func TestReactiveSearchRequestContext_AdaptsSession(t *testing.T) {
	fakeCtx := "stand-in for a *gin.Context"
	session := &reactivesearchdefs.ReactiveSearchActionSession{Ctx: fakeCtx}
	adapter := reactiveSearchRequestContext{session: session}

	if adapter.GetGinCtx() != interface{}(fakeCtx) {
		t.Fatalf("expected GetGinCtx to return the session's Ctx, got %v", adapter.GetGinCtx())
	}
	if adapter.GetCliCtx() != nil {
		t.Fatalf("expected GetCliCtx to always be nil, got %v", adapter.GetCliCtx())
	}
}

// TestDefaultAuthorize_NoGinContext verifies the fireback.ResolveActionContext-based
// fallback used when ReactiveSearchModuleConfig.Authorize is left nil degrades to a
// zero QueryDSL (no panic, no error) when there's no real gin context to resolve -
// matching fireback.ResolveActionContext's own documented behavior for a nilish
// context.
func TestDefaultAuthorize_NoGinContext(t *testing.T) {
	query, err := defaultAuthorize(reactiveSearchRequestContext{session: &reactivesearchdefs.ReactiveSearchActionSession{}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if query.WorkspaceId != "" || query.UserId != "" {
		t.Fatalf("expected a zero-value QueryDSL with no gin context, got %+v", query)
	}
}
