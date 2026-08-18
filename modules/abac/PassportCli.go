//go:build !wasm

package abac

import (
	"context"
	"fmt"
	"log"

	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
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
	Action: func(ctx context.Context, c *cli.Command) error {

		f := fireback.QueryDSL{
			UserId: c.String("user-id"),
		}

		email := c.String("email")
		password := c.String("password")

		session, err := PassportAppendEmailToUser(&abacdefs.ClassicAuthDto{
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
	Commands: []*cli.Command{
		&AppendEmailPassportToUser,
		{
			Name: "method",
			Commands: []*cli.Command{
				abacdefs.PassportMethodBrowseActionCliHandler(PassportMethodBrowseAction),
				abacdefs.PassportMethodGetActionCliHandler(PassportMethodGetAction),
				abacdefs.PassportMethodCreateActionCliHandler(PassportMethodCreateAction),
				abacdefs.PassportMethodUpdateActionCliHandler(PassportMethodUpdateAction),
				abacdefs.PassportMethodAwareDeletePreviewActionCliHandler(PassportMethodAwareDeletePreviewAction),
				abacdefs.PassportMethodAwareDeleteActionCliHandler(PassportMethodAwareDeleteAction),
			},
		},
		abacdefs.OsLoginAuthenticateActionCliHandler(OsLoginAuthenticateAction),
		abacdefs.CheckPassportMethodsActionCliHandler(CheckPassportMethodsAction),
		abacdefs.UserPassportsActionCliHandler(UserPassportsAction),
		abacdefs.OauthAuthenticateActionCliHandler(OauthAuthenticateAction),
		abacdefs.PassportBrowseActionCliHandler(PassportBrowseAction),
		abacdefs.PassportGetActionCliHandler(PassportGetAction),
		abacdefs.PassportCreateActionCliHandler(PassportCreateAction),
		abacdefs.PassportUpdateActionCliHandler(PassportUpdateAction),
		abacdefs.PassportAwareDeletePreviewActionCliHandler(PassportAwareDeletePreviewAction),
		abacdefs.PassportAwareDeleteActionCliHandler(PassportAwareDeleteAction),
	},
}
