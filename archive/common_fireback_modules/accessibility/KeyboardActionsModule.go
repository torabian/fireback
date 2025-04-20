package accessibility

import (
	"embed"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func AccessibilityModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name:        "accessibility",
		Definitions: &Module3Definitions,
	}

	module.ProvideMockImportHandler(func() {
		KeyboardShortcutSyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		KeyboardShortcutWriteQueryMock(fireback.MockQueryContext{Languages: languages})
	})

	module.ProvideCliHandlers([]cli.Command{
		KeyboardShortcutCliFn(),
	})

	module.ProvidePermissionHandler(ALL_KEYBOARD_SHORTCUT_PERMISSIONS)

	module.Actions = [][]fireback.Module3Action{
		GetKeyboardShortcutModule3Actions(),
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
