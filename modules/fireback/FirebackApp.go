package fireback

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/connmonitor"
	"github.com/torabian/fireback/modules/fireback/gintools"
	statics "github.com/torabian/fireback/modules/fireback/static"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
)

var SERVER_INSTANCE string = UUID_Long()

/*

This function would read a web server modules, and recursively convert them into a
cli application

*/

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

func ExecuteMockImport(x *application.Application) {

	for _, item := range x.Modules {
		if item.MockHandler != nil {
			item.MockHandler()
		}
	}

	for _, item := range x.Modules {
		for _, entity := range item.EntityBundles {
			if entity.MockProvider != nil {
				entity.MockProvider()
			}
		}
	}

	if x.SeedersSync != nil {
		x.SeedersSync()
	}

	if x.MockSync != nil {
		x.MockSync()
	}

}

func hasSuffix(path string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
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

	r.Use(GinMiddleware())

	// Range requests (used for resumable downloads and media seeking, e.g. an
	// HTML5 <video> tag) return a slice of the underlying resource picked
	// after the byte length is known. Gzip-encoding just that slice produces
	// a self-contained stream whose Content-Range still describes offsets in
	// the *original* file, a combination browsers' media/range handling
	// doesn't tolerate (Chrome fails such responses with
	// ERR_CONTENT_DECODING_FAILED). So skip compression whenever a Range
	// header is present, the same way nginx/Apache do.
	gzipMiddleware := gzip.Gzip(gzip.DefaultCompression)
	r.Use(func(c *gin.Context) {
		if c.GetHeader("Range") != "" {
			c.Next()
			return
		}
		gzipMiddleware(c)
	})

	r.Use(func(c *gin.Context) {
		cacheableSuffixes := []string{
			".js",
			".css",
			".svg",
			".png",
			".jpg",
			".woff",
			".woff2",
			".ttf",
		}

		if c.Request.Method == http.MethodGet && hasSuffix(c.Request.URL.Path, cacheableSuffixes) {
			c.Header("Cache-Control", "public, max-age=604800") // 1 year cache
		}

		c.Next()
	})

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

type GooseZapLogger struct {
	Logger *zap.Logger
}

func (l GooseZapLogger) Printf(format string, v ...interface{}) {
	l.Logger.Sugar().Infof(format, v...)
}

func (l GooseZapLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Sugar().Fatalf(format, v...)
}

// Logs all files under a given directory in an embedded FS
func LogAllFiles(fsys embed.FS, dir string) error {
	subFS, err := fs.Sub(fsys, dir)
	if err != nil {
		return err
	}

	return fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println("File:", path)
		}
		return nil
	})
}

func RunApp(xapp *application.Application) {

	app := &cli.Command{
		EnableShellCompletion: true,
		Name:                  xapp.Title,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "al",
				Usage: "Set's the language of the query, equal to accept-language header in http requests",
				Value: "en-us",
			},
		},
		Commands: GetCliCommands(xapp),
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
