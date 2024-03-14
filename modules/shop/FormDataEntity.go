package shop
import "github.com/torabian/fireback/modules/workspaces"
func FormDataActionCreate(
	dto *FormDataEntity, query workspaces.QueryDSL,
) (*FormDataEntity, *workspaces.IError) {
	return FormDataActionCreateFn(dto, query)
}
func FormDataActionUpdate(
	query workspaces.QueryDSL,
	fields *FormDataEntity,
) (*FormDataEntity, *workspaces.IError) {
	return FormDataActionUpdateFn(query, fields)
}