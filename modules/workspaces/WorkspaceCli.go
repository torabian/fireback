package workspaces

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"
)

var GetUserAccessScope cli.Command = cli.Command{

	Name:  "scope",
	Usage: "Returns the access level, roles, and scopes that an specific user has access to",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Value:    "",
			Required: true,
			Usage:    "User specific id",
		},
	},
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		query.UserId = c.String("id")
		access, err := GetUserAccessLevels(query)
		HandleActionInCli(c, access, err, map[string]map[string]string{})

		return nil
	},
}
var LspCmd cli.Command = cli.Command{

	Name:  "lsp",
	Usage: "Runs the language server",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {

		return BeginLspServer(c)

	},
}

func PermissionInfoFromString(items []string) []PermissionInfo {
	res := []PermissionInfo{}

	for _, item := range items {
		res = append(res, PermissionInfo{
			CompleteKey: item,
		})
	}

	return res
}

var CheckUserMeetsAPermissionCmd cli.Command = cli.Command{

	Name:  "meets",
	Usage: "By given a user id, to will check if user has the capabilities asked for",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Value:    "",
			Required: true,
			Usage:    "User specific id",
		},
		&cli.StringFlag{
			Name:     "capabilities",
			Value:    "",
			Required: true,
			Usage:    "Capabilities list, separated by , aka: ROOT_BOOKS_CREATE,ROOT_BOOKS_DELETE",
		},
	},
	Action: func(c *cli.Context) error {
		f := CommonCliQueryDSLBuilder(c)
		f.UserId = c.String("id")
		access, _ := GetUserAccessLevels(f)

		query := QueryDSL{
			UserHas:        access.Capabilities,
			ActionRequires: PermissionInfoFromString(strings.Split(c.String("capabilities"), ",")),
		}

		meets, missing := MeetsAccessLevel(query, false)

		if !meets {
			fmt.Println("Not enough access level. Missing:")
			fmt.Println(strings.Join(missing, ","))
		} else {
			fmt.Println("User has access :)")
		}

		return nil
	},
}

var ConfigWorkspaceCmd cli.Command = cli.Command{

	Name:  "config",
	Usage: "Sets the configuration for an specific workspace",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Value:    "",
			Required: true,
			Usage:    "Workspace id",
		},

		&cli.StringFlag{
			Name:     "zoom-client-id",
			Value:    "",
			Required: false,
			Usage:    "Zoom thirdparty client id",
		},
		&cli.StringFlag{
			Name:     "zoom-client-secret",
			Value:    "",
			Required: false,
			Usage:    "Zoom thirdparty secret",
		},
		&cli.BoolFlag{
			Name:     "allow-public",
			Required: false,
			Usage:    "Allow anonymouse people to signup into the workspace",
		},
	},
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		id := c.String("id")
		config := &WorkspaceConfigEntity{
			WorkspaceId: &id,
		}

		if c.IsSet("zoom-client-id") {
			val := c.String("zoom-client-id")
			config.ZoomClientId = &val
		}

		if c.IsSet("zoom-client-secret") {
			val := c.String("zoom-client-secret")
			config.ZoomClientSecret = &val
		}

		if c.IsSet("allow-public") {
			val := c.Bool("allow-public")
			config.AllowPublicToJoinTheWorkspace = &val
		}

		conf, err := UpdateWorkspaceConfigurationAction(query, config)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(conf)

		return nil
	},
}

var WorkspaceRemoveCmd cli.Command = cli.Command{

	Name:    "remove",
	Aliases: []string{"r", "del", "delete"},
	Usage:   "Deletes a workspace",
	Action: func(c *cli.Context) error {

		fmt.Printf("Delete workspace")

		return nil
	},
}

var WorkspaceTestCmd cli.Command = cli.Command{

	Name:  "tests",
	Usage: "Tests related to the workspace cli",
	Subcommands: cli.Commands{
		cli.Command{

			Name:  "dbx",
			Usage: "Tests the database integrity",
			Action: func(c *cli.Context) error {
				f := CommonCliQueryDSLBuilder(c)

				RunTests(f)

				return nil
			},
		},
		cli.Command{

			Name:  "new",
			Usage: "Tests the new project generation",
			Action: func(c *cli.Context) error {

				testing := &TestContext{}

				TestRunner(testing, []Test{
					TestNewModuleProjectGen,
				})

				return nil
			},
		},
	},
}

var WorkspaceAsCmd cli.Command = cli.Command{

	Name:  "as",
	Usage: "Set the workspace in terminal",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "wid",
			Required: true,
			Usage:    "Workspace id that you want to act on behalf",
		},
		&cli.StringFlag{
			Name:     "token",
			Required: true,
			Usage:    "Selected token that you are using as authorization her in cli",
		},
	},
	Action: func(c *cli.Context) error {
		wid := c.String("wid")
		token := c.String("token")
		cfg := GetAppConfig()
		cfg.WorkspaceAs = wid
		cfg.Token = token
		WriteAppConfig(cfg)
		fmt.Println("Set workspace to:", wid, "and token", token)
		return nil
	},
}

var ViewAuthorize cli.Command = cli.Command{

	Name:  "view",
	Usage: "Shows the authorization result for current user",

	Action: func(c *cli.Context) error {
		cfg := GetAppConfig()

		fmt.Println("Workspace::", cfg.WorkspaceAs)
		fmt.Println("Token::", cfg.Token)

		result, err := CliAuth()
		if err != nil {
			log.Fatalln(err)
		} else {
			result.JsonPrint()
		}

		return err
	},
}

var CliConfigCmd cli.Command = cli.Command{

	Name:  "cli",
	Usage: "Set some configuration for cli, such as language, region, etc",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lang",
			Required: false,
			Usage:    "Set the language of the cli, does not affect other protocols",
		},
		&cli.StringFlag{
			Name:     "region",
			Required: false,
			Usage:    "Sets the default region in the entire cli context",
		},
	},
	Action: func(c *cli.Context) error {
		cfg := GetAppConfig()
		if c.IsSet("lang") {
			ws := c.String("lang")
			cfg.CliLanguage = ws
			fmt.Println("Cli response language has been changed to:", ws)
		}
		if c.IsSet("region") {
			ws := c.String("region")
			cfg.CliRegion = ws
			fmt.Println("Cli region has been changed to:", ws)
		}
		WriteAppConfig(cfg)

		return nil
	},
}

func init() {
	WorkspaceCliCommands = append(
		WorkspaceCliCommands,
		GetUserAccessScope,
		LspCmd,
		CliConfigCmd,
		ViewAuthorize,
		ConfigWorkspaceCmd,
		CheckUserMeetsAPermissionCmd,
		WorkspaceAsCmd,
		WorkspaceTestCmd,
		WorkspaceTypeCliFn(),
		WorkspaceConfigCliFn(),
		WorkspaceRoleCliFn(),
		UserWorkspaceCliFn(),
	)

}
