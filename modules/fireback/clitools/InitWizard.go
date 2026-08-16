//go:build !wasm

package clitools

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

// config is a pointer to the core package's package-level config instance
// (via fireback.GetConfigRef()), so mutating config.Field here mutates the
// same storage the core package reads from - and, just as importantly,
// reflects whatever LoadFirebackAppConfiguration loaded from .env before
// this package's init() ran. A plain fireback.GetConfig() (a value copy,
// taken at package-variable-init time - i.e. before .env is ever read) was
// the actual bug behind "defaults from .env aren't picked up": it captured
// nothing but the struct's hardcoded zero values, permanently.
var config = fireback.GetConfigRef()

func init() {
	fireback.CLIInit = cliInit
	fireback.AskProjectDatabase = askProjectDatabase
	fireback.AskSSL = askSSL
	fireback.AskSqlLogLevel = askSqlLogLevel
	fireback.AskPortName = askPortName
	fireback.AskFolderName = askFolderName
}

var tryToSolve = "Let me retry to configurate the database parameters"
var forceContinue = "Use the configuration without connection test"

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

func cliInit(xapp *application.Application) *cli.Command {

	return &cli.Command{
		Name:  "init",
		Usage: "Creates a environment for project, by configurating database connection, http port, etc.",
		Flags: fireback.GetConfigCliFlags(),
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NumFlags() > 0 {
				fireback.CastConfigFromCli(config, c)

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

func dataBaseConfigEnv(xapp *application.Application) error {

	// 2. Determine the database type, test the connection, create tables
	for {
		databaseData, err := askProjectDatabase(config.Name)
		if err != nil {
			log.Fatalln("cannot determine the database config", err)
			return nil
		}

		// 3. Check if the database could be connected, if not show error and move on.
		// DbDsn is the only thing persisted for the connection (see
		// modules/fireback/dbdsn) - askProjectDatabase already folded
		// whatever host/port/username/... it gathered into that Dsn.
		config.DbVendor = databaseData.Vendor
		config.DbDsn = databaseData.Dsn

		db, err := fireback.DirectConnectToDb(*config)
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

func executeSeeders(xapp *application.Application) error {
	if r := fireback.AskForSelect("Do you want to add the seed data, menu items, etc?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		fireback.ExecuteSeederImport(xapp)
	}

	return nil
}

func envRunMigration(xapp *application.Application) error {
	if r := fireback.AskForSelect("Do you want to run migration, adding tables or columns to database (both automigrate, and manual)?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		fireback.ApplyMigration(xapp, 0)
	}

	return nil
}

func initEnvironment(xapp *application.Application, envFileName string, c *cli.Command) error {

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

	askSqlLogLevel(config)

	// 4. Ask for the ports, it's important.
	po, _ := strconv.Atoi(askPortName("Http port which fireback will be lifted:", fmt.Sprintf("%v", config.Port)))
	config.Port = int64(po)

	// microserivce has lighter migration

	envRunMigration(xapp)

	executeSeeders(xapp)

	askSSL(config)

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
