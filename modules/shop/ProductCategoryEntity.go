package shop
import "github.com/torabian/fireback/modules/workspaces"
func ProductCategoryActionCreate(
	dto *ProductCategoryEntity, query workspaces.QueryDSL,
) (*ProductCategoryEntity, *workspaces.IError) {
	return ProductCategoryActionCreateFn(dto, query)
}
func ProductCategoryActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductCategoryEntity,
) (*ProductCategoryEntity, *workspaces.IError) {
	return ProductCategoryActionUpdateFn(query, fields)
}