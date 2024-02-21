package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoLocationTypeTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoLocationType",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoLocationTypeCliCommands = append(GeoLocationTypeCliCommands, GeoLocationTypeTestCmd)
}
