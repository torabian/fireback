package shop
import "github.com/torabian/fireback/modules/workspaces"
func PaymentStatusActionCreate(
	dto *PaymentStatusEntity, query workspaces.QueryDSL,
) (*PaymentStatusEntity, *workspaces.IError) {
	return PaymentStatusActionCreateFn(dto, query)
}
func PaymentStatusActionUpdate(
	query workspaces.QueryDSL,
	fields *PaymentStatusEntity,
) (*PaymentStatusEntity, *workspaces.IError) {
	return PaymentStatusActionUpdateFn(query, fields)
}