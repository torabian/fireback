package cms
import "github.com/torabian/fireback/modules/workspaces"
func PageActionCreate(
	dto *PageEntity, query workspaces.QueryDSL,
) (*PageEntity, *workspaces.IError) {
	return PageActionCreateFn(dto, query)
}
func PageActionUpdate(
	query workspaces.QueryDSL,
	fields *PageEntity,
) (*PageEntity, *workspaces.IError) {
	return PageActionUpdateFn(query, fields)
}