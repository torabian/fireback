package cms
import "github.com/torabian/fireback/modules/workspaces"
func PostTagActionCreate(
	dto *PostTagEntity, query workspaces.QueryDSL,
) (*PostTagEntity, *workspaces.IError) {
	return PostTagActionCreateFn(dto, query)
}
func PostTagActionUpdate(
	query workspaces.QueryDSL,
	fields *PostTagEntity,
) (*PostTagEntity, *workspaces.IError) {
	return PostTagActionUpdateFn(query, fields)
}