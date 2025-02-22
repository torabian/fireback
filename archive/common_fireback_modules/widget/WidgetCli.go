package widget

import (
	"fmt"

	"github.com/urfave/cli"
)

var WidgetTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the widget",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	WidgetCliCommands = append(WidgetCliCommands, WidgetTestCmd)
}
