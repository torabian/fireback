package workspaces

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/manifoldco/promptui"
	systemconfigs "github.com/torabian/fireback/modules/workspaces/systemconfigs"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var DATABASE_TYPE_MYSQL string = "mysql"
var DATABASE_TYPE_SQLITE string = "sqlite"
var DATABASE_TYPE_SQLITE_MEMORY string = "sqlite (:memory:)"
var DATABASE_TYPE_POSTGRES string = "postgres"

var USE_DSN_OPTION = "I have dsn query string for connection"
var USE_MANUAL_OPTION = "I enter port, host, username of database manually"

var TRY_TO_SOLVE = "Let me retry to configurate the database parameters"
var FORCE_CONTINUE = "Use the configuration without connection test"

func askRetry() bool {
	promptVariable := promptui.Select{
		Label: "Database connection failed, do you want retry again?",
		Items: []string{TRY_TO_SOLVE,
			FORCE_CONTINUE},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	if actionType == TRY_TO_SOLVE {
		return true
	}

	return false
}

func askMysqlDetails(db *Database) (*Database, error) {

	promptVariable := promptui.Select{
		Label: "Do you have dsn string or port, host , username?",
		Items: []string{USE_DSN_OPTION, USE_MANUAL_OPTION},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == USE_DSN_OPTION {
		value, err := askMysqlDsn()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	if actionType == USE_MANUAL_OPTION {

		db.Host = askHostName()
		db.Port = askHostPort("3306")
		db.Database = askDatabaseName()
		db.Username = askHostUsername("root")
		db.Password = askHostPassword()
	}

	return db, nil
}

func askPostgresDetails(db *Database) (*Database, error) {

	promptVariable := promptui.Select{
		Label: "Do you have dsn string or port, host , username?",
		Items: []string{USE_DSN_OPTION, USE_MANUAL_OPTION},
	}

	_, actionType, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return db, err
	}

	if actionType == USE_DSN_OPTION {
		value, err := askPostgresDsn()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return db, err
		}

		db.Dsn = value

		return db, nil
	}

	if actionType == USE_MANUAL_OPTION {

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

func askProjectDatabase(projectName string) (Database, error) {
	db := Database{}

	promptVariable := promptui.Select{
		Label: "Database type",
		Items: []string{
			DATABASE_TYPE_SQLITE_MEMORY,
			DATABASE_TYPE_SQLITE,
			DATABASE_TYPE_MYSQL,
			// Postgres is not well tested yet, we are not adding production ready
			// features anymore in fireback at all.
			// DATABASE_TYPE_POSTGRES,
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
	} else if db.Vendor == DATABASE_TYPE_SQLITE_MEMORY {
		db.Database = ":memory:"
		db.Vendor = "sqlite"
	} else if db.Vendor == DATABASE_TYPE_MYSQL {

		askMysqlDetails(&db)
	} else if db.Vendor == DATABASE_TYPE_POSTGRES {
		askPostgresDetails(&db)
	}

	return db, nil
}

type InteractiveInitData struct {
	ProjectName string
	Database    Database
}

type LicenseData struct {
	License string `yaml:"license"`
}

var CLIActivate cli.Command = cli.Command{

	Name:  "activate-license",
	Usage: "Activates the license when you get it from activator or pixelplux website.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "key",
			Value:    "",
			Required: true,
			Usage:    "The long activation key you got from web activator",
		},
	},
	Action: func(c *cli.Context) error {

		data := &LicenseData{
			License: c.String("key"),
		}

		body, err := yaml.Marshal(data)
		if err != nil {
			log.Fatal(err)
			return err
		}

		os.WriteFile("fireback-license.yml", body, 0644)

		return nil
	},
}

func (x *AppConfig) Save() error {
	file, err := os.OpenFile(GetEnvironmentUris().CurrentDirectoryConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Cannot create project file yml file here.: %v You might need to check the permissions to write in this directory.", err)

		return nil
	}

	defer file.Close()
	enc := yaml.NewEncoder(file)
	err = enc.Encode(x)
	if err != nil {
		log.Fatalf("Error encoding configuration: %v", err)
		return nil
	}

	return nil
}

func InitProject(xapp *XWebServer) error {
	if _, err := os.Stat(os.Getenv("CONFIG_PATH")); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("There is a ", os.Getenv("CONFIG_PATH"), " in this directory. Only one fireback app per directory is allowed.")
		return nil
	}

	datum := ""
	var err error

	// 1. Determine the project name
	datum, err = askProjectName()
	if err != nil {
		log.Fatalln("cannot determine the project name", err)
		return nil
	}
	config.Name = datum
	config.DebianIdentifier = datum
	config.MacIdentifier = datum
	config.WindowsIdentifier = datum

	// 2. Determine the database type, test the connection, create tables
	for {
		databaseData, err := askProjectDatabase(config.Name)
		if err != nil {
			log.Fatalln("canno determine the project name", err)
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

		db, err := DirectConnectToDb(config)
		if err == nil && db.Exec("select 1").Error == nil {
			fmt.Println("✔ connection is successful")
			break
		}

		fmt.Println(err)

		if !askRetry() {
			break
		}
	}

	// 4. Ask for the ports, it's important.
	po, _ := strconv.Atoi(askPortName("Http port which fireback will be lifted:", fmt.Sprintf("%v", config.Port)))
	config.Port = int64(po)

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Remember to use always absolute path for database, and storage.
	config.Storage = askFolderName("Storage folder (all upload files from users will go here)", filepath.Join(workingDirectory, "storage"))
	config.TusPort = askPortName("TUS File upload port", "4506")

	// 5. Ask for the storage folder as well

	config.Save(".env")

	fmt.Println("Creating storage directory, where all files will be uploaded to:", config.Storage)
	if err := os.Mkdir(config.Storage, os.ModePerm); err != nil {
		fmt.Println("Folder for storage exists or inaccessible.")
	}

	fmt.Println("Your new project has been created successfully.")
	fmt.Println("\nIf you want to start the project with HTTP Server, run:")
	fmt.Println("$ " + GetExePath() + " start \n ")
	fmt.Println("You can also run the project on daemon, as a system server to presist the connection: (good for production)")
	fmt.Println("$ " + GetExePath() + " service load \n ")

	if r := AskForSelect("Do you want to run migration, adding tables or columns to database?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		ApplyMigration(xapp, 2)
		RepairTheWorkspaces()
	} else {
		return nil
	}

	if r := AskForSelect("Do you want to add the seed data, menu items, etc?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		ExecuteSeederImport(xapp)
		AppMenuSyncSeeders()
	} else {
		return nil
	}

	if r := AskForSelect("Do you want to create a root admin for project?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		if err := InteractiveUserAdmin(QueryDSL{
			WorkspaceHas: []string{"root/*"},
			WorkspaceId:  "system",
			ItemsPerPage: 10,
		}); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func CLIInit(xapp *XWebServer) cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "Initialize the project, adds yaml configuration in the folder.",
		Action: func(c *cli.Context) error { InitProject(xapp); return nil },
	}
}

