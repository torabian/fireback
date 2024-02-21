package currency

import "github.com/torabian/fireback/modules/workspaces"

func PriceTagActionCreate(
	dto *PriceTagEntity, query workspaces.QueryDSL,
) (*PriceTagEntity, *workspaces.IError) {
	return PriceTagActionCreateFn(dto, query)
}

func PriceTagActionUpdate(
	query workspaces.QueryDSL,
	fields *PriceTagEntity,
) (*PriceTagEntity, *workspaces.IError) {
	return PriceTagActionUpdateFn(query, fields)
}
