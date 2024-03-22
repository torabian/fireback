package shop
import "github.com/torabian/fireback/modules/workspaces"
func ShoppingCartActionCreate(
	dto *ShoppingCartEntity, query workspaces.QueryDSL,
) (*ShoppingCartEntity, *workspaces.IError) {
	return ShoppingCartActionCreateFn(dto, query)
}
func ShoppingCartActionUpdate(
	query workspaces.QueryDSL,
	fields *ShoppingCartEntity,
) (*ShoppingCartEntity, *workspaces.IError) {
	return ShoppingCartActionUpdateFn(query, fields)
}