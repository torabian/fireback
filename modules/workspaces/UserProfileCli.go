package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var UserProfileTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the userProfile",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	UserProfileCliCommands = append(UserProfileCliCommands, UserProfileTestCmd)
}
