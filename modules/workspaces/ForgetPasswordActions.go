package workspaces

func ForgetPasswordActionCreate(
	dto *ForgetPasswordEntity, query QueryDSL,
) (*ForgetPasswordEntity, *IError) {
	return ForgetPasswordActionCreateFn(dto, query)
}

func ForgetPasswordActionUpdate(
	query QueryDSL,
	fields *ForgetPasswordEntity,
) (*ForgetPasswordEntity, *IError) {
	return ForgetPasswordActionUpdateFn(query, fields)
}
