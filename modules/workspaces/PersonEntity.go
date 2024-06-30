package workspaces

func PersonActionCreate(
	dto *PersonEntity, query QueryDSL,
) (*PersonEntity, *IError) {
	return PersonActionCreateFn(dto, query)
}

func PersonActionUpdate(
	query QueryDSL,
	fields *PersonEntity,
) (*PersonEntity, *IError) {
	return PersonActionUpdateFn(query, fields)
}
