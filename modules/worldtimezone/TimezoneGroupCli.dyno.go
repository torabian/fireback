//go:build !omit_cli

package worldtimezone

import (
	"embed"
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/urfave/cli"

	"github.com/torabian/fireback/modules/workspaces"
	seeders "github.com/torabian/fireback/modules/worldtimezone/seeders/TimezoneGroup"

	metas "github.com/torabian/fireback/modules/worldtimezone/metas"
)

var TimezoneGroupCreateCmd cli.Command = cli.Command{

	Name:    "create",
	Aliases: []string{"c"},
	Flags:   TimezoneGroupCommonCliFlags,
	Usage:   "Create a new timezoneGroup",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastTimezoneGroupFromCli(c)

		if entity, err := TimezoneGroupActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var TimezoneGroupCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new timezoneGroup, using requied fields in an interactive name",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)

		entity := &TimezoneGroupEntity{}

		for _, item := range TimezoneGroupCommonInteractiveCliFlags {

			if !item.Required && c.Bool("all") == false {
				continue
			}

			result := workspaces.AskForInput(item.Name, "")

			workspaces.SetFieldString(entity, item.StructField, result)

		}

		if entity, err := TimezoneGroupActionCreate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

var TimezoneGroupUpdateCmd cli.Command = cli.Command{

	Name:    "update",
	Aliases: []string{"u"},
	Flags:   TimezoneGroupCommonCliFlagsOptional,
	Usage:   "Updates a timezoneGroup by passing the parameters",
	Action: func(c *cli.Context) error {

		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastTimezoneGroupFromCli(c)

		if entity, err := TimezoneGroupActionUpdate(query, entity); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}

		return nil
	},
}

var TimezoneGroupCommonCliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "TimezoneGroup uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.StringFlag{
		Name:     "value",
		Required: false,
		Usage:    "value",
	},

	&cli.StringFlag{
		Name:     "abbr",
		Required: false,
		Usage:    "abbr",
	},

	&cli.Int64Flag{
		Name:     "offset",
		Required: false,
		Usage:    "offset",
	},

	&cli.BoolFlag{
		Name:     "isdst",
		Required: false,
		Usage:    "isdst",
	},

	&cli.StringFlag{
		Name:     "text",
		Required: false,
		Usage:    "text",
	},

	&cli.StringSliceFlag{
		Name:     "utcItems",
		Required: false,
		Usage:    "utcItems",
	},
}

var TimezoneGroupCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{

	{
		Name:        "value",
		StructField: "Value",
		Required:    false,
		Usage:       "value",
		Type:        "string",
	},

	{
		Name:        "abbr",
		StructField: "Abbr",
		Required:    false,
		Usage:       "abbr",
		Type:        "string",
	},

	{
		Name:        "offset",
		StructField: "Offset",
		Required:    false,
		Usage:       "offset",
		Type:        "int64",
	},

	{
		Name:        "isdst",
		StructField: "Isdst",
		Required:    false,
		Usage:       "isdst",
		Type:        "bool",
	},

	{
		Name:        "text",
		StructField: "Text",
		Required:    false,
		Usage:       "text",
		Type:        "string",
	},
}

var TimezoneGroupCommonCliFlagsOptional = []cli.Flag{
	&cli.StringFlag{
		Name:     "workspaceId",
		Required: false,
		Usage:    "Provide workspace id, if you want to change the data workspace",
	},
	&cli.StringFlag{
		Name:     "uniqueId",
		Required: false,
		Usage:    "TimezoneGroup uniqueId (primary key)",
	},
	&cli.StringFlag{
		Name:     "parentId",
		Required: false,
		Usage:    " Parent record id of the same type",
	},

	&cli.StringFlag{
		Name:     "value",
		Required: false,
		Usage:    "value",
	},

	&cli.StringFlag{
		Name:     "abbr",
		Required: false,
		Usage:    "abbr",
	},

	&cli.Int64Flag{
		Name:     "offset",
		Required: false,
		Usage:    "offset",
	},

	&cli.BoolFlag{
		Name:     "isdst",
		Required: false,
		Usage:    "isdst",
	},

	&cli.StringFlag{
		Name:     "text",
		Required: false,
		Usage:    "text",
	},

	&cli.StringSliceFlag{
		Name:     "utcItems",
		Required: false,
		Usage:    "utcItems",
	},
}

