package workspaces

import "sort"

func init() {
	// Override the implementation with our actual code.
	QueryUserRoleWorkspacesActionImp = QueryUserRoleWorkspacesAction
}

func QueryUserRoleWorkspacesAction(
	q QueryDSL) ([]*QueryUserRoleWorkspacesActionResDto,
	*QueryResultMeta,
	*IError,
) {

	items := []*QueryUserRoleWorkspacesActionResDto{}

	if q.UserAccessPerWorkspace != nil {

		for workspaceId, content := range *q.UserAccessPerWorkspace {
			roles := []*QueryUserRoleWorkspacesResDtoRoles{}

			for roleId, roleContent := range content.UserRoles {
				roles = append(roles, &QueryUserRoleWorkspacesResDtoRoles{
					Name:         roleContent.Name,
					Capabilities: roleContent.Accesses,
					UniqueId:     roleId,
				})
			}

			items = append(items, &QueryUserRoleWorkspacesActionResDto{
				Name:         workspaceId,
				UniqueId:     workspaceId,
				Roles:        roles,
				Capabilities: content.WorkspacesAccesses,
			})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	return items, nil, nil
}
