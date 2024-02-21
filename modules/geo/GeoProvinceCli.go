package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var GeoProvinceTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the geoProvince",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	GeoProvinceCliCommands = append(GeoProvinceCliCommands, GeoProvinceTestCmd)
}
