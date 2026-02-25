package fireback

import (
	"database/sql"
	"embed"
	"os"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func ApplyMigration(xapp *FirebackApp, level int64) {
	db := GetDbRef()

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,            // Slow SQL threshold
	// 		LogLevel:                  logger.LogLevel(level), // Log level
	// 		IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,                   // Don't include params in the SQL log
	// 		Colorful:                  false,                  // Disable color
	// 	},
	// )

	// db.Config.Logger = newLogger
	SyncDatabase(xapp, db)

	// This is a fireback data managemnt issue - we do not want it
	// on projects which are not using fireback permission system
	if os.Getenv("DISABLE_FIREBACK_DATA_MANAGEMENT") != "true" {
		SyncPermissionsInDatabase(xapp, db)
	}
}

func ApplyGoMigration(db *sql.DB, fs embed.FS) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}

	source, err := iofs.New(fs, ".")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", source, "sqlite", driver)
	if err != nil {
		return err
	}

	return ignoreNoChange(m.Up())
}

func ignoreNoChange(err error) error {
	if err == migrate.ErrNoChange {
		return nil
	}
	return err
}
