package geo

import "pixelplux.com/fireback/modules/workspaces"

func GeoProvinceActionCreate(
	dto *GeoProvinceEntity, query workspaces.QueryDSL,
) (*GeoProvinceEntity, *workspaces.IError) {
	return GeoProvinceActionCreateFn(dto, query)
}

func GeoProvinceActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoProvinceEntity,
) (*GeoProvinceEntity, *workspaces.IError) {
	return GeoProvinceActionUpdateFn(query, fields)
}
