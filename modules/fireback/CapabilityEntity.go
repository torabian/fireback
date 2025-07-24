package fireback

import (
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

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

type PermissionInfo struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	CompleteKey string `yaml:"completeKey,omitempty" json:"completeKey,omitempty"`
	GoVariable  string `yaml:"-" json:"-"`
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

func GetStats(xapp *FirebackApp) cli.Command {
	return cli.Command{
		Name:        "stats",
		Flags:       CommonQueryFlags,
		Usage:       "Some stats regarding the application will go here",
		Description: "Some stats regarding the application will go here",
		Action: func(c *cli.Context) error {

			return nil
		},
	}

}

func SyncPermissionsInDatabase(x *FirebackApp, db *gorm.DB) {

	for _, item := range x.Modules {

		// if item.BackupTables != nil && len(item.BackupTables) > 0 {
		// 	for _, table := range item.BackupTables {

		// 		GetDbRef().Model(&BackupTableMetaEntity{}).Create(&BackupTableMetaEntity{
		// 			UniqueId:      table.EntityName,
		// 			TableNameInDb: table.TableNameInDb,
		// 		})
		// 	}
		// }

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
