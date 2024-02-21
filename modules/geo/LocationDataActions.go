package geo

import "pixelplux.com/fireback/modules/workspaces"

func LocationDataActionCreate(
	dto *LocationDataEntity, query workspaces.QueryDSL,
) (*LocationDataEntity, *workspaces.IError) {
	return LocationDataActionCreateFn(dto, query)
}

func LocationDataActionUpdate(
	query workspaces.QueryDSL,
	fields *LocationDataEntity,
) (*LocationDataEntity, *workspaces.IError) {
	return LocationDataActionUpdateFn(query, fields)
}
