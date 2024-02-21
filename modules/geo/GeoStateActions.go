package geo

import "github.com/torabian/fireback/modules/workspaces"

func GeoStateActionCreate(
	dto *GeoStateEntity, query workspaces.QueryDSL,
) (*GeoStateEntity, *workspaces.IError) {
	return GeoStateActionCreateFn(dto, query)
}

func GeoStateActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoStateEntity,
) (*GeoStateEntity, *workspaces.IError) {
	return GeoStateActionUpdateFn(query, fields)
}
