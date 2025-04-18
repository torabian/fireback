package abac

import "github.com/torabian/fireback/modules/workspaces"

func EmailConfirmationActionCreate(
	dto *EmailConfirmationEntity, query workspaces.QueryDSL,
) (*EmailConfirmationEntity, *workspaces.IError) {
	return EmailConfirmationActionCreateFn(dto, query)
}

func EmailConfirmationActionUpdate(
	query workspaces.QueryDSL,
	fields *EmailConfirmationEntity,
) (*EmailConfirmationEntity, *workspaces.IError) {
	return EmailConfirmationActionUpdateFn(query, fields)
}
