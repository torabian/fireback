//go:build !wasm

package fireback

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/connmonitor"
	"github.com/torabian/fireback/modules/fireback/envm"
	"github.com/torabian/fireback/modules/fireback/gintools"
	statics "github.com/torabian/fireback/modules/fireback/static"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// AskProjectDatabase, AskSSL and AskSqlLogLevel are interactive terminal
// wizards. Their real implementation lives in modules/fireback/clitools
// (tagged !wasm) and registers itself here via init() - see CliActions.go
// for the same pattern applied to the simpler Ask* prompt helpers.
var AskProjectDatabase func(projectName string) (Database, error)
var AskSSL func(config *Config)
var AskSqlLogLevel func(cfg *Config)
var AskPortName func(label string, defaultPort string) string
var AskFolderName func(label string, defaultFolder string) string

// QueryStringCastCli and CliInteractiveSearchAndSelect back the interactive
// bubbletea-based CLI search/select UI. Real implementation lives in
// modules/fireback/clitools (tagged !wasm) and registers itself here via
// init().
var QueryStringCastCli func(searchFields []string, keyword string, page int) QueryDSL
var CliInteractiveSearchAndSelect func(title string, fn func(keyword string, page int) ([]string, *QueryResultMeta, *IError)) []string

// SystemServiceHandler and GetMacDaemon manage the OS-level system service
// (systemd/launchd) via os/exec, unavailable under wasm. Real implementation
// lives in modules/fireback/clitools (tagged !wasm) and registers itself
// here via init().
var SystemServiceHandler func(action string, c *cli.Command)
var GetMacDaemon func() string

// MeetsAccessLevel checks whether a QueryDSL's granted access (query.UserAccessPerWorkspace)
// satisfies query.ActionRequires - abac wires its own implementation here (see
// AbacModule.go's setup), since the actual access model lives there, not in fireback
// itself. Used both directly (e.g. WorkspaceCli.go's own permission checks) and by
// modules/eventbus, which defaults to this if a project doesn't override
// EventBusModuleConfig.MeetsAccessLevel.
var MeetsAccessLevel func(query QueryDSL, onlyRoot bool) (bool, []string)

// CLIInit is the interactive `fireback init` wizard. Its real implementation
// lives in modules/fireback/clitools (tagged !wasm) and registers itself
// here via init().
var CLIInit func(xapp *application.Application) *cli.Command

// SSLCliCommand builds the `ssl` command group (`ssl enable`, `ssl status`)
// backing config.SslProvider/CertFile/KeyFile/AcmeDomains/AcmeEmail/AcmeCacheDir
// - detects any certificate already configured, offers to keep or replace
// it, and supports Let's Encrypt (via golang.org/x/crypto/acme/autocert -
// see NewAcmeManager in SSL.go), a manual cert+key pair, or a persisted
// self-signed certificate. Its real implementation lives in
// modules/fireback/clitools/SSLCli.go (tagged !wasm) and registers itself
// here via init() - same pattern as CLIInit above.
var SSLCliCommand func() *cli.Command

// Fireback actions run in cli as well. In this file, we place tools and helpers for that.

var CommonQueryFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "verbose",
		Usage: "Verbose query, show fireback columns as well such as workspace, etc",
	},
	&cli.StringFlag{
		Name:     "x-select",
		Required: false,
		Usage:    `Select only specific fields to be queried and returned`,
	},
	&cli.BoolFlag{
		Name:  "minimal",
		Usage: "Make a minimal query, skips printing some of the fields",
	},
	&cli.IntFlag{
		Name:  "offset",
		Usage: "Add the start index",
		Value: 0,
	},
	&cli.StringFlag{
		Name:  "cursor",
		Usage: "Cursor value from the pagination",
	},
	&cli.IntFlag{
		Name:  "limit",
		Usage: "Items per page",
		Value: 0,
	},
	&cli.StringFlag{
		Name:  "x-accept",
		Usage: "Change the return results to such as as 'yaml'",
	},
	&cli.StringFlag{
		Name:  "sort",
		Usage: "Sorting strategy",
	},
	&cli.BoolFlag{
		Name:  "deep",
		Usage: "Should query the arrays, nested objects, relations of the entity",
	},
	&cli.StringFlag{
		Name:  "query",
		Usage: "Query DSL which filters out the results.",
		Value: "",
	},
	&cli.StringFlag{
		Name:  "wp",
		Usage: "withPreloads The sub or nested entities to be loaded with. Comma separated",
		Value: "",
	},
	&cli.StringFlag{
		Name:  "lang",
		Usage: "define the language in 2 char code, aka: en, de",
		Value: "en",
	},
}

