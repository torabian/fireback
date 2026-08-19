package abac

// Like analytics (see AnalyticsActionImplementation.go's own comment), internal
// stats has no entity of its own - modules/internalstats is a separate module that
// never imports abac, so it can't define an abac-style permission itself. This lives
// here instead, and is wired into internalstats.ModuleSetup's Authorize closure (and
// the "internal_stats" AppMenu entry's CapabilityId) from cmd/fireback/main.go, the
// one place that already imports both packages. Query is the only verb: internal
// stats is read-only, same as analytics.
var internalStatsPerms = NewCrudPermissionSet("root.modules", "internal-stats", "internal stats")
var PERM_ROOT_INTERNAL_STATS_QUERY = internalStatsPerms.Query
var ALL_INTERNAL_STATS_PERMISSIONS = internalStatsPerms.All
