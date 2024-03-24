package widget
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_WIDGET_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/widget/*",
}
var ALL_PERM_WIDGET_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_WIDGET_EVERYTHING,
}
