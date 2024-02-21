package workspaces

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// func MoveToRoot(node *RoleEntity, items []*RoleEntity, path string) string {

// 	name := node.Name

// 	if name == "" {
// 		name = "Unnamed Role"
// 	}

// 	title := fmt.Sprintf("%s (%s)", name, node.UniqueId)

// 	if node.ParentRoleId == nil || *node.ParentRoleId == "" {
// 		path = "/" + title + path
// 		return path
// 	} else {
// 		for _, item := range items {
// 			if item.UniqueId == *node.ParentRoleId {
// 				if strings.Contains(path, title) {
// 					return "~ with circular"
// 				}

// 				path = "/" + title + path
// 				return MoveToRoot(item, items, path)
// 			}
// 		}
// 	}

// 	return ""
// }

// var RoleTreeCmd cli.Command = cli.Command{

// 	Name:    "tree",
// 	Aliases: []string{"t"},
// 	Flags:   CommonQueryFlags,
// 	Usage:   "Queries all the roles, and prints them as tree",
// 	Action: func(c *cli.Context) error {

// 		f := CommonCliQueryDSLBuilder(c)
// 		if results, _, err := RoleActionQuery(f); err != nil {
// 			fmt.Println(err)
// 		} else {

// 			tree := Tree{}

// 			for _, item := range results {

// 				tree.Add(MoveToRoot(item, results, ""))
// 			}

// 			tree.Fprint(os.Stdout, true, "")

// 		}
// 		return nil
// 	},
// }

// var RoleCli cli.Command = cli.Command{
// 	Name:        "roles",
// 	Description: "roles module actions",
// 	Usage:       "Actions related to the roles module",

// 	Subcommands: []cli.Command{
// 		RoleQueryCmd,
// 		RoleUpdateCmd,
// 		InteractiveCMD,
// 		RoleCreateCmd,
// 		RoleWipeCmd,
// 		RoleSeederCmd,
// 		RoleRemoveCmd,
// 		RoleUnlinkCmd,
// 		RoleTreeCmd,
// 	},
// }

var InteractiveCMD cli.Command = cli.Command{

	Name:  "i",
	Usage: "Opens interactive cli",
	Action: func(c *cli.Context) error {
		validate := func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			return err
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}

		prompt := promptui.Prompt{
			Label:     "Spicy Level",
			Templates: templates,
			Validate:  validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		fmt.Printf("You answered %s\n", result)
		return nil
	},
}
