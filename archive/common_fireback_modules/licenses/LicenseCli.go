package licenses

import (
	"fmt"
	"log"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

var GenerateLicenseCmd cli.Command = cli.Command{

	Name:    "issue-plan",
	Aliases: []string{"pi"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "planId",
			Required: true,
			Usage:    "Plan id, which you will will be redeemed",
		},
		&cli.StringFlag{
			Name:     "mid",
			Required: true,
			Usage:    "User machine id, which will be issuing the license for",
		},

		&cli.StringFlag{
			Name:  "email",
			Usage: "User email address",
		},
		&cli.StringFlag{
			Name:  "owner",
			Usage: "License owner, you can type the full name of user, or their business",
		},
		&cli.StringFlag{
			Name:  "workspace-id",
			Usage: "User workspace-id that you are issuing the license for",
		},
		&cli.StringFlag{
			Name:  "user-id",
			Usage: "User id that you are issuing the license for",
		},
	},
	Usage: "Generates a new license for given product id",
	Action: func(c *cli.Context) error {

		Email := c.String("email")
		Owner := c.String("owner")
		MachineId := c.String("mid")
		dto := &LicenseFromPlanIdDto{
			Email:     &Email,
			Owner:     &Owner,
			MachineId: &MachineId,
		}

		query := workspaces.CommonCliQueryDSLBuilder(c)
		query.UniqueId = c.String("planId")
		query.WorkspaceId = c.String("workspace-id")
		query.UserId = c.String("user-id")

		if license, err := LicenseActionFromPlanId(dto, query); err == nil {
			fmt.Println(license.SignedLicense)
		} else {
			log.Fatalln(err)
		}

		return nil
	},
}

var LicenseFromActivationKeyCmd cli.Command = cli.Command{

	Name:    "issue-key",
	Aliases: []string{"ki"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "mid",
			Required: true,
			Usage:    "User machine id, which will be issuing the license for",
		},
		&cli.StringFlag{
			Name:     "id",
			Required: true,
			Usage:    "The activation key id to be used for generating the license",
		},
	},
	Usage: "Creates a new license, by getting information from activation key, and mark that key as used",
	Action: func(c *cli.Context) {
		// f := workspaces.CommonCliQueryDSLBuilder(c)

		// license := LicenseFromActivationKeyDto{}

		// if c.IsSet("id") {
		// 	license.ActivationKeyId = c.String("id")
		// }

		// if c.IsSet("mid") {
		// 	license.MachineId = c.String("mid")
		// }

		// if entity, err := LicenseActionFromActivationKey(&license, f); err != nil {
		// 	fmt.Println(err.Error())
		// } else {

		// 	f, _ := json.MarshalIndent(entity, "", "  ")
		// 	fmt.Println(string(f))
		// }
	},
}

var InitPrivateKeyCmd cli.Command = cli.Command{

	Name:  "keygen",
	Usage: "Prints out private and public key generated. It won't be stored anywhere.",
	Action: func(c *cli.Context) error {
		info, err := GenertePrivatePublicKeySet()

		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println("This information is not stored anywhere, this is a sample private/public key. You might want to use this when editing a product on low level.")
		fmt.Println("")
		fmt.Println("Private key:")
		fmt.Println("")
		fmt.Println(info.PrivateKey)
		fmt.Println("")
		fmt.Println("Public key:")
		fmt.Println("")
		fmt.Println(info.PublicKey)
		fmt.Println("")

		return nil
	},
}

func init() {
	LicensableProductCliCommands = append(LicensableProductCliCommands, LicensableProductGenerateCmd)
	ActivationKeyCliCommands = append(ActivationKeyCliCommands, CreateBulkActivationKeyCmd)
	ProductPlanCliCommands = append(ProductPlanCliCommands)

	LicenseCliCommands = append(LicenseCliCommands,
		LicenseFromActivationKeyCmd,
		GenerateLicenseCmd,
		InitPrivateKeyCmd,
		LicensableProductCliFn(),
		ActivationKeyCliFn(),
		ProductPlanCliFn(),
	)
}
