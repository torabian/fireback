package abac

import (
	"fmt"
	"log"
	"reflect"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

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

		f := fireback.QueryDSL{
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
		OsLoginAuthenticateActionDef.ToCli(),
		PassportMethodCliFn(),
		CheckPassportMethodsActionDef.ToCli(),
		UserPassportsActionDef.ToCli(),
		OauthAuthenticateActionDef.ToCli(),
		PassportWipeCmd,
		PassportUpdateCmd,
		fireback.GetCommonRemoveQuery(
			reflect.ValueOf(&PassportEntity{}).Elem(),
			PassportActions.RemoveEnqueue,
		),
		PASSPORT_ACTION_QUERY.ToCli(),
	}, fireback.FirebackCustomActionsCli...),
}
