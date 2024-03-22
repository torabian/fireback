package shop
import "github.com/torabian/fireback/modules/workspaces"
func DiscountScopeActionCreate(
	dto *DiscountScopeEntity, query workspaces.QueryDSL,
) (*DiscountScopeEntity, *workspaces.IError) {
	return DiscountScopeActionCreateFn(dto, query)
}
func DiscountScopeActionUpdate(
	query workspaces.QueryDSL,
	fields *DiscountScopeEntity,
) (*DiscountScopeEntity, *workspaces.IError) {
	return DiscountScopeActionUpdateFn(query, fields)
}