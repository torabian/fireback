package abac

import abacdefs "github.com/torabian/fireback/modules/abac/defs"

var ANONYMOUS_AUTHENTICATION = "anonymous"

var ROOT_ALL_ACCESS = "root.*"
var ROOT_ALL_MODULES = "root.modules.*"

var OS_SIGNIN_CAPABILITIES []*abacdefs.CapabilityEntity = []*abacdefs.CapabilityEntity{
	{UniqueId: ROOT_ALL_ACCESS, Name: "Root"},
}
