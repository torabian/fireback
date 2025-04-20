package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoStateActionCreate(
	dto *GeoStateEntity, query fireback.QueryDSL,
) (*GeoStateEntity, *fireback.IError) {
	return GeoStateActionCreateFn(dto, query)
}
func GeoStateActionUpdate(
	query fireback.QueryDSL,
	fields *GeoStateEntity,
) (*GeoStateEntity, *fireback.IError) {
	return GeoStateActionUpdateFn(query, fields)
}
