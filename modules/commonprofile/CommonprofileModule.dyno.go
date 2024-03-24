package commonprofile
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_COMMONPROFILE_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/commonprofile/*",
}
var ALL_PERM_COMMONPROFILE_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_COMMONPROFILE_EVERYTHING,
}
