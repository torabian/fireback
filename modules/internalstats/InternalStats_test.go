package internalstats

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

// TestCollectSnapshot_ReturnsPopulatedStats verifies CollectSnapshot (the one function
// every transport - HTTP snapshot, reactive stream, `internalstats watch` - ultimately
// calls) actually returns a real, non-trivial set of measurements with every field
// filled in, on whatever machine the test runs on.
func TestCollectSnapshot_ReturnsPopulatedStats(t *testing.T) {
	snap := CollectSnapshot()

	if snap.Hostname == "" {
		t.Fatal("expected a non-empty hostname")
	}
	if _, err := time.Parse(time.RFC3339, snap.GeneratedAt); err != nil {
		t.Fatalf("expected GeneratedAt to be RFC3339, got %q: %v", snap.GeneratedAt, err)
	}

	items := snap.Items.Items
	// "around 30-40" per InternalStats.emi.yml's description - loosely bounded so a
	// platform-specific field going missing/n/a doesn't make this flaky.
	if len(items) < 25 {
		t.Fatalf("expected at least 25 stats, got %d", len(items))
	}

	seenCategories := map[string]bool{}
	for _, item := range items {
		if item.Key == "" || item.Label == "" || item.Category == "" || item.Value == "" {
			t.Fatalf("expected every field filled in, got %+v", item)
		}
		switch item.Severity {
		case SeverityOk, SeverityWarn, SeverityCritical, SeverityInfo:
		default:
			t.Fatalf("unexpected severity %q on %+v", item.Severity, item)
		}
		seenCategories[item.Category] = true
	}

	for _, want := range []string{categoryHost, categoryCPU, categoryMemory, categoryDisk, categoryNetwork, categoryRuntime} {
		if !seenCategories[want] {
			t.Fatalf("expected a %q stat, got categories %v", want, seenCategories)
		}
	}
}

// TestPercentSeverity_Thresholds locks in the ok/warn/critical cutoffs so a future
// tweak to warnPercent/criticalPercent is a deliberate change, not an accident.
func TestPercentSeverity_Thresholds(t *testing.T) {
	cases := []struct {
		pct  float64
		want string
	}{
		{0, SeverityOk},
		{74.9, SeverityOk},
		{75, SeverityWarn},
		{89.9, SeverityWarn},
		{90, SeverityCritical},
		{100, SeverityCritical},
	}
	for _, c := range cases {
		if got := percentSeverity(c.pct); got != c.want {
			t.Errorf("percentSeverity(%v) = %q, want %q", c.pct, got, c.want)
		}
	}
}

