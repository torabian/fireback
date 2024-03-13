package shop
import "github.com/torabian/fireback/modules/workspaces"
func ProductTagActionCreate(
	dto *ProductTagEntity, query workspaces.QueryDSL,
) (*ProductTagEntity, *workspaces.IError) {
	return ProductTagActionCreateFn(dto, query)
}
func ProductTagActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductTagEntity,
) (*ProductTagEntity, *workspaces.IError) {
	return ProductTagActionUpdateFn(query, fields)
}