package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var WorkspaceTypeTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the workspaceType",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	WorkspaceTypeCliCommands = append(WorkspaceTypeCliCommands, WorkspaceTypeTestCmd)
}
