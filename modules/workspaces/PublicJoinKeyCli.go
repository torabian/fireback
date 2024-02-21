package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var PublicJoinKeyTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the publicJoinKey",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PublicJoinKeyCliCommands = append(PublicJoinKeyCliCommands, PublicJoinKeyTestCmd)
}
