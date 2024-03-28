package geo
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_GEO_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/geo/*",
}
var ALL_PERM_GEO_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_GEO_EVERYTHING,
}
