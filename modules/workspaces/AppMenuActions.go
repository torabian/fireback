package workspaces

func AppMenuActionCreate(
	dto *AppMenuEntity, query QueryDSL,
) (*AppMenuEntity, *IError) {
	return AppMenuActionCreateFn(dto, query)
}

func AppMenuActionUpdate(
	query QueryDSL,
	fields *AppMenuEntity,
) (*AppMenuEntity, *IError) {
	return AppMenuActionUpdateFn(query, fields)
}
