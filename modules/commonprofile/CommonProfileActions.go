package commonprofile

import "github.com/torabian/fireback/modules/workspaces"

func CommonProfileActionCreate(
	dto *CommonProfileEntity, query workspaces.QueryDSL,
) (*CommonProfileEntity, *workspaces.IError) {
	return CommonProfileActionCreateFn(dto, query)
}

func CommonProfileActionUpdate(
	query workspaces.QueryDSL,
	fields *CommonProfileEntity,
) (*CommonProfileEntity, *workspaces.IError) {
	return CommonProfileActionUpdateFn(query, fields)
}
