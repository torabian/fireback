package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoLocationTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoLocation",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoLocationCliCommands = append(GeoLocationCliCommands, GeoLocationTestCmd)
}
