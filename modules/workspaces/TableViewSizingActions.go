package workspaces

func TableViewSizingActionCreate(
	dto *TableViewSizingEntity, query QueryDSL,
) (*TableViewSizingEntity, *IError) {
	return TableViewSizingActionCreateFn(dto, query)
}

func TableViewSizingActionUpdate(
	query QueryDSL,
	fields *TableViewSizingEntity,
) (*TableViewSizingEntity, *IError) {
	return TableViewSizingActionUpdateFn(query, fields)
}
