package internalstats

import (
	"context"
	"encoding/json"
	"time"

	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

// streamRequestContext adapts an InternalStatsStreamActionSession into
// emigo.EmiRequestContexts, mirroring reactivesearch's own
// reactiveSearchRequestContext (see modules/reactivesearch/ReactiveSearchActionImplementation.go)
// so AuthorizeFn implementations can reuse fireback.ResolveActionContext without
// internalstats needing its own gin-specific plumbing. The stream only ever runs over
// a live websocket connection, never CLI, so GetCliCtx is always nil.
type streamRequestContext struct {
	session *internalstatsdefs.InternalStatsStreamActionSession
}

func (r streamRequestContext) GetGinCtx() interface{} { return r.session.Ctx }
func (r streamRequestContext) GetCliCtx() interface{} { return nil }

// createStreamHandler builds the InternalStatsStreamAction reactive factory
// ModuleSetup wires up: authorize once per connection, then push a JSON-encoded
// snapshot from StreamSnapshots every interval until the client disconnects
// (session.Done) or the write side closes.
func createStreamHandler(authorize AuthorizeFn, interval time.Duration) func(
	session internalstatsdefs.InternalStatsStreamActionSession,
) (chan []byte, error) {
	if authorize == nil {
		authorize = defaultAuthorize
	}

	return func(session internalstatsdefs.InternalStatsStreamActionSession) (chan []byte, error) {
		if _, err := authorize(streamRequestContext{session: &session}); err != nil {
			return nil, err
		}

		out := make(chan []byte)

		go func() {
			defer close(out)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// session.Done fires when the client disconnects (see
			// InternalStatsStreamActionGin.go's generated handler) - stop the
			// underlying ticker goroutine (StreamSnapshots) the instant that
			// happens, rather than leaking it until the next tick discovers the
			// write side is gone.
			go func() {
				<-session.Done
				cancel()
			}()

			for snapshot := range StreamSnapshots(ctx, interval) {
				b, err := json.Marshal(snapshot)
				if err != nil {
					continue
				}
				select {
				case out <- b:
				case <-ctx.Done():
					return
				}
			}
		}()

		return out, nil
	}
}
