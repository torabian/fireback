package workspaces

func EmailConfirmationActionCreate(
	dto *EmailConfirmationEntity, query QueryDSL,
) (*EmailConfirmationEntity, *IError) {
	return EmailConfirmationActionCreateFn(dto, query)
}

func EmailConfirmationActionUpdate(
	query QueryDSL,
	fields *EmailConfirmationEntity,
) (*EmailConfirmationEntity, *IError) {
	return EmailConfirmationActionUpdateFn(query, fields)
}
