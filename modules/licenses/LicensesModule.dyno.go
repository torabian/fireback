package licenses
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_LICENSES_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/licenses/*",
}
var ALL_PERM_LICENSES_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_LICENSES_EVERYTHING,
}
