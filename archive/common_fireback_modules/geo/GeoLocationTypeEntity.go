package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoLocationTypeActionCreate(
	dto *GeoLocationTypeEntity, query fireback.QueryDSL,
) (*GeoLocationTypeEntity, *fireback.IError) {
	return GeoLocationTypeActionCreateFn(dto, query)
}
func GeoLocationTypeActionUpdate(
	query fireback.QueryDSL,
	fields *GeoLocationTypeEntity,
) (*GeoLocationTypeEntity, *fireback.IError) {
	return GeoLocationTypeActionUpdateFn(query, fields)
}
