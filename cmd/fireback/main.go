package main

import (
	"os"

	"github.com/torabian/emi/lib/gorunner"
	"github.com/urfave/cli/v3"

	"github.com/torabian/fireback/modules/abac"
	"github.com/torabian/fireback/modules/abac/interfacetools"
	"github.com/torabian/fireback/modules/backup"
	"github.com/torabian/fireback/modules/eventbus"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/reactivesearch"

	// clitools registers every terminal/CLI-interactive fireback feature
	// (promptui/bubbletea prompts, os/exec service management, asynq
	// workers, graceful HTTP shutdown) via init(). It's tagged !wasm and
	// deliberately not imported by cmd/fireback-wasm.
	_ "github.com/torabian/fireback/modules/fireback/clitools"
	"github.com/torabian/fireback/modules/fireback/gintools"
	FBManage "github.com/torabian/fireback/modules/interfaces/fireback-manage"
	FbSelfService "github.com/torabian/fireback/modules/interfaces/selfservice"
	project_generator "github.com/torabian/fireback/modules/project-generator"
	"github.com/torabian/fireback/modules/storage"
)

var PRODUCT_NAMESPACENAME = "fireback"
var PRODUCT_DESCRIPTION = "Fireback core microservice - v" + fireback.FIREBACK_VERSION
var PRODUCT_LANGUAGES = []string{"fa", "en"}

// //go:embed all:ui
// var ui embed.FS

func main() {

	// wal-g is embedded directly in this binary (see modules/backup/Exec.go)
	// rather than shelling out to a separate installed executable. This
	// must run before any normal CLI parsing below, since it re-execs into
	// wal-g's own command tree when invoked with its hidden marker arg.
	backup.MaybeRunEmbeddedWalg()

	emiCommand := &cli.Command{
		Name:     "emi",
		Usage:    "Emi compiler - Backend-for-Frontend with automatic SDK generation.",
		Commands: gorunner.BuildCommands(),
	}

	modules := []*fireback.ModuleProvider{
		fireback.FirebackModuleSetup(nil),
		{
			CliHandlers: []*cli.Command{
				project_generator.NewProjectCli(),
				emiCommand,
			},
		},
		{
			CliHandlers: []*cli.Command{
				{
					Name:     "backup",
					Usage:    "Take and restore point-in-time Postgres backups via wal-g",
					Commands: backup.Commands(),
				},
			},
		},
		storage.StorageModuleSetup(nil),
		eventbus.ModuleSetup(nil),
		reactivesearch.ModuleSetup(&reactivesearch.ReactiveSearchModuleConfig{
			SearchProviders: []reactivesearch.SearchProviderFn{
				abac.QueryMenusReact,
				abac.QueryRolesReact,
			},
		}),
	}

	// For fireback we have abac module added.
	modules = append(modules, abac.AbacCompleteModules()...)

	var xapp = &fireback.FirebackApp{
		Title:              PRODUCT_DESCRIPTION,
		SupportedLanguages: PRODUCT_LANGUAGES,
		SeedersSync: func() {
			abac.PassportMethodSyncSeeders()
			interfacetools.AppMenuSyncSeeders()
		},
		PublicFolders: []gintools.PublicFolderInfo{
			{Fs: &FBManage.FirebackManageTmpl, Folder: ".", Prefix: "/manage"},
			{Fs: &FbSelfService.FbSelfService, Folder: ".", Prefix: "/selfservice"},
		},
		Modules: modules,
	}

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// If anything needs to be done after database initialized
		// fireback.RegionalContentSyncSeeders()
		// fireback.AppMenuSyncSeeders()
	})
}
