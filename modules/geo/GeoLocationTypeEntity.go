package geo

import "github.com/torabian/fireback/modules/workspaces"

func GeoLocationTypeActionCreate(
	dto *GeoLocationTypeEntity, query workspaces.QueryDSL,
) (*GeoLocationTypeEntity, *workspaces.IError) {
	return GeoLocationTypeActionCreateFn(dto, query)
}
func GeoLocationTypeActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoLocationTypeEntity,
) (*GeoLocationTypeEntity, *workspaces.IError) {
	return GeoLocationTypeActionUpdateFn(query, fields)
}
