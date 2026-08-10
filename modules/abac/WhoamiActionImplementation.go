package abac

import (
	"sort"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

// WhoamiAction reports the currently authenticated user: their own userId, plus
// every workspace they belong to along with their role(s) and capabilities in
// each - the same UserRoleWorkspace shape QueryUserRoleWorkspacesAction already
// returns (see QueryUserRoleWorkspacesActionImplementation.go), just wrapped as a
// single item (GResponseSingleItem) instead of a list, since this is always
// exactly one caller's own info, never a paginated collection.
func WhoamiAction(c abacdefs.WhoamiActionRequest) (*abacdefs.WhoamiActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	workspaces := []abacdefs.WhoamiActionResWorkspaces{}

	if q.UserAccessPerWorkspace != nil {
		for workspaceId, content := range *q.UserAccessPerWorkspace {
			roles := []abacdefs.WhoamiActionResWorkspacesRoles{}

			for roleId, roleContent := range content.UserRoles {
				roles = append(roles, abacdefs.WhoamiActionResWorkspacesRoles{
					Name:         roleContent.Name,
					Capabilities: roleContent.Accesses,
					UniqueId:     roleId,
				})
			}

			workspaces = append(workspaces, abacdefs.WhoamiActionResWorkspaces{
				Name:         content.Name,
				UniqueId:     workspaceId,
				Roles:        emigo.ArrayReplace(roles),
				Capabilities: content.WorkspacesAccesses,
			})
		}
	}

	sort.Slice(workspaces, func(i, j int) bool {
		return workspaces[i].UniqueId < workspaces[j].UniqueId
	})

	res := abacdefs.WhoamiActionRes{
		UserId:     q.UserId,
		Workspaces: emigo.ArrayReplace(workspaces),
	}

	return &abacdefs.WhoamiActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}
