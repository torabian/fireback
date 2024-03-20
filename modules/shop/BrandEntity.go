package shop
import "github.com/torabian/fireback/modules/workspaces"
func BrandActionCreate(
	dto *BrandEntity, query workspaces.QueryDSL,
) (*BrandEntity, *workspaces.IError) {
	return BrandActionCreateFn(dto, query)
}
func BrandActionUpdate(
	query workspaces.QueryDSL,
	fields *BrandEntity,
) (*BrandEntity, *workspaces.IError) {
	return BrandActionUpdateFn(query, fields)
}