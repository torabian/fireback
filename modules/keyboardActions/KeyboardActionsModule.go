package keyboardActions

import (
	"embed"

	"github.com/urfave/cli"
	"gorm.io/gorm"
	"pixelplux.com/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func KeyboardActionsModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "keyboardActions",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		KeyboardShortcutSyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		KeyboardShortcutWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
	})

	module.ProvideCliHandlers([]cli.Command{
		KeyboardShortcutCliFn(),
	})

	module.ProvidePermissionHandler(ALL_KEYBOARDSHORTCUT_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetKeyboardShortcutModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&KeyboardShortcutEntity{})
		dbref.AutoMigrate(&KeyboardShortcutDefaultCombination{})
		dbref.AutoMigrate(&KeyboardShortcutUserCombination{})
		dbref.AutoMigrate(&KeyboardShortcutEntityPolyglot{})
	})

	return module
}