func GetCommonRemoveQuery(el reflect.Value, fn ActionDeleteSignature) *cli.Command {

	return &cli.Command{

		Name:    "remove",
		Aliases: []string{"r", "del", "delete"},
		Usage:   "Deletes a given record using it's unique_id string hash field from database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "uid",
				Value:    "",
				Usage:    "String unique id of the record which will be deleted.",
				Required: true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			f := QueryDSL{
				UniqueId: c.String("uid"),
			}

			res, err := fn(DeleteRequest{
				Query: "unique_id = " + c.String("id"),
			}, f)

			if err != nil {
				fmt.Println("Delete operation encountered an error on database level: %w", err)
				return err
			}

			if res.Data.Item.RowsAffected == 0 {
				fmt.Println("Query executed successfully, but no Rows have been deleted.")
			} else {
				fmt.Printf("Deleted rows by delete operation: %d\r\n", res.Data.Item.RowsAffected)
			}

			return nil

		},
	}

}

type CliInteractiveFlag struct {
	Name        string
	StructField string
	// It is required on database level
	Required bool

	// Its recommended on the cli level to make it easier.
	Recommended bool
	Usage       string
	Type        string
}

// AskForSelect, AskBoolean, AskForInputOptional, AskForInput and AskForPassword
// are interactive terminal prompts. Their real implementation lives in the
// modules/fireback/terminal package (tagged !wasm) and registers itself here
// via init(). Building without importing that package (e.g. cmd/fireback-wasm)
// leaves these nil - callers only reach them from CLI flows that aren't part
// of a wasm build.
var AskForSelect func(label string, items []string) string
var AskBoolean func(label string) bool
var AskForInputOptional func(label string, defaultV string) (string, bool, error)
var AskForInput func(label string, defaultV string) string
var AskForPassword func(label string, defaultV string) string

// Use the actions bundle for ease and provide it to the ModuleProvider
// and it would gather all actions in the module level, it's to make it easier
// for intelisense
type CliActionsBundle = cli.Command

// Represents both http, and cli actions in one single object
type ModuleActionsBundle struct {

	// cli.Command which has Subcommands of all actions
	CliAction *cli.Command
}

var Cliversion cli.Command = cli.Command{

	Name:  "version",
	Usage: "Returns the version of the fireback: " + FIREBACK_VERSION,

	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println(FIREBACK_VERSION)
		fmt.Println("Written with love by Ali Torabi")
		return nil
	},
}

