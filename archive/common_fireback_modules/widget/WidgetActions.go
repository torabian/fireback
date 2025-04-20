package widget

import "github.com/torabian/fireback/modules/fireback"

func WidgetActionCreate(
	dto *WidgetEntity, query fireback.QueryDSL,
) (*WidgetEntity, *fireback.IError) {
	return WidgetActionCreateFn(dto, query)
}

func WidgetActionUpdate(
	query fireback.QueryDSL,
	fields *WidgetEntity,
) (*WidgetEntity, *fireback.IError) {
	return WidgetActionUpdateFn(query, fields)
}
