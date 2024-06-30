package licenses

import (
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

// This goes under the license category
//
//	func init() {
//		ActivationKeyCliCommands = append(ActivationKeyCliCommands, ActivationKeyTestCmd)
//	}
var CreateBulkActivationKeyCmd cli.Command = cli.Command{

	Name: "bulk",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "length",
			Usage: "length of the activation key",
			Value: 16,
		},
		&cli.IntFlag{
			Name:  "count",
			Usage: "how many activation key do you need to be generated and stored in database",
			Value: 10,
		},
		&cli.StringFlag{
			Name:  "series",
			Usage: "an string identifier, something unique, to label the series of activation code. For example you want mark 100 activation code as 'microsoft', which you sold to them",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "plan-id",
			Usage: "The plan of product, which will be activated using these activation codes",
			Value: "",
		},
	},
	Usage: "Generate a list of new activation codes, to be distributed, or printed on paper",
	Action: func(c *cli.Context) error {
		f := workspaces.QueryDSL{}

		LicenseActionSeederActivationKey(f, c.String("series"), c.Int("count"), c.Int("length"), c.String("plan-id"))

		return nil
	},
}
