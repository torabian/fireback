package cms
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_CMS_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/cms/*",
}
var ALL_PERM_CMS_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_CMS_EVERYTHING,
}
