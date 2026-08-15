//go:build wasm

package application

import (
	"syscall/js"

	"github.com/torabian/fireback/modules/fireback/wasmpgdriver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectWasmPostgres is the wasm counterpart of fireback.DirectConnectToDb.
// The non-wasm path hands gorm.io/driver/postgres a DSN and lets it dial a
// real TCP connection; here there is no socket, so instead we hand it a
// *sql.DB whose driver forwards every statement to queryFunc — normally
// js.Global().Get("queryDatabase"), injected by browser/database-bridge.js
// in front of an in-browser pglite instance (see the emi in-browser-server
// example). Everything above the Dialector — gorm.Open, AutoMigrate, Find,
// Create, transactions, hooks — is completely unaware of the difference.
func ConnectWasmPostgres(queryFunc js.Value, gormConfig *gorm.Config) (*gorm.DB, error) {
	if queryFunc.IsUndefined() || queryFunc.IsNull() {
		panic("ConnectWasmPostgres: queryFunc is not defined — is database-bridge.js loaded before the wasm module starts?")
	}

	if gormConfig == nil {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	sqlDB := wasmpgdriver.Open(queryFunc)

	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
		// pglite speaks the extended protocol on a single in-process session;
		// there's no wire-level prepared statement cache to reuse across
		// "connections" the way a pooled TCP driver would, so keep gorm on
		// the simple protocol to avoid relying on server-side statement
		// naming that the bridge never actually implements.
		PreferSimpleProtocol: true,
	})

	return gorm.Open(dialector, gormConfig)
}