var ConfigCommand cli.Command = cli.Command{

	Name:  "config",
	Usage: "Set of tools to configurate the product",
	Subcommands: []cli.Command{
		{

			Name:  "db",
			Usage: "Configurates the database of the project",
			Action: func(c *cli.Context) error {

				databaseData, err := askProjectDatabase(config.Name)
				if err != nil {
					log.Fatalln("Database could not be determined after all", err)
					return nil
				}

				config.DbUsername = databaseData.Username
				p, _ := strconv.Atoi(databaseData.Port)
				config.DbPort = int64(p)
				config.DbHost = databaseData.Host
				config.DbPassword = databaseData.Password
				config.DbName = databaseData.Database
				config.DbVendor = databaseData.Vendor
				config.DbDsn = databaseData.Dsn

				if _, err := DirectConnectToDb(config); err != nil {
					fmt.Println("Connection to database failed:", err)
					return nil
				} else {
					fmt.Println("Database connected")
				}

				config.Save(".env")

				return nil
			},
		},
	},
}

var CLIServiceCommand cli.Command = cli.Command{

	Name:  "service",
	Usage: "Manages the system service on operating system",
	Subcommands: []cli.Command{
		{

			Name:    "unload",
			Aliases: []string{"u"},
			Usage:   "Unloads the system service",
			Action: func(c *cli.Context) error {
				SystemServiceHandler("unload", c)

				return nil
			},
		},
		{

			Name:    "reload",
			Aliases: []string{"r"},
			Usage:   "Unloads the service, and basically loads it once again.",
			Action: func(c *cli.Context) error {
				SystemServiceHandler("reload", c)

				return nil
			},
		},
		{

			Name:    "mac-daemon",
			Aliases: []string{"mac"},
			Usage:   "Shows the mac daemon path",
			Action: func(c *cli.Context) error {
				fmt.Println("Daemon path:", GetMacDaemon())
				return nil
			},
		},
		{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "stderr",
					Value: "",
					Usage: "Where to log the error messages, such as /tmp/fireback-err.log",
				},
				&cli.StringFlag{
					Name:  "stdout",
					Value: "",
					Usage: "Where to log the standard output messages, such as /tmp/fireback.log",
				},
			},
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "Starts the system service",
			Action: func(c *cli.Context) error {
				SystemServiceHandler("load", c)

				return nil
			},
		},
	},
}

