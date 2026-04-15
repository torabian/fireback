package abac

import (
	"sort"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	QueryUserRoleWorkspacesImpl = QueryUserRoleWorkspacesAction
}

func QueryUserRoleWorkspacesAction(c QueryUserRoleWorkspacesActionRequest, q fireback.QueryDSL) (*QueryUserRoleWorkspacesActionResponse, error) {

	items := []*QueryUserRoleWorkspacesActionRes{}

	if q.UserAccessPerWorkspace != nil {

		for workspaceId, content := range *q.UserAccessPerWorkspace {
			roles := []QueryUserRoleWorkspacesActionResRoles{}

			for roleId, roleContent := range content.UserRoles {
				roles = append(roles, QueryUserRoleWorkspacesActionResRoles{
					Name:         roleContent.Name,
					Capabilities: roleContent.Accesses,
					UniqueId:     roleId,
				})
			}

			items = append(items, &QueryUserRoleWorkspacesActionRes{
				Name:         content.Name,
				UniqueId:     workspaceId,
				Roles:        roles,
				Capabilities: content.WorkspacesAccesses,
			})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	return &QueryUserRoleWorkspacesActionResponse{
		Payload: fireback.GResponseQuery(items, nil, &q),
	}, nil
}
