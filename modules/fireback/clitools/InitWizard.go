//go:build !wasm

package clitools

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

// config is a pointer to the core package's package-level config instance
// (via fireback.GetConfigRef()), so mutating config.Field here mutates the
// same storage the core package reads from.
var config = fireback.GetConfig()

func init() {
	fireback.CLIInit = cliInit
	fireback.AskProjectDatabase = askProjectDatabase
	fireback.AskSSL = askSSL
	fireback.AskSqlLogLevel = askSqlLogLevel
	fireback.AskPortName = askPortName
	fireback.AskFolderName = askFolderName
}

var useDsnOption = "I have dsn query string for connection"
var useManualOption = "I enter port, host, username of database manually"

var tryToSolve = "Let me retry to configurate the database parameters"
var forceContinue = "Use the configuration without connection test"

func askProjectDatabase(projectName string) (fireback.Database, error) {
	db := fireback.Database{}

	promptVariable := promptui.Select{
		Label: "Database type",
		Items: []string{
			fireback.DATABASE_TYPE_SQLITE,
			fireback.DATABASE_TYPE_SQLITE_MEMORY,
			fireback.DATABASE_TYPE_MYSQL,
			fireback.DATABASE_TYPE_MARIADB,
			fireback.DATABASE_TYPE_POSTGRES,
		},
	}

	_, databaseType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	db.Vendor = databaseType

	if db.Vendor == "sqlite" {
		path, err := askSQLiteDatabaseLocation(projectName)
		if err != nil {
			fmt.Printf("cannot access the sqlite database, or cannot create it %v\n", err)
			return db, err
		}
		db.Database = path
	} else if db.Vendor == fireback.DATABASE_TYPE_SQLITE_MEMORY {
		db.Database = ":memory:"
		db.Vendor = "sqlite"
	} else if db.Vendor == fireback.DATABASE_TYPE_MYSQL || db.Vendor == fireback.DATABASE_TYPE_MARIADB {
		askMysqlDetails(&db)
	} else if db.Vendor == fireback.DATABASE_TYPE_POSTGRES {
		askPostgresDetails(&db)
	}

	return db, nil
}

