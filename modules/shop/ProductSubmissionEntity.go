package shop
import "github.com/torabian/fireback/modules/workspaces"
func ProductSubmissionActionCreate(
	dto *ProductSubmissionEntity, query workspaces.QueryDSL,
) (*ProductSubmissionEntity, *workspaces.IError) {
	return ProductSubmissionActionCreateFn(dto, query)
}
func ProductSubmissionActionUpdate(
	query workspaces.QueryDSL,
	fields *ProductSubmissionEntity,
) (*ProductSubmissionEntity, *workspaces.IError) {
	return ProductSubmissionActionUpdateFn(query, fields)
}