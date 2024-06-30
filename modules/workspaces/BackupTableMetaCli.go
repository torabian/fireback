package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var BackupTableMetaTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the backupTableMeta",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	BackupTableMetaCliCommands = append(BackupTableMetaCliCommands, BackupTableMetaTestCmd)
}
