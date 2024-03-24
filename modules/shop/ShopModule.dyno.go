package shop
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_SHOP_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/shop/*",
}
var ALL_PERM_SHOP_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_SHOP_EVERYTHING,
}
