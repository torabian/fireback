package workspaces

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var CreateRootUser cli.Command = cli.Command{
	Name:  "new",
	Usage: "Creates a user interactively, and sets that credential into the workspace config",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "in-root",
			Usage: "Append this user to root group",
		},
		&cli.StringFlag{
			Name:  "value",
			Usage: "value",
		},
		&cli.StringFlag{
			Name:  "workspace-type-id",
			Usage: "The workspace type id, you can use 'root'",
		},
		&cli.StringFlag{
			Name:  "session-secrent",
			Usage: "The secret generated through the otp authentication process.",
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "One of: 'phonenumber', 'email'",
		},
		&cli.StringFlag{
			Name:  "password",
			Usage: "password",
		},
		&cli.StringFlag{
			Name:  "first-name",
			Usage: "firstName",
		},
		&cli.StringFlag{
			Name:  "last-name",
			Usage: "lastName",
		},
	},
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)

		if c.NumFlags() == 0 {
			// This is gonna be an interactive, there are no flags
			if err := InteractiveUserAdmin(query); err != nil {
				fmt.Println(err)
			}

		} else {
			dto := CastClassicSignupFromCli(c)
			if err := CreateAdminTransaction(dto, c.Bool("in-root"), query); err != nil {
				fmt.Println(err)
			}
		}

	},
}

var AuthorizeUserInteractively cli.Command = cli.Command{
	Name:  "auth",
	Usage: "Signins the user with passport and password, and stores the credentials in env for cli usage",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "value",
			Usage: "value",
		},
		&cli.StringFlag{
			Name:  "password",
			Usage: "password",
		},
		&cli.StringFlag{
			Name:  "wid",
			Usage: "Workspace id, if user is in multiple workspaces to change the cli. It will ask interactively if not set.",
		},
	},
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
		var session *ClassicSigninActionResDto
		preferedWorkspaceId := c.String("wid")
		var err *IError = nil
		if c.NumFlags() == 0 {

			dto := &ClassicSigninActionReqDto{}
			if result := AskForInput("Passport (email, phonenumber,...)", "admin"); result != "" {
				dto.Value = &result
			}

			if result := AskForInput("Password", "admin"); result != "" {
				dto.Password = &result
			}

			// This is gonna be an interactive, there are no flags
			session, err = ClassicSigninAction(dto, query)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			dto := CastClassicSigninFromCli(c)
			session, err = ClassicSigninAction(dto, query)
			if err != nil {
				fmt.Println(err)
			}
		}

		if session == nil {
			log.Fatal("Session could not be retrieved, no change to cli token or workspaces made.")
			return
		}
		workspaces := []string{}
		for _, item := range session.Session.UserWorkspaces {
			workspaces = append(workspaces, *item.WorkspaceId)
		}
		if len(workspaces) > 1 && preferedWorkspaceId == "" {
			preferedWorkspaceId = AskForSelect("Which workspaces to be selected?", workspaces)
		} else if len(workspaces) == 1 {
			preferedWorkspaceId = workspaces[0]
		}
		config.CliWorkspace = preferedWorkspaceId
		config.CliToken = *session.Session.Token
		config.Save(".env")

	},
}

var AuthorizeOsCmd cli.Command = cli.Command{
	Name:  "os",
	Usage: "Authorizes the user, as os owner. Useful for desktop offline apps or mobile apps",

	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
		result, err := PassportActionAuthorizeOs2(&EmptyRequest{}, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}

var AppendEmailPassportToUser cli.Command = cli.Command{

	Name:  "append-email",
	Usage: "Appends a new passport to an specific user, given by userid in the system",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "email",
			Usage:    "E-mail address",
			Value:    "",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "Password",
			Value:    "",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user-id",
			Usage:    "Userid, which the passport will be append to",
			Value:    "",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		f := QueryDSL{
			UserId: c.String("user-id"),
		}

		email := c.String("email")
		password := c.String("password")

		session, err := PassportAppendEmailToUser(&ClassicAuthDto{
			Value: &email, Password: &password,
		}, f)

		if err != nil {
			log.Fatal(err)
		} else {
			out, err := yaml.Marshal(session)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(out))
		}

		return nil
	},
}

var PassportCli cli.Command = cli.Command{
	Name:  "passport",
	Usage: "Manage the methods of authentication in the app, as well as users passports (root only)",
	Subcommands: append([]cli.Command{
		AppendEmailPassportToUser,
		PassportUpdateCmd,
		AuthorizeOsCmd,
		CreateRootUser,
		AuthorizeUserInteractively,
		PassportMethodCliFn(),
		CheckPassportMethodsActionCmd,
		PassportWipeCmd,
		PassportUpdateCmd,
		GetCommonQuery(PassportActionQuery),
	}, WorkspacesCustomActionsCli...),
}
