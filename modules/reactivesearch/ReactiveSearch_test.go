package reactivesearch

import (
	"encoding/json"
	"sort"
	"testing"
	"time"

	"github.com/torabian/fireback/modules/fireback"
)

// collect drains a chan *ReactiveSearchResultDto until it's closed or the deadline
// passes, returning whatever arrived.
func collect(t *testing.T, ch chan *ReactiveSearchResultDto) []*ReactiveSearchResultDto {
	t.Helper()

	var got []*ReactiveSearchResultDto
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
	menus := func(query fireback.QueryDSL, out chan *ReactiveSearchResultDto) {
		out <- &ReactiveSearchResultDto{UniqueId: "menu-1", Group: "menus"}
		out <- &ReactiveSearchResultDto{UniqueId: "menu-2", Group: "menus"}
	}
	roles := func(query fireback.QueryDSL, out chan *ReactiveSearchResultDto) {
		time.Sleep(10 * time.Millisecond) // slower provider must not block/drop the others
		out <- &ReactiveSearchResultDto{UniqueId: "role-1", Group: "roles"}
	}
	empty := func(query fireback.QueryDSL, out chan *ReactiveSearchResultDto) {
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
	provider := func(query fireback.QueryDSL, out chan *ReactiveSearchResultDto) {
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
	in := make(chan *ReactiveSearchResultDto, 2)
	in <- &ReactiveSearchResultDto{UniqueId: "1", Phrase: "first"}
	in <- &ReactiveSearchResultDto{UniqueId: "2", Phrase: "second"}
	close(in)

	out := AdaptResultsToBytes(in)

	var got []ReactiveSearchResultDto
	deadline := time.After(2 * time.Second)
	for i := 0; i < 2; i++ {
		select {
		case b := <-out:
			var r ReactiveSearchResultDto
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
