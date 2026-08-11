// package migrations_mysql mirrors modules/storage/migrations_sqlite, but
// holds the MySQL/MariaDB-flavored schema (see
// 00001_create_tus_uploads.sql's own comment for why it isn't just the same
// SQL). Selected instead of migrations.MigrationsFs/migrations_sqlite's
// whenever the app's configured DB vendor is mysql/mariadb. See
// StorageModuleSetup (StorageModule.go) and Migrate.go's MigrateMySQL.
package migrations_mysql

import "embed"

//go:embed *.sql
var MigrationsFs embed.FS
