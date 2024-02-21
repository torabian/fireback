package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var TableViewSizingTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the tableViewSizing",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	TableViewSizingCliCommands = append(TableViewSizingCliCommands, TableViewSizingTestCmd)
}