type NginxConf struct {
	Location string
	Host     string
	Port     string
}

func NginxCommand() cli.Command {

	return cli.Command{

		Name:  "nginx",
		Usage: "Prints out the config for nginx",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "location",
				Value: "",
				Usage: "Prefix location of fireback, given http://x.com/auth => enter --location 'auth' without slashes",
			},
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "Host which you have installed the fireback, such as localhost",
			},
			&cli.StringFlag{
				Name:  "port",
				Value: "",
				Usage: "Port which you are running your http fireback, defaults to your configuration",
			},
		},
		Action: func(c *cli.Context) error {
			td := NginxConf{
				Host:     config.Host,
				Port:     fmt.Sprintf("%v", config.Port),
				Location: c.String("location"),
			}

			if strings.HasPrefix(td.Location, "/") || strings.HasSuffix(td.Location, "/") {
				log.Fatalln("Location cannot start with / or end with /")
				return nil
			}

			td.Location = "/" + td.Location + "/"

			if td.Location == "//" {
				td.Location = "/"
			}

			t, err := template.ParseFS(systemconfigs.SystemConfigs, "nginx.conf.tpl")
			if err != nil {
				panic(err)
			}
			var tpl bytes.Buffer
			err = t.Execute(&tpl, td)
			if err != nil {
				panic(err)
			}

			result := tpl.String()
			fmt.Println(result)

			return nil
		},
	}
}

func GetHttpCommand(engineFn func() *gin.Engine) cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "port",
				Value: 4500,
				Usage: "The port that the server will come up. Defaults to the user configuration file",
			},
		},
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Starts http server only",
		Action: func(c *cli.Context) error {
			Doctor()
			engine := engineFn()
			CreateHttpServer(engine)

			return nil
		},
	}
}

func CreateBackup(actions []TableMetaData, path string) {
	dictionary := map[string]interface{}{}

	q := QueryDSL{ItemsPerPage: 9999999}
	for _, act := range actions {
		v, _ := act.ExportStream(q)
		data := []interface{}{}
	L:
		for {
			select {

			case row, more := <-v:

				data = append(data, row...)
				if !more {
					dictionary[act.ExportKey] = data
					break L
				}

			}
		}

	}

	j, err := yaml.Marshal(dictionary)
	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile(path, j, 0644)

}

func CreateBackupToStream(actions []TableMetaData) (chan []byte, *IError) {

	ret := make(chan []byte)
	go func() {

		dictionary := map[string]interface{}{}

		q := QueryDSL{ItemsPerPage: 9999999}
		for _, act := range actions {
			v, _ := act.ExportStream(q)
			data := []interface{}{}
		L:
			for {
				select {

				case row, more := <-v:

					data = append(data, row...)
					if !more {
						dictionary[act.ExportKey] = data
						break L
					}

				}
			}

		}

		j, err := yaml.Marshal(dictionary)
		if err != nil {
			fmt.Println(err)
		}
		ret <- j

		close(ret)
	}()

	return ret, nil

}