func TestFormatBytes(t *testing.T) {
	cases := []struct {
		in   uint64
		want string
	}{
		{0, "0 B"},
		{1023, "1023 B"},
		{1024, "1.0 KiB"},
		{1536, "1.5 KiB"},
		{1024 * 1024 * 1024, "1.0 GiB"},
	}
	for _, c := range cases {
		if got := formatBytes(c.in); got != c.want {
			t.Errorf("formatBytes(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	cases := []struct {
		in   time.Duration
		want string
	}{
		{0, "0s"},
		{45 * time.Second, "45s"},
		{90 * time.Second, "1m 30s"},
		{2*time.Hour + 5*time.Minute, "2h 5m"},
		{25*time.Hour + 3*time.Minute, "1d 1h 3m"},
	}
	for _, c := range cases {
		if got := formatDuration(c.in); got != c.want {
			t.Errorf("formatDuration(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestInternalStatsSnapshot_UsesProvidedAuthorize verifies a custom Authorize callback
// is what actually gates InternalStatsSnapshotAction - proving createSnapshotHandler
// doesn't need fireback.ResolveActionContext (and therefore abac) at all once
// Authorize is supplied, the same guarantee reactivesearch's equivalent test makes.
func TestInternalStatsSnapshot_UsesProvidedAuthorize(t *testing.T) {
	var gotReq emigo.EmiRequestContexts
	authorize := func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		gotReq = req
		return fireback.QueryDSL{}, nil
	}

	handler := createSnapshotHandler(authorize)
	resp, err := handler(internalstatsdefs.InternalStatsSnapshotActionRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotReq == nil {
		t.Fatal("expected Authorize to be called with a request context")
	}

	envelope, ok := resp.Payload.(fireback.GoogleResponse[internalstatsdefs.InternalStatsSnapshotActionRes])
	if !ok {
		t.Fatalf("expected Payload to be a GoogleResponse[InternalStatsSnapshotActionRes], got %T", resp.Payload)
	}
	if envelope.Data.Item.Hostname == "" {
		t.Fatal("expected the snapshot's hostname to be filled in")
	}
}

// TestInternalStatsSnapshot_PropagatesAuthorizeError verifies a rejected/failed
// Authorize call short-circuits before ever collecting a snapshot.
func TestInternalStatsSnapshot_PropagatesAuthorizeError(t *testing.T) {
	wantErr := errors.New("nope")
	handler := createSnapshotHandler(func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		return fireback.QueryDSL{}, wantErr
	})

	resp, err := handler(internalstatsdefs.InternalStatsSnapshotActionRequest{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected Authorize's error to propagate, got %v", err)
	}
	if resp != nil {
		t.Fatalf("expected a nil response on auth failure, got %+v", resp)
	}
}

// TestInternalStatsStream_PushesFramesUntilDone verifies a successful Authorize leads
// to at least one JSON-encoded snapshot on the returned channel, and that closing
// session.Done stops further pushes (the channel eventually closes).
func TestInternalStatsStream_PushesFramesUntilDone(t *testing.T) {
	handler := createStreamHandler(func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		return fireback.QueryDSL{}, nil
	}, 10*time.Millisecond)

	session := internalstatsdefs.InternalStatsStreamActionSession{Done: make(chan bool)}
	out, err := handler(session)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	select {
	case frame, ok := <-out:
		if !ok {
			t.Fatal("expected at least one frame before the channel closed")
		}
		var snap internalstatsdefs.InternalStatsSnapshotActionRes
		if err := json.Unmarshal(frame, &snap); err != nil {
			t.Fatalf("decoding frame: %v", err)
		}
		if snap.Hostname == "" {
			t.Fatal("expected the decoded frame's hostname to be filled in")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for the first frame")
	}

	close(session.Done)

	select {
	case _, ok := <-out:
		if ok {
			// A second frame that was already in flight when Done closed is fine;
			// drain until the channel actually closes.
			for range out {
			}
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for the channel to close after Done")
	}
}

// TestInternalStatsStream_PropagatesAuthorizeError verifies a rejected/failed
// Authorize call short-circuits before ever touching the socket or starting the
// streaming goroutine.
func TestInternalStatsStream_PropagatesAuthorizeError(t *testing.T) {
	wantErr := errors.New("nope")
	handler := createStreamHandler(func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		return fireback.QueryDSL{}, wantErr
	}, time.Second)

	// Ctx/Socket/Done deliberately left unset - if the handler tried to use them
	// before returning the error, this would panic instead of erroring cleanly.
	_, err := handler(internalstatsdefs.InternalStatsStreamActionSession{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected Authorize's error to propagate, got %v", err)
	}
}

// TestStreamRequestContext_AdaptsSession verifies the emigo.EmiRequestContexts adapter
// exposes the session's gin context and always reports no CLI context (the reactive
// stream never runs over CLI).
func TestStreamRequestContext_AdaptsSession(t *testing.T) {
	fakeCtx := "stand-in for a *gin.Context"
	session := &internalstatsdefs.InternalStatsStreamActionSession{Ctx: fakeCtx}
	adapter := streamRequestContext{session: session}

	if adapter.GetGinCtx() != interface{}(fakeCtx) {
		t.Fatalf("expected GetGinCtx to return the session's Ctx, got %v", adapter.GetGinCtx())
	}
	if adapter.GetCliCtx() != nil {
		t.Fatalf("expected GetCliCtx to always be nil, got %v", adapter.GetCliCtx())
	}
}

// TestDefaultAuthorize_NoGinContext verifies the fireback.ResolveActionContext-based
// fallback used when InternalStatsModuleConfig.Authorize is left nil degrades to a
// zero QueryDSL (no panic, no error) when there's no real gin/cli context to resolve -
// matching fireback.ResolveActionContext's own documented behavior for a nilish
// context (AllowOnRoot never gets evaluated in that case, since the security check
// only runs inside the gin/cli branches).
func TestDefaultAuthorize_NoGinContext(t *testing.T) {
	query, err := defaultAuthorize(streamRequestContext{session: &internalstatsdefs.InternalStatsStreamActionSession{}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if query.WorkspaceId != "" || query.UserId != "" {
		t.Fatalf("expected a zero-value QueryDSL with no gin context, got %+v", query)
	}
}

// TestStreamSnapshots_EmitsImmediately verifies the first snapshot arrives right away
// rather than waiting a full interval - important for both the reactive websocket
// (a client shouldn't stare at nothing for `interval` seconds after connecting) and
// the `internalstats watch` table (first paint should be immediate).
func TestStreamSnapshots_EmitsImmediately(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := StreamSnapshots(ctx, time.Hour)
	select {
	case snap := <-ch:
		if snap == nil {
			t.Fatal("expected a non-nil snapshot")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for the immediate first snapshot")
	}
}

// TestStreamSnapshots_StopsOnContextCancel verifies the channel closes (and the
// underlying goroutine exits) once ctx is cancelled - the only way this is meant to
// stop, for both the reactive handler and the CLI watch loop.
func TestStreamSnapshots_StopsOnContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := StreamSnapshots(ctx, 5*time.Millisecond)
	<-ch // drain the immediate first snapshot
	cancel()

	select {
	case _, ok := <-ch:
		if ok {
			// A snapshot already in flight when cancel() ran is fine - keep
			// draining until the channel actually closes.
			for range ch {
			}
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for the channel to close after cancel")
	}
}
