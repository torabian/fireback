package workspaces

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

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

func GetWorkspaceAndUserAccesses(query QueryDSL) ([]string, []string) {
	if query.UserAccessPerWorkspace == nil {
		return []string{}, []string{}
	}

	data := *query.UserAccessPerWorkspace
	workspaceAccesses := []string{}
	rolesPermission := []string{}
	if data[query.WorkspaceId] != nil {
		workspaceAccesses = data[query.WorkspaceId].WorkspacesAccesses

		// Now we are checking with all the roles user has, but need to have access to role id
		// and only look for that.
		for _, role := range data[query.WorkspaceId].UserRoles {
			rolesPermission = append(rolesPermission, role.Accesses...)
		}
	}

	return workspaceAccesses, rolesPermission
}

func CapabilityActionGetTree(query QueryDSL) (*CapabilitiesResult, *IError) {

	// Read the comments inside CapabilityActionQuery
	items, _, err := CapabilityActions.Query(query)
	itemsFiltered := []*CapabilityEntity{}

	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(query)
	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := Tree{}

	for _, item := range items {
		if item.UniqueId == "" {
			continue
		}

		// Filter based on the workspace and role and not allow to create more access than the user has.
		meetsUser := meetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, rolesPermission)
		meetsWorkspace := meetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, item)
		if strings.HasSuffix(item.UniqueId, ".*") {
			tree.Add(strings.TrimRight(item.UniqueId, ".*"), ".")
		} else {
			tree.Add(item.UniqueId, ".")
		}
	}
	itemsa := tree.ToObject(true)

	return &CapabilitiesResult{
		Capabilities: itemsFiltered,
		Nested:       treeToCapabilityChild(itemsa),
	}, GormErrorToIError(err)
}

type PermissionInfo struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	CompleteKey string `yaml:"completeKey,omitempty" json:"completeKey,omitempty"`
	GoVariable  string `yaml:"-" json:"-"`
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

		if items, _, err := CapabilityActions.Query(f); err != nil {
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

func SyncPermissionsInDatabase(x *FirebackApp, db *gorm.DB) {

	for _, item := range x.Modules {

		if item.BackupTables != nil && len(item.BackupTables) > 0 {
			for _, table := range item.BackupTables {

				GetDbRef().Model(&BackupTableMetaEntity{}).Create(&BackupTableMetaEntity{
					UniqueId:      table.EntityName,
					TableNameInDb: table.TableNameInDb,
				})
			}
		}

		// Insert the permissions into the database
		item.PermissionsProvider = append(item.PermissionsProvider, PermissionInfo{
			CompleteKey: ROOT_ALL_ACCESS,
		}, PermissionInfo{
			CompleteKey: ROOT_ALL_MODULES,
		})

		for _, perm := range item.PermissionsProvider {
			hasChildren := HasChildren(perm.CompleteKey, PermissionInfoToString(item.PermissionsProvider))
			UpsertPermission(&perm, hasChildren, db)
		}

		for _, bundle := range item.EntityBundles {
			for _, perm := range bundle.Permissions {
				hasChildren := HasChildren(perm.CompleteKey, PermissionInfoToString(bundle.Permissions))
				UpsertPermission(&perm, hasChildren, db)
			}
		}

	}

}

func PermissionInfoToString(items []PermissionInfo) []string {
	res := []string{}

	for _, j := range items {
		res = append(res, j.CompleteKey)
	}

	return res
}
