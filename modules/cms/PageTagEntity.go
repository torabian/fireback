package cms
import "github.com/torabian/fireback/modules/workspaces"
func PageTagActionCreate(
	dto *PageTagEntity, query workspaces.QueryDSL,
) (*PageTagEntity, *workspaces.IError) {
	return PageTagActionCreateFn(dto, query)
}
func PageTagActionUpdate(
	query workspaces.QueryDSL,
	fields *PageTagEntity,
) (*PageTagEntity, *workspaces.IError) {
	return PageTagActionUpdateFn(query, fields)
}