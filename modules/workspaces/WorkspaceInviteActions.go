package workspaces

func WorkspaceInviteActionCreate(
	dto *WorkspaceInviteEntity, query QueryDSL,
) (*WorkspaceInviteEntity, *IError) {
	return WorkspaceInviteActionCreateFn(dto, query)
}

func WorkspaceInviteActionUpdate(
	query QueryDSL,
	fields *WorkspaceInviteEntity,
) (*WorkspaceInviteEntity, *IError) {
	return WorkspaceInviteActionUpdateFn(query, fields)
}
