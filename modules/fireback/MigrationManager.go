package fireback

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm/logger"
)

func RunMigrationBasedOnGoose(db *sql.DB, moduleName string, dialect string, mdir *embed.FS) error {

	goose.SetTableName("goose_db_module_" + moduleName)
	goose.SetBaseFS(mdir)
	goose.SetDialect(dialect)

	// Set a logger to capture all output
	// Use our custom logger
	goose.SetLogger(GooseZapLogger{Logger: LOG})

	// Optionally, enable verbose mode
	goose.SetVerbose(true)

	// Run all migrations
	err := goose.Up(db, ".")
	if err != nil {
		if strings.Contains(err.Error(), "no migration files found") {
			// handle empty migration directory gracefully
			return nil
		}
		return err
	}

	return nil
}

func ApplyMigration(xapp *FirebackApp, level int64) {
	db := GetDbRef()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,            // Slow SQL threshold
			LogLevel:                  logger.LogLevel(level), // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			Colorful:                  false,                  // Disable color
		},
	)

	db.Config.Logger = newLogger

	originalDb, _ := db.DB()
	for _, item := range xapp.Modules {
		if item.GoMigrateDirectory != nil {

			name := item.Name
			if item.DatabaseMigrationHistoryName != "" {
				name = item.DatabaseMigrationHistoryName
			}

			if err := RunMigrationBasedOnGoose(
				originalDb,
				name,
				db.Dialector.Name(),
				item.GoMigrateDirectory,
			); err != nil {
				LOG.Sugar().Fatalln(err)

			}
		}

		// Customize the entity provider, specifically
		if item.EntityProvider != nil {
			item.EntityProvider(db)
		}

		for _, bundle := range item.EntityBundles {
			if err := dbref.AutoMigrate(bundle.AutoMigrationEntities...); err != nil {
				fmt.Println("There is an error on migrating:", bundle)
				log.Fatalln(err.Error())
			}

			for _, funx := range bundle.MigrationScripts {
				if funx.Exec != nil {
					fmt.Println(funx.Exec())
					os.Exit(0)
				}
			}
		}

		if item.MigrationFunction != nil {
			item.MigrationFunction(xapp, db)
		}
	}

}
