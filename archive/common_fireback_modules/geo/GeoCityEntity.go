package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoCityActionCreate(
	dto *GeoCityEntity, query fireback.QueryDSL,
) (*GeoCityEntity, *fireback.IError) {
	return GeoCityActionCreateFn(dto, query)
}
func GeoCityActionUpdate(
	query fireback.QueryDSL,
	fields *GeoCityEntity,
) (*GeoCityEntity, *fireback.IError) {
	return GeoCityActionUpdateFn(query, fields)
}
