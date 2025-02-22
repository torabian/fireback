package licenses

import "github.com/torabian/fireback/modules/workspaces"

func LicenseActionCreate(
	dto *LicenseEntity, query workspaces.QueryDSL,
) (*LicenseEntity, *workspaces.IError) {
	return LicenseActionCreateFn(dto, query)
}

func LicenseActionUpdate(
	query workspaces.QueryDSL,
	fields *LicenseEntity,
) (*LicenseEntity, *workspaces.IError) {
	return LicenseActionUpdateFn(query, fields)
}
