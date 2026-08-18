//go:build wasm

package fireback

import "gorm.io/gorm"

// wasm counterpart of DatabaseConnection.go's dbref/GetDbRef - CrudCoreActions.go
// (used by every entity, not just abac's) calls GetDbRef() to get "the current
// database" without a gorm.DB threaded through every call. The non-wasm side sets
// it inside DirectConnectToDb once a real TCP dial succeeds; there's no equivalent
// single connect-and-set entry point for wasm (application.ConnectWasmPostgres just
// returns a *gorm.DB to whoever calls it - see that function's own doc comment), so
// cmd/fireback-wasm/main.go calls SetDbRef itself right after ConnectWasmPostgres
// succeeds.
var dbref *gorm.DB

func GetDbRef() *gorm.DB {
	if dbref == nil {
		panic("Database connection is not available - call fireback.SetDbRef(db) after application.ConnectWasmPostgres succeeds (see cmd/fireback-wasm/main.go).")
	}

	return dbref
}

// SetDbRef registers the *gorm.DB every GetDbRef() call after this one will
// return. Wasm-only counterpart to DirectConnectToDb's own dbref assignment (see
// DatabaseConnection.go) - there's no single "connect" entry point on this side to
// hook that into instead.
func SetDbRef(db *gorm.DB) {
	dbref = db
}
