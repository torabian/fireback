package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var WorkspaceInviteTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the workspaceInvite",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	WorkspaceInviteCliCommands = append(WorkspaceInviteCliCommands, WorkspaceInviteTestCmd)
}
