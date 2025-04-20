package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoProvinceActionCreate(
	dto *GeoProvinceEntity, query fireback.QueryDSL,
) (*GeoProvinceEntity, *fireback.IError) {
	return GeoProvinceActionCreateFn(dto, query)
}
func GeoProvinceActionUpdate(
	query fireback.QueryDSL,
	fields *GeoProvinceEntity,
) (*GeoProvinceEntity, *fireback.IError) {
	return GeoProvinceActionUpdateFn(query, fields)
}
