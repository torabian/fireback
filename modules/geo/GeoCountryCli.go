package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoCountryTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoCountry",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoCountryCliCommands = append(GeoCountryCliCommands, GeoCountryTestCmd)
}
