package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPubliclyActionImp = QueryWorkspaceTypesPubliclyAction
}
func QueryWorkspaceTypesPubliclyAction(
	q workspaces.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto,
	*workspaces.QueryResultMeta,
	*workspaces.IError,
) {
	// Implement the logic here.
	return nil, nil, nil
}
