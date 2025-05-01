package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	CreateWorkspaceActionImp = CreateWorkspaceAction
}

/**
*	Creates a workspace, considering the parent workspace,
*	Who creates it, and might accept even manager and roles in the first
**/
func CreateWorkspaceAction(req *CreateWorkspaceActionReqDto, q fireback.QueryDSL) (*WorkspaceEntity, *fireback.IError) {

	context := &GenerateUserDto{
		createUser:      false,
		createWorkspace: true,
		workspace: &WorkspaceEntity{
			Name: req.Name,
		},
		user: &UserEntity{
			UniqueId: q.UserId,
			UserId:   fireback.NewString(q.UserId),
		},
		restricted: true,
		// createRole: true,
		// role: &RoleEntity{
		// 	Name: "role",
		// },
	}
	session := &UserSessionDto{}
	if err := CreateWorkspaceAndAssignUser(context, q, session); err != nil {
		return nil, err
	} else {
		return context.workspace, nil
	}

}
