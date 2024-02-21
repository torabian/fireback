package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoStateTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoState",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoStateCliCommands = append(GeoStateCliCommands, GeoStateTestCmd)
}
