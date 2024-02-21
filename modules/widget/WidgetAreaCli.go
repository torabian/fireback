package widget

import (
	"fmt"

	"github.com/urfave/cli"
)

var WidgetAreaTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the widgetArea",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	WidgetAreaCliCommands = append(WidgetAreaCliCommands, WidgetAreaTestCmd)
}
