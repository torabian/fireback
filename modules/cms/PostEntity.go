package cms
import "github.com/torabian/fireback/modules/workspaces"
func PostActionCreate(
	dto *PostEntity, query workspaces.QueryDSL,
) (*PostEntity, *workspaces.IError) {
	return PostActionCreateFn(dto, query)
}
func PostActionUpdate(
	query workspaces.QueryDSL,
	fields *PostEntity,
) (*PostEntity, *workspaces.IError) {
	return PostActionUpdateFn(query, fields)
}