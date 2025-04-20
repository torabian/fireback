package licenses

import "github.com/torabian/fireback/modules/fireback"

func init() {

	AppendProductPlanRouter = func(r *[]fireback.Module3Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */

		// r.DELETE("/productPlan/:uniqueId", fireback.WithAuthorization([]string{PERM_ROOT_PRODUCTPLAN_DELETE}), HttpRemoveProductPlan)
	}
}
