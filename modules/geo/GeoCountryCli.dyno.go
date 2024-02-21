//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	seeders "github.com/torabian/fireback/modules/geo/seeders/GeoCountry"
	"github.com/torabian/fireback/modules/workspaces"

	metas "github.com/torabian/fireback/modules/geo/metas"
)

var GeoCountryCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoCountryCommonCliFlags,
	Usage:   "Create a new geoCountry",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoCountryFromCli(c)

		if entity, err := GeoCountryActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoCountryCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoCountry, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoCountryEntity{}

		for _, item := range GeoCountryCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoCountryActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoCountryUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoCountryCommonCliFlagsOptional,
	Usage:   "Updates a geoCountry by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoCountryFromCli(c)

		if entity, err := GeoCountryActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoCountryCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoCountry uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
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
		Name:     "commonName",
		Required: false,
		Usage:    "commonName",
	},

	&cli.StringFlag{
		Name:     "officialName",
		Required: false,
		Usage:    "officialName",
	},
}

var GeoCountryCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

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
		Name:        "commonName",
		StructField: "CommonName",
		Required:    false,
		Usage:       "commonName",
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

var GeoCountryCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoCountry uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
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
		Name:     "commonName",
		Required: false,
		Usage:    "commonName",
	},

	&cli.StringFlag{
		Name:     "officialName",
		Required: false,
		Usage:    "officialName",
	},
}

func CastGeoCountryFromCli(c *cli.Context) *GeoCountryEntity {
	geoCountry := &GeoCountryEntity{}

	if c.IsSet("id") {
		geoCountry.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoCountry.ParentId = &x
	}

	if c.IsSet("status") || false {
		value := c.String("status")
		geoCountry.Status = &value
	}

	if c.IsSet("flag") || false {
		value := c.String("flag")
		geoCountry.Flag = &value
	}

	if c.IsSet("commonName") || false {
		value := c.String("commonName")
		geoCountry.CommonName = &value
	}

	if c.IsSet("officialName") || false {
		value := c.String("officialName")
		geoCountry.OfficialName = &value
	}

	return geoCountry
}

func GeoCountrySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoCountryActionCreate,
		reflect.ValueOf(&GeoCountryEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoCountrySyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoCountryActionCreate,
		reflect.ValueOf(&GeoCountryEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func GeoCountryWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoCountryActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoCountry", result)
	}
}

var GeoCountryWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoCountrys ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoCountryActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoCountryImportExportCommands = []cli.Command{
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
			GeoCountryActionSeeder(query, c.Int("count"))

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
				Value: "geoCountry-seeder.yml",
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
			GeoCountryActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoCountry-seeder-geoCountry.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoCountrys, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoCountryEntity{}
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
				GeoCountryActionCreate,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
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
				GeoCountryActionQuery,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoCountryFieldMap.yml",
				GeoCountryPreloadRelations,
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
				GeoCountryActionCreate,
				reflect.ValueOf(&GeoCountryEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoCountryCliCommands []cli.Command = []cli.Command{

	GeoCountryCreateCmd,
	GeoCountryUpdateCmd,
	GeoCountryCreateInteractiveCmd,

	GeoCountryWipeCmd,

	workspaces.GetCommonQuery(GeoCountryActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoCountryEntity{}).Elem(), GeoCountryActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoCountryEntity{}).Elem(), GeoCountryActionRemove),
}

func GeoCountryCliFn() cli.Command {
	GeoCountryCliCommands = append(GeoCountryCliCommands, GeoCountryImportExportCommands...)

	return cli.Command{
		Name:        "country",
		Description: "GeoCountrys module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoCountrys module (" + fmt.Sprintf("%v", len(GeoCountryCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoCountryCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoCountryPreloadRelations []string = []string{}
