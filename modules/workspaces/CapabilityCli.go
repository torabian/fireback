package workspaces

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// import (
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"

// 	"github.com/urfave/cli"
// )

// var BuiltInCapabilities []string = []string{}

// func CastToCapabilityTree(children []CapabilityGroup, prefix string, content *[]string) {
// 	for index, item := range children {
// 		if index == 0 {
// 			*content = append(*content, prefix)
// 		}
// 		name := prefix + "/" + item.Name
// 		if len(item.Children) > 0 {
// 			CastToCapabilityTree(item.Children, name, content)
// 		} else {
// 			*content = append(*content, name)
// 		}
// 	}
// }

type PermissionInfo struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	CompleteKey string `yaml:"completeKey,omitempty" json:"completeKey,omitempty"`
}

var CapabilityTreeCmd cli.Command = cli.Command{

	Name:    "tree",
	Aliases: []string{"t"},
	Flags:   CommonQueryFlags,
	Usage:   "Queries all the roles, and prints them as tree",
	Action: func(c *cli.Context) error {

		tree := Tree{}

		f := QueryDSL{
			ItemsPerPage: -1,
		}

		if items, _, err := CapabilityActionQuery(f); err != nil {
			fmt.Println(err)
		} else {
			for _, item := range items {
				tree.Add(item.UniqueId)
			}

			tree.Fprint(os.Stdout, true, "")
			fmt.Println(tree.Json())
		}

		return nil
	},
}

func init() {
	CapabilityCliCommands = append(CapabilityCliCommands, CapabilityTreeCmd)
}
