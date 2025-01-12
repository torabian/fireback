package licenses

import "github.com/torabian/fireback/modules/workspaces"

func init() {

	AppendActivationKeyRouter = func(r *[]workspaces.Module3Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */

		// r.DELETE("/activationKey/:uniqueId", workspaces.WithAuthorization([]string{PERM_ROOT_ACTIVATIONKEY_DELETE}), HttpRemoveActivationKey)
	}
}
