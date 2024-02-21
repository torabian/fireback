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

var GeoStateCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   GeoStateCommonCliFlags,
	Usage:   "Create a new geoState",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoStateFromCli(c)

		if entity, err := GeoStateActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoStateCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new geoState, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &GeoStateEntity{}

		for _, item := range GeoStateCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := GeoStateActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var GeoStateUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   GeoStateCommonCliFlagsOptional,
	Usage:   "Updates a geoState by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastGeoStateFromCli(c)

		if entity, err := GeoStateActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var GeoStateCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoState uniqueId (primary key)",
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

var GeoStateCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "name",
		StructField: "Name",
		Required:    false,
		Usage:       "name",
		Type:        "string",
	},
}

var GeoStateCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "GeoState uniqueId (primary key)",
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

func CastGeoStateFromCli(c *cli.Context) *GeoStateEntity {
	geoState := &GeoStateEntity{}

	if c.IsSet("id") {
		geoState.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		geoState.ParentId = &x
	}

	if c.IsSet("name") || false {
		value := c.String("name")
		geoState.Name = &value
	}

	return geoState
}

func GeoStateSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		GeoStateActionCreate,
		reflect.ValueOf(&GeoStateEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func GeoStateWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := GeoStateActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "GeoState", result)
	}
}

var GeoStateWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire geoStates ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := GeoStateActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var GeoStateImportExportCommands = []cli.Command{
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
			GeoStateActionSeeder(query, c.Int("count"))

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
				Value: "geoState-seeder.yml",
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
			GeoStateActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "geoState-seeder-geoState.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of geoStates, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]GeoStateEntity{}
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
				GeoStateActionQuery,
				reflect.ValueOf(&GeoStateEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"GeoStateFieldMap.yml",
				GeoStatePreloadRelations,
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
				GeoStateActionCreate,
				reflect.ValueOf(&GeoStateEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var GeoStateCliCommands []cli.Command = []cli.Command{

	GeoStateCreateCmd,
	GeoStateUpdateCmd,
	GeoStateCreateInteractiveCmd,

	GeoStateWipeCmd,

	workspaces.GetCommonQuery(GeoStateActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&GeoStateEntity{}).Elem(), GeoStateActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&GeoStateEntity{}).Elem(), GeoStateActionRemove),
}

func GeoStateCliFn() cli.Command {
	GeoStateCliCommands = append(GeoStateCliCommands, GeoStateImportExportCommands...)

	return cli.Command{
		Name:        "state",
		Description: "GeoStates module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the geoStates module (" + fmt.Sprintf("%v", len(GeoStateCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: GeoStateCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var GeoStatePreloadRelations []string = []string{}
