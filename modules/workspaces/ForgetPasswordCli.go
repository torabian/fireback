package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var ForgetPasswordTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the forgetPassword",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	ForgetPasswordCliCommands = append(ForgetPasswordCliCommands, ForgetPasswordTestCmd)
}
