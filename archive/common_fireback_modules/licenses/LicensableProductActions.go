package licenses

import "github.com/torabian/fireback/modules/fireback"

func LicensableProductActionCreate(
	dto *LicensableProductEntity, query fireback.QueryDSL,
) (*LicensableProductEntity, *fireback.IError) {
	return LicensableProductActionCreateFn(dto, query)
}

func LicensableProductActionUpdate(
	query fireback.QueryDSL,
	fields *LicensableProductEntity,
) (*LicensableProductEntity, *fireback.IError) {
	return LicensableProductActionUpdateFn(query, fields)
}

/**
*	This action generates the product publicly
**/
func ProductActionGenerate(
	dto *LicensableProductEntity, query fireback.QueryDSL,
) (*LicensableProductEntity, *fireback.IError) {

	info, err := GenertePrivatePublicKeySet()

	if err != nil {

		return nil, fireback.GormErrorToIError(err)
	}

	dto.PrivateKey = &info.PrivateKey
	dto.PublicKey = &info.PublicKey

	return LicensableProductActionCreateFn(dto, query)
}
