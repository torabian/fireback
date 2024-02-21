package licenses

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	"pixelplux.com/fireback/modules/workspaces"
)

var LicensableProductTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the licensableProduct",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

var LicensableProductGenerateCmd cli.Command = cli.Command{

	Name:    "generate",
	Aliases: []string{"gen"},
	Flags:   LicensableProductCommonCliFlagsOptional,
	Usage:   "Create a new product with given name from user, and fills private/public key automatically",
	Action: func(c *cli.Context) {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		entity := CastLicensableProductFromCli(c)

		if entity, err := ProductActionGenerate(entity, query); err != nil {
			fmt.Println(err.Error())
		} else {

			f, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(f))
		}
	},
}

func init() {
	LicensableProductCliCommands = append(LicensableProductCliCommands, LicensableProductTestCmd, LicensableProductGenerateCmd)
}
