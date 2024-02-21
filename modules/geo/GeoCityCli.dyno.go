//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	seeders "pixelplux.com/fireback/modules/geo/seeders/GeoCity"
	"pixelplux.com/fireback/modules/workspaces"

	metas "pixelplux.com/fireback/modules/geo/metas"
)

var GeoCityCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoCityCommonCliFlags,
	Usage:   "Create a new geoCity",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoCityFromCli(c)

		if entity, err := GeoCityActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoCityCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoCity, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoCityEntity{}

		for _, item := range GeoCityCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoCityActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoCityUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoCityCommonCliFlagsOptional,
	Usage:   "Updates a geoCity by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoCityFromCli(c)

		if entity, err := GeoCityActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoCityCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoCity uniqueId (primary key)",
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
		Name:     "provinceId",
		Required: false,
		Usage:    "province",
	},

	&cli.StringFlag{
		Name:     "stateId",
		Required: false,
		Usage:    "state",
	},

	&cli.StringFlag{
		Name:     "countryId",
		Required: false,
		Usage:    "country",
	},
}

var GeoCityCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
}

var GeoCityCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoCity uniqueId (primary key)",
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
		Name:     "provinceId",
		Required: false,
		Usage:    "province",
	},

	&cli.StringFlag{
		Name:     "stateId",
		Required: false,
		Usage:    "state",
	},

	&cli.StringFlag{
		Name:     "countryId",
		Required: false,
		Usage:    "country",
	},
}

func CastGeoCityFromCli(c *cli.Context) *GeoCityEntity {
	geoCity := &GeoCityEntity{}

	if c.IsSet("id") {
		geoCity.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoCity.ParentId = &x
	}

	if c.IsSet("name") || false {
		value := c.String("name")
		geoCity.Name = &value
	}

	if c.IsSet("provinceId") || false {
		id := c.String("provinceId")
		geoCity.ProvinceId = &id
	}

	if c.IsSet("stateId") || false {
		id := c.String("stateId")
		geoCity.StateId = &id
	}

	if c.IsSet("countryId") || false {
		id := c.String("countryId")
		geoCity.CountryId = &id
	}

	return geoCity
}

func GeoCitySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoCityActionCreate,
		reflect.ValueOf(&GeoCityEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoCitySyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoCityActionCreate,
		reflect.ValueOf(&GeoCityEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func GeoCityWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoCityActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoCity", result)
	}
}

var GeoCityWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoCitys ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoCityActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoCityImportExportCommands = []cli.Command{
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
			GeoCityActionSeeder(query, c.Int("count"))

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
				Value: "geoCity-seeder.yml",
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
			GeoCityActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoCity-seeder-geoCity.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoCitys, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoCityEntity{}
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
				GeoCityActionCreate,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
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
				GeoCityActionQuery,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoCityFieldMap.yml",
				GeoCityPreloadRelations,
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
				GeoCityActionCreate,
				reflect.ValueOf(&GeoCityEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoCityCliCommands []cli.Command = []cli.Command{

	GeoCityCreateCmd,
	GeoCityUpdateCmd,
	GeoCityCreateInteractiveCmd,

	GeoCityWipeCmd,

	workspaces.GetCommonQuery(GeoCityActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoCityEntity{}).Elem(), GeoCityActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoCityEntity{}).Elem(), GeoCityActionRemove),
}

func GeoCityCliFn() cli.Command {
	GeoCityCliCommands = append(GeoCityCliCommands, GeoCityImportExportCommands...)

	return cli.Command{
		Name:        "city",
		Description: "GeoCitys module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoCitys module (" + fmt.Sprintf("%v", len(GeoCityCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoCityCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoCityPreloadRelations []string = []string{}
