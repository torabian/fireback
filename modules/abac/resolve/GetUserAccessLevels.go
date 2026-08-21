package resolve

import (
	queries "github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
)

func GetUserAccessLevels(query QueryDSL) (*UserAccessLevelDto, *fireback.IError) {

	access := &UserAccessLevelDto{}
	query.ItemsPerPage = 1000

	items, _, err := fireback.UnsafeQuerySqlFromFs[UserRoleWorkspacePermissionDto](
		&queries.QueriesFs, "UserRolePermission", fireback.QueryDSL{
			UserId:        query.UserId,
			WorkspaceId:   query.WorkspaceId,
			InternalQuery: query.InternalQuery,
		},
	)

	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	ws := UserAccessPerWorkspaceDto{}

	for _, item := range items {
		if ws[item.WorkspaceId] == nil {
			ws[item.WorkspaceId] = &struct {
				Name               string
				WorkspacesAccesses []string
				UserRoles          map[string]*struct {
					Name     string
					Accesses []string
				}
			}{}
		}

		ws[item.WorkspaceId].Name = item.WorkspaceName

		if item.Type == "account_restrict" {
			// Bug fix: this used to re-init the whole UserRoles map (wiping every
			// other role already accumulated for this same workspace) whenever it hit
			// a *second* role_id it hadn't seen yet - the map-is-nil check and the
			// key-is-nil check were conflated into one condition. In practice this
			// meant a user holding more than one role in the same workspace (e.g. via
			// a Workspaceabacdefs.RoleEntity granting an extra role alongside their normal one)
			// silently lost every role but the last one processed - for root, that
			// could drop the seeded "root.*" wildcard role itself, causing
			// MeetsAccessLevel to reject actions root should always be allowed.
			if ws[item.WorkspaceId].UserRoles == nil {
				ws[item.WorkspaceId].UserRoles = map[string]*struct {
					Name     string
					Accesses []string
				}{}
			}
			if ws[item.WorkspaceId].UserRoles[item.RoleId] == nil {
				ws[item.WorkspaceId].UserRoles[item.RoleId] = &struct {
					Name     string
					Accesses []string
				}{}
			}
			ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses = append(ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses, item.CapabilityId)
			ws[item.WorkspaceId].UserRoles[item.RoleId].Name = item.RoleName
		}

		if item.Type == "workspace_restrict" {
			ws[item.WorkspaceId].WorkspacesAccesses = append(ws[item.WorkspaceId].WorkspacesAccesses, item.CapabilityId)
		}
	}

	access.UserAccessPerWorkspace = &ws

	return access, nil
}
