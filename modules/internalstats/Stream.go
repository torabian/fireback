package internalstats

import (
	"context"
	"time"

	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

// DefaultInterval is used whenever InternalStatsModuleConfig.Interval is left zero.
const DefaultInterval = 2 * time.Second

// StreamSnapshots is the one place that decides "push a fresh snapshot on a ticker" -
// StreamActionImplementation.go marshals each one to JSON for the reactive websocket,
// and the `internalstats watch` CLI table (Cli.go) renders each one directly in-process,
// so both are really just two different renderers over this same feed rather than
// separate implementations of "collect on an interval".
//
// The first snapshot is pushed immediately (not after waiting a full interval), and the
// channel is closed once ctx is done - the only way this ever stops.
func StreamSnapshots(ctx context.Context, interval time.Duration) <-chan *internalstatsdefs.InternalStatsSnapshotActionRes {
	if interval <= 0 {
		interval = DefaultInterval
	}

	out := make(chan *internalstatsdefs.InternalStatsSnapshotActionRes)

	go func() {
		defer close(out)

		emit := func() bool {
			select {
			case out <- CollectSnapshot():
				return true
			case <-ctx.Done():
				return false
			}
		}

		if !emit() {
			return
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if !emit() {
					return
				}
			}
		}
	}()

	return out
}
