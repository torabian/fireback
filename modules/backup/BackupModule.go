package backup

import (
	"github.com/gin-gonic/gin"
	backupdefs "github.com/torabian/fireback/modules/backup/defs"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

// ModuleSetup wires this module into a fireback app the same way every
// other module does (abac.WorkspaceModuleSetup, storage.StorageModuleSetup,
// eventbus.ModuleSetup, ...): a project registers it once in main.go
// (`backup.ModuleSetup(nil)` in the Modules slice) and gets its CLI
// commands, HTTP routes, and config all wired automatically - rather than
// main.go hand-assembling a bare CliHandlers-only ModuleProvider and this
// module reading its own env vars directly, which is what it used to do.
//
// cfg is currently unused (nil is fine) - reserved the same way
// EventBusModuleConfig/InterfaceToolsModuleConfig are, for options a
// project might want to override later (e.g. a custom ApiToken check)
// without changing this function's signature again.
type BackupModuleConfig struct{}

func ModuleSetup(cfg *BackupModuleConfig) *application.ModuleProvider {
	return &application.ModuleProvider{
		Name: "backup",

		CliHandlers: []*cli.Command{
			{
				Name:  "backup",
				Usage: "Take/restore Postgres point-in-time backups (wal-g), or dump/restore a single database (postgres/mysql/sqlite)",
				Commands: append([]*cli.Command{
					// Per-field get/set for Backup.emi.yml's `config:` block - nested
					// here under "backup" the same way abacdefs.GetConfigCli() is
					// nested under "abac" (AbacModule.go), to avoid a name collision
					// with fireback's own top-level "config" command.
					{
						Name:     "config",
						Usage:    "Set of tools to configure the backup module",
						Commands: backupdefs.GetConfigCli(),
					},
				}, Commands()...),
			},
		},

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				MountDumpHttp(g)
				return nil
			},
		},

		// Folds backupdefs.Config into fireback's combined "config list" -
		// see ConfigRegistry.go/GetCombinedConfigInfo, and abacdefs'
		// identical use of this same hook in AbacModule.go.
		ConfigProvider: func() interface{} {
			return backupdefs.LoadConfiguration()
		},
	}
}
