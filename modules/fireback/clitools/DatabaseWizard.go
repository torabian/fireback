//go:build !wasm

// DatabaseWizard.go is the interactive half of fireback's database setup
// (`fireback init` / `fireback config db`): it asks which vendor, probes
// whether that vendor is already listening on its default port, and either
// takes a DSN directly or builds one field-by-field - listing (and
// optionally creating) real databases on the server along the way instead
// of asking the user to type a name blind.
//
// It never assembles or stores connection details itself: every piece it
// gathers (host/port/username/password/database/ssl) is handed to
// modules/fireback/dbdsn to become the one Dsn string that actually gets
// persisted, and any defaults it suggests come from dbdsn parsing whatever
// Dsn is already configured. See DatabaseCommon.go's Database struct doc
// for why that split exists.
package clitools

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // registers the "mysql" database/sql driver, used to list/create databases interactively
	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" database/sql driver, used to list/create databases interactively
	"github.com/manifoldco/promptui"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/dbdsn"
)

var useDsnOption = "I have dsn query string for connection"
var useManualOption = "I enter port, host, username of database manually"

var createNewDatabaseOption = "➕ Create a new database"
var typeDatabaseNameOption = "✏️  Type the database name"

// defaultDatabasePort returns the well-known default port for a database
// vendor, used both as the prompt suggestion and as the address probed by
// detectDefaultDatabasePort.
func defaultDatabasePort(vendor string) string {
	if vendor == fireback.DATABASE_TYPE_POSTGRES {
		return "5432"
	}
	return "3306"
}

