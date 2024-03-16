package shop
import "github.com/torabian/fireback/modules/workspaces"
func ProductActionCreate(
	dto *ProductEntity, query workspaces.QueryDSL,
) (*ProductEntity, *workspaces.IError) {
	return ProductActionCreateFn(dto, query)
}
func ProductActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductEntity,
) (*ProductEntity, *workspaces.IError) {
	return ProductActionUpdateFn(query, fields)
}