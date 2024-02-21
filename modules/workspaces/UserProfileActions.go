package workspaces

func UserProfileActionCreate(
	dto *UserProfileEntity, query QueryDSL,
) (*UserProfileEntity, *IError) {
	return UserProfileActionCreateFn(dto, query)
}

func UserProfileActionUpdate(
	query QueryDSL,
	fields *UserProfileEntity,
) (*UserProfileEntity, *IError) {
	return UserProfileActionUpdateFn(query, fields)
}
