//go:build !omit_cli

package geo

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	seeders "pixelplux.com/fireback/modules/geo/seeders/GeoLocationType"
	"pixelplux.com/fireback/modules/workspaces"

	metas "pixelplux.com/fireback/modules/geo/metas"
)

var GeoLocationTypeCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoLocationTypeCommonCliFlags,
	Usage:   "Create a new geoLocationType",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationTypeFromCli(c)

		if entity, err := GeoLocationTypeActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoLocationTypeCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoLocationType, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoLocationTypeEntity{}

		for _, item := range GeoLocationTypeCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoLocationTypeActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoLocationTypeUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoLocationTypeCommonCliFlagsOptional,
	Usage:   "Updates a geoLocationType by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoLocationTypeFromCli(c)

		if entity, err := GeoLocationTypeActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoLocationTypeCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoLocationType uniqueId (primary key)",
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
}

var GeoLocationTypeCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
}

var GeoLocationTypeCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoLocationType uniqueId (primary key)",
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
}

func CastGeoLocationTypeFromCli(c *cli.Context) *GeoLocationTypeEntity {
	geoLocationType := &GeoLocationTypeEntity{}

	if c.IsSet("id") {
		geoLocationType.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoLocationType.ParentId = &x
	}

	if c.IsSet("name") || false {
		value := c.String("name")
		geoLocationType.Name = &value
	}

	return geoLocationType
}

func GeoLocationTypeSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoLocationTypeActionCreate,
		reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoLocationTypeSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoLocationTypeActionCreate,
		reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func GeoLocationTypeWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoLocationTypeActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoLocationType", result)
	}
}

var GeoLocationTypeWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoLocationTypes ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoLocationTypeActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoLocationTypeImportExportCommands = []cli.Command{
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
			GeoLocationTypeActionSeeder(query, c.Int("count"))

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
				Value: "geoLocationType-seeder.yml",
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
			GeoLocationTypeActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoLocationType-seeder-geoLocationType.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoLocationTypes, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoLocationTypeEntity{}
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
				GeoLocationTypeActionCreate,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
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
				GeoLocationTypeActionQuery,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoLocationTypeFieldMap.yml",
				GeoLocationTypePreloadRelations,
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
				GeoLocationTypeActionCreate,
				reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoLocationTypeCliCommands []cli.Command = []cli.Command{

	GeoLocationTypeCreateCmd,
	GeoLocationTypeUpdateCmd,
	GeoLocationTypeCreateInteractiveCmd,

	GeoLocationTypeWipeCmd,

	workspaces.GetCommonQuery(GeoLocationTypeActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(), GeoLocationTypeActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoLocationTypeEntity{}).Elem(), GeoLocationTypeActionRemove),
}

func GeoLocationTypeCliFn() cli.Command {
	GeoLocationTypeCliCommands = append(GeoLocationTypeCliCommands, GeoLocationTypeImportExportCommands...)

	return cli.Command{
		Name:        "type",
		Description: "GeoLocationTypes module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoLocationTypes module (" + fmt.Sprintf("%v", len(GeoLocationTypeCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoLocationTypeCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoLocationTypePreloadRelations []string = []string{}
