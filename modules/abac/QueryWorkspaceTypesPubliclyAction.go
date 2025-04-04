package workspaces

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPubliclyActionImp = QueryWorkspaceTypesPubliclyAction
}
func QueryWorkspaceTypesPubliclyAction(
	q QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto,
	*QueryResultMeta,
	*IError,
) {
	// Implement the logic here.
	return nil, nil, nil
}
