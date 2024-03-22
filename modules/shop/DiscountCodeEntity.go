package shop
import "github.com/torabian/fireback/modules/workspaces"
func DiscountCodeActionCreate(
	dto *DiscountCodeEntity, query workspaces.QueryDSL,
) (*DiscountCodeEntity, *workspaces.IError) {
	return DiscountCodeActionCreateFn(dto, query)
}
func DiscountCodeActionUpdate(
	query workspaces.QueryDSL,
	fields *DiscountCodeEntity,
) (*DiscountCodeEntity, *workspaces.IError) {
	return DiscountCodeActionUpdateFn(query, fields)
}