package worldtimezone
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_WORLDTIMEZONE_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/worldtimezone/*",
}
var ALL_PERM_WORLDTIMEZONE_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_WORLDTIMEZONE_EVERYTHING,
}
