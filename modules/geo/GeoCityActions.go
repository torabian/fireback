package geo

import "pixelplux.com/fireback/modules/workspaces"

func GeoCityActionCreate(
	dto *GeoCityEntity, query workspaces.QueryDSL,
) (*GeoCityEntity, *workspaces.IError) {
	return GeoCityActionCreateFn(dto, query)
}

func GeoCityActionUpdate(
	query workspaces.QueryDSL,
	fields *GeoCityEntity,
) (*GeoCityEntity, *workspaces.IError) {
	return GeoCityActionUpdateFn(query, fields)
}
