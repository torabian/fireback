package main

import (
	"os"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/lib/gorunner"
	"github.com/urfave/cli/v3"

	"github.com/torabian/fireback/modules/abac"
	"github.com/torabian/fireback/modules/abac/interfacetools"
	"github.com/torabian/fireback/modules/backup"
	"github.com/torabian/fireback/modules/eventbus"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/internalstats"
	"github.com/torabian/fireback/modules/reactivesearch"

	// clitools registers every terminal/CLI-interactive fireback feature
	// (promptui/bubbletea prompts, os/exec service management, asynq
	// workers, graceful HTTP shutdown) via init(). It's tagged !wasm and
	// deliberately not imported by cmd/fireback-wasm.
	"github.com/torabian/fireback/modules/fireback/application"
	_ "github.com/torabian/fireback/modules/fireback/clitools"
	"github.com/torabian/fireback/modules/fireback/gintools"
	FBManage "github.com/torabian/fireback/modules/interfaces/fireback-manage"
	FbSelfService "github.com/torabian/fireback/modules/interfaces/selfservice"
	"github.com/torabian/fireback/modules/storage"
)

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

	modules := []*application.ModuleProvider{
		fireback.FirebackModuleSetup(nil),
		{
			CliHandlers: []*cli.Command{
				emiCommand,
			},
		},
		// backup.ModuleSetup wires up backup's CLI commands (wal-g push/restore/...,
		// plus dump/restore-dump), its config: block (WALG_*/BACKUP_*, see
		// Backup.emi.yml) into fireback's combined "config list", and the HTTP
		// endpoints `backup dump --hash` registers a job against - instead of main.go
		// hand-assembling a bare CliHandlers-only ModuleProvider the way this used to.
		backup.ModuleSetup(nil),
		storage.StorageModuleSetup(nil),
		eventbus.ModuleSetup(nil),
		reactivesearch.ModuleSetup(&reactivesearch.ReactiveSearchModuleConfig{
			SearchProviders: []reactivesearch.SearchProviderFn{
				abac.QueryMenusReact,
				abac.QueryRolesReact,
			},
		}),
		// internalstats never imports abac - Authorize below is built entirely out of
		// fireback.ResolveActionContext/SecurityModel, the same generic contract every
		// module's own actions already use. It ends up enforcing abac's root-workspace
		// check anyway, transitively: fireback.AuthorizeRequest is nil until some auth
		// provider's module setup assigns it, and abac.AbacCompleteModules() (appended
		// to modules below) is what does that here (fireback.AuthorizeRequest =
		// abac.AuthorizeRequest - see modules/abac/AbacModule.go). Swap this whole
		// Authorize func for your own token check to drop the abac dependency for real.
		internalstats.ModuleSetup(&internalstats.InternalStatsModuleConfig{
			Authorize: func(req emigo.EmiRequestContexts) (fireback.QueryDSL, error) {
				query, err := fireback.ResolveActionContext(req, &fireback.SecurityModel{
					ResolveStrategy: fireback.ResolveStrategyWorkspace,
					AllowOnRoot:     true,
				})
				if err != nil {
					return fireback.QueryDSL{}, err
				}
				return *query, nil
			},
		}),
		// interfacetools.ModuleSetup is constructed here, not inside
		// abac.AbacCompleteModules(), so abac itself never has to import
		// interfacetools (only interfacetoolsdefs, for abac.Menu's plain data - see
		// modules/abac/Menu.go). abac.Menu is injected as ExtraAppMenus, which
		// interfacetools syncs (delete-by-uniqueId-then-recreate) as part of its own
		// migration on every `migration apply`, instead of requiring the manual
		// "seeders" command the way this used to (fireback-manage-menu.yml/
		// fireback-menu-cloud.yml/fireback-personal-menu.yml, before they moved here).
		interfacetools.ModuleSetup(&interfacetools.InterfaceToolsModuleConfig{
			ExtraAppMenus: abac.Menu,
		}),
	}

	// For fireback we have abac module added.
	modules = append(modules, abac.AbacCompleteModules()...)

	var xapp = &application.Application{
		Title: "Fireback core microservice - v" + fireback.FIREBACK_VERSION,
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
	os.Setenv("PRODUCT_UNIQUE_NAME", "fireback")

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	fireback.CommonHeadlessAppStart(xapp, func() {
		// If anything needs to be done after database initialized
		// fireback.RegionalContentSyncSeeders()
		// fireback.AppMenuSyncSeeders()
	})
}
