package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPubliclyActionImp = QueryWorkspaceTypesPubliclyAction
}
func QueryWorkspaceTypesPubliclyAction(
	q fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto,
	*fireback.QueryResultMeta,
	*fireback.IError,
) {
	// Implement the logic here.
	return nil, nil, nil
}
