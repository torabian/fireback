package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var EmailConfirmationTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the emailConfirmation",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	EmailConfirmationCliCommands = append(EmailConfirmationCliCommands, EmailConfirmationTestCmd)
}
