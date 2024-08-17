package workspaces

func PhoneConfirmationActionCreate(
	dto *PhoneConfirmationEntity, query QueryDSL,
) (*PhoneConfirmationEntity, *IError) {
	return PhoneConfirmationActionCreateFn(dto, query)
}

func PhoneConfirmationActionUpdate(
	query QueryDSL,
	fields *PhoneConfirmationEntity,
) (*PhoneConfirmationEntity, *IError) {
	return PhoneConfirmationActionUpdateFn(query, fields)
}
