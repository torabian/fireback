package internalstats

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

// AuthorizeFn resolves/authorizes an incoming request for both InternalStatsSnapshot
// and InternalStatsStream. It's handed anything satisfying emigo.EmiRequestContexts
// (GetGinCtx/GetCliCtx) - the exact contract every generated *ActionRequest and
// reactive *ActionSession-derived context already implements (see
// streamRequestContext in StreamActionImplementation.go) - so it plugs straight into
// fireback.ResolveActionContext the same way a normal action would, or into a
// project's own auth logic instead: internalstats itself never has to know which, and
// never imports abac (or any other auth provider) to get a working default - see
// defaultAuthorize below and InternalStatsModuleConfig.Authorize.
type AuthorizeFn func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error)

// defaultAuthorize is used whenever InternalStatsModuleConfig.Authorize is left nil.
// Server health data is sensitive (hostnames, IPs indirectly, resource pressure a
// would-be attacker could use to time an attack) so, unlike reactivesearch's
// workspace-scoped default, this requires the caller's token to belong to the root
// workspace (AllowOnRoot) - fireback.ResolveActionContext enforces that by calling
// fireback.AuthorizeRequest, which is nil until some auth provider's own module setup
// assigns it (e.g. abac.WorkspaceModuleSetup sets fireback.AuthorizeRequest =
// abac.AuthorizeRequest - see modules/abac/AbacModule.go). internalstats never imports
// abac to get this; a project's main.go wires abac's root-only resolution in by
// passing an Authorize func built the same way defaultAuthorize is here (see
// cmd/fireback/main.go's internalstats.ModuleSetup call).
func defaultAuthorize(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
	query, err := fireback.ResolveActionContext(req, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyWorkspace,
		AllowOnRoot:     true,
	})
	if err != nil {
		return fireback.QueryDSL{}, err
	}
	return *query, nil
}

// createSnapshotHandler builds the InternalStatsSnapshotAction implementation
// ModuleSetup wires up - a closure (rather than a plain package-level function like
// abac's own actions) because which AuthorizeFn it enforces depends on
// InternalStatsModuleConfig, decided per ModuleSetup call, the same reason
// reactivesearch's createReactiveSearchHandler is a closure too.
func createSnapshotHandler(authorize AuthorizeFn) internalstatsdefs.InternalStatsSnapshotActionRequestSig {
	if authorize == nil {
		authorize = defaultAuthorize
	}

	return func(c internalstatsdefs.InternalStatsSnapshotActionRequest) (*internalstatsdefs.InternalStatsSnapshotActionResponse, error) {
		if _, err := authorize(c); err != nil {
			return nil, err
		}

		return &internalstatsdefs.InternalStatsSnapshotActionResponse{
			Payload: fireback.GResponseSingleItem(*CollectSnapshot()),
		}, nil
	}
}
