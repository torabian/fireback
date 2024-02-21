package geo

import (
	"fmt"

	"github.com/urfave/cli"
)

var LocationDataTestCmd cli.Command = cli.Command{

	Name:  "test",
	Usage: "Tests the locationData",
	Action: func(c *cli.Context) error {

		fmt.Printf("Implement the test logic here")

		return nil
	},
}

func init() {
	LocationDataCliCommands = append(LocationDataCliCommands, LocationDataTestCmd)
}
