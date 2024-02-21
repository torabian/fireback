package geo

import (
	"fmt"

	"github.com/urfave/cli"
	"gorm.io/gorm"
	"pixelplux.com/fireback/modules/workspaces"
)

func GeoModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "geo",
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
		ALL_GEOCITY_PERMISSIONS,
		ALL_GEOPROVINCE_PERMISSIONS,
		ALL_GEOSTATE_PERMISSIONS,
		ALL_LOCATIONDATA_PERMISSIONS,
		ALL_GEOLOCATIONTYPE_PERMISSIONS,
		ALL_GEOLOCATION_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetGeoCityModule2Actions(),
		GetGeoProvinceModule2Actions(),
		GetGeoStateModule2Actions(),
		GetLocationDataModule2Actions(),
		GetGeoCountryModule2Actions(),
		GetGeoLocationModule2Actions(),
		GetGeoLocationTypeModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(&GeoCityEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoProvinceEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoProvinceEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoStateEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoStateEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&LocationDataEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoCountryEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&GeoCountryEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(
			&GeoLocationEntityPolyglot{},
			&GeoLocationTypeEntityPolyglot{},
			&GeoLocationEntity{},
			&GeoLocationTypeEntity{},
		); err != nil {
			fmt.Println(err.Error())
		}
	})

	module.ProvideCliHandlers([]cli.Command{
		{
			Name:  "geo",
			Usage: "Geo location tools, and data set, cities, and provinces",
			Subcommands: cli.Commands{
				GeoCityCliFn(),
				GeoProvinceCliFn(),
				GeoStateCliFn(),
				LocationDataCliFn(),
				GeoCountryCliFn(),
				GeoLocationTypeCliFn(),
				GeoLocationCliFn(),
			},
		},
	})

	return module
}
