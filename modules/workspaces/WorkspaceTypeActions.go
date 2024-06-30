package workspaces

func WorkspaceTypeActionCreate(
	dto *WorkspaceTypeEntity, query QueryDSL,
) (*WorkspaceTypeEntity, *IError) {
	return WorkspaceTypeActionCreateFn(dto, query)
}

func WorkspaceTypeActionUpdate(
	query QueryDSL,
	fields *WorkspaceTypeEntity,
) (*WorkspaceTypeEntity, *IError) {
	return WorkspaceTypeActionUpdateFn(query, fields)
}

func WorkspaceTypeActionPublicQuery(query QueryDSL) ([]*WorkspaceTypeEntity, *QueryResultMeta, error) {
	// Make this API public, so the signup screen can get it.
	// At this moment, we just move things back as are, but maybe later we need
	// to add some limits on what kind of information is going out.
	query.WorkspaceId = "root"
	query.UserId = "root"
	return WorkspaceTypeActionQuery(query)
}
