package worldtimezone

import "github.com/torabian/fireback/modules/workspaces"

func init() {

	AppendTimezoneGroupRouter = func(r *[]workspaces.Module2Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */

		// r.DELETE("/TimezoneGroup/:uniqueId", workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_DELETE}), HttpRemoveTimezoneGroup)
	}
}
