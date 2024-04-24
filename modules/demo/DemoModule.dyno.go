package demo
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_DEMO_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/demo/*",
}
var ALL_PERM_DEMO_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_DEMO_EVERYTHING,
}
