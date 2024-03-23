package shop
import "github.com/torabian/fireback/modules/workspaces"
func PaymentMethodActionCreate(
	dto *PaymentMethodEntity, query workspaces.QueryDSL,
) (*PaymentMethodEntity, *workspaces.IError) {
	return PaymentMethodActionCreateFn(dto, query)
}
func PaymentMethodActionUpdate(
	query workspaces.QueryDSL,
	fields *PaymentMethodEntity,
) (*PaymentMethodEntity, *workspaces.IError) {
	return PaymentMethodActionUpdateFn(query, fields)
}