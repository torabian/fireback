//go:build !wasm

package workspaces

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
 */
func excludeDatabaseConnection() bool {
	if Contains(os.Args, "gen") ||
		Contains(os.Args, "config") ||
		Contains(os.Args, "init") ||
		Contains(os.Args, "new") ||
		Contains(os.Args, "doctor") {
		return true
	}

	return false
}

func (x *XWebServer) CommonHeadlessAppStart(onDatabaseCompleted func()) {
	if !excludeDatabaseConnection() {

		db, dbErr := CreateDatabasePool()
		if db != nil && dbErr == nil {
			SyncDatabase(x, db)
			SyncPermissionsInDatabase(x, db)
		} else {
			log.Fatalln("Database error", dbErr)
		}

		if onDatabaseCompleted != nil {
			onDatabaseCompleted()
		}
	}

	x.CliActions = func() []cli.Command {
		return GetCommonWebServerCliActions(x)
	}

	RunApp(x)
}

func GetDatabaseDsn(info Database) (vendor string, dsn string) {
	uris := GetEnvironmentUris()
	vendor = info.Vendor

	if info.Vendor == "mysql" {
		dsn = info.Dsn
		if dsn == "" {
			dsn = info.Username + ":" + info.Password + "@tcp(" + info.Host + ":" + info.Port + ")/" + info.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
		}
	} else if info.Vendor == "postgres" {
		dsn = info.Dsn
		if dsn == "" {
			dsn = "host=" + info.Host + " user=" + info.Username + " password=" + info.Password + " dbname=" + info.Database + " port=" + info.Port + " sslmode=disable"

		}
	} else if info.Vendor == "sqlite" {
		var path = info.Database
		if path == "" {
			path = OsGetDefaultDatabase()
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

func DirectConnectToDb(info Database) (*gorm.DB, error) {

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "fb_",
		},
	}

	vendor, dsn := GetDatabaseDsn(info)

	var dialector gorm.Dialector

	if vendor == "mysql" {
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

	config := GetAppConfig()
	db, err := DirectConnectToDb(config.Database)
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(1)
	return db, nil

}

var DB_ORDER_DESC = "Created desc"
