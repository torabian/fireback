//go:build !wasm

package fireback

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	// This is go version (glebarez), use it for desktop and server, and use gorm.io/driver/sqlite for mobile
	// on Android and IOS the golang version does not work. On the server, it's better not use cgo
	// to make it portable on more operating systems.

	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dbref *gorm.DB

var databaseConnectionError error

func GetDbRef() *gorm.DB {
	// I think it's a wise choise that expect this would always return a database
	// reference. Use other functions to see if the database connection exists, if ever
	// needed.

	if dbref == nil {
		log.Fatalln("Database connection is not available. It might be because of exclusion for database for some commands, look for excludeDatabaseConnection function")
	}

	return dbref
}

func GetExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	return exPath + fmt.Sprintf("%c", os.PathSeparator)
}

/*
* Cli would need database for almost everything, excerpt few things
* We exclude them here, because those commands actually needed for setup the database``
 */
func excludeDatabaseConnection() bool {
	if len(os.Args) == 1 {
		return true
	}

	// We need a bit better check here, because they might match to other
	// configuration that user generates and database connection will be terminated and they have
	// no idea why.
	if os.Args[1] == "gen" ||
		os.Args[1] == "config" ||
		os.Args[1] == "init" ||
		os.Args[1] == "new" ||
		os.Args[1] == "doctor" {
		return true
	}

	return false
}

func (x *FirebackApp) CommonHeadlessAppStart(onDatabaseCompleted func()) {
	commonHeadlessStarter(x, onDatabaseCompleted, true)
}
func (x *FirebackApp) CommonHeadlessMsStart(onDatabaseCompleted func()) {
	commonHeadlessStarter(x, onDatabaseCompleted, false)
}

func commonHeadlessStarter(x *FirebackApp, onDatabaseCompleted func(), completeTool bool) {

	// Use the logger
	initLogger()

	if !excludeDatabaseConnection() {
		// This is experimental to start the event bus
		// Honestly maybe this needs to be running also even without database scenario,
		// But I am not sure yet how it should work on cli
		StartEventBus()

		db, dbErr := CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		if onDatabaseCompleted != nil {
			onDatabaseCompleted()
		}
	}

	x.CliActions = func() []cli.Command {
		if completeTool {
			return GetCommonWebServerCliActions(x)
		}
		return GetCommonMicroserviceCliActions(x)
	}

	RunApp(x)
}

func GetDatabaseDsn(config Config) (vendor string, dsn string) {
	uris := GetEnvironmentUris()
	vendor = config.DbVendor

	if vendor == "mysql" || vendor == "mariadb" {
		dsn = config.DbDsn
		if dsn == "" {
			dsn = config.DbUsername + ":" + config.DbPassword + "@tcp(" + config.DbHost + ":" + fmt.Sprintf("%v", config.DbPort) + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		}
	} else if vendor == "postgres" {
		dsn = config.DbDsn
		if dsn == "" {
			dsn = "host=" + config.DbHost + " user=" + config.DbUsername + " password=" + config.DbPassword + " dbname=" + config.DbName + " port=" + fmt.Sprintf("%v", config.DbPort) + " sslmode=disable"
		}
	} else if vendor == "sqlite" {
		var path = config.DbName
		if path == "" {
			path = ":memory:"
		}

		forceDn := os.Getenv("FORCE_DATABASE_PATH")
		if forceDn != "" {
			path = forceDn
			fmt.Println("Using forced databse path:", path)
		}

		dsn = strings.ReplaceAll(path, "{appDataDirectory}", uris.AppDataDirectory)
	}
	return
}

func logLevelToNumber(level string) int {
	if level == "silent" {
		return 1
	}
	if level == "error" {
		return 2
	}
	if level == "warn" {
		return 3
	}
	if level == "info" {
		return 4
	}

	return 1
}

func DirectConnectToDb(config Config) (*gorm.DB, error) {

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevelToNumber(config.DbLogLevel))),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.TablePrefix,
		},
	}

	vendor, dsn := GetDatabaseDsn(config)

	var dialector gorm.Dialector

	if vendor == "mysql" || vendor == "mariadb" {
		dialector = mysql.Open(dsn)
	} else if vendor == "postgres" {
		dialector = postgres.Open(dsn)
	} else if vendor == "sqlite" {
		dialector = GetSQLiteDialector(dsn)
	}

	db, err := gorm.Open(dialector, gormConfig)

	if err != nil {
		return nil, err
	}

	dbref = db

	// Database couldn't be connected, requires someone to take care of it
	if databaseConnectionError != nil {
		panic("failed to connect database" + databaseConnectionError.Error())
	}

	return dbref, nil
}

func CreateDatabasePool() (*gorm.DB, error) {

	db, err := DirectConnectToDb(config)
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	// In SQLite for some bizare reason if there are more than 1 connection
	// it would freeze the app on some envrionments
	if config.DbVendor == DATABASE_TYPE_SQLITE || config.DbVendor == DATABASE_TYPE_SQLITE_MEMORY {
		sqlDb.SetMaxOpenConns(1)
		sqlDb.SetMaxIdleConns(1)
		sqlDb.SetConnMaxLifetime(time.Minute)
		db.Exec("PRAGMA busy_timeout = 5000")
		db.Exec("PRAGMA foreign_keys = ON")
	}

	return db, nil

}

var DB_ORDER_DESC = "Created desc"

type ReplicaRef struct {
	Clickhouse driver.Conn
}

type ReplicaService = string

const (
	ClickHouse ReplicaService = "clickhouse"
)

var clickhouseInstance driver.Conn = nil

func GetReplica(name ReplicaService) (*ReplicaRef, error) {

	if name == ClickHouse {
		if clickhouseInstance == nil {
			conn, err := connectToClickHouse()
			if err != nil {
				return nil, err
			}
			clickhouseInstance = conn
		}

		return &ReplicaRef{
			Clickhouse: clickhouseInstance,
		}, nil

	}

	return nil, errors.New("Not implemented yet")
}
