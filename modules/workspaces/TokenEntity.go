package workspaces

func TokenActionCreate(
	dto *TokenEntity, query QueryDSL,
) (*TokenEntity, *IError) {
	return TokenActionCreateFn(dto, query)
}

func TokenActionUpdate(
	query QueryDSL,
	fields *TokenEntity,
) (*TokenEntity, *IError) {
	return TokenActionUpdateFn(query, fields)
}
