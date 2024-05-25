package book
import "github.com/torabian/fireback/modules/workspaces"
func BookActionCreate(
	dto *BookEntity, query workspaces.QueryDSL,
) (*BookEntity, *workspaces.IError) {
	return BookActionCreateFn(dto, query)
}
func BookActionUpdate(
	query workspaces.QueryDSL,
	fields *BookEntity,
) (*BookEntity, *workspaces.IError) {
	return BookActionUpdateFn(query, fields)
}