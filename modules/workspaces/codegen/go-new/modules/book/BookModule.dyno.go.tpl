package book
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_BOOK_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/book/*",
}
var ALL_PERM_BOOK_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_BOOK_EVERYTHING,
}
