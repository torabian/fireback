package currency

import "pixelplux.com/fireback/modules/workspaces"

func CurrencyActionCreate(
	dto *CurrencyEntity, query workspaces.QueryDSL,
) (*CurrencyEntity, *workspaces.IError) {
	return CurrencyActionCreateFn(dto, query)
}

func CurrencyActionUpdate(
	query workspaces.QueryDSL,
	fields *CurrencyEntity,
) (*CurrencyEntity, *workspaces.IError) {
	return CurrencyActionUpdateFn(query, fields)
}
