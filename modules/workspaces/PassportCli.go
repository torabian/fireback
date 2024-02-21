package workspaces

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var UserWithPassportCreateInteractiveCmd cli.Command = cli.Command{
	Name:  "ic",
	Usage: "Creates a new user in the system, using an interactive question builder, and adding a passport to it",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "Interactively asks for all inputs, not only required ones",
		},
	},
	Action: func(c *cli.Context) {
		user := InteractiveCreateUserInCli()

		email := AskForInput("Enter the email address (primary passport)", "")
		password := AskForPassword("Enter the password", "")

		CreateEmailPassportForUser(user.UniqueId, email, password, nil)

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
		UserWithPassportCreateInteractiveCmd,
		AppendEmailPassportToUser,
		PassportUpdateCmd,
		PassportMethodCliFn(),
		PassportWipeCmd,
		PassportUpdateCmd,
		GetCommonQuery(PassportActionQuery),
	}, WorkspacesCustomActionsCli...),
}
