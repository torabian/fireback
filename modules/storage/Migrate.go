package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/ncruces/go-sqlite3/driver"
	"github.com/pressly/goose/v3"

	"github.com/torabian/fireback/modules/storage/migrations"
	migrationsmysql "github.com/torabian/fireback/modules/storage/migrations_mysql"
	migrationssqlite "github.com/torabian/fireback/modules/storage/migrations_sqlite"
)

// Migrate applies the migrations in migrations/ against connString. It opens
// its own *sql.DB because goose works over database/sql, while the rest of
// this module talks to postgres through pgxpool for the large object API.
func Migrate(connString string) error {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(migrations.MigrationsFs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(db, ".")
}

// sqliteDSN turns a plain file path (as resolved by fireback.GetDatabaseDsn)
// into a "file:" URI carrying the pragmas this module wants on every
// connection it opens to a sqlite file - see OpenSQLiteStore's own doc
// comment for why a separate connection (distinct from GORM's own) exists
// at all. Order matters per the driver's own doc comment: busy_timeout
// before journal_mode.
func sqliteDSN(path string) string {
	return fmt.Sprintf(
		"file:%s?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)",
		path,
	)
}

// MigrateSQLite is Migrate's sqlite counterpart: it applies the migrations
// in migrations_sqlite/ against the sqlite database at path instead (see
// migrations_sqlite/00001_create_tus_uploads.sql's own comment for why the
// schema itself differs from Postgres').
func MigrateSQLite(path string) error {
	db, err := sql.Open("sqlite3", sqliteDSN(path))
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(migrationssqlite.MigrationsFs)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	return goose.Up(db, ".")
}

// OpenSQLiteStore opens a dedicated *sql.DB against the sqlite file at path
// (with the WAL/busy_timeout/foreign_keys pragmas sqliteDSN sets) and wraps
// it as a Store, having already run MigrateSQLite against it. Mirrors
// NewPgxPool's own shape/contract (separate connection from whatever GORM
// itself uses, migrations run first) for the sqlite vendor.
func OpenSQLiteStore(path string) (Store, error) {
	if err := MigrateSQLite(path); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", sqliteDSN(path))
	if err != nil {
		return nil, err
	}

	return NewSQLiteStore(db), nil
}

// MigrateMySQL is Migrate's MySQL/MariaDB counterpart: it applies the
// migrations in migrations_mysql/ against dsn instead (see
// migrations_mysql/00001_create_tus_uploads.sql's own comment for why the
// schema itself differs from Postgres'/sqlite's). dsn is whatever
// fireback.GetDatabaseDsn already builds for the mysql/mariadb vendor
// ("user:pass@tcp(host:port)/dbname?charset=...") - no extra pragma-style
// wrapping needed the way sqlite's DSN gets, since MySQL doesn't have an
// equivalent single-writer-database caveat to work around.
func MigrateMySQL(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(migrationsmysql.MigrationsFs)
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	return goose.Up(db, ".")
}

// OpenMySQLStore opens a dedicated *sql.DB against dsn and wraps it as a
// Store, having already run MigrateMySQL against it. Mirrors
// NewPgxPool/OpenSQLiteStore's own shape/contract (separate connection from
// whatever GORM itself uses, migrations run first) for the mysql/mariadb
// vendor.
func OpenMySQLStore(dsn string) (Store, error) {
	if err := MigrateMySQL(dsn); err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return NewMySQLStore(db), nil
}
