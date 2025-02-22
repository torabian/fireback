package geo

import "github.com/torabian/fireback/modules/workspaces"

func GeoLocationActionCreate(
	dto *GeoLocationEntity, query workspaces.QueryDSL,
) (*GeoLocationEntity, *workspaces.IError) {
	return GeoLocationActionCreateFn(dto, query)
}
func GeoLocationActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoLocationEntity,
) (*GeoLocationEntity, *workspaces.IError) {
	return GeoLocationActionUpdateFn(query, fields)
}
