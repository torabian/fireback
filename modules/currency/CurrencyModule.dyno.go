package currency
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_CURRENCY_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/currency/*",
}
var ALL_PERM_CURRENCY_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_CURRENCY_EVERYTHING,
}
