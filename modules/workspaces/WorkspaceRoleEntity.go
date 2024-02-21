package workspaces

func WorkspaceRoleActionCreate(
	dto *WorkspaceRoleEntity, query QueryDSL,
) (*WorkspaceRoleEntity, *IError) {
	return WorkspaceRoleActionCreateFn(dto, query)
}
func WorkspaceRoleActionUpdate(
	query QueryDSL,
	fields *WorkspaceRoleEntity,
) (*WorkspaceRoleEntity, *IError) {
	return WorkspaceRoleActionUpdateFn(query, fields)
}
