package worldtimezone

import "pixelplux.com/fireback/modules/workspaces"

func TimezoneGroupActionCreate(
	dto *TimezoneGroupEntity, query workspaces.QueryDSL,
) (*TimezoneGroupEntity, *workspaces.IError) {
	return TimezoneGroupActionCreateFn(dto, query)
}

func TimezoneGroupActionUpdate(
	query workspaces.QueryDSL,
	fields *TimezoneGroupEntity,
) (*TimezoneGroupEntity, *workspaces.IError) {
	return TimezoneGroupActionUpdateFn(query, fields)
}
