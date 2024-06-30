package workspaces

func WorkspaceConfigActionCreate(
	dto *WorkspaceConfigEntity, query QueryDSL,
) (*WorkspaceConfigEntity, *IError) {
	return WorkspaceConfigActionCreateFn(dto, query)
}

func WorkspaceConfigActionUpdate(
	query QueryDSL,
	fields *WorkspaceConfigEntity,
) (*WorkspaceConfigEntity, *IError) {
	return WorkspaceConfigActionUpdateFn(query, fields)
}
