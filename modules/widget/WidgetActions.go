package widget

import "github.com/torabian/fireback/modules/workspaces"

func WidgetActionCreate(
	dto *WidgetEntity, query workspaces.QueryDSL,
) (*WidgetEntity, *workspaces.IError) {
	return WidgetActionCreateFn(dto, query)
}

func WidgetActionUpdate(
	query workspaces.QueryDSL,
	fields *WidgetEntity,
) (*WidgetEntity, *workspaces.IError) {
	return WidgetActionUpdateFn(query, fields)
}