// detectDefaultDatabasePort does a quick TCP dial to see if something is
// already listening on host:port. It's only meant to catch the common case
// (the database installed and running on its default port) - scanning every
// port to find a non-default install would be slow and unreliable, so we
// don't attempt that.
func detectDefaultDatabasePort(host string, port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 400*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// announceDatabasePortDetection probes 127.0.0.1:port and, if something
// answers, tells the user - the suggested port itself doesn't change, since
// it's already the vendor default.
func announceDatabasePortDetection(label string, port string) {
	if detectDefaultDatabasePort("127.0.0.1", port) {
		fmt.Printf("✔ Found a %s server listening on 127.0.0.1:%s, using it as the suggested port below.\n", label, port)
	}
}

// indexOf returns the position of target in items, or 0 (the first item) if
// it isn't there - used to point promptui.Select's cursor at whatever value
// is already configured, falling back to the usual first-item default.
func indexOf(items []string, target string) int {
	if target == "" {
		return 0
	}
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return 0
}

func orDefault(value string, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
}

// configuredDsn returns the already-loaded DB_DSN, but only when it was
// saved for the same vendor currently being configured.
func configuredDsn(vendor string) string {
	if config.DbVendor != vendor {
		return ""
	}
	return config.DbDsn
}

// configuredDbDefaults parses whatever DB_DSN is already loaded (from
// .env) into its individual pieces, so the wizard can prefill them - but
// only when that Dsn was saved for the same vendor we're currently
// configuring; a postgres port or mysql username from a previous setup
// isn't a sane suggestion once the user switches vendors.
func configuredDbDefaults(vendor string) dbdsn.ConnectionInfo {
	if config.DbVendor != vendor || config.DbDsn == "" {
		return dbdsn.ConnectionInfo{}
	}
	info, err := dbdsn.Parse(vendor, config.DbDsn)
	if err != nil {
		return dbdsn.ConnectionInfo{}
	}
	return info
}

func askProjectDatabase(projectName string) (fireback.Database, error) {
	db := fireback.Database{}

	items := []string{
		fireback.DATABASE_TYPE_SQLITE,
		fireback.DATABASE_TYPE_SQLITE_MEMORY,
		fireback.DATABASE_TYPE_MYSQL,
		fireback.DATABASE_TYPE_MARIADB,
		fireback.DATABASE_TYPE_POSTGRES,
	}

	// If a .env is already loaded, point the cursor at the vendor it
	// configured, so re-running the wizard doesn't start back at the top.
	configuredVendor := config.DbVendor
	if configuredVendor == fireback.DATABASE_TYPE_SQLITE && config.DbDsn == ":memory:" {
		configuredVendor = fireback.DATABASE_TYPE_SQLITE_MEMORY
	}

	promptVariable := promptui.Select{
		Label:     "Database type",
		Items:     items,
		CursorPos: indexOf(items, configuredVendor),
	}

	_, databaseType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	db.Vendor = databaseType

	if db.Vendor == fireback.DATABASE_TYPE_SQLITE {
		defaultPath := ""
		if config.DbVendor == fireback.DATABASE_TYPE_SQLITE && config.DbDsn != "" && config.DbDsn != ":memory:" {
			defaultPath = config.DbDsn
		}
		path, err := askSQLiteDatabaseLocation(projectName, defaultPath)
		if err != nil {
			fmt.Printf("cannot access the sqlite database, or cannot create it %v\n", err)
			return db, err
		}
		// sqlite has no DSN to speak of beyond the file path itself - see
		// GetDatabaseDsn.
		db.Dsn = path
	} else if db.Vendor == fireback.DATABASE_TYPE_SQLITE_MEMORY {
		db.Dsn = ":memory:"
		db.Vendor = fireback.DATABASE_TYPE_SQLITE
	} else if db.Vendor == fireback.DATABASE_TYPE_MYSQL || db.Vendor == fireback.DATABASE_TYPE_MARIADB {
		askMysqlDetails(&db)
	} else if db.Vendor == fireback.DATABASE_TYPE_POSTGRES {
		askPostgresDetails(&db)
	}

	return db, nil
}

func askMysqlDetails(db *fireback.Database) (*fireback.Database, error) {

	defaults := configuredDbDefaults(db.Vendor)
	defaultPort := orDefault(defaults.Port, defaultDatabasePort(db.Vendor))

	label := "mysql"
	if db.Vendor == fireback.DATABASE_TYPE_MARIADB {
		label = "mariadb"
	}
	announceDatabasePortDetection(label, defaultPort)

	defaultDsn := configuredDsn(db.Vendor)
	actionItems := []string{useDsnOption, useManualOption}
	defaultAction := useManualOption
	if defaultDsn != "" {
		defaultAction = useDsnOption
	}

	promptVariable := promptui.Select{
		Label:     "Do you have dsn string or port, host , username?",
		Items:     actionItems,
		CursorPos: indexOf(actionItems, defaultAction),
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == useDsnOption {
		value, err := askMysqlDsn(defaultDsn)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	db.Host = askHostName(orDefault(defaults.Host, "127.0.0.1"))
	db.Port = askHostPort(defaultPort)
	db.Username = askHostUsername(orDefault(defaults.Username, "root"))
	db.Password = askHostPassword(defaults.Password)
	db.Database = askDatabaseSelection(db.Vendor, db, defaults.Database)

	db.Dsn, _ = dbdsn.Build(db.Vendor, dbdsn.ConnectionInfo{
		Host: db.Host, Port: db.Port, Username: db.Username, Password: db.Password, Database: db.Database,
	})

	return db, nil
}

func askPostgresDetails(db *fireback.Database) (*fireback.Database, error) {

	defaults := configuredDbDefaults(db.Vendor)
	defaultPort := orDefault(defaults.Port, defaultDatabasePort(db.Vendor))
	announceDatabasePortDetection("postgres", defaultPort)

	defaultDsn := configuredDsn(db.Vendor)
	actionItems := []string{useDsnOption, useManualOption}
	defaultAction := useManualOption
	if defaultDsn != "" {
		defaultAction = useDsnOption
	}

	promptVariable := promptui.Select{
		Label:     "Do you have dsn string or port, host , username?",
		Items:     actionItems,
		CursorPos: indexOf(actionItems, defaultAction),
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == useDsnOption {
		value, err := askPostgresDsn(defaultDsn)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	db.Host = askHostName(orDefault(defaults.Host, "127.0.0.1"))
	db.Port = askHostPort(defaultPort)
	db.Username = askHostUsername(orDefault(defaults.Username, "postgres"))
	db.Password = askHostPassword(defaults.Password)
	db.SSL = askPostgresSSL(defaults.SSL)
	db.Database = askDatabaseSelection(db.Vendor, db, defaults.Database)

	db.Dsn = dbdsn.BuildPostgres(dbdsn.ConnectionInfo{
		Host: db.Host, Port: db.Port, Username: db.Username, Password: db.Password, Database: db.Database, SSL: db.SSL,
	})

	return db, nil
}

func askPostgresSSL(defaultSSL bool) bool {
	items := []string{"no", "yes"}
	defaultAnswer := "no"
	if defaultSSL {
		defaultAnswer = "yes"
	}

	promptVariable := promptui.Select{
		Label:     "Use SSL for the postgres connection?",
		Items:     items,
		CursorPos: indexOf(items, defaultAnswer),
	}

	_, answer, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return defaultSSL
	}

	return answer == "yes"
}

// askDatabaseSelection tries to connect with the credentials gathered so
// far and list the databases already on the server, offering them as a
// pick-list instead of making the user type a name blind. Creating a new
// database and typing a name manually are always the first two choices.
// If listing fails (bad credentials, no network access, unsupported
// vendor, ...) it falls back to the plain text prompt.
func askDatabaseSelection(vendor string, db *fireback.Database, defaultName string) string {
	names, err := listDatabases(vendor, db.Host, db.Port, db.Username, db.Password, db.SSL)
	if err != nil {
		fmt.Printf("Could not list existing databases (%v), please type the database name.\n", err)
		return askDatabaseName(defaultName)
	}

	items := append([]string{createNewDatabaseOption, typeDatabaseNameOption}, names...)

	// If the .env already named a database that still exists on the server,
	// point the cursor straight at it instead of "create new".
	promptVariable := promptui.Select{
		Label:     "Select the database to use",
		Items:     items,
		CursorPos: indexOf(items, defaultName),
	}

	_, selection, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return askDatabaseName(defaultName)
	}

	if selection == typeDatabaseNameOption {
		return askDatabaseName(defaultName)
	}

	if selection == createNewDatabaseOption {
		name := askDatabaseName(defaultName)
		if err := createDatabase(vendor, db.Host, db.Port, db.Username, db.Password, name, db.SSL); err != nil {
			fmt.Printf("Could not create database %q: %v\n", name, err)
		} else {
			fmt.Printf("✔ Database %q created\n", name)
		}
		return name
	}

	return selection
}

// listDatabases connects to the server (not to any particular application
// database - postgres always has its "postgres" maintenance database,
// mysql/mariadb are happy with no database selected at all) and returns the
// databases visible on it.
func listDatabases(vendor string, host string, port string, username string, password string, ssl bool) ([]string, error) {
	switch vendor {
	case fireback.DATABASE_TYPE_POSTGRES:
		conn, err := sql.Open("pgx", postgresAdminDsn(host, port, username, password, ssl))
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		rows, err := conn.Query(`select datname from pg_database where not datistemplate and datallowconn order by datname`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return scanDatabaseNames(rows)

	case fireback.DATABASE_TYPE_MYSQL, fireback.DATABASE_TYPE_MARIADB:
		conn, err := sql.Open("mysql", mysqlAdminDsn(host, port, username, password))
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		rows, err := conn.Query(`select schema_name from information_schema.schemata where schema_name not in ('information_schema','mysql','performance_schema','sys') order by schema_name`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return scanDatabaseNames(rows)
	}

	return nil, fmt.Errorf("listing databases is not supported for vendor %q", vendor)
}

// createDatabase issues a CREATE DATABASE against the server using the
// admin connection (no target database selected).
func createDatabase(vendor string, host string, port string, username string, password string, name string, ssl bool) error {
	switch vendor {
	case fireback.DATABASE_TYPE_POSTGRES:
		conn, err := sql.Open("pgx", postgresAdminDsn(host, port, username, password, ssl))
		if err != nil {
			return err
		}
		defer conn.Close()

		_, err = conn.Exec(fmt.Sprintf("create database %s", pgQuoteIdentifier(name)))
		return err

	case fireback.DATABASE_TYPE_MYSQL, fireback.DATABASE_TYPE_MARIADB:
		conn, err := sql.Open("mysql", mysqlAdminDsn(host, port, username, password))
		if err != nil {
			return err
		}
		defer conn.Close()

		_, err = conn.Exec(fmt.Sprintf("create database `%s`", strings.ReplaceAll(name, "`", "``")))
		return err
	}

	return fmt.Errorf("creating databases is not supported for vendor %q", vendor)
}

func scanDatabaseNames(rows *sql.Rows) ([]string, error) {
	var out []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		out = append(out, name)
	}
	return out, rows.Err()
}

// postgresAdminDsn connects to postgres's own always-present "postgres"
// maintenance database, so we have somewhere to run pg_database/CREATE
// DATABASE queries from before an application database exists. It's built
// through dbdsn.BuildPostgres like every other postgres dsn in this file -
// no separate hand-rolled fmt.Sprintf here, so the empty-password handling
// (see BuildPostgres's doc comment) only has to be right once.
func postgresAdminDsn(host string, port string, username string, password string, ssl bool) string {
	return dbdsn.BuildPostgres(dbdsn.ConnectionInfo{
		Host: host, Port: port, Username: username, Password: password, Database: "postgres", SSL: ssl,
	})
}

// mysqlAdminDsn connects without selecting a database, which mysql/mariadb
// are fine with for information_schema queries and CREATE DATABASE.
func mysqlAdminDsn(host string, port string, username string, password string) string {
	return dbdsn.BuildMysql(dbdsn.ConnectionInfo{
		Host: host, Port: port, Username: username, Password: password,
	})
}

func pgQuoteIdentifier(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

func askHostUsername(defaultUsername string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("enter database username")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "Database username",
		Validate: validate,
		Default:  defaultUsername,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func promptInput(label, defaultValue string, validate func(string) error) string {
	promptVariable := promptui.Prompt{
		Label:    label,
		Default:  defaultValue,
		Validate: validate,
	}

	result, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return ""
	}

	return result
}

func askDatabaseName(defaultName string) string {
	validateDatabaseName := func(input string) error {
		if input == "" {
			return errors.New("database name is required on this type of databse.")
		}
		return nil
	}

	return promptInput("Database name", defaultName, validateDatabaseName)
}

func askHostName(defaultHost string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the mysql host, for example localhost")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "The host, ip which mysql is installed. (eg. 127.0.0.1 or localhost or 210.231.20.30",
		Validate: validate,
		Default:  defaultHost,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askHostPassword(defaultPassword string) string {

	promptVariable := promptui.Prompt{
		Label:   "password",
		Default: defaultPassword,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askHostPort(defaultp string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the database port")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "port",
		Validate: validate,
		Default:  defaultp,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askSQLiteDatabaseLocation(projectName string, defaultPath string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the database path on file system, eg. /tmp/database1.db")
		}
		return nil
	}

	if defaultPath == "" {
		workingDirectory, err := filepath.Abs(".")
		if err != nil {
			log.Println(err)
		}
		defaultPath = filepath.Join(workingDirectory, projectName+"-database.db")
	}

	promptVariable := promptui.Prompt{
		Label:    "Database file location (.db)",
		Validate: validate,
		Default:  defaultPath,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}

func askMysqlDsn(defaultDsn string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: username:password@protocol(address)/dbname?param=value)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: username:password@protocol(address)/dbname?param=value)",
		Validate: validate,
		Default:  defaultDsn,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}

func askPostgresDsn(defaultDsn string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)",
		Validate: validate,
		Default:  defaultDsn,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}
