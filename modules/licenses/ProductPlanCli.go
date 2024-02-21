package licenses

import (
	"fmt"

	"github.com/urfave/cli"
)

var ProductPlanTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the productPlan",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	ProductPlanCliCommands = append(ProductPlanCliCommands, ProductPlanTestCmd)
}
