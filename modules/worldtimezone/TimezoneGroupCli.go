package worldtimezone

import (
	"fmt"

	"github.com/urfave/cli"
)

var TimezoneGroupTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the TimezoneGroup",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	TimezoneGroupCliCommands = append(TimezoneGroupCliCommands, TimezoneGroupTestCmd)
}
