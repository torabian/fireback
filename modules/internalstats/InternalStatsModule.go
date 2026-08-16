package internalstats

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
	"github.com/urfave/cli/v3"
)

// InternalStatsModuleConfig lets a project override how InternalStatsSnapshot/
// InternalStatsStream authorize their callers, and how often the reactive stream (and
// the `internalstats watch` CLI table) refreshes - the same shape reactivesearch's own
// ReactiveSearchModuleConfig uses, for the same reason: cfg is decided once, at
// ModuleSetup time, and closed over by the handlers built in
// SnapshotActionImplementation.go/StreamActionImplementation.go.
type InternalStatsModuleConfig struct {
	// Authorize resolves/authorizes each incoming request (both the JSON snapshot and
	// the reactive stream). Left nil (the default), it falls back to defaultAuthorize -
	// fireback.ResolveActionContext with SecurityModel{ResolveStrategy:
	// ResolveStrategyWorkspace, AllowOnRoot: true}, which requires a root-workspace
	// token and depends on abac (or whatever else wires up fireback.AuthorizeRequest)
	// having already run its own module setup. See defaultAuthorize's own comment for
	// why internalstats itself never imports abac to get this.
	Authorize AuthorizeFn

	// Interval between pushes on InternalStatsStream, and between redraws of the
	// `internalstats watch` CLI table (both are just renderers over the same
	// StreamSnapshots feed - see Stream.go). Defaults to DefaultInterval (2s) when
	// left zero.
	Interval time.Duration
}

// ModuleSetup registers the internalstats module - only a project that actually calls
// this from its own main.go gets the /internal-stats/* routes and the `internalstats`
// CLI command tree, the same opt-in convention every other module here uses
// (backup.ModuleSetup, reactivesearch.ModuleSetup, ...).
func ModuleSetup(cfg *InternalStatsModuleConfig) *application.ModuleProvider {
	if cfg == nil {
		cfg = &InternalStatsModuleConfig{}
	}

	interval := cfg.Interval
	if interval <= 0 {
		interval = DefaultInterval
	}

	snapshotHandler := createSnapshotHandler(cfg.Authorize)
	streamHandler := createStreamHandler(cfg.Authorize, interval)

	return &application.ModuleProvider{
		Name: "internalstats",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				internalstatsdefs.InternalStatsSnapshotActionGin(g, snapshotHandler)

				meta := internalstatsdefs.InternalStatsStreamActionMeta()
				g.GET(meta.URL, internalstatsdefs.InternalStatsStreamActionReactiveHandler(streamHandler))
				return nil
			},
		},

		CliHandlers: []*cli.Command{
			{
				Name:  "internalstats",
				Usage: "Server/process health stats - CPU, memory, disk, network, Go runtime (see modules/internalstats/Collector.go)",
				Commands: []*cli.Command{
					internalstatsdefs.InternalStatsSnapshotActionCliHandler(snapshotHandler),
					watchCommand(interval),
				},
			},
		},
	}
}
