//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	seeders "pixelplux.com/fireback/modules/geo/seeders/GeoLocation"
	"pixelplux.com/fireback/modules/workspaces"

	metas "pixelplux.com/fireback/modules/geo/metas"
)

var GeoLocationCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoLocationCommonCliFlags,
	Usage:   "Create a new geoLocation",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationFromCli(c)

		if entity, err := GeoLocationActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoLocationCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoLocation, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoLocationEntity{}

		for _, item := range GeoLocationCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoLocationActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoLocationUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoLocationCommonCliFlagsOptional,
	Usage:   "Updates a geoLocation by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationFromCli(c)

		if entity, err := GeoLocationActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoLocationCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoLocation uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.StringFlag{
		Name:     "name",
		Required: false,
		Usage:    "name",
	},

	&cli.StringFlag{
		Name:     "code",
		Required: false,
		Usage:    "code",
	},

	&cli.StringFlag{
		Name:     "typeId",
		Required: false,
		Usage:    "type",
	},

	&cli.StringFlag{
		Name:     "status",
		Required: false,
		Usage:    "status",
	},

	&cli.StringFlag{
		Name:     "flag",
		Required: false,
		Usage:    "flag",
	},

	&cli.StringFlag{
		Name:     "officialName",
		Required: false,
		Usage:    "officialName",
	},
}

var GeoLocationCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},

	{
		Name:        "code",
		StructField: "Code",
		Required:    false,
		Usage:       "code",
		Type:        "string",
	},

	{
		Name:        "status",
		StructField: "Status",
		Required:    false,
		Usage:       "status",
		Type:        "string",
	},

	{
		Name:        "flag",
		StructField: "Flag",
		Required:    false,
		Usage:       "flag",
		Type:        "string",
	},

	{
		Name:        "officialName",
		StructField: "OfficialName",
		Required:    false,
		Usage:       "officialName",
		Type:        "string",
	},
}

var GeoLocationCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoLocation uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.StringFlag{
		Name:     "name",
		Required: false,
		Usage:    "name",
	},

	&cli.StringFlag{
		Name:     "code",
		Required: false,
		Usage:    "code",
	},

	&cli.StringFlag{
		Name:     "typeId",
		Required: false,
		Usage:    "type",
	},

	&cli.StringFlag{
		Name:     "status",
		Required: false,
		Usage:    "status",
	},

	&cli.StringFlag{
		Name:     "flag",
		Required: false,
		Usage:    "flag",
	},

	&cli.StringFlag{
		Name:     "officialName",
		Required: false,
		Usage:    "officialName",
	},
}

func CastGeoLocationFromCli(c *cli.Context) *GeoLocationEntity {
	geoLocation := &GeoLocationEntity{}

	if c.IsSet("id") {
		geoLocation.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoLocation.ParentId = &x
	}

	if c.IsSet("name") || false {
		value := c.String("name")
		geoLocation.Name = &value
	}

	if c.IsSet("code") || false {
		value := c.String("code")
		geoLocation.Code = &value
	}

	if c.IsSet("typeId") || false {
		id := c.String("typeId")
		geoLocation.TypeId = &id
	}

	if c.IsSet("status") || false {
		value := c.String("status")
		geoLocation.Status = &value
	}

	if c.IsSet("flag") || false {
		value := c.String("flag")
		geoLocation.Flag = &value
	}

	if c.IsSet("officialName") || false {
		value := c.String("officialName")
		geoLocation.OfficialName = &value
	}

	return geoLocation
}

func GeoLocationSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoLocationActionCreate,
		reflect.ValueOf(&GeoLocationEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoLocationSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoLocationActionCreate,
		reflect.ValueOf(&GeoLocationEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func GeoLocationWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoLocationActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoLocation", result)
	}
}

var GeoLocationWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoLocations ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoLocationActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoLocationImportExportCommands = []cli.Command{
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
			GeoLocationActionSeeder(query, c.Int("count"))

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
				Value: "geoLocation-seeder.yml",
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
			GeoLocationActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoLocation-seeder-geoLocation.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoLocations, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoLocationEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)

			fmt.Println(data)
			return nil
		},
	},

	cli.Command{
		Name:  "list",
		Usage: "Prints the list of files attached to this module for syncing or bootstrapping project",
		Action: func(c *cli.Context) error {
			if entity, err := workspaces.GetSeederFilenames(&seeders.ViewsFs, ""); err != nil {
				fmt.Println(err.Error())
			} else {

				f, _ := json.MarshalIndent(entity, "", "  ")
				fmt.Println(string(f))
			}

			return nil
		},
	},
	cli.Command{
		Name:  "sync",
		Usage: "Tries to sync the embedded content into the database, the list could be seen by 'list' command",
		Action: func(c *cli.Context) error {

			workspaces.CommonCliImportEmbedCmd(c,
				GeoLocationActionCreate,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				&seeders.ViewsFs,
			)

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
				GeoLocationActionQuery,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoLocationFieldMap.yml",
				GeoLocationPreloadRelations,
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
				GeoLocationActionCreate,
				reflect.ValueOf(&GeoLocationEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoLocationCliCommands []cli.Command = []cli.Command{

	GeoLocationCreateCmd,
	GeoLocationUpdateCmd,
	GeoLocationCreateInteractiveCmd,

	GeoLocationWipeCmd,

	workspaces.GetCommonQuery(GeoLocationActionQuery),

	workspaces.GetCommonCteQuery(GeoLocationActionCteQuery),
	workspaces.GetCommonPivotQuery(GeoLocationActionCommonPivotQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoLocationEntity{}).Elem(), GeoLocationActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoLocationEntity{}).Elem(), GeoLocationActionRemove),
}

func GeoLocationCliFn() cli.Command {
	GeoLocationCliCommands = append(GeoLocationCliCommands, GeoLocationImportExportCommands...)

	return cli.Command{
		Name:        "location",
		Description: "GeoLocations module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoLocations module (" + fmt.Sprintf("%v", len(GeoLocationCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoLocationCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoLocationPreloadRelations []string = []string{}
