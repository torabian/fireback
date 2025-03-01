package workspaces

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func CapabilityActionCreate(
	dto *CapabilityEntity, query QueryDSL,
) (*CapabilityEntity, *IError) {
	return CapabilityActionCreateFn(dto, query)
}

func CapabilityActionUpdate(
	query QueryDSL,
	fields *CapabilityEntity,
) (*CapabilityEntity, *IError) {
	return CapabilityActionUpdateFn(query, fields)
}

type CapabilityChild struct {
	UniqueId   string             `json:"uniqueId,omitempty"`
	Children   []*CapabilityChild `json:"children,omitempty"`
	Visibility *string            `json:"visibility,omitempty" yaml:"visibility"`
	Updated    int64              `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created    int64              `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
}
type CapabilitiesResult struct {
	Capabilities []*CapabilityEntity `json:"capabilities,omitempty"`
	Nested       []*CapabilityChild  `json:"nested,omitempty"`
	Visibility   *string             `json:"visibility,omitempty" yaml:"visibility"`
	Updated      int64               `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created      int64               `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
}

func treeToCapabilityChild(items []NestedNode) []*CapabilityChild {
	data := []*CapabilityChild{}

	for _, item := range items {

		children := []*CapabilityChild{}

		if len(item.Children) > 0 {
			children = treeToCapabilityChild(item.Children)
		}

		data = append(data, &CapabilityChild{
			UniqueId: item.UniqueId,
			Children: children,
		})
	}

	return data
}

func CapabilityActionGetTree(query QueryDSL) (*CapabilitiesResult, *IError) {

	// Read the comments inside CapabilityActionQuery
	items, _, err := CapabilityActionQuery(query)

	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := Tree{}

	for _, item := range items {
		if item.UniqueId == "" {
			continue
		}
		if strings.HasSuffix(item.UniqueId, ".*") {
			tree.Add(strings.TrimRight(item.UniqueId, ".*"), ".")
		} else {
			tree.Add(item.UniqueId, ".")
		}
	}
	itemsa := tree.ToObject(true)

	return &CapabilitiesResult{
		Capabilities: items,
		Nested:       treeToCapabilityChild(itemsa),
	}, GormErrorToIError(err)
}

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
				tree.Add(item.UniqueId, ".")
			}

			tree.Fprint(os.Stdout, true, "")
			fmt.Println(tree.Json())
		}

		return nil
	},
}

func GetCapabilityRefreshCommand(xapp *FirebackApp) cli.Command {
	return cli.Command{

		Name:        "capsync",
		Flags:       CommonQueryFlags,
		Usage:       "Idemponent sync the modules capabilities into the database again.",
		Description: "Fireback and sub projects need to have permissions as capability strings into database to create role or check. This is happening on env startup, but after project updates needs to be refreshed, or if you have deleted them from database.",
		Action: func(c *cli.Context) error {

			SyncPermissionsInDatabase(xapp, GetDbRef())
			return nil
		},
	}

}

func init() {
	CapabilityCliCommands = append(CapabilityCliCommands, CapabilityTreeCmd, ListCapabilitiesActionCmd)
	AppendCapabilityRouter = func(r *[]Module3Action) {

		*r = append(*r, Module3Action{
			Method: "GET",
			Url:    "/capabilitiesTree",
			Handlers: []gin.HandlerFunc{
				WithAuthorization(&SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
				}),
				func(c *gin.Context) {
					HttpGetEntity(c, CapabilityActionGetTree)
				},
			},
			Action:         CapabilityActionGetTree,
			Format:         "GET_ONE",
			ResponseEntity: &CapabilitiesResult{},
		})

	}
}
