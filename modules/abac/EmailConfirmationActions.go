package abac

import "github.com/torabian/fireback/modules/fireback"

func EmailConfirmationActionCreate(
	dto *EmailConfirmationEntity, query fireback.QueryDSL,
) (*EmailConfirmationEntity, *fireback.IError) {
	return EmailConfirmationActionCreateFn(dto, query)
}

func EmailConfirmationActionUpdate(
	query fireback.QueryDSL,
	fields *EmailConfirmationEntity,
) (*EmailConfirmationEntity, *fireback.IError) {
	return EmailConfirmationActionUpdateFn(query, fields)
}
