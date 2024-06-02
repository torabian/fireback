package workspaces

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/exp/maps"
	"gorm.io/gorm"
)

type PublicFolderInfo struct {
	Fs     *embed.FS
	Folder string
}

type SearchProviderFn = func(query QueryDSL, chanStream chan *ReactiveSearchResultDto)

type XWebServer struct {
	Title              string
	SupportedLanguages []string
	Modules            []*ModuleProvider
	CliActions         func() []cli.Command
	RunTus             func()
	RunSocket          func(*gin.Engine)
	RunSearch          func(*gin.Engine, *XWebServer)
	SetupWebServerHook func(*gin.Engine, *XWebServer)
	SearchProviders    []SearchProviderFn
	SeedersSync        func()
	MockSync           func()
	PublicFolders      []PublicFolderInfo
}

func GetCliCommands(x *XWebServer) []cli.Command {
	commands := []cli.Command{}

	for _, module := range x.Modules {
		commands = append(commands, module.CliHandlers...)
		for _, bundle := range module.EntityBundles {
			commands = append(commands, bundle.CliCommands...)
		}
	}

	commands = append(commands, x.CliActions()...)

	return commands
}

func GetReportCommands(x *XWebServer) []cli.Command {
	commands := []cli.Command{}

	for _, item := range x.Modules {
		commands = append(commands, item.CliHandlers...)
	}

	commands = append(commands, x.CliActions()...)

	return commands
}

func ExecuteMockImport(x *XWebServer) {

	for _, item := range x.Modules {
		if item.MockHandler != nil {
			item.MockHandler()
		}

	}

	if x.SeedersSync != nil {
		x.SeedersSync()
	}

	if x.MockSync != nil {
		x.MockSync()
	}

}
func ExecuteSeederImport(x *XWebServer) {

	for _, item := range x.Modules {
		if item.SeederHandler != nil {

			item.SeederHandler()
		}

	}

	if x.SeedersSync != nil {
		x.SeedersSync()
	}
}

func GetAppReportsString(items []Report) ([]string, error) {

	result := []string{}
	for _, entity := range items {
		result = append(result, entity.UniqueId+" >>> "+entity.Title+" ("+entity.Description+")")
	}
	return result, nil
}

func ExecuteMockWriter(x *XWebServer) {

	for _, item := range x.Modules {
		if item.MockWriterHandler != nil {
			item.MockWriterHandler(x.SupportedLanguages)
		}

	}

}

func SetupHttpServer(x *XWebServer) *gin.Engine {

	r := gin.New()

	translations := map[string]map[string]string{}
	for _, item := range x.Modules {
		maps.Copy(translations, item.Translations)
	}
	maps.Copy(translations, BasicTranslations)

	if x.RunTus != nil {
		go x.RunTus()
	}

	if x.RunSocket != nil {
		go x.RunSocket(r)
	}
	if x.RunSearch != nil {
		go x.RunSearch(r, x)
	}
	if x.SetupWebServerHook != nil {
		x.SetupWebServerHook(r, x)
	}

	r.GET("/docs", func(c *gin.Context) {

		c.Header("content-type", "text/html")
		c.String(200, `<!doctype html>
		<html lang="en">
		  <head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<title>Elements in HTML</title>
			<!-- Embed elements Elements via Web Component -->
			<script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
			<link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css">
		  </head>
		  <body>
		
			<elements-api
			  apiDescriptionUrl="/openapi.yml"
			  router="hash"
			  layout="sidebar"
			/>
		
		  </body>
		</html>
		`)
	})

	r.GET("/openapi.yml", func(c *gin.Context) {

		data, _ := ConvertStructToOpenAPIYaml(x)
		c.Header("content-type", "application/json")
		c.String(200, data)
	})

	r.Use(GinPostTranslateErrorMessages(translations))
	r.Use(GinMiddleware())

	r.GET("/ping", func(c *gin.Context) {

		if BundledConfig != nil && BundledConfig.SelfHosted {
			c.JSON(200, gin.H{
				"data": GetAppConfig(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": gin.H{
					"pong": "yes",
				},
			})
		}

	})

	for _, item := range x.PublicFolders {
		EmbedFolderForGin(item.Fs, item.Folder, r)
	}

	// Enable the mvc app from here, if it's needed. Work on your static website on
	// website.go instead of here, and only uncomment line below
	// ServeMVCWebsite(r)

	for _, item := range x.Modules {
		for _, actions := range item.Actions {
			CastRoutes(actions, r)
		}

		for _, bundle := range item.EntityBundles {
			CastRoutes(bundle.Actions, r)
		}
	}

	return r

}

func SyncDatabase(x *XWebServer, db *gorm.DB) {

	for _, item := range x.Modules {
		if item.EntityProvider != nil {
			item.EntityProvider(db)
		}

		for _, bundle := range item.EntityBundles {
			if err := dbref.AutoMigrate(bundle.AutoMigrationEntities...); err != nil {
				fmt.Println("There is an error on migrating:", bundle)
				log.Fatalln(err.Error())
			}
		}
	}

}

func RunApp(xapp *XWebServer) {

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 xapp.Title,
		Flags:                cliGlobalFlags,
		Commands:             GetCliCommands(xapp),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
