package geo

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func GeoModuleSetup() *workspaces.ModuleProvider {

	module := &workspaces.ModuleProvider{
		Name:          "geo",
		Definitions:   &Module3Definitions,
		ActionsBundle: GetGeoActionsBundle(),
	}

	module.ProvideMockImportHandler(func() {
		GeoCountrySyncSeeders()
		GeoProvinceSyncSeeders()
		GeoCitySyncSeeders()
		GeoLocationTypeSyncSeeders()
		GeoLocationSyncSeeders()
	})

	module.ProvideSeederImportHandler(func() {
		GeoCountrySyncSeeders()
		GeoProvinceSyncSeeders()
		GeoCitySyncSeeders()
		GeoLocationTypeSyncSeeders()
		GeoLocationSyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		GeoCountryWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 20})
		GeoProvinceWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 20})
		GeoCityWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 50})
		GeoLocationWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 50})
		GeoLocationTypeWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 50})
	})

	module.ProvidePermissionHandler(
		ALL_GEO_CITY_PERMISSIONS,
		ALL_GEO_PROVINCE_PERMISSIONS,
		ALL_GEO_STATE_PERMISSIONS,
		ALL_GEO_LOCATION_TYPE_PERMISSIONS,
		ALL_GEO_LOCATION_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module3Action{
		GetGeoCityModule3Actions(),
		GetGeoProvinceModule3Actions(),
		GetGeoStateModule3Actions(),
		GetGeoCountryModule3Actions(),
		GetGeoLocationModule3Actions(),
		GetGeoLocationTypeModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&GeoCityEntity{},
			&GeoProvinceEntity{},
			&GeoProvinceEntityPolyglot{},
			&GeoStateEntity{},
			&GeoStateEntityPolyglot{},
			&GeoCountryEntity{},
			&GeoCountryEntityPolyglot{},
			&GeoLocationEntity{},
			&GeoLocationEntityPolyglot{},
			&GeoLocationTypeEntity{},
			&GeoLocationTypeEntityPolyglot{},
		)
	})

	return module
}
