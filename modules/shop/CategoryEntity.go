package shop
import "github.com/torabian/fireback/modules/workspaces"
func CategoryActionCreate(
	dto *CategoryEntity, query workspaces.QueryDSL,
) (*CategoryEntity, *workspaces.IError) {
	return CategoryActionCreateFn(dto, query)
}
func CategoryActionUpdate(
	query workspaces.QueryDSL,
	fields *CategoryEntity,
) (*CategoryEntity, *workspaces.IError) {
	return CategoryActionUpdateFn(query, fields)
}