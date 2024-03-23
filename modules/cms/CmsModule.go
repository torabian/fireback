package cms

import (
	"embed"
	"fmt"

	// "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func CmsModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "cms",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {

	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	module.ProvidePermissionHandler(
		ALL_PAGECATEGORY_PERMISSIONS,
		ALL_PAGETAG_PERMISSIONS,
		ALL_PAGE_PERMISSIONS,
		ALL_POSTCATEGORY_PERMISSIONS,
		ALL_POSTTAG_PERMISSIONS,
		ALL_POST_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetPageCategoryModule2Actions(),
		GetPageTagModule2Actions(),
		GetPageModule2Actions(),

		GetPostCategoryModule2Actions(),
		GetPostTagModule2Actions(),
		GetPostModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(&PageCategoryEntity{}, PageCategoryEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&PageEntity{}, &PageTagEntity{}); err != nil {
			fmt.Println(err.Error())
		}

		if err := dbref.AutoMigrate(&PostCategoryEntity{}, PostCategoryEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&PostEntity{}, &PostTagEntity{}); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		PageCategoryCliFn(),
		PageCliFn(),
		PageTagCliFn(),
		PostCategoryCliFn(),
		PostCliFn(),
		PostTagCliFn(),
	})

	return module
}
