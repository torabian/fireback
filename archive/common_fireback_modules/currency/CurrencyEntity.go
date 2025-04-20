package currency

import "github.com/torabian/fireback/modules/fireback"

func CurrencyActionCreate(
	dto *CurrencyEntity, query fireback.QueryDSL,
) (*CurrencyEntity, *fireback.IError) {
	return CurrencyActionCreateFn(dto, query)
}
func CurrencyActionUpdate(
	query fireback.QueryDSL,
	fields *CurrencyEntity,
) (*CurrencyEntity, *fireback.IError) {
	return CurrencyActionUpdateFn(query, fields)
}
