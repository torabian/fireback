//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoProvince"
	"github.com/torabian/fireback/modules/workspaces"

	metas "github.com/torabian/fireback/modules/geo/metas"
)

var GeoProvinceCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoProvinceCommonCliFlags,
	Usage:   "Create a new geoProvince",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoProvinceFromCli(c)

		if entity, err := GeoProvinceActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoProvinceCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoProvince, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoProvinceEntity{}

		for _, item := range GeoProvinceCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoProvinceActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoProvinceUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoProvinceCommonCliFlagsOptional,
	Usage:   "Updates a geoProvince by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoProvinceFromCli(c)

		if entity, err := GeoProvinceActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoProvinceCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoProvince uniqueId (primary key)",
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
		Name:     "countryId",
		Required: false,
		Usage:    "country",
	},
}

var GeoProvinceCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
}

var GeoProvinceCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoProvince uniqueId (primary key)",
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
		Name:     "countryId",
		Required: false,
		Usage:    "country",
	},
}

func CastGeoProvinceFromCli(c *cli.Context) *GeoProvinceEntity {
	geoProvince := &GeoProvinceEntity{}

	if c.IsSet("id") {
		geoProvince.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoProvince.ParentId = &x
	}

	if c.IsSet("name") || false {
		value := c.String("name")
		geoProvince.Name = &value
	}

	if c.IsSet("countryId") || false {
		id := c.String("countryId")
		geoProvince.CountryId = &id
	}

	return geoProvince
}

func GeoProvinceSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoProvinceActionCreate,
		reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoProvinceSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoProvinceActionCreate,
		reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func GeoProvinceWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoProvinceActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoProvince", result)
	}
}

var GeoProvinceWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoProvinces ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoProvinceActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoProvinceImportExportCommands = []cli.Command{
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
			GeoProvinceActionSeeder(query, c.Int("count"))

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
				Value: "geoProvince-seeder.yml",
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
			GeoProvinceActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoProvince-seeder-geoProvince.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoProvinces, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoProvinceEntity{}
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
				GeoProvinceActionCreate,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
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
				GeoProvinceActionQuery,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoProvinceFieldMap.yml",
				GeoProvincePreloadRelations,
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
				GeoProvinceActionCreate,
				reflect.ValueOf(&GeoProvinceEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoProvinceCliCommands []cli.Command = []cli.Command{

	GeoProvinceCreateCmd,
	GeoProvinceUpdateCmd,
	GeoProvinceCreateInteractiveCmd,

	GeoProvinceWipeCmd,

	workspaces.GetCommonQuery(GeoProvinceActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoProvinceEntity{}).Elem(), GeoProvinceActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoProvinceEntity{}).Elem(), GeoProvinceActionRemove),
}

func GeoProvinceCliFn() cli.Command {
	GeoProvinceCliCommands = append(GeoProvinceCliCommands, GeoProvinceImportExportCommands...)

	return cli.Command{
		Name:        "province",
		Description: "GeoProvinces module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoProvinces module (" + fmt.Sprintf("%v", len(GeoProvinceCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoProvinceCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoProvincePreloadRelations []string = []string{}
