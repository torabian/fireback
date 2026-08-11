package interfacetools

import (
	"github.com/gin-gonic/gin"
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// InterfaceToolsModuleConfig lets a caller inject its own AppMenu definitions - e.g.
// abac.Menu - without this package importing anything back from that caller (abac
// deliberately only imports interfacetoolsdefs, the plain generated types, not this
// interfacetools package itself - see abac/Menu.go's own comment). main.go is what
// actually wires the two together, passing abac.Menu in here when it constructs the
// module list.
type InterfaceToolsModuleConfig struct {
	// ExtraAppMenus are synced (SyncAppMenus - delete-by-uniqueId-then-recreate, so
	// re-running migrations doesn't pile up duplicates) as part of this module's own
	// migration (ProvideEntityHandlers, called from ApplyMigration), rather than
	// requiring a separate manual "seeders" step the way this module's own
	// AppMenuSyncSeeders/seeders/AppMenu directory still does for anything else that
	// wants to land menu items via that older, plain-insert mechanism.
	ExtraAppMenus []*interfacetoolsdefs.AppMenuEntity
}

// ModuleSetup registers appMenu, tableViewSizing and timezoneGroup (plus the
// /cte-app-menus action) as their own application.ModuleProvider - split out of abac since
// none of them need Role/Workspace/User types directly. They have no dedicated
// front-end management screens (unlike e.g. modules/abac/messaging's providers), so only
// the backend moved here - see AbacCompleteModules.
func ModuleSetup(config *InterfaceToolsModuleConfig) *application.ModuleProvider {
	if config == nil {
		config = &InterfaceToolsModuleConfig{}
	}
	module := &application.ModuleProvider{
		Name: "abac",

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				interfacetoolsdefs.AppMenuBrowseActionGin(g, AppMenuBrowseAction)
				interfacetoolsdefs.AppMenuGetActionGin(g, AppMenuGetAction)
				interfacetoolsdefs.AppMenuCreateActionGin(g, AppMenuCreateAction)
				interfacetoolsdefs.AppMenuUpdateActionGin(g, AppMenuUpdateAction)
				interfacetoolsdefs.AppMenuAwareDeletePreviewActionGin(g, AppMenuAwareDeletePreviewAction)
				interfacetoolsdefs.AppMenuAwareDeleteActionGin(g, AppMenuAwareDeleteAction)
				interfacetoolsdefs.CteAppMenusActionGin(g, CteAppMenusAction)

				interfacetoolsdefs.TableViewSizingBrowseActionGin(g, TableViewSizingBrowseAction)
				interfacetoolsdefs.TableViewSizingGetActionGin(g, TableViewSizingGetAction)
				interfacetoolsdefs.TableViewSizingCreateActionGin(g, TableViewSizingCreateAction)
				interfacetoolsdefs.TableViewSizingUpdateActionGin(g, TableViewSizingUpdateAction)
				interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionGin(g, TableViewSizingAwareDeletePreviewAction)
				interfacetoolsdefs.TableViewSizingAwareDeleteActionGin(g, TableViewSizingAwareDeleteAction)

				interfacetoolsdefs.TimezoneGroupBrowseActionGin(g, TimezoneGroupBrowseAction)
				interfacetoolsdefs.TimezoneGroupGetActionGin(g, TimezoneGroupGetAction)
				interfacetoolsdefs.TimezoneGroupCreateActionGin(g, TimezoneGroupCreateAction)
				interfacetoolsdefs.TimezoneGroupUpdateActionGin(g, TimezoneGroupUpdateAction)
				interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionGin(g, TimezoneGroupAwareDeletePreviewAction)
				interfacetoolsdefs.TimezoneGroupAwareDeleteActionGin(g, TimezoneGroupAwareDeleteAction)

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
		if err := dbref.AutoMigrate(
			&interfacetoolsdefs.AppMenuEntity{},
			&interfacetoolsdefs.TableViewSizingEntity{},
			&interfacetoolsdefs.TimezoneGroupEntity{},
		); err != nil {
			return err
		}

		SyncAppMenus(config.ExtraAppMenus)

		return nil
	})

	module.ProvideSeederImportHandler(func() {
		TimezoneGroupSyncSeeders()
	})

	module.ProvideCliHandlers([]*cli.Command{
		{
			Name:        "interface",
			Aliases:     []string{"intf", "if"},
			Description: "Stores information about interface, which is useful for front-end application, such as menus, table sizes, sort orders",
			Commands: []*cli.Command{

				AppMenuCliFn(),
				{
					Name:        "tableviewsizing",
					Aliases:     []string{"tvs"},
					Description: "Stores the table column configuration per user in database",
					Commands: []*cli.Command{
						interfacetoolsdefs.TableViewSizingBrowseActionCliHandler(TableViewSizingBrowseAction),
						interfacetoolsdefs.TableViewSizingGetActionCliHandler(TableViewSizingGetAction),
						interfacetoolsdefs.TableViewSizingCreateActionCliHandler(TableViewSizingCreateAction),
						interfacetoolsdefs.TableViewSizingUpdateActionCliHandler(TableViewSizingUpdateAction),
						interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionCliHandler(TableViewSizingAwareDeletePreviewAction),
						interfacetoolsdefs.TableViewSizingAwareDeleteActionCliHandler(TableViewSizingAwareDeleteAction),
					},
				},
				{
					Name:        "timezone",
					Aliases:     []string{"tz"},
					Description: "Contains common timezones information",
					Commands: []*cli.Command{
						interfacetoolsdefs.TimezoneGroupBrowseActionCliHandler(TimezoneGroupBrowseAction),
						interfacetoolsdefs.TimezoneGroupGetActionCliHandler(TimezoneGroupGetAction),
						interfacetoolsdefs.TimezoneGroupCreateActionCliHandler(TimezoneGroupCreateAction),
						interfacetoolsdefs.TimezoneGroupUpdateActionCliHandler(TimezoneGroupUpdateAction),
						interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionCliHandler(TimezoneGroupAwareDeletePreviewAction),
						interfacetoolsdefs.TimezoneGroupAwareDeleteActionCliHandler(TimezoneGroupAwareDeleteAction),
					},
				},
			},
		},
	})

	return module
}
