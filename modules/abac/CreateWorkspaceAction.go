package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

/**
*	Creates a workspace, considering the parent workspace,
*	Who creates it, and might accept even manager and roles in the first
**/
func CreateWorkspaceAction(c CreateWorkspaceActionRequest) (*CreateWorkspaceActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}
	q := *query

	context := &GenerateUserDto{
		createUser:      false,
		createWorkspace: true,
		workspace: &WorkspaceEntity{
			Name: c.Body.Name,
		},
		user: &UserEntity{
			UniqueId: q.UserId,
			UserId:   emigo.NullableOf(q.UserId),
		},
		restricted: true,
	}

	session := &UserSessionDto{}
	if err := CreateWorkspaceAndAssignUser(context, q, session); err != nil {
		return nil, err
	} else {
		return &CreateWorkspaceActionResponse{
			Payload: fireback.GResponseSingleItem(&CreateWorkspaceActionRes{
				WorkspaceId: context.workspace.UniqueId,
			}),
		}, nil
	}

}
