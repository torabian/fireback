package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoCityTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoCity",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoCityCliCommands = append(GeoCityCliCommands, GeoCityTestCmd)
}
