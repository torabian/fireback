package geo

import "pixelplux.com/fireback/modules/workspaces"

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
