package eventbus

import (
	"github.com/gin-gonic/gin"
	eventbusdefs "github.com/torabian/fireback/modules/eventbus/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// EventBusModuleConfig configures the parts of the event bus a project actually
// needs, instead of fireback injecting them unconditionally (see EventBus.emi.yml's
// description).
type EventBusModuleConfig struct {

	// Redis address used to distribute events across instances (e.g.
	// "127.0.0.1:6379"). Empty (the default) means events stay local to this
	// instance - the in-process event manager is used and nothing is distributed
	// across other instances.
	RedisEventsUrl string

	// Filters which of a connected user's socket messages they're allowed to see,
	// based on Event.Security.ActionRequires/AllowOnRoot. Left nil (the default),
	// falls back to fireback.MeetsAccessLevel (which abac wires up, see
	// AbacModule.go) - only override this if you want eventbus to use a different
	// check than the rest of the app.
	MeetsAccessLevel func(query fireback.QueryDSL, onlyRoot bool) (bool, []string)
}

// ModuleSetup registers the eventbus module - only a project that actually calls this
// from its own main.go gets the /ws route wired and the event bus goroutine started
// (via OnAppStart, see modules/fireback/DatabaseConnection.go's commonHeadlessStarter).
func ModuleSetup(cfg *EventBusModuleConfig) *application.ModuleProvider {
	if cfg == nil {
		cfg = &EventBusModuleConfig{}
	}

	meetsAccessLevel = cfg.MeetsAccessLevel

	module := &application.ModuleProvider{
		Name: "eventbus",

		OnAppStart: func(x *application.Application) error {
			StartEventBus(cfg.RedisEventsUrl)
			return nil
		},

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				meta := eventbusdefs.EventBusSubscriptionActionMeta()
				g.GET(meta.URL, eventbusdefs.EventBusSubscriptionActionReactiveHandler(EventBusSubscriptionActionSig))
				return nil
			},
		},
	}

	return module
}
