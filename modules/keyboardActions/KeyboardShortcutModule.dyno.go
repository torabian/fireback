package keyboardActions
import "github.com/torabian/fireback/modules/workspaces"
// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_KEYBOARD_SHORTCUT_EVERYTHING = workspaces.PermissionInfo{
  CompleteKey: "root/keyboard-shortcut/*",
}
var ALL_PERM_KEYBOARD_SHORTCUT_MODULE = []workspaces.PermissionInfo{
  PERM_ROOT_KEYBOARD_SHORTCUT_EVERYTHING,
}
