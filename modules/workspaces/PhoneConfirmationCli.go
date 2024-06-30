package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var PhoneConfirmationTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the phoneConfirmation",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PhoneConfirmationCliCommands = append(PhoneConfirmationCliCommands, PhoneConfirmationTestCmd)
}
