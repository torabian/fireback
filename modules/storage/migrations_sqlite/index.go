// package migrations_sqlite mirrors modules/storage/migrations, but holds the
// sqlite-flavored schema (see 00001_create_tus_uploads.sql's own comment for
// why it isn't just the same SQL) - selected instead of migrations.MigrationsFs
// whenever the app's configured DB vendor is sqlite, not postgres. See
// StorageModuleSetup (StorageModule.go) and Migrate.go's MigrateSQLite.
package migrations_sqlite

import "embed"

//go:embed *.sql
var MigrationsFs embed.FS
