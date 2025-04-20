package licenses

import (
	currency "github.com/torabian/fireback/modules/currency"
	"github.com/torabian/fireback/modules/fireback"
)

func ProductPlanActionCreate(
	dto *ProductPlanEntity, query fireback.QueryDSL,
) (*ProductPlanEntity, *fireback.IError) {
	return ProductPlanActionCreateFn(dto, query)
}

func ProductPlanActionUpdate(
	query fireback.QueryDSL,
	fields *ProductPlanEntity,
) (*ProductPlanEntity, *fireback.IError) {

	// Added extra logic, probably should move to fireback
	if fields.PriceTag != nil {
		currency.PriceTagActionUpdate(query, fields.PriceTag)
	}

	return ProductPlanActionUpdateFn(query, fields)
}