func FindExportInfo(name string, actions []TableMetaData) *TableMetaData {
	for _, item := range actions {
		if item.ExportKey == name {
			return &item
		}
	}
	return nil
}

func ImportBackup(actions []TableMetaData, file string, f QueryDSL) *IError {

	var data map[string][]interface{}
	ReadYamlFile(file, &data)

	for k, v := range data {

		imp := FindExportInfo(k, actions)
		if imp == nil {
			continue
		}
		for _, record := range v {
			fmt.Println(k, record)

			err := imp.ImportQuery(record, f)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return nil

}

func GetCommonWebServerCliActions(xapp *XWebServer) cli.Commands {

	// I do not undestand, why vscode does not accept a single argument
	// to start `fireback lsp`. In vscode, we start it using `LSP=true fireback`
	// unfortunately this is what it is, and perhaps will remain.
	lsp := os.Getenv("LSP")
	if lsp == "yes" || lsp == "true" {
		BeginLspServer(nil)
		return []cli.Command{}
	}

	return cli.Commands{
		CLIInit(xapp),
		CLIAboutCommand,
		Cliversion,
		LSPSerever,
		CodeGenTools(xapp),
		GetApplicationTests(xapp),
		GetApplicationTasks(xapp),
		CLIDoctor,
		ManifestTools(),
		CLIServiceCommand,
		NewProjectCli(),
		ConfigCommand,
		GetMigrationCommand(xapp),
		GetHttpCommand(func() *gin.Engine {
			return SetupHttpServer(xapp)
		}),
		GetCliMockTools(xapp),
		GetSeeder(xapp),
		GetReportsTool(xapp),
	}
}
func GetCommonMicroserviceCliActions(xapp *XWebServer) cli.Commands {

	return cli.Commands{
		CLIInit(xapp),
		GetApplicationTasks(xapp),
		CLIDoctor,
		ManifestTools(),
		CLIServiceCommand,
		ConfigCommand,
		GetMigrationCommand(xapp),
		GetHttpCommand(func() *gin.Engine {
			return SetupHttpServer(xapp)
		}),
		GetCliMockTools(xapp),
		GetSeeder(xapp),
		GetReportsTool(xapp),
	}
}

func GetCliMockTools(xapp *XWebServer) cli.Command {
	return cli.Command{
		Name:  "mock",
		Usage: "Generates or export mocks based on all available content inside the database",
		Subcommands: cli.Commands{
			{

				Name:  "import",
				Usage: "Execute the mock services, and populates the entire backend with data and instructions",
				Action: func(c *cli.Context) error {

					fmt.Println("This function would create a virtual product, by first running mock data into database, and then run some actions as specified")
					ExecuteMockImport(xapp)
					return nil
				},
			},
			{
				Name:  "write",
				Usage: "Writes the instructions and module mock data into the sample json files. Clean system before, run mock-import, and then execute this to keep data safe",
				Action: func(c *cli.Context) error {

					fmt.Println("Writing all mocks into artifacts folder...")
					ExecuteMockWriter(xapp)
					return nil
				},
			},
		},
	}
}

var Cliversion cli.Command = cli.Command{

	Name:  "version",
	Usage: "Returns the version of the fireback: " + FIREBACK_VERSION,

	Action: func(c *cli.Context) error {
		fmt.Println(FIREBACK_VERSION)
		fmt.Println("Written with love by Ali Torabi")
		return nil
	},
}
var LSPSerever cli.Command = cli.Command{

	Name:  "lsp",
	Usage: "Runs the lsp language server for fireback and fireback projects on std",

	Action: func(c *cli.Context) error {
		BeginLspServer(nil)
		return nil
	},
}

var CLIAboutCommand cli.Command = cli.Command{

	Name:  "about",
	Usage: "About Fireback, the author of software, support and contact :)",

	Action: func(c *cli.Context) error {

		fmt.Println("Written by passion by Ali Torabi, distributed under PixelPlux Sp. z.o.o - reach me on 0048783538796 or https://github.com/torabian")

		fmt.Println(",.. .     /  .(  (                 .,/***/(%#(#####%&&%@@@@@@&&&&&&&&&&&&&&&&@@@")
		fmt.Println(",.. , (/ **,. , ,      .  ..  ,. ./(#((%###((((###(&&#@@@@@  ,&&&&&&&&&&&&&@@@@&")
		fmt.Println("*...   .   .                   ,*(#@&&#&@%&@&((###%#(#@@@@@%%%&&&&&&&&&&&&@@@&&&")
		fmt.Println("/*,,......,,.,,,,,,,,,,***#%%##%&&@@@@@&@@@@@@@&%%#(#&&@@@@&&&&&&&&&&&&@@@@&&&&&")
		fmt.Println("####(((((((((((((#/*,*//(#@@@&%&&&#%%%@@@@@@@@@@@&%#%%#&@&%%%%%%%%%%%%@@@&&&&&&@")
		fmt.Println("((((((((//(((((#(#*****/(#%%%##%&@&&@&%#&@@&@@@@&%&%%#%%#%%&@@@@@@@@@@@@&&%&@@@@")
		fmt.Println("((((((((((((#( *(#(((((##%&&#%#%%@@@&%@@&@@@@@%&&@@@#####%%##&@@@@@@@@&&&@@@@@@@")
		fmt.Println("#######%#(#(/*///(((#((&&%@@@@@&&&@@@@@&@@&@@@@@@@&&&%%%%&%&%%%%%%%&%%@@@@@@@@@@")
		fmt.Println("//,.,((((#/,&#/(((/((#%&&%#&&&&@@@@@&&&@@@@@@@&&&&&&%%&&&&%%%#%##%%%@@@@@@@@@&&&")
		fmt.Println("#((%%&#&%#%#(((/##/((##&@&&%&&&&@%@@@@@&@@@@@@@@&@&&&&&&&&&%%%%&%%@@@@@@@@&%%%&@")
		fmt.Println("#%##&&&&%&#&%#####(/(/(&&@@@@@@@&@@@@&@@@@@@@&@&&&&&&@@@&####%&&&%%(%&%%%%&@@@@@")
		fmt.Println("//(#%%##(##%%%/**(#///#%&@@@@@@@@@@@@@@@@@&&%%%%&&&&&&@%%&&%#(%%(/(((@@@@@@@@@@(")
		fmt.Println("####(//(**(/#%/, ,/**/(%%@@@@@@@@@@@@@@@%#((####%&&&(#&&&&&(///(/////%%%&&&&&&&/")
		fmt.Println("%#(#/(///((,(** ./(*/#/((##%%@@@&&&#(%@&%###%%%%%#(##%#&*##*,*##%%###//(%&&@@&%#")
		fmt.Println("(.*(./**(* ,/*,,,(/##((((######@@(###((#%%%%%%%%##(((((.     /%&@/       .*@&*..")
		fmt.Println("%%%##(@/****%/#*,/(((#######&&&(@&((((((#####%%%####((((((((((#&&(*.    #@@(#@%%")
		fmt.Println("//////(/*//*//#%##(#%&&%&&&&@@@@%##(((((((##((####((((#@@@@%##%%&(/#.,. ,.*,##(*")
		fmt.Println("				   ##(((((###((%%%%%(#/  ####%%%%##////(#(,.,/&%")
		fmt.Println("/#(#/*#(#*/(*((**((*,/((..((*  #&@@@%%#(%#%##(((##(#&@@%&%%#%%%%%##(///((((((###")
		fmt.Println("/( *//, #/(./.(*/../*/.,/@@@@@@@@@@@@#&%#(((((##((#%@%%@%%@&#&#/////**/(((((%%%%")
		fmt.Println("##(###########((/,.*@@@@@@@@@@@@@@@@@@#(%%###((##%&@@(%%&&&&&&&&@&,,,,,(****//((")
		fmt.Println("############%&%#///*,#@@@@@@@@@@@@@@@@#(#(((&###%@@%##.@&&%&&%@@@&&(///#/(#((##%")
		fmt.Println("((###%%%##(/&&&&(/(((/**,..,,%@@@@@@@@@(%/(/(((#@((#%%/&@%@@@&&@@&@&********,*/,")
		fmt.Println("/%#(/(((((###%%%/.  ......,,*(#((#%%&@@%#((//,/(%%%(%%,%&%&@%%&@%@&@****/(((((#&")
		fmt.Println("////(###(/(%,*#///,,,,**//#/...,/@@@@@@(/,..*(%#(#&%%#*%@&&@&%&%&@@&,***********")
		fmt.Println("/////(///(*/#((((///(###/,,.*(#/**@@@@@@((%#(((#%(%#%%(/@@@&#&@@&@@@,,,,,,,*****")
		fmt.Println("////////&%###%%%%##(((///(,...,,.,@@@@@@###(%###%(%#%&#/@@@@&&@@@@@@,,,,,,*,,***")
		fmt.Println("#&####%&&&#%%%##((((#%#/,,****../@@@@@@@&##%&%#&%###%&((@%&@@@@@@@@@/,,,,,,*,**,")
		fmt.Println("#&&@&%@&&&#(((((#####(/*/###**(@@@@@@@@@&&##%&%#%%#&%&(#@@@@@@@@@@@@(,,,,*,,*,**")
		fmt.Println("%&&@&@@&%%%%#(###((((/*/**/*%@@@@@@@@@@@@&&##%&#%%#&&@@@@&&@@@@@@@@@/..*,,,*,,*/")
		fmt.Println("@&&&&&&&&&&&#%%/**///******(@@@@@@@@@@@@@%%&%&@%%&%&@@@@@@@@&@@@@@@@@*,,,**,,*/*")
		fmt.Println("&&&&&&&&&&&&#&&%%#((//****#@@@@@@@@@@@@@@&%&%&@&&&&@@@@@@@&@@@@@@&@@#(**,*****/*")
		return nil
	},
}

var CLIMIDCommand cli.Command = cli.Command{

	Name:  "mid",
	Usage: "Gives you computer unique identifier, can be used to get a license for product.",

	Action: func(c *cli.Context) error {

		fmt.Println(UNIQUE_MACHINE_ID)
		return nil
	},
}

const (
	Reset     = "\033[0m"
	Green     = "\033[32m"
	Orange    = "\033[31m"
	Bold      = "\033[1m"
	GreenBold = "\033[32;1m"
)

func FormatYamlKeys(yamlStr string) string {
	// Regular expression to match YAML keys
	keyRegex := regexp.MustCompile(`(?m)^( *)([^:]+):`)

	// Replace keys with green bold text
	formattedYaml := keyRegex.ReplaceAllString(yamlStr, fmt.Sprintf("$1%s$2%s:", GreenBold, Reset))

	return formattedYaml
}
func Doctor() {

	fmt.Println("Fireback version: " + Orange + Bold + FIREBACK_VERSION + Reset)
	fmt.Println()
	uri, _ := ResolveConfigurationUri()
	fmt.Println(Bold + "Configuration will be read from:" + Reset)
	fmt.Println(uri)
	fmt.Println()

	vendor, dsn := GetDatabaseDsn(config)
	fmt.Println(Bold + "Database connection vender:" + Reset)
	fmt.Println(vendor)
	fmt.Println()

	fmt.Println(Bold + "Computed dsn for database connection:" + Reset)
	fmt.Println(dsn)

	fmt.Println()
	fmt.Println(Bold + "Environment urls:" + Reset)
	fmt.Println(FormatYamlKeys(GetEnvironmentUris().Yaml()))

	fmt.Println()
	fmt.Println(Bold + "Configuration:" + Reset)
	fmt.Println(FormatYamlKeys(config.Yaml()))
}

var CLIDoctor cli.Command = cli.Command{

	Name:  "doctor",
	Usage: "Gives some information about the app, operating system, for remote debugging",

	Action: func(c *cli.Context) error {
		Doctor()
		return nil
	},
}
