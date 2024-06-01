package geo

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func GeoModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "geo",
		Definitions: &Module2Definitions,
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

	module.Actions = [][]workspaces.Module2Action{
		GetGeoCityModule2Actions(),
		GetGeoProvinceModule2Actions(),
		GetGeoStateModule2Actions(),
		GetGeoCountryModule2Actions(),
		GetGeoLocationModule2Actions(),
		GetGeoLocationTypeModule2Actions(),
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

	module.ProvideCliHandlers([]cli.Command{
		{
			Name:  "geo",
			Usage: "Geo location tools, and data set, cities, and provinces",
			Subcommands: cli.Commands{
				GeoCityCliFn(),
				GeoProvinceCliFn(),
				GeoStateCliFn(),
				GeoCountryCliFn(),
				GeoLocationTypeCliFn(),
				GeoLocationCliFn(),
			},
		},
	})

	return module
}
