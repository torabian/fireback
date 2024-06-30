package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var AppMenuTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the appMenu",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	AppMenuCliCommands = append(AppMenuCliCommands, AppMenuTestCmd)
}
