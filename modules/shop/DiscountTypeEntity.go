package shop
import "github.com/torabian/fireback/modules/workspaces"
func DiscountTypeActionCreate(
	dto *DiscountTypeEntity, query workspaces.QueryDSL,
) (*DiscountTypeEntity, *workspaces.IError) {
	return DiscountTypeActionCreateFn(dto, query)
}
func DiscountTypeActionUpdate(
	query workspaces.QueryDSL,
	fields *DiscountTypeEntity,
) (*DiscountTypeEntity, *workspaces.IError) {
	return DiscountTypeActionUpdateFn(query, fields)
}