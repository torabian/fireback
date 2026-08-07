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
					ReactiveSearchActionReactiveHandler(createReactiveSearchHandler(cfg.SearchProviders)),
				)
				return nil
			},
		},
	}

	return module
}
