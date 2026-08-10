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
	// A nil SecurityModel here (as opposed to an empty &fireback.SecurityModel{}, which
	// still requires *some* valid token - see e.g. TimezoneGroupBrowseAction) skips
	// AuthorizeRequest entirely, so q.UserId never actually gets populated from the
	// caller's token (see ExtractQueryDslFromGinContext's user_id read, which only ever
	// has anything to read once AuthorizeRequest has run). That silently broke this
	// action's whole premise ("Who creates it", below): every caller - authenticated or
	// not - got q.UserId == "", so CreateWorkspaceAndAssignUser never linked the new
	// workspace to anyone, leaving it permanently inaccessible to its own creator.
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
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
