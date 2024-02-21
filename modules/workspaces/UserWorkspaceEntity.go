package workspaces

func UserWorkspaceActionCreate(
	dto *UserWorkspaceEntity, query QueryDSL,
) (*UserWorkspaceEntity, *IError) {
	return UserWorkspaceActionCreateFn(dto, query)
}
func UserWorkspaceActionUpdate(
	query QueryDSL,
	fields *UserWorkspaceEntity,
) (*UserWorkspaceEntity, *IError) {
	return UserWorkspaceActionUpdateFn(query, fields)
}
