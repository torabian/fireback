package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/torabian/fireback/modules/storage/migrations"
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
