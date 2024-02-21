//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	metas "pixelplux.com/fireback/modules/geo/metas"
	"pixelplux.com/fireback/modules/workspaces"
)

var LocationDataCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   LocationDataCommonCliFlags,
	Usage:   "Create a new locationData",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLocationDataFromCli(c)

		if entity, err := LocationDataActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var LocationDataCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new locationData, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &LocationDataEntity{}

		for _, item := range LocationDataCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := LocationDataActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var LocationDataUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   LocationDataCommonCliFlagsOptional,
	Usage:   "Updates a locationData by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLocationDataFromCli(c)

		if entity, err := LocationDataActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var LocationDataCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "LocationData uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.Float64Flag{
		Name:     "lat",
		Required: false,
		Usage:    "lat",
	},

	&cli.Float64Flag{
		Name:     "lng",
		Required: false,
		Usage:    "lng",
	},

	&cli.StringFlag{
		Name:     "physicalAddress",
		Required: false,
		Usage:    "physicalAddress",
	},
}

var LocationDataCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "lat",
		StructField: "Lat",
		Required:    false,
		Usage:       "lat",
		Type:        "float64",
	},

	{
		Name:        "lng",
		StructField: "Lng",
		Required:    false,
		Usage:       "lng",
		Type:        "float64",
	},

	{
		Name:        "physicalAddress",
		StructField: "PhysicalAddress",
		Required:    false,
		Usage:       "physicalAddress",
		Type:        "string",
	},
}

var LocationDataCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "LocationData uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.Float64Flag{
		Name:     "lat",
		Required: false,
		Usage:    "lat",
	},

	&cli.Float64Flag{
		Name:     "lng",
		Required: false,
		Usage:    "lng",
	},

	&cli.StringFlag{
		Name:     "physicalAddress",
		Required: false,
		Usage:    "physicalAddress",
	},
}

func CastLocationDataFromCli(c *cli.Context) *LocationDataEntity {
	locationData := &LocationDataEntity{}

	if c.IsSet("id") {
		locationData.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		locationData.ParentId = &x
	}

	if c.IsSet("lat") || false {
		val := c.Float64("lat")
		locationData.Lat = &val
	}

	if c.IsSet("lng") || false {
		val := c.Float64("lng")
		locationData.Lng = &val
	}

	if c.IsSet("physicalAddress") || false {
		value := c.String("physicalAddress")
		locationData.PhysicalAddress = &value
	}

	return locationData
}

func LocationDataSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		LocationDataActionCreate,
		reflect.ValueOf(&LocationDataEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func LocationDataWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := LocationDataActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "LocationData", result)
	}
}

var LocationDataWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire locationDatas ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := LocationDataActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var LocationDataImportExportCommands = []cli.Command{
	{

		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
		},
		Action: func(c *cli.Context) error {
			query := workspaces.CommonCliQueryDSLBuilder(c)
			LocationDataActionSeeder(query, c.Int("count"))

			return nil
		},
	},
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The address of file you want the csv be exported to",
				Value: "locationData-seeder.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
			f := workspaces.CommonCliQueryDSLBuilder(c)
			LocationDataActionSeederInit(f, c.String("file"), c.String("format"))
			return nil
		},
	},
	{
		Name:    "validate",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Validates an import file, such as yaml, json, csv, and gives some insights how the after import it would look like",
				Value: "locationData-seeder-locationData.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of locationDatas, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]LocationDataEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)

			fmt.Println(data)
			return nil
		},
	},

	cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json be exported to",
				Required: true,
			}),
		Usage: "Exports a query results into the csv/yaml/json format",
		Action: func(c *cli.Context) error {

			workspaces.CommonCliExportCmd(c,
				LocationDataActionQuery,
				reflect.ValueOf(&LocationDataEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"LocationDataFieldMap.yml",
				LocationDataPreloadRelations,
			)

			return nil
		},
	},
	cli.Command{
		Name: "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {

			workspaces.CommonCliImportCmd(c,
				LocationDataActionCreate,
				reflect.ValueOf(&LocationDataEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var LocationDataCliCommands []cli.Command = []cli.Command{

	LocationDataCreateCmd,
	LocationDataUpdateCmd,
	LocationDataCreateInteractiveCmd,

	LocationDataWipeCmd,

	workspaces.GetCommonQuery(LocationDataActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&LocationDataEntity{}).Elem(), LocationDataActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&LocationDataEntity{}).Elem(), LocationDataActionRemove),
}

func LocationDataCliFn() cli.Command {
	LocationDataCliCommands = append(LocationDataCliCommands, LocationDataImportExportCommands...)

	return cli.Command{
		Name:        "location",
		Description: "LocationDatas module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the locationDatas module (" + fmt.Sprintf("%v", len(LocationDataCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: LocationDataCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var LocationDataPreloadRelations []string = []string{}
