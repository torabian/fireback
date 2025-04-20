package geo

import "github.com/torabian/fireback/modules/fireback"

func GeoLocationActionCreate(
	dto *GeoLocationEntity, query fireback.QueryDSL,
) (*GeoLocationEntity, *fireback.IError) {
	return GeoLocationActionCreateFn(dto, query)
}
func GeoLocationActionUpdate(
	query fireback.QueryDSL,
	fields *GeoLocationEntity,
) (*GeoLocationEntity, *fireback.IError) {
	return GeoLocationActionUpdateFn(query, fields)
}
