package abac

import "github.com/torabian/fireback/modules/fireback"

func PreferenceActionCreate(
	dto *PreferenceEntity, query fireback.QueryDSL,
) (*PreferenceEntity, *fireback.IError) {
	return PreferenceActionCreateFn(dto, query)
}

func PreferenceActionUpdate(
	query fireback.QueryDSL,
	fields *PreferenceEntity,
) (*PreferenceEntity, *fireback.IError) {
	return PreferenceActionUpdateFn(query, fields)
}
