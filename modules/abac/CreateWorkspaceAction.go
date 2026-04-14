package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	CreateWorkspaceImpl = CreateWorkspaceAction
}

/**
*	Creates a workspace, considering the parent workspace,
*	Who creates it, and might accept even manager and roles in the first
**/
func CreateWorkspaceAction(c CreateWorkspaceActionRequest, q fireback.QueryDSL) (*CreateWorkspaceActionResponse, error) {

	context := &GenerateUserDto{
		createUser:      false,
		createWorkspace: true,
		workspace: &WorkspaceEntity{
			Name: c.Body.Name,
		},
		user: &UserEntity{
			UniqueId: q.UserId,
			UserId:   fireback.NewString(q.UserId),
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
