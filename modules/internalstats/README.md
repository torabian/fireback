# internalstats

Read-only server/process health stats - CPU, memory, disk, network, and this
process's own Go runtime - around 40 individual measurements (see
`Collector.go`'s `CollectSnapshot` for the full list). One function
(`CollectSnapshot`) is the single source of truth; every transport below is
just a different renderer over it:

- **`GET /internal-stats/snapshot`** - one point-in-time snapshot, as JSON
  (`InternalStatsSnapshot` in `InternalStats.emi.yml`).
- **`GET /internal-stats/stream`** (websocket, emi `method: reactive`) -
  the same snapshot shape, pushed on an interval until the client
  disconnects (`InternalStatsStream`).
- **`internalstats snapshot`** (CLI) - the generated one-shot JSON command,
  same auth path as the HTTP route.
- **`internalstats watch`** (CLI) - a live-refreshing, colored, two-column
  (label | value) terminal table, redrawn every tick. It consumes
  `StreamSnapshots` directly in-process (no network hop, no `Authorize`
  check - same trust boundary as `backup dump`/`config list`: an operator
  already holding a shell on the box) rather than reimplementing the
  refresh loop, so it's always exactly what a websocket client would see.

`gopsutil` (`github.com/shirou/gopsutil/v3`) does the actual OS-level
reading, cross-platform (linux/darwin/windows) - see `Collector.go`.

## Wiring this into a project

Only a project whose `main.go` calls `internalstats.ModuleSetup(...)` gets
the routes/CLI commands - same opt-in convention as every other module here
(`backup.ModuleSetup`, `reactivesearch.ModuleSetup`, ...).

```go
internalstats.ModuleSetup(&internalstats.InternalStatsModuleConfig{
	// Authorize gates both InternalStatsSnapshot and InternalStatsStream.
	// Left nil, defaultAuthorize requires a root-workspace token via
	// fireback.ResolveActionContext - which only actually enforces anything
	// once some auth provider's module setup has assigned
	// fireback.AuthorizeRequest (e.g. abac.WorkspaceModuleSetup does this -
	// see modules/abac/AbacModule.go). internalstats itself never imports
	// abac; a project wires abac's root-only check in explicitly, the same
	// way cmd/fireback/main.go does:
	Authorize: func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
		query, err := fireback.ResolveActionContext(req, &fireback.SecurityModel{
			ResolveStrategy: fireback.ResolveStrategyWorkspace,
			AllowOnRoot:     true,
		})
		if err != nil {
			return fireback.QueryDSL{}, err
		}
		return *query, nil
	},

	// How often InternalStatsStream pushes, and how often `watch` redraws.
	// Defaults to 2s.
	Interval: 2 * time.Second,
}),
```

## Regenerating defs after editing InternalStats.emi.yml

```
./app emi compile --path modules/internalstats/InternalStats.emi.yml
```

(also wired into the repo-wide `make defs` target).
