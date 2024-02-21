package licenses

import (
	currency "pixelplux.com/fireback/modules/currency"
	"pixelplux.com/fireback/modules/workspaces"
)

func ProductPlanActionCreate(
	dto *ProductPlanEntity, query workspaces.QueryDSL,
) (*ProductPlanEntity, *workspaces.IError) {
	return ProductPlanActionCreateFn(dto, query)
}

func ProductPlanActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductPlanEntity,
) (*ProductPlanEntity, *workspaces.IError) {

	// Added extra logic, probably should move to fireback
	if fields.PriceTag != nil {
		currency.PriceTagActionUpdate(query, fields.PriceTag)
	}

	return ProductPlanActionUpdateFn(query, fields)
}
