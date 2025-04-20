package currency

import "github.com/torabian/fireback/modules/fireback"

func PriceTagActionCreate(
	dto *PriceTagEntity, query fireback.QueryDSL,
) (*PriceTagEntity, *fireback.IError) {
	return PriceTagActionCreateFn(dto, query)
}
func PriceTagActionUpdate(
	query fireback.QueryDSL,
	fields *PriceTagEntity,
) (*PriceTagEntity, *fireback.IError) {
	return PriceTagActionUpdateFn(query, fields)
}
