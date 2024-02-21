package currency

import (
	"fmt"

	"github.com/urfave/cli"
)

var PriceTagTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the priceTag",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PriceTagCliCommands = append(PriceTagCliCommands, PriceTagTestCmd)
}
