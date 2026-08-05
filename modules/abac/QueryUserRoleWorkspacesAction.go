package abac

import (
	"sort"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func QueryUserRoleWorkspacesAction(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

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
				Roles:        emigo.ArrayReplace(roles),
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
