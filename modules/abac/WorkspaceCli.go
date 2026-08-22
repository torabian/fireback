//go:build !wasm

package abac

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/interfacetools"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
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
	Action: func(ctx context.Context, c *cli.Command) error {
		query := fireback.CommonCliQueryDSLBuilder(c)
		query.UserId = c.String("id")
		fmt.Println(GetUserAccessLevels(query))
		// fireback.HandleActionInCli(c, access, err, map[string]map[string]string{})

		return nil
	},
}

func PermissionInfoFromString(items []string) []application.PermissionInfo {
	res := []application.PermissionInfo{}

	for _, item := range items {
		res = append(res, application.PermissionInfo{
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
	Action: func(ctx context.Context, c *cli.Command) error {
		f := fireback.CommonCliQueryDSLBuilder(c)
		f.UserId = c.String("id")
		access, _ := GetUserAccessLevels(f)

		query := fireback.QueryDSL{
			UserAccessPerWorkspace: access.UserAccessPerWorkspace,
			ActionRequires:         PermissionInfoFromString(strings.Split(c.String("capabilities"), ",")),
		}

		meets, missing := fireback.MeetsAccessLevel(query, false)

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
	Action: func(ctx context.Context, c *cli.Command) error {

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
	Action: func(ctx context.Context, c *cli.Command) error {
		appConfig := fireback.GetConfig()
		wid := c.String("wid")
		token := c.String("token")
		appConfig.CliWorkspace = wid
		appConfig.CliToken = token
		appConfig.Save(".env")
		fmt.Println("Set workspace to:", wid, "and token", token)
		return nil
	},
}

var ViewAuthorize cli.Command = cli.Command{

	Name:  "view",
	Usage: "Shows the authorization result for current user",

	Action: func(ctx context.Context, c *cli.Command) error {
		appConfig := fireback.GetConfig()
		fmt.Println("Workspace:", appConfig.CliWorkspace)
		fmt.Println("Token:", appConfig.CliToken)

		result, err := fireback.CliAuth(nil)
		if err != nil {
			log.Fatalln(err)
		}
		result.JsonPrint()

		// Explicit untyped nil, not `return err`: err is a *fireback.IError, and
		// by this point it's a nil *pointer* - wrapping that in the `error`
		// interface return value produces a non-nil interface (the classic Go
		// typed-nil gotcha), which made urfave/cli treat every successful `ws
		// view` as a failed command and made FirebackApp.go's log.Fatal(err)
		// call Error() on a nil *IError, printing "null" and exiting 1.
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
		&cli.StringFlag{
			Name:     "region",
			Required: false,
			Usage:    "Sets the default region in the entire cli context",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		appConfig := fireback.GetConfig()
		if c.IsSet("lang") {
			ws := c.String("lang")
			appConfig.CliLanguage = ws
			fmt.Println("Cli response language has been changed to:", ws)
		}
		if c.IsSet("region") {
			ws := c.String("region")
			appConfig.CliRegion = ws
			fmt.Println("Cli region has been changed to:", ws)
		}

		appConfig.Save(".env")

		return nil
	},
}

var MiscCli cli.Command = cli.Command{

	Name:  "misc",
	Usage: "Managing the application related content, thirdparty configs such as email, sms, or ui data",
	Commands: []*cli.Command{
		&cli.Command{
			Name:        "regionalcontent",
			Aliases:     []string{"rc"},
			Description: `Email templates, sms templates or other textual content which can be accessed.`,
			Usage:       `Email templates, sms templates or other textual content which can be accessed.`,
			Commands: []*cli.Command{
				abacdefs.RegionalContentBrowseActionCliHandler(RegionalContentBrowseAction),
				abacdefs.RegionalContentGetActionCliHandler(RegionalContentGetAction),
				abacdefs.RegionalContentCreateActionCliHandler(RegionalContentCreateAction),
				abacdefs.RegionalContentUpdateActionCliHandler(RegionalContentUpdateAction),
				abacdefs.RegionalContentAwareDeletePreviewActionCliHandler(RegionalContentAwareDeletePreviewAction),
				abacdefs.RegionalContentAwareDeleteActionCliHandler(RegionalContentAwareDeleteAction),
				&RegionalContentGetCmd,
			},
		},
		interfacetools.AppMenuCliFn(),
		getCssMinCombineCli(),
		&cli.Command{
			Name:        "totp",
			Description: "Generates a time based code (6 digit) from a totp secret, simulates the mobile app",
			Usage:       "Generates a time based code (6 digit) from a totp secret, simulates the mobile app",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "secret",
					Value:    "",
					Usage:    "The secret from the totp system which is given.",
					Required: true,
				},
			},
			Action: func(ctx context.Context, c *cli.Command) error {
				secret := c.String("secret")

				for {
					now := time.Now()

					code, err := totp.GenerateCode(secret, now)
					if err != nil {
						return err
					}

					remaining := 30 - now.Second()%30

					fmt.Printf("\rCode: %s (expires in %2ds)", code, remaining)

					time.Sleep(1 * time.Second)
				}
			},
		},
	},
}

var WorkspaceCliCommands = []*cli.Command{
	&GetUserAccessScope,
	&CliConfigCmd,
	&ViewAuthorize,
	abacdefs.QueryWorkspaceTypesPubliclyActionCliHandler(QueryWorkspaceTypesPubliclyAction),
	abacdefs.QueryUserRoleWorkspacesActionCliHandler(QueryUserRoleWorkspacesAction),
	&CheckUserMeetsAPermissionCmd,
	&WorkspaceAsCmd,
	&cli.Command{
		Name:        "publicauthentication",
		Aliases:     []string{"pa"},
		Description: `Keeps information about user onboarding, otp state, and other things which are necessary for onboarding new users in multiple endpoints`,
		Usage:       `Keeps information about user onboarding, otp state, and other things which are necessary for onboarding new users in multiple endpoints`,
		Commands: []*cli.Command{
			abacdefs.PublicAuthenticationBrowseActionCliHandler(PublicAuthenticationBrowseAction),
			abacdefs.PublicAuthenticationGetActionCliHandler(PublicAuthenticationGetAction),
			abacdefs.PublicAuthenticationCreateActionCliHandler(PublicAuthenticationCreateAction),
			abacdefs.PublicAuthenticationUpdateActionCliHandler(PublicAuthenticationUpdateAction),
			abacdefs.PublicAuthenticationAwareDeletePreviewActionCliHandler(PublicAuthenticationAwareDeletePreviewAction),
			abacdefs.PublicAuthenticationAwareDeleteActionCliHandler(PublicAuthenticationAwareDeleteAction),
		},
	},

	&cli.Command{
		Name:        "type",
		Aliases:     []string{"wt"},
		Description: `Workspace types - the role/capability template a workspace gets created with.`,
		Usage:       `Workspace types - the role/capability template a workspace gets created with.`,
		Commands: []*cli.Command{
			abacdefs.WorkspaceTypeBrowseActionCliHandler(WorkspaceTypeBrowseAction),
			abacdefs.WorkspaceTypeGetActionCliHandler(WorkspaceTypeGetAction),
			abacdefs.WorkspaceTypeCreateActionCliHandler(WorkspaceTypeCreateAction),
			abacdefs.WorkspaceTypeUpdateActionCliHandler(WorkspaceTypeUpdateAction),
			abacdefs.WorkspaceTypeAwareDeletePreviewActionCliHandler(WorkspaceTypeAwareDeletePreviewAction),
			abacdefs.WorkspaceTypeAwareDeleteActionCliHandler(WorkspaceTypeAwareDeleteAction),
		},
	},
	&cli.Command{
		Name:        "config",
		Aliases:     []string{"wc"},
		Description: `Application-wide configuration (recaptcha, otp enforcement, and so on) - a single record, root workspace only.`,
		Usage:       `Application-wide configuration (recaptcha, otp enforcement, and so on) - a single record, root workspace only.`,
		Commands: []*cli.Command{
			abacdefs.WorkspaceConfigBrowseActionCliHandler(WorkspaceConfigBrowseAction),
			abacdefs.WorkspaceConfigGetActionCliHandler(WorkspaceConfigGetAction),
			abacdefs.WorkspaceConfigCreateActionCliHandler(WorkspaceConfigCreateAction),
			abacdefs.WorkspaceConfigUpdateActionCliHandler(WorkspaceConfigUpdateAction),
			abacdefs.WorkspaceConfigAwareDeletePreviewActionCliHandler(WorkspaceConfigAwareDeletePreviewAction),
			abacdefs.WorkspaceConfigAwareDeleteActionCliHandler(WorkspaceConfigAwareDeleteAction),
		},
	},
	&cli.Command{
		Name:        "invite",
		Aliases:     []string{"wi"},
		Description: `Pending invitations for someone to join a workspace under a given role.`,
		Usage:       `Pending invitations for someone to join a workspace under a given role.`,
		Commands: []*cli.Command{
			abacdefs.WorkspaceInviteBrowseActionCliHandler(WorkspaceInviteBrowseAction),
			abacdefs.WorkspaceInviteGetActionCliHandler(WorkspaceInviteGetAction),
			abacdefs.WorkspaceInviteCreateActionCliHandler(WorkspaceInviteCreateAction),
			abacdefs.WorkspaceInviteUpdateActionCliHandler(WorkspaceInviteUpdateAction),
			abacdefs.WorkspaceInviteAwareDeletePreviewActionCliHandler(WorkspaceInviteAwareDeletePreviewAction),
			abacdefs.WorkspaceInviteAwareDeleteActionCliHandler(WorkspaceInviteAwareDeleteAction),
		},
	},
	&cli.Command{
		Name:        "role",
		Aliases:     []string{"wr"},
		Description: `Roles scoped to a specific workspace (as opposed to a role usable across workspaces).`,
		Usage:       `Roles scoped to a specific workspace (as opposed to a role usable across workspaces).`,
		Commands: []*cli.Command{
			abacdefs.WorkspaceRoleBrowseActionCliHandler(WorkspaceRoleBrowseAction),
			abacdefs.WorkspaceRoleGetActionCliHandler(WorkspaceRoleGetAction),
			abacdefs.WorkspaceRoleCreateActionCliHandler(WorkspaceRoleCreateAction),
			abacdefs.WorkspaceRoleUpdateActionCliHandler(WorkspaceRoleUpdateAction),
			abacdefs.WorkspaceRoleAwareDeletePreviewActionCliHandler(WorkspaceRoleAwareDeletePreviewAction),
			abacdefs.WorkspaceRoleAwareDeleteActionCliHandler(WorkspaceRoleAwareDeleteAction),
		},
	},
	&cli.Command{
		Name:        "user-workspace",
		Aliases:     []string{"uw"},
		Description: `Membership rows tying a user to a workspace (and, through it, to a role).`,
		Usage:       `Membership rows tying a user to a workspace (and, through it, to a role).`,
		Commands: []*cli.Command{
			abacdefs.UserWorkspaceBrowseActionCliHandler(UserWorkspaceBrowseAction),
			abacdefs.UserWorkspaceGetActionCliHandler(UserWorkspaceGetAction),
			abacdefs.UserWorkspaceCreateActionCliHandler(UserWorkspaceCreateAction),
			abacdefs.UserWorkspaceUpdateActionCliHandler(UserWorkspaceUpdateAction),
			abacdefs.UserWorkspaceAwareDeletePreviewActionCliHandler(UserWorkspaceAwareDeletePreviewAction),
			abacdefs.UserWorkspaceAwareDeleteActionCliHandler(UserWorkspaceAwareDeleteAction),
		},
	},
	&cli.Command{
		Name:        "public-join-key",
		Aliases:     []string{"pjk"},
		Description: `Shareable links that let anyone join a workspace without an explicit per-person invite.`,
		Usage:       `Shareable links that let anyone join a workspace without an explicit per-person invite.`,
		Commands: []*cli.Command{
			abacdefs.PublicJoinKeyBrowseActionCliHandler(PublicJoinKeyBrowseAction),
			abacdefs.PublicJoinKeyGetActionCliHandler(PublicJoinKeyGetAction),
			abacdefs.PublicJoinKeyCreateActionCliHandler(PublicJoinKeyCreateAction),
			abacdefs.PublicJoinKeyUpdateActionCliHandler(PublicJoinKeyUpdateAction),
			abacdefs.PublicJoinKeyAwareDeletePreviewActionCliHandler(PublicJoinKeyAwareDeletePreviewAction),
			abacdefs.PublicJoinKeyAwareDeleteActionCliHandler(PublicJoinKeyAwareDeleteAction),
		},
	},
}

// WorkspaceCliFn mirrors the old Module3-generated grouped "workspace" cli command
// (minus the import/export/dev commands, which had no hand-written equivalent to
// recover), plus the "cte" subcommand for the recursive tree query. WorkspaceCliCommands
// (see WorkspaceCli.go) carries every entity-scoped cli group that doesn't have its own
// top-level command elsewhere (publicAuthentication, timezoneGroup, workspaceType,
// workspaceConfig, workspaceInvite, workspaceRole, userWorkspace, publicJoinKey, ...).
func WorkspaceCliFn() *cli.Command {
	commands := []*cli.Command{
		abacdefs.WorkspaceBrowseActionCliHandler(WorkspaceBrowseAction),
		abacdefs.WorkspaceGetActionCliHandler(WorkspaceGetAction),
		abacdefs.WorkspaceCreateActionCliHandler(WorkspaceCreateAction),
		abacdefs.WorkspaceUpdateActionCliHandler(WorkspaceUpdateAction),
		abacdefs.WorkspaceAwareDeletePreviewActionCliHandler(WorkspaceAwareDeletePreviewAction),
		abacdefs.WorkspaceAwareDeleteActionCliHandler(WorkspaceAwareDeleteAction),
	}
	commands = append(commands, WorkspaceCliCommands...)
	return &cli.Command{
		Name:        "workspace",
		Aliases:     []string{"ws"},
		Description: `Fireback general user role, workspaces services.`,
		Usage:       `Fireback general user role, workspaces services.`,
		Commands:    commands,
	}
}
