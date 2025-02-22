package workspaces

func TimezoneGroupActionCreate(
	dto *TimezoneGroupEntity, query QueryDSL,
) (*TimezoneGroupEntity, *IError) {
	return TimezoneGroupActionCreateFn(dto, query)
}
func TimezoneGroupActionUpdate(
	query QueryDSL,
	fields *TimezoneGroupEntity,
) (*TimezoneGroupEntity, *IError) {
	return TimezoneGroupActionUpdateFn(query, fields)
}
