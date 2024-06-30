package workspaces

func RoleActionCreate(
	dto *RoleEntity, query QueryDSL,
) (*RoleEntity, *IError) {
	return RoleActionCreateFn(dto, query)
}

func RoleActionUpdate(
	query QueryDSL,
	fields *RoleEntity,
) (*RoleEntity, *IError) {
	return RoleActionUpdateFn(query, fields)
}
