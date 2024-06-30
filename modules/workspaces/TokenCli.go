package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var TokenTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the token",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	TokenCliCommands = append(TokenCliCommands, TokenTestCmd)
}
