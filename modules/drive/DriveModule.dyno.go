package drive
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_DRIVE_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/drive/*",
}
var ALL_PERM_DRIVE_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_DRIVE_EVERYTHING,
}
