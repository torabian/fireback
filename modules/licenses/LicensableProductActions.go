package licenses

import "pixelplux.com/fireback/modules/workspaces"

func LicensableProductActionCreate(
	dto *LicensableProductEntity, query workspaces.QueryDSL,
) (*LicensableProductEntity, *workspaces.IError) {
	return LicensableProductActionCreateFn(dto, query)
}

func LicensableProductActionUpdate(
	query workspaces.QueryDSL,
	fields *LicensableProductEntity,
) (*LicensableProductEntity, *workspaces.IError) {
	return LicensableProductActionUpdateFn(query, fields)
}

/**
*	This action generates the product publicly
**/
func ProductActionGenerate(
	dto *LicensableProductEntity, query workspaces.QueryDSL,
) (*LicensableProductEntity, *workspaces.IError) {

	info, err := GenertePrivatePublicKeySet()

	if err != nil {

		return nil, workspaces.GormErrorToIError(err)
	}

	dto.PrivateKey = &info.PrivateKey
	dto.PublicKey = &info.PublicKey

	return LicensableProductActionCreateFn(dto, query)
}