func CastTimezoneGroupFromCli(c *cli.Context) *TimezoneGroupEntity {
	timezoneGroup := &TimezoneGroupEntity{}

	if c.IsSet("id") {
		timezoneGroup.UniqueId = c.String("id")
	}

	if c.IsSet("parentId") {
		x := c.String("parentId")
		timezoneGroup.ParentId = &x
	}

	if c.IsSet("value") || false {
		value := c.String("value")
		timezoneGroup.Value = &value
	}

	if c.IsSet("abbr") || false {
		value := c.String("abbr")
		timezoneGroup.Abbr = &value
	}

	if c.IsSet("offset") || false {
		val := c.Int64("offset")
		timezoneGroup.Offset = &val
	}

	if c.IsSet("isdst") || false {
		value := c.Bool("isdst")
		timezoneGroup.Isdst = &value
	}

	if c.IsSet("text") || false {
		value := c.String("text")
		timezoneGroup.Text = &value
	}

	if c.IsSet("utcItems") {
		timezoneGroup.UtcItems = []*TimezoneGroupUtcItemsEntity{}
		for _, item := range c.StringSlice("utcItems") {

			// if only one field, and it's string, then just add string
			timezoneGroup.UtcItems = append(timezoneGroup.UtcItems, &TimezoneGroupUtcItemsEntity{
				Name: &item,
			})

		}
	}

	return timezoneGroup
}

func TimezoneGroupSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		TimezoneGroupActionCreate,
		reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
		fsRef,
		fileNames,
		true,
	)
}

func TimezoneGroupSyncSeeders() {
	workspaces.SeederFromFSImport(
		workspaces.QueryDSL{},
		TimezoneGroupActionCreate,
		reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

func TimezoneGroupWriteQueryMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := TimezoneGroupActionQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "TimezoneGroup", result)
	}
}

var TimezoneGroupWipeCmd cli.Command = cli.Command{

	Name:  "wipe",
	Usage: "Wipes entire timezoneGroups ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := TimezoneGroupActionWipeClean(query)

		fmt.Println("Removed", count, "of entities")

		return nil
	},
}

var TimezoneGroupImportExportCommands = []cli.Command{
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
			TimezoneGroupActionSeeder(query, c.Int("count"))

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
				Value: "timezoneGroup-seeder.yml",
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
			TimezoneGroupActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "timezoneGroup-seeder-timezoneGroup.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of timezoneGroups, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {

			data := &[]TimezoneGroupEntity{}
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
				TimezoneGroupActionCreate,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
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
				TimezoneGroupActionQuery,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
				c.String("file"),
				&metas.MetaFs,
				"TimezoneGroupFieldMap.yml",
				TimezoneGroupPreloadRelations,
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
				TimezoneGroupActionCreate,
				reflect.ValueOf(&TimezoneGroupEntity{}).Elem(),
				c.String("file"),
			)

			return nil
		},
	},
}

var TimezoneGroupCliCommands []cli.Command = []cli.Command{

	TimezoneGroupCreateCmd,
	TimezoneGroupUpdateCmd,
	TimezoneGroupCreateInteractiveCmd,

	TimezoneGroupWipeCmd,

	workspaces.GetCommonQuery(TimezoneGroupActionQuery),

	workspaces.GetCommonTableQuery(reflect.ValueOf(&TimezoneGroupEntity{}).Elem(), TimezoneGroupActionQuery),

	workspaces.GetCommonRemoveQuery(reflect.ValueOf(&TimezoneGroupEntity{}).Elem(), TimezoneGroupActionRemove),
}

func TimezoneGroupCliFn() cli.Command {
	TimezoneGroupCliCommands = append(TimezoneGroupCliCommands, TimezoneGroupImportExportCommands...)

	return cli.Command{
		Name:        "tz",
		Description: "TimezoneGroups module actions (sample module to handle complex entities)",
		Usage:       "Actions related to the timezoneGroups module (" + fmt.Sprintf("%v", len(TimezoneGroupCliCommands)) + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language",
				Value: "en",
			},
		},
		Subcommands: TimezoneGroupCliCommands,
	}
}

// At this moment, we do not detect this automatically yet. Append to this in the cli
var TimezoneGroupPreloadRelations []string = []string{}
