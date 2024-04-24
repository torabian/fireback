package main

import (
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/cms"
	"github.com/torabian/fireback/modules/commonprofile"
	"github.com/torabian/fireback/modules/currency"
	"github.com/torabian/fireback/modules/demo"
	"github.com/torabian/fireback/modules/geo"
	"github.com/torabian/fireback/modules/keyboardActions"
	"github.com/torabian/fireback/modules/licenses"
	"github.com/torabian/fireback/modules/shop"
	"github.com/torabian/fireback/modules/widget"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/torabian/fireback/modules/worldtimezone"
	zayshop "github.com/torabian/fireback/themes/zay-shop"
	"github.com/urfave/cli"
)

/////go:embed all:ui
// var ui embed.FS

var PRODUCT_NAMESPACENAME = "fireback"
var PRODUCT_DESCRIPTION = "Fireback core microservice"
var PRODUCT_LANGUAGES = []string{"fa", "en"}

type QueryableAction func(query workspaces.QueryDSL) ([]*shop.ProductSubmissionEntity, *workspaces.QueryResultMeta, error)

func QueryHelper(fn QueryableAction, query workspaces.QueryDSL) gin.H {
	products, qrm, err := shop.ProductSubmissionActionQuery(query)
	return gin.H{
		"items": products,
		"qrm":   qrm,
		"err":   err,
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var xapp = &workspaces.XWebServer{
	Title: PRODUCT_DESCRIPTION,

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
		shop.QueryProductSubmissionsReact,
	},
	RunTus: func() {
		workspaces.LiftTusServer()
	},
	RunSocket: func(e *gin.Engine) {
		workspaces.HandleSocket(e)
	},
	RunSearch: workspaces.InjectReactiveSearch,
	PublicFolders: []workspaces.PublicFolderInfo{
		// {Fs: &ui, Folder: "ui"},
		// {Fs: &ui.UI, Folder: "."},
		{Fs: &zayshop.UI, Folder: "."},
	},
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.XWebServer) {
		// ui.Bootstrap(e)
		zayshop.Bootstrap(e)

		// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// failOnError(err, "Failed to connect to RabbitMQ")
		// defer conn.Close()

		// ch, err := conn.Channel()
		// failOnError(err, "Failed to open a channel")
		// defer ch.Close()

		// q, err := ch.QueueDeclare(
		// 	"hello", // name
		// 	false,   // durable
		// 	false,   // delete when unused
		// 	false,   // exclusive
		// 	false,   // no-wait
		// 	nil,     // arguments
		// )
		// failOnError(err, "Failed to declare a queue")

		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()

		// body := "Hello World!"

		// for i := 0; i <= 100000; i++ {

		// 	err = ch.PublishWithContext(ctx,
		// 		"",     // exchange
		// 		q.Name, // routing key
		// 		false,  // mandatory
		// 		false,  // immediate
		// 		amqp.Publishing{
		// 			ContentType: "text/plain",
		// 			Body:        []byte(body),
		// 		})
		// 	// failOnError(err, "Failed to publish a message")
		// 	// log.Printf(" [x] Sent %s\n", body)
		// }

	},
	Modules: []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		geo.GeoModuleSetup(),
		keyboardActions.KeyboardActionsModuleSetup(),
		workspaces.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
		widget.WidgetModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		cms.CmsModuleSetup(),
		currency.CurrencyModuleSetup(),
		licenses.LicensesModuleSetup(),
		shop.ShopModuleSetup(),
		demo.DemoModuleSetup(),
		worldtimezone.LicensesModuleSetup(),
	},
}

func main() {
	numCPU := runtime.NumCPU()
	maxProcs := int(float64(numCPU) * 0.9)
	runtime.GOMAXPROCS(maxProcs)

	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	db, dbErr := workspaces.CreateDatabasePool()

	if db != nil && dbErr == nil {
		workspaces.SyncDatabase(xapp, db)
		workspaces.SyncPermissionsInDatabase(xapp, db)
	} else {
		log.Fatalln("Database error", dbErr)
	}

	xapp.CliActions = func() []cli.Command {
		return workspaces.GetCommonWebServerCliActions(xapp)
	}

	workspaces.RunApp(xapp)
}
