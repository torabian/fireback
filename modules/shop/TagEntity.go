package shop
import "github.com/torabian/fireback/modules/workspaces"
func TagActionCreate(
	dto *TagEntity, query workspaces.QueryDSL,
) (*TagEntity, *workspaces.IError) {
	return TagActionCreateFn(dto, query)
}
func TagActionUpdate(
	query workspaces.QueryDSL,
	fields *TagEntity,
) (*TagEntity, *workspaces.IError) {
	return TagActionUpdateFn(query, fields)
}