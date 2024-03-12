package device
import "github.com/torabian/fireback/modules/workspaces"
func DeviceActionCreate(
	dto *DeviceEntity, query workspaces.QueryDSL,
) (*DeviceEntity, *workspaces.IError) {
	return DeviceActionCreateFn(dto, query)
}
func DeviceActionUpdate(
	query workspaces.QueryDSL,
	fields *DeviceEntity,
) (*DeviceEntity, *workspaces.IError) {
	return DeviceActionUpdateFn(query, fields)
}