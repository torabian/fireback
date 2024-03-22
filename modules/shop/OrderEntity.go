package shop
import "github.com/torabian/fireback/modules/workspaces"
func OrderActionCreate(
	dto *OrderEntity, query workspaces.QueryDSL,
) (*OrderEntity, *workspaces.IError) {
	return OrderActionCreateFn(dto, query)
}
func OrderActionUpdate(
	query workspaces.QueryDSL,
	fields *OrderEntity,
) (*OrderEntity, *workspaces.IError) {
	return OrderActionUpdateFn(query, fields)
}