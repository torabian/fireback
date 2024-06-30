package workspaces

import (
	"fmt"

	"github.com/urfave/cli"
)

var PreferenceTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the preference",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	PreferenceCliCommands = append(PreferenceCliCommands, PreferenceTestCmd)
}
