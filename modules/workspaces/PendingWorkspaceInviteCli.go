package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var PendingWorkspaceInviteTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the pendingWorkspaceInvite",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PendingWorkspaceInviteCliCommands = append(PendingWorkspaceInviteCliCommands, PendingWorkspaceInviteTestCmd)
}
