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
				log.Fatalln(err)
			}

		} else {
			dto := CastClassicSignupFromCli(c)
			if err := CreateAdminTransaction(dto, c.Bool("in-root"), query); err != nil {
				log.Fatalln(err)
			}
		}

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
			Value: email, Password: password,
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
		PassportMethodCliFn(),
		CheckPassportMethodsActionCmd,
		PassportWipeCmd,
		PassportUpdateCmd,
		GetCommonQuery(PassportActionQuery),
	}, WorkspacesCustomActionsCli...),
}
