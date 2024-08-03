package accessibility

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func AccessibilityModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "accessibility",
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

	module.ProvidePermissionHandler(ALL_KEYBOARD_SHORTCUT_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetKeyboardShortcutModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&KeyboardShortcutEntity{},
			&KeyboardShortcutDefaultCombination{},
			&KeyboardShortcutUserCombination{},
			&KeyboardShortcutEntityPolyglot{},
		)
	})

	return module
}
