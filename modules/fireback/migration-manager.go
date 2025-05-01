package fireback

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func ApplyMigration(xapp *FirebackApp, level int64) {
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

	db := GetDbRef()
	db.Config.Logger = newLogger
	SyncDatabase(xapp, db)

	// This is a fireback data managemnt issue - we do not want it
	// on projects which are not using fireback permission system
	if os.Getenv("DISABLE_FIREBACK_DATA_MANAGEMENT") != "true" {
		SyncPermissionsInDatabase(xapp, db)
	}
}
