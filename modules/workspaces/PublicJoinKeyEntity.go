package workspaces

func PublicJoinKeyActionCreate(
	dto *PublicJoinKeyEntity, query QueryDSL,
) (*PublicJoinKeyEntity, *IError) {
	return PublicJoinKeyActionCreateFn(dto, query)
}

func PublicJoinKeyActionUpdate(
	query QueryDSL,
	fields *PublicJoinKeyEntity,
) (*PublicJoinKeyEntity, *IError) {
	return PublicJoinKeyActionUpdateFn(query, fields)
}