func askMysqlDetails(db *fireback.Database) (*fireback.Database, error) {

	promptVariable := promptui.Select{
		Label: "Do you have dsn string or port, host , username?",
		Items: []string{useDsnOption, useManualOption},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == useDsnOption {
		value, err := askMysqlDsn()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	if actionType == useManualOption {

		db.Host = askHostName()
		db.Port = askHostPort("3306")
		db.Database = askDatabaseName()
		db.Username = askHostUsername("root")
		db.Password = askHostPassword()
	}

	return db, nil
}

func askPostgresDetails(db *fireback.Database) (*fireback.Database, error) {

	promptVariable := promptui.Select{
		Label: "Do you have dsn string or port, host , username?",
		Items: []string{useDsnOption, useManualOption},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == useDsnOption {
		value, err := askPostgresDsn()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	if actionType == useManualOption {

		db.Host = askHostName()
		db.Port = askHostPort("5432")
		db.Database = askDatabaseName()
		db.Username = askHostUsername("postgres")
		db.Password = askHostPassword()
	}

	return db, nil
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

func askDatabaseName() string {
	validateDatabaseName := func(input string) error {
		if input == "" {
			return errors.New("database name is required on this type of databse.")
		}
		return nil
	}

	return promptInput("Database name", "", validateDatabaseName)
}

func askHostName() string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the mysql host, for example localhost")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "The host, ip which mysql is installed. (eg. 127.0.0.1 or localhost or 210.231.20.30",
		Validate: validate,
		Default:  "localhost",
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askPortName(label string, defaultPort string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("port should be between 0 to 65536")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultPort,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askFolderName(label string, defaultFolder string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this folder is necessary for file uploads")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultFolder,
	}

	hostname, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return hostname
}

func askHostPassword() string {

	promptVariable := promptui.Prompt{
		Label:   "password",
		Default: "",
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

func askSqlLogLevel(cfg *fireback.Config) {

	SILENT_PICK := "Silent - shows nothing, useful for production environment"
	ERROR_PICK := "Error - Show sql errors"
	WARNING_PICK := "Warning - show only warnings"
	INFO_PICK := "Info - prints all queries to the database"

	level := fireback.AskForSelect("Select the database log level for SQL queries", []string{
		SILENT_PICK,
		ERROR_PICK,
		WARNING_PICK,
		INFO_PICK,
	})

	if level == SILENT_PICK {
		cfg.DbLogLevel = "silent"
	}
	if level == INFO_PICK {
		cfg.DbLogLevel = "info"
	}
	if level == WARNING_PICK {
		cfg.DbLogLevel = "warning"
	}
	if level == ERROR_PICK {
		cfg.DbLogLevel = "error"
	}

	cfg.Save(".env")
}

func askSSL(cfg *fireback.Config) {

	if r := fireback.AskForSelect("Use SSL instead of Plain Http?", []string{"no", "yes"}); r == "yes" {
		cfg.UseSSL = true

		cfg.CertFile = askFolderName("Certfile address", "/etc/letsencrypt/live/")
		cfg.KeyFile = askFolderName("Keyfile address", "/etc/letsencrypt/live/")

	} else {
		cfg.UseSSL = false
	}

	cfg.Save(".env")
}

func askRetry() bool {
	promptVariable := promptui.Select{
		Label: "Database connection failed, do you want retry again?",
		Items: []string{tryToSolve,
			forceContinue},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	if actionType == tryToSolve {
		return true
	}

	return false
}

func cliInit(xapp *fireback.FirebackApp) *cli.Command {

	return &cli.Command{
		Name:  "init",
		Usage: "Creates a environment for project, by configurating database connection, http port, etc.",
		Flags: fireback.GetConfigCliFlags(),
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NumFlags() > 0 {
				fireback.CastConfigFromCli(&config, c)

				if !c.IsSet("mac-identifier") {
					config.MacIdentifier = config.Name
				}

				if !c.IsSet("debian-identifier") {
					config.DebianIdentifier = config.Name
				}

				if !c.IsSet("windows-identifier") {
					config.WindowsIdentifier = config.Name
				}

				config.Save(".env")
			} else {
				initEnvironment(xapp, ".env", c)
			}
			return nil
		},
	}
}

func dataBaseConfigEnv(xapp *fireback.FirebackApp) error {

	// 2. Determine the database type, test the connection, create tables
	for {
		databaseData, err := askProjectDatabase(config.Name)
		if err != nil {
			log.Fatalln("cannot determine the database config", err)
			return nil
		}

		// 3. Check if the database could be connected, if not show error and move on
		config.DbUsername = databaseData.Username
		p, _ := strconv.Atoi(databaseData.Port)
		config.DbPort = int64(p)
		config.DbHost = databaseData.Host
		config.DbPassword = databaseData.Password
		config.DbName = databaseData.Database
		config.DbVendor = databaseData.Vendor
		config.DbDsn = databaseData.Dsn

		db, err := fireback.DirectConnectToDb(config)
		if err == nil && db.Exec("select 1").Error == nil {
			config.Save(".env")
			fmt.Println("✔ connection is successful")
			break
		}

		fmt.Println(err)

		if !askRetry() {
			break
		}
	}

	return nil
}

func executeSeeders(xapp *fireback.FirebackApp) error {
	if r := fireback.AskForSelect("Do you want to add the seed data, menu items, etc?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		fireback.ExecuteSeederImport(xapp)
	}

	return nil
}

func envRunMigration(xapp *fireback.FirebackApp) error {
	if r := fireback.AskForSelect("Do you want to run migration, adding tables or columns to database (both automigrate, and manual)?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		fireback.ApplyMigration(xapp, 0)
	}

	return nil
}

func initEnvironment(xapp *fireback.FirebackApp, envFileName string, c *cli.Command) error {

	if fireback.AskForSelect("This command is to generate a .env file, for existing project, or standalone fireback installation. You need to use `fireback new` to create new project. Agree?", []string{"yes", "no"}) == "no" {
		return nil
	}

	datum := ""
	var err error

	// 1. Determine the project name
	datum, err = askEnvironmentName(config.Name)
	if err != nil {
		log.Fatalln("cannot determine the project name", err)
		return nil
	}
	config.Name = datum
	config.DebianIdentifier = datum
	config.MacIdentifier = datum
	config.WindowsIdentifier = datum

	if isProd := fireback.AskForSelect("Is this a production environment?", []string{"no", "yes"}); isProd == "yes" {
		if isProd == "yes" {
			config.Production = true
		}
	}

	if err := dataBaseConfigEnv(xapp); err != nil {
		return err
	}

	askSqlLogLevel(&config)

	// 4. Ask for the ports, it's important.
	po, _ := strconv.Atoi(askPortName("Http port which fireback will be lifted:", fmt.Sprintf("%v", config.Port)))
	config.Port = int64(po)

	// microserivce has lighter migration

	envRunMigration(xapp)

	executeSeeders(xapp)

	askSSL(&config)

	config.Save(".env")

	for _, module := range xapp.Modules {
		if module.OnEnvInit != nil {
			module.OnEnvInit(c)
		}
	}

	return nil
}

func askEnvironmentName(originalName string) (string, error) {
	validate := func(input string) error {
		re := regexp.MustCompile(`^[a-z0-9-]*$`)

		if strings.Trim(input, " ") == "" || input == "" || !re.MatchString(input) {
			return errors.New("environment name can only be lowercase and dash, and not empty")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "Environment name (used for system service files and launchctl, and switching env)",
		Validate: validate,
		Default:  originalName,
	}

	variable, err := promptVariable.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return variable, nil
}

func askSQLiteDatabaseLocation(projectName string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the database path on file system, eg. /tmp/database1.db")
		}
		return nil
	}

	workingDirectory, err := filepath.Abs(".")
	if err != nil {
		log.Println(err)
	}

	promptVariable := promptui.Prompt{
		Label:    "Database file location (.db)",
		Validate: validate,
		Default:  filepath.Join(workingDirectory, projectName+"-database.db"),
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}

func askMysqlDsn() (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: username:password@protocol(address)/dbname?param=value)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: username:password@protocol(address)/dbname?param=value)",
		Validate: validate,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}

func askPostgresDsn() (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)",
		Validate: validate,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}
