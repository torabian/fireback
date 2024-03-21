package main

import (
	"embed"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/cms"
	"github.com/torabian/fireback/modules/commonprofile"
	"github.com/torabian/fireback/modules/currency"
	"github.com/torabian/fireback/modules/drive"
	"github.com/torabian/fireback/modules/geo"
	"github.com/torabian/fireback/modules/keyboardActions"
	"github.com/torabian/fireback/modules/licenses"
	"github.com/torabian/fireback/modules/shop"
	"github.com/torabian/fireback/modules/widget"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/torabian/fireback/modules/worldtimezone"
	"github.com/urfave/cli"
)

//go:embed all:ui
var ui embed.FS

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

var xapp = &workspaces.XWebServer{
	Title: PRODUCT_DESCRIPTION,

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
		shop.QueryProductSubmissionsReact,
	},
	RunTus: func() {
		drive.LiftTusServer()
	},
	RunSocket: func(e *gin.Engine) {
		workspaces.HandleSocket(e)
	},
	RunSearch: workspaces.InjectReactiveSearch,
	PublicFolders: []workspaces.PublicFolderInfo{
		{Fs: &ui, Folder: "ui"},
	},
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.XWebServer) {

		e.GET("/", func(ctx *gin.Context) {
			query := workspaces.TemplateQueryDSL(ctx)
			query.Deep = true
			workspaces.RenderTemplateToGin(ctx, "ui/index.tpl", ui, gin.H{
				"products": QueryHelper(shop.ProductSubmissionActionQuery, query),
			})
		})
		e.GET("/products-inline", func(ctx *gin.Context) {
			query := workspaces.TemplateQueryDSL(ctx)
			query.Deep = true
			workspaces.RenderTemplateToGin(ctx, "ui/partials/products-inline.tpl", ui, QueryHelper(shop.ProductSubmissionActionQuery, query))
		})
	},
	Modules: []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		geo.GeoModuleSetup(),
		keyboardActions.KeyboardActionsModuleSetup(),
		drive.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
		widget.WidgetModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		cms.CmsModuleSetup(),
		currency.CurrencyModuleSetup(),
		licenses.LicensesModuleSetup(),
		shop.ShopModuleSetup(),
		worldtimezone.LicensesModuleSetup(),
	},
}

func main() {

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
