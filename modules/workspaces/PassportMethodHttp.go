package workspaces

func init() {

	AppendPassportMethodRouter = func(r *[]Module2Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */

		// r.DELETE("/passportMethod/:uniqueId", WithAuthorization([]string{PERM_ROOT_PASSPORTMETHOD_DELETE}), HttpRemovePassportMethod)
	}
}
