package licenses

import "github.com/torabian/fireback/modules/fireback"

func LicenseActionCreate(
	dto *LicenseEntity, query fireback.QueryDSL,
) (*LicenseEntity, *fireback.IError) {
	return LicenseActionCreateFn(dto, query)
}

func LicenseActionUpdate(
	query fireback.QueryDSL,
	fields *LicenseEntity,
) (*LicenseEntity, *fireback.IError) {
	return LicenseActionUpdateFn(query, fields)
}
