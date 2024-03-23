package cms
import "github.com/torabian/fireback/modules/workspaces"
func PageCategoryActionCreate(
	dto *PageCategoryEntity, query workspaces.QueryDSL,
) (*PageCategoryEntity, *workspaces.IError) {
	return PageCategoryActionCreateFn(dto, query)
}
func PageCategoryActionUpdate(
	query workspaces.QueryDSL,
	fields *PageCategoryEntity,
) (*PageCategoryEntity, *workspaces.IError) {
	return PageCategoryActionUpdateFn(query, fields)
}