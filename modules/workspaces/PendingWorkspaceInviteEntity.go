package workspaces

func PendingWorkspaceInviteActionCreate(
	dto *PendingWorkspaceInviteEntity, query QueryDSL,
) (*PendingWorkspaceInviteEntity, *IError) {
	return PendingWorkspaceInviteActionCreateFn(dto, query)
}

func PendingWorkspaceInviteActionUpdate(
	query QueryDSL,
	fields *PendingWorkspaceInviteEntity,
) (*PendingWorkspaceInviteEntity, *IError) {
	return PendingWorkspaceInviteActionUpdateFn(query, fields)
}
