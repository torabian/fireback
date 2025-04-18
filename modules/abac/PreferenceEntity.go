package abac

import "github.com/torabian/fireback/modules/workspaces"

func PreferenceActionCreate(
	dto *PreferenceEntity, query workspaces.QueryDSL,
) (*PreferenceEntity, *workspaces.IError) {
	return PreferenceActionCreateFn(dto, query)
}

func PreferenceActionUpdate(
	query workspaces.QueryDSL,
	fields *PreferenceEntity,
) (*PreferenceEntity, *workspaces.IError) {
	return PreferenceActionUpdateFn(query, fields)
}
