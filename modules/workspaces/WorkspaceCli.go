package workspaces

import (
	"fmt"
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
			ActionRequires: strings.Split(c.String("capabilities"), ","),
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
	Action: func(c *cli.Context) error {
		f := CommonCliQueryDSLBuilder(c)

		RunTests(f)

		return nil
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
			Name:     "uid",
			Required: true,
			Usage:    "User id that you want to act on their behalf",
		},
	},
	Action: func(c *cli.Context) error {
		wid := c.String("wid")
		uid := c.String("uid")
		cfg := GetAppConfig()
		cfg.WorkspaceAs = wid
		cfg.UserAs = uid
		WriteAppConfig(cfg)
		fmt.Println("Set workspace to:", wid, "and user", uid)
		return nil
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
	},
	Action: func(c *cli.Context) error {
		if c.IsSet("lang") {
			ws := c.String("lang")
			cfg := GetAppConfig()
			cfg.CliLanguage = ws
			WriteAppConfig(cfg)
			fmt.Println("Cli response language has been changed to:", ws)
		}

		return nil
	},
}

func init() {
	WorkspaceCliCommands = append(
		WorkspaceCliCommands,
		GetUserAccessScope,
		CliConfigCmd,
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
