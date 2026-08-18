package fireback

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	"go.uber.org/zap"
)

// FirebackAppToGin wires every registered module's GinWebServerInitHooks onto g.
// Split out of CliActions.go (!wasm) into its own, untagged file so it can also
// be called from cmd/fireback-wasm/main.go: unlike the rest of that file (real
// TLS/ACME listener setup, connmonitor, embedded static assets, ...), this loop
// has no wasm-unsafe dependency - it only touches gin, zap and application,
// all of which already compile fine under GOOS=js (see ApplicationWasm.go's
// own doc comment).
func FirebackAppToGin(x *application.Application, g *gin.RouterGroup, prefix string) {

	for _, item := range x.Modules {

		for _, hook := range item.GinWebServerInitHooks {
			if err := hook(g, x); err != nil {
				LOG.Error("Error %w", zap.Error(err))
				LOG.Fatal("Gin server failed to run a hook on module", zap.String("module", item.Name))
			}
		}

	}
}
