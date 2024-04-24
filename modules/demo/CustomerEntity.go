package demo
import "github.com/torabian/fireback/modules/workspaces"
func CustomerActionCreate(
	dto *CustomerEntity, query workspaces.QueryDSL,
) (*CustomerEntity, *workspaces.IError) {
	return CustomerActionCreateFn(dto, query)
}
func CustomerActionUpdate(
	query workspaces.QueryDSL,
	fields *CustomerEntity,
) (*CustomerEntity, *workspaces.IError) {
	return CustomerActionUpdateFn(query, fields)
}