package workspaces

func EmailSenderActionCreate(
	dto *EmailSenderEntity, query QueryDSL,
) (*EmailSenderEntity, *IError) {
	return EmailSenderActionCreateFn(dto, query)
}

func EmailSenderActionUpdate(
	query QueryDSL,
	fields *EmailSenderEntity,
) (*EmailSenderEntity, *IError) {
	return EmailSenderActionUpdateFn(query, fields)
}
