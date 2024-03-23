package shop
import "github.com/torabian/fireback/modules/workspaces"
func OrderStatusActionCreate(
	dto *OrderStatusEntity, query workspaces.QueryDSL,
) (*OrderStatusEntity, *workspaces.IError) {
	return OrderStatusActionCreateFn(dto, query)
}
func OrderStatusActionUpdate(
	query workspaces.QueryDSL,
	fields *OrderStatusEntity,
) (*OrderStatusEntity, *workspaces.IError) {
	return OrderStatusActionUpdateFn(query, fields)
}