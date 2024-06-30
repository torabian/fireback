package workspaces

func PreferenceActionCreate(
	dto *PreferenceEntity, query QueryDSL,
) (*PreferenceEntity, *IError) {
	return PreferenceActionCreateFn(dto, query)
}

func PreferenceActionUpdate(
	query QueryDSL,
	fields *PreferenceEntity,
) (*PreferenceEntity, *IError) {
	return PreferenceActionUpdateFn(query, fields)
}
