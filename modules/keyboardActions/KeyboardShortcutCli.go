package keyboardActions

import (
	"fmt"

	"github.com/urfave/cli"
)

var KeyboardShortcutTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the keyboardShortcut",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	KeyboardShortcutCliCommands = append(KeyboardShortcutCliCommands, KeyboardShortcutTestCmd)
}
