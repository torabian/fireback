package geo

import "pixelplux.com/fireback/modules/workspaces"

func GeoCountryActionCreate(
	dto *GeoCountryEntity, query workspaces.QueryDSL,
) (*GeoCountryEntity, *workspaces.IError) {
	return GeoCountryActionCreateFn(dto, query)
}

func GeoCountryActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoCountryEntity,
) (*GeoCountryEntity, *workspaces.IError) {
	return GeoCountryActionUpdateFn(query, fields)
}
