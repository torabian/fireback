package interfacetools

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// ModuleSetup registers appMenu, tableViewSizing and timezoneGroup (plus the
// /cte-app-menus action) as their own fireback.ModuleProvider - split out of abac since
// none of them need Role/Workspace/User types directly. They have no dedicated
// front-end management screens (unlike e.g. modules/abac/messaging's providers), so only
// the backend moved here - see AbacCompleteModules.
func ModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name: "abac",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *fireback.FirebackApp) error{
			func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
				AppMenuBrowseActionGin(g, AppMenuBrowseAction)
				AppMenuGetActionGin(g, AppMenuGetAction)
				AppMenuCreateActionGin(g, AppMenuCreateAction)
				AppMenuUpdateActionGin(g, AppMenuUpdateAction)
				AppMenuAwareDeletePreviewActionGin(g, AppMenuAwareDeletePreviewAction)
				AppMenuAwareDeleteActionGin(g, AppMenuAwareDeleteAction)
				CteAppMenusActionGin(g, CteAppMenusAction)

				TableViewSizingBrowseActionGin(g, TableViewSizingBrowseAction)
				TableViewSizingGetActionGin(g, TableViewSizingGetAction)
				TableViewSizingCreateActionGin(g, TableViewSizingCreateAction)
				TableViewSizingUpdateActionGin(g, TableViewSizingUpdateAction)
				TableViewSizingAwareDeletePreviewActionGin(g, TableViewSizingAwareDeletePreviewAction)
				TableViewSizingAwareDeleteActionGin(g, TableViewSizingAwareDeleteAction)

				TimezoneGroupBrowseActionGin(g, TimezoneGroupBrowseAction)
				TimezoneGroupGetActionGin(g, TimezoneGroupGetAction)
				TimezoneGroupCreateActionGin(g, TimezoneGroupCreateAction)
				TimezoneGroupUpdateActionGin(g, TimezoneGroupUpdateAction)
				TimezoneGroupAwareDeletePreviewActionGin(g, TimezoneGroupAwareDeletePreviewAction)
				TimezoneGroupAwareDeleteActionGin(g, TimezoneGroupAwareDeleteAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_APP_MENU_PERMISSIONS,
		ALL_TABLE_VIEW_SIZING_PERMISSIONS,
		ALL_TIMEZONE_GROUP_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&AppMenuEntity{},
			&TableViewSizingEntity{},
			&TimezoneGroupEntity{},
		)
	})

	module.ProvideSeederImportHandler(func() {
		TimezoneGroupSyncSeeders()
	})

	module.ProvideCliHandlers([]*cli.Command{
		AppMenuCliFn(),
		TableViewSizingBrowseActionCliHandler(TableViewSizingBrowseAction),
		TableViewSizingGetActionCliHandler(TableViewSizingGetAction),
		TableViewSizingCreateActionCliHandler(TableViewSizingCreateAction),
		TableViewSizingUpdateActionCliHandler(TableViewSizingUpdateAction),
		TableViewSizingAwareDeletePreviewActionCliHandler(TableViewSizingAwareDeletePreviewAction),
		TableViewSizingAwareDeleteActionCliHandler(TableViewSizingAwareDeleteAction),
		TimezoneGroupBrowseActionCliHandler(TimezoneGroupBrowseAction),
		TimezoneGroupGetActionCliHandler(TimezoneGroupGetAction),
		TimezoneGroupCreateActionCliHandler(TimezoneGroupCreateAction),
		TimezoneGroupUpdateActionCliHandler(TimezoneGroupUpdateAction),
		TimezoneGroupAwareDeletePreviewActionCliHandler(TimezoneGroupAwareDeletePreviewAction),
		TimezoneGroupAwareDeleteActionCliHandler(TimezoneGroupAwareDeleteAction),
	})

	return module
}
