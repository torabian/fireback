package cms
import "github.com/torabian/fireback/modules/workspaces"
func PostCategoryActionCreate(
	dto *PostCategoryEntity, query workspaces.QueryDSL,
) (*PostCategoryEntity, *workspaces.IError) {
	return PostCategoryActionCreateFn(dto, query)
}
func PostCategoryActionUpdate(
	query workspaces.QueryDSL,
	fields *PostCategoryEntity,
) (*PostCategoryEntity, *workspaces.IError) {
	return PostCategoryActionUpdateFn(query, fields)
}