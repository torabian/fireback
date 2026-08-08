package reactivesearch

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

// ReactiveSearchModuleConfig lets a project pass the search sources it actually wants
// searched, instead of fireback reaching into a FirebackApp.SearchProviders field
// every project carried whether or not it used reactive search (see
// ReactiveSearch.emi.yml's description).
type ReactiveSearchModuleConfig struct {
	SearchProviders []SearchProviderFn

	// Authorize resolves/authorizes each incoming reactive search request. Left nil
	// (the default), it falls back to defaultAuthorize - fireback.ResolveActionContext
	// with ResolveStrategyWorkspace, which requires abac (or whatever else wires up
	// fireback.AuthorizeRequest/fireback.MeetsAccessLevel) to already be set up.
	// Override this to make reactivesearch's authorization fully self-contained -
	// your own token check, no fireback/abac auth machinery involved at all - by
	// implementing AuthorizeReactiveSearchFn against the given
	// emigo.EmiRequestContexts (GetGinCtx/GetCliCtx), the same contract every other
	// emi action request already satisfies.
	Authorize AuthorizeReactiveSearchFn
}

// ModuleSetup registers the reactivesearch module - only a project that actually
// calls this from its own main.go, passing the sources it wants searched, gets the
// /reactive-search route wired.
func ModuleSetup(cfg *ReactiveSearchModuleConfig) *fireback.ModuleProvider {
	if cfg == nil {
		cfg = &ReactiveSearchModuleConfig{}
	}

	module := &fireback.ModuleProvider{
		Name: "reactivesearch",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *fireback.FirebackApp) error{
			func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
				meta := ReactiveSearchActionMeta()
				g.GET(
					meta.URL,
					ReactiveSearchActionReactiveHandler(createReactiveSearchHandler(cfg.SearchProviders, cfg.Authorize)),
				)
				return nil
			},
		},
	}

	return module
}
