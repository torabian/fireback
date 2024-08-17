package workspaces

func PassportMethodActionCreate(
	dto *PassportMethodEntity, query QueryDSL,
) (*PassportMethodEntity, *IError) {
	return PassportMethodActionCreateFn(dto, query)
}

func PassportMethodActionUpdate(
	query QueryDSL,
	fields *PassportMethodEntity,
) (*PassportMethodEntity, *IError) {
	return PassportMethodActionUpdateFn(query, fields)
}
