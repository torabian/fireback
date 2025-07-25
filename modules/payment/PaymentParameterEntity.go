package payment

/*
*	Generated by fireback 1.1.27
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import "github.com/torabian/fireback/modules/fireback"

func PaymentParameterActionCreate(
	dto *PaymentParameterEntity, query fireback.QueryDSL,
) (*PaymentParameterEntity, *fireback.IError) {
	return PaymentParameterActionCreateFn(dto, query)
}
func PaymentParameterActionUpdate(
	query fireback.QueryDSL,
	fields *PaymentParameterEntity,
) (*PaymentParameterEntity, *fireback.IError) {
	return PaymentParameterActionUpdateFn(query, fields)
}
