package currency

import (
	"fmt"

	"github.com/urfave/cli"
)

var CurrencyTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the currency",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	CurrencyCliCommands = append(CurrencyCliCommands, CurrencyTestCmd)
}
