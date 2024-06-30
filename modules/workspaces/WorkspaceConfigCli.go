package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var WorkspaceConfigTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the workspaceConfig",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	WorkspaceConfigCliCommands = append(WorkspaceConfigCliCommands, WorkspaceConfigTestCmd)
}
