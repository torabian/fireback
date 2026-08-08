package abac

import (
	"context"
	"fmt"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func GetWorkspaceAndUserAccesses(query fireback.QueryDSL) ([]string, []string) {

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

	fmt.Println(5, workspaceAccesses, rolesPermission)
	return workspaceAccesses, rolesPermission
}

func GetCapabilityRefreshCommand(xapp *fireback.FirebackApp) *cli.Command {
	return &cli.Command{

		Name:        "capsync",
		Flags:       fireback.CommonQueryFlags,
		Usage:       "Idemponent sync the modules capabilities into the database again.",
		Description: "Fireback and sub projects need to have permissions as capability strings into database to create role or check. This is happening on env startup, but after project updates needs to be refreshed, or if you have deleted them from database.",
		Action: func(ctx context.Context, c *cli.Command) error {

			SyncPermissionsInDatabase(xapp, fireback.GetDbRef())
			return nil
		},
	}

}

func GetStats(xapp *fireback.FirebackApp) cli.Command {
	return cli.Command{
		Name:        "stats",
		Flags:       fireback.CommonQueryFlags,
		Usage:       "Some stats regarding the application will go here",
		Description: "Some stats regarding the application will go here",
		Action: func(ctx context.Context, c *cli.Command) error {

			return nil
		},
	}

}

func SyncPermissionsInDatabase(x *fireback.FirebackApp, db *gorm.DB) {

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
		item.PermissionsProvider = append(item.PermissionsProvider, fireback.PermissionInfo{
			CompleteKey: ROOT_ALL_ACCESS,
		}, fireback.PermissionInfo{
			CompleteKey: ROOT_ALL_MODULES,
		})

		for _, perm := range item.PermissionsProvider {
			hasChildren := fireback.HasChildren(perm.CompleteKey, PermissionInfoToString(item.PermissionsProvider))
			CapabilityUpsertPermissionFn(&perm, hasChildren, db)
		}

		for _, bundle := range item.EntityBundles {
			for _, perm := range bundle.Permissions {
				hasChildren := fireback.HasChildren(perm.CompleteKey, PermissionInfoToString(bundle.Permissions))
				CapabilityUpsertPermissionFn(&perm, hasChildren, db)
			}
		}

	}

}

func PermissionInfoToString(items []fireback.PermissionInfo) []string {
	res := []string{}

	for _, j := range items {
		res = append(res, j.CompleteKey)
	}

	return res
}