var CLIAboutCommand cli.Command = cli.Command{

	Name:  "about",
	Usage: "About Fireback, the author of software, support and contact :)",

	Action: func(ctx context.Context, c *cli.Command) error {

		fmt.Println("Written with passion by Ali Torabi, distributed under torabian.github.io, use this to build something which lasts for decades.")

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

var CLIServiceCommand cli.Command = cli.Command{

	Name:  "service",
	Usage: "Manages the system service on operating system",
	Commands: []*cli.Command{
		{

			Name:    "unload",
			Aliases: []string{"u"},
			Usage:   "Unloads the system service",
			Action: func(ctx context.Context, c *cli.Command) error {
				SystemServiceHandler("unload", c)

				return nil
			},
		},
		{

			Name:    "reload",
			Aliases: []string{"r"},
			Usage:   "Unloads the service, and basically loads it once again.",
			Action: func(ctx context.Context, c *cli.Command) error {
				SystemServiceHandler("reload", c)

				return nil
			},
		},
		{

			Name:    "mac-daemon",
			Aliases: []string{"mac"},
			Usage:   "Shows the mac daemon path",
			Action: func(ctx context.Context, c *cli.Command) error {
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
			Action: func(ctx context.Context, c *cli.Command) error {
				SystemServiceHandler("load", c)

				return nil
			},
		},
	},
}

var ConfigCommand cli.Command = cli.Command{

	Name:  "config",
	Usage: "Set of tools to configurate the product",
	Commands: append([]*cli.Command{
		{

			Name:  "db",
			Usage: "Configurates the database of the project",
			Action: func(ctx context.Context, c *cli.Command) error {

				databaseData, err := AskProjectDatabase(config.Name)
				if err != nil {
					log.Fatalln("Database could not be determined after all", err)
					return nil
				}

				// DbDsn is the only thing persisted for the connection - see
				// modules/fireback/dbdsn. databaseData's other fields
				// (Host/Port/Username/...) only existed to build that Dsn
				// interactively; AskProjectDatabase already folded them in.
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
		{

			Name:  "dbdsn",
			Usage: "Returns the database dsn which will be used",
			Action: func(ctx context.Context, c *cli.Command) error {
				fmt.Println(GetDatabaseDsn(config))

				return nil
			},
		},
		{

			Name:  "ssl",
			Usage: "Wizard to configurate the ssl on the server",
			Action: func(ctx context.Context, c *cli.Command) error {

				AskSSL(&config)

				config.Save(".env")

				return nil
			},
		},
		{

			Name:  "sqllog",
			Usage: "Change the sql log level",
			Action: func(ctx context.Context, c *cli.Command) error {

				AskSqlLogLevel(&config)

				config.Save(".env")

				return nil
			},
		},
	}, GetConfigCli()...),
}

func GetHttpCommand(engineFn func(cfg2 HttpServerInstanceConfig) *gin.Engine) *cli.Command {
	return &cli.Command{
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "port",
				Usage: "The port that the server will come up. Defaults to the user configuration file",
			},
			&cli.BoolFlag{
				Name:  "watch",
				Usage: "Monitor server stat using charts and interactive graphics",
			},
			&cli.StringFlag{
				Name:  "domains",
				Usage: "Mimics the runtime on specific domains locally, for example: test.com,test.pl. If ssl enabled, the certificate needs to include the domains.",
			},
			&cli.BoolFlag{
				Name:  "ssl",
				Usage: "Runs ssl server on 443 port",
			},
			&cli.BoolFlag{
				Name:  "slow",
				Usage: "Makes a delay on serving xattach files to mimic slow server, might slow down also API calls",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "routes",
				Usage: "Returns the count of the http requests in the application",
				Action: func(ctx context.Context, c *cli.Command) error {
					engine := engineFn(HttpServerInstanceConfig{})
					fmt.Println("Total routes:", len(engine.Routes()))
					for _, route := range engine.Routes() {
						fmt.Println(route.Method, route.Path)
					}
					return nil
				},
			},
		},
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Starts http(s) server only, has few configuration",
		Action: func(ctx context.Context, c *cli.Command) error {

			initLogger()
			if config.ShowDoctor {
				Doctor()
			}

			cfg2 := HttpServerInstanceConfig{
				Monitor:        c.Bool("watch"),
				Port:           c.Int64("port"),
				SSL:            c.Bool("ssl"),
				Slow:           c.Bool("slow"),
				VirtualDomains: strings.Split(c.String("domains"), ","),
			}

			engine := engineFn(cfg2)
			CreateHttpServer(engine, cfg2)

			return nil
		},
	}
}

var CLIDoctor cli.Command = cli.Command{

	Name:  "doctor",
	Usage: "Gives some information about the app, operating system, for remote debugging",

	Action: func(ctx context.Context, c *cli.Command) error {
		Doctor()
		return nil
	},
}

type DoctorHeadData struct {
	ConfigurationUrl    string `yaml:"configurationUrl"`
	FirebackVersion     string `yaml:"firebackVersion"`
	DatabaseVendor      string `yaml:"databaseVendor"`
	ComputedDataBaseDsn string `yaml:"computedDataBaseDsn"`
}
type GeneralDoctorData struct {
	General         DoctorHeadData        `yaml:"general"`
	EnvironmentUris *envm.EnvironmentUris `yaml:"environmentUris"`
	Config          Config                `yaml:"config"`
}

func (x *GeneralDoctorData) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
}
func Doctor() {

	uri, _ := envm.ResolveConfigurationUri()
	vendor, dsn := GetDatabaseDsn(config)
	data := GeneralDoctorData{
		General: DoctorHeadData{
			ConfigurationUrl:    uri,
			FirebackVersion:     FIREBACK_VERSION,
			DatabaseVendor:      vendor,
			ComputedDataBaseDsn: dsn,
		},
		EnvironmentUris: envm.GetEnvironmentUris(),
		Config:          config,
	}

	fmt.Println(Bold + "General information:" + Reset)
	fmt.Println(FormatYamlKeys(data.Yaml()))

	fmt.Println()
	fmt.Println(Bold + "Environment urls:" + Reset)
	fmt.Println(FormatYamlKeys(envm.GetEnvironmentUris().Yaml()))

	fmt.Println()
	fmt.Println(Bold + "Configuration:" + Reset)
	fmt.Println(FormatYamlKeys(config.Yaml()))
}

func GetCommonWebServerCliActions(xapp *application.Application) []*cli.Command {

	// A local copy rather than mutating the package-level ConfigCommand var directly -
	// "list" needs xapp (to reach every module's own ConfigProvider), which isn't
	// available at ConfigCommand's own package-init time, unlike its per-field
	// get/set subcommands which only ever touch fireback's own package-level config.
	combinedConfigCommand := ConfigCommand
	combinedConfigCommand.Commands = append(
		append([]*cli.Command{}, ConfigCommand.Commands...),
		GetCombinedConfigListCli(xapp),
	)

	return []*cli.Command{
		CLIInit(xapp),
		envm.EnvManagement(xapp),
		GetApplicationTasks(xapp),
		&CLIDoctor,
		&CLIServiceCommand,
		SSLCliCommand(),
		&combinedConfigCommand,
		GetHttpCommand(func(cfg HttpServerInstanceConfig) *gin.Engine {
			return SetupHttpServer(xapp, cfg)
		}),
		GetSeeder(xapp),

		// Keep these in the last
		&CLIAboutCommand,
		&Cliversion,
	}
}

func GetCliCommands(x *application.Application) []*cli.Command {
	var commands []*cli.Command

	// Helper function to recursively collect CLI commands
	var collectCommands func(modules []*application.ModuleProvider) []*cli.Command
	collectCommands = func(modules []*application.ModuleProvider) []*cli.Command {
		var collected []*cli.Command
		for _, module := range modules {

			// Add CLI handlers from the module
			collected = append(collected, module.CliHandlers...)

			if module.AppendCli != nil {
				collected = append(collected, module.AppendCli(x)...)
			}

			// Add commands from entity bundles
			for _, bundle := range module.EntityBundles {
				collected = append(collected, bundle.CliCommands...)
			}

			// Recursively collect from children

		}
		return collected
	}

	// Collect commands from the root modules
	commands = collectCommands(x.Modules)

	// Add any CLI actions from x itself
	commands = append(commands, x.CliActions()...)

	return commands
}

func GetReportCommands(x *application.Application) []*cli.Command {
	commands := []*cli.Command{}

	for _, item := range x.Modules {
		commands = append(commands, item.CliHandlers...)
	}

	commands = append(commands, x.CliActions()...)

	return commands
}

type HttpServerInstanceConfig struct {

	// Shows some charts and keeps track of active connections
	Monitor bool

	// Override the port
	Port int64

	SSL bool

	Slow bool

	VirtualDomains []string
}

func SetupHttpServer(x *application.Application, cfg HttpServerInstanceConfig) *gin.Engine {

	if config.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		if config.GinMode == "release" {
			gin.SetMode(gin.ReleaseMode)
		}

		if config.GinMode == "test" {
			gin.SetMode(gin.TestMode)
		}

		if config.GinMode == "gin" {
			gin.SetMode(gin.EnvGinMode)
		}

		if config.GinMode == "debug" {
			gin.SetMode(gin.DebugMode)
		}
	}

	r := gin.New()

	r.Use(gintools.GinAllowEverything())

	// Gzip compression is handled entirely inside gintools.EmbedFoldersForGin,
	// not here. It would be tempting to add a single blanket "gzip everything
	// under an embedded folder's Prefix" middleware at this level, but that's
	// wrong the moment any folder is mounted at Prefix "/" (the common case
	// for a project's default SPA, e.g. this project's own "public" folder in
	// cmd/nima-server/main.go): every request path is a string-prefix match
	// for "/", so a prefix-based check would compress every route in the app
	// - including routes mounted outside PublicFolders entirely, like the
	// storage module's tus upload endpoints and its manual Range/streaming
	// download handler, whose protocols depend on exact byte counts and
	// streaming semantics (tus's Content-Length/Upload-Offset contract, HTTP
	// Range/206 responses) that a gzip-wrapped ResponseWriter can violate.
	// gintools instead only compresses a response once a specific
	// PublicFolderInfo has already confirmed *it* is about to write it (a
	// real embedded file existed, or this is its own SPA/index fallback) -
	// see mountEmbedFolder's doc comment.
	r.Use(connmonitor.TrackConnectionsMiddleware())

	if config.WithTaskServer {
		go taskServerLifter(x)
	}

	if x.SetupWebServerHook != nil {
		x.SetupWebServerHook(r, x)
	}

	r.GET("/stoplight.js", func(c *gin.Context) {
		file, err := statics.StaticFs.ReadFile("stoplight.js")
		if err != nil {
			c.String(http.StatusInternalServerError, "File not found")
			return
		}
		c.Data(http.StatusOK, "application/javascript", file)
	})

	r.GET("/stoplight.css", func(c *gin.Context) {
		file, err := statics.StaticFs.ReadFile("stoplight.css")
		if err != nil {
			c.String(http.StatusInternalServerError, "File not found")
			return
		}
		c.Data(http.StatusOK, "text/css", file)
	})

	r.GET("/ping", func(c *gin.Context) {
		if config.Production {
			c.JSON(200, gin.H{
				"data": gin.H{
					"pong": "yes",
				},
			})
		} else {
			// Used to also report eventbus.GetEventBusInstance()'s connected socket
			// users/snapshot here directly - eventbus is now an opt-in module (see
			// modules/eventbus), so this core diagnostic endpoint no longer assumes
			// it's registered. A project that registers eventbus can extend its own
			// /ping (or add another route) if it wants that detail back.
			c.JSON(200, gin.H{
				"data": gin.H{
					"pong":       "yes",
					"instanceId": SERVER_INSTANCE,
				},
			})
		}
	})

	if len(x.PublicFolders) > 0 {
		gintools.EmbedFoldersForGin(x.PublicFolders, r)
	}

	// Enable the mvc app from here, if it's needed. Work on your static website on
	// website.go instead of here, and only uncomment line below
	// ServeMVCWebsite(r)

	group := r.Group(x.ApiPrefix)
	FirebackAppToGin(x, group, "")

	return r
}

func FirebackAppToGin(x *application.Application, g *gin.RouterGroup, prefix string) {

	for _, item := range x.Modules {

		for _, hook := range item.GinWebServerInitHooks {
			if err := hook(g, x); err != nil {
				LOG.Error("Error %w", zap.Error(err))
				LOG.Fatal("Gin server failed to run a hook on module", zap.String("module", item.Name))
			}
		}

	}
}
