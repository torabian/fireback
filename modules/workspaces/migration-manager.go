package workspaces

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func ApplyMigration(xapp *XWebServer, level int64) {
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
	SyncPermissionsInDatabase(xapp, db)
}
