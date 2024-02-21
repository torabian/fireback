package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var PassportMethodTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the passportMethod",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PassportMethodCliCommands = append(PassportMethodCliCommands, PassportMethodTestCmd)
}
