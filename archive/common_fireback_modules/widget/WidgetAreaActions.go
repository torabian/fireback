package widget

import "github.com/torabian/fireback/modules/fireback"

func WidgetAreaActionCreate(
	dto *WidgetAreaEntity, query fireback.QueryDSL,
) (*WidgetAreaEntity, *fireback.IError) {
	return WidgetAreaActionCreateFn(dto, query)
}

func WidgetAreaActionUpdate(
	query fireback.QueryDSL,
	fields *WidgetAreaEntity,
) (*WidgetAreaEntity, *fireback.IError) {
	return WidgetAreaActionUpdateFn(query, fields)
}
