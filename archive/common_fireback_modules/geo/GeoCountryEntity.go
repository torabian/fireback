package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoCountryActionCreate(
	dto *GeoCountryEntity, query fireback.QueryDSL,
) (*GeoCountryEntity, *fireback.IError) {
	return GeoCountryActionCreateFn(dto, query)
}
func GeoCountryActionUpdate(
	query fireback.QueryDSL,
	fields *GeoCountryEntity,
) (*GeoCountryEntity, *fireback.IError) {
	return GeoCountryActionUpdateFn(query, fields)
}
