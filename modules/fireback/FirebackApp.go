package fireback

import (
	"go.uber.org/zap"
)

// SERVER_INSTANCE - a random id minted once per process, not just !wasm builds
// (modules/eventbus's own EventBus.go/EventBusSubscriptionActionImplementation.go
// use it to tell this instance's own published events apart from another
// server's on the same bus) - genuinely needed everywhere, unlike RunApp below
// (moved to RunAppCli.go, !wasm-only: CLI dispatch the wasm binary never does).
var SERVER_INSTANCE string = UUID_Long()

// GooseZapLogger adapts *zap.Logger to goose's own logger interface (used by
// MigrationManager.go, which does run under wasm - migrations apply against
// pglite the same as any other gorm.DB) - kept here, not in the !wasm-only
// RunAppCli.go, for that reason.
type GooseZapLogger struct {
	Logger *zap.Logger
}

func (l GooseZapLogger) Printf(format string, v ...interface{}) {
	l.Logger.Sugar().Infof(format, v...)
}

func (l GooseZapLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Sugar().Fatalf(format, v...)
}
