package commonprofile

import "github.com/torabian/fireback/modules/fireback"

func CommonProfileActionCreate(
	dto *CommonProfileEntity, query fireback.QueryDSL,
) (*CommonProfileEntity, *fireback.IError) {
	return CommonProfileActionCreateFn(dto, query)
}

func CommonProfileActionUpdate(
	query fireback.QueryDSL,
	fields *CommonProfileEntity,
) (*CommonProfileEntity, *fireback.IError) {
	return CommonProfileActionUpdateFn(query, fields)
}
