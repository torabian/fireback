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

var WorkspaceRemoveCmd cli.Command = cli.Command{

	Name:    "remove",
	Aliases: []string{"r", "del", "delete"},
	Usage:   "Deletes a workspace",
	Action: func(c *cli.Context) error {

		fmt.Printf("Delete workspace")

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
			Name:     "token",
			Required: true,
			Usage:    "Selected token that you are using as authorization her in cli",
		},
	},
	Action: func(c *cli.Context) error {
		wid := c.String("wid")
		token := c.String("token")
		config.CliWorkspace = wid
		config.CliToken = token
		config.Save(".env")
		fmt.Println("Set workspace to:", wid, "and token", token)
		return nil
	},
}

var ViewAuthorize cli.Command = cli.Command{

	Name:  "view",
	Usage: "Shows the authorization result for current user",

	Action: func(c *cli.Context) error {

		fmt.Println("Workspace:", config.CliWorkspace)
		fmt.Println("Token:", config.CliToken)

		result, err := CliAuth(nil)
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
		if c.IsSet("lang") {
			ws := c.String("lang")
			config.CliLanguage = ws
			fmt.Println("Cli response language has been changed to:", ws)
		}
		if c.IsSet("region") {
			ws := c.String("region")
			config.CliRegion = ws
			fmt.Println("Cli region has been changed to:", ws)
		}

		config.Save(".env")

		return nil
	},
}

var MiscCli cli.Command = cli.Command{

	Name:  "misc",
	Usage: "Managing the application related content, thirdparty configs such as email, sms, or ui data",
	Subcommands: []cli.Command{
		TableViewSizingCliFn(),
		RegionalContentCliFn(),
		AppMenuCliFn(),
	},
}

func init() {
	WorkspaceCliCommands = append(
		WorkspaceCliCommands,
		GetUserAccessScope,
		CliConfigCmd,
		ViewAuthorize,
		QueryWorkspaceTypesPubliclyActionCmd,
		CheckUserMeetsAPermissionCmd,
		WorkspaceAsCmd,
		PublicAuthenticationCliFn(),
		TimezoneGroupCliFn(),
		WorkspaceTypeCliFn(),
		WorkspaceConfigCliFn(),
		WorkspaceInviteCliFn(),
		WorkspaceRoleCliFn(),
		UserWorkspaceCliFn(),
		WorkspaceInviteCliFn(),
		PublicJoinKeyCliFn(),
	)

}
