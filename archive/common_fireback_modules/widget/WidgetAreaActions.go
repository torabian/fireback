package widget

import "github.com/torabian/fireback/modules/workspaces"

func WidgetAreaActionCreate(
	dto *WidgetAreaEntity, query workspaces.QueryDSL,
) (*WidgetAreaEntity, *workspaces.IError) {
	return WidgetAreaActionCreateFn(dto, query)
}

func WidgetAreaActionUpdate(
	query workspaces.QueryDSL,
	fields *WidgetAreaEntity,
) (*WidgetAreaEntity, *workspaces.IError) {
	return WidgetAreaActionUpdateFn(query, fields)
}
