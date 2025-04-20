package abac

import "github.com/torabian/fireback/modules/fireback"

// Module dynamic things comes here. Don't touch it :D
var PERM_ROOT_DRIVE_EVERYTHING = fireback.PermissionInfo{
	CompleteKey: "root.drive.*",
}
var ALL_PERM_DRIVE_MODULE = []fireback.PermissionInfo{
	PERM_ROOT_DRIVE_EVERYTHING,
}
