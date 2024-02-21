package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var EmailSenderTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the emailSender",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	EmailSenderCliCommands = append(EmailSenderCliCommands, EmailSenderTestCmd)
}
