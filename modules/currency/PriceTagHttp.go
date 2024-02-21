package currency

import "github.com/torabian/fireback/modules/workspaces"

func init() {

	AppendPriceTagRouter = func(r *[]workspaces.Module2Action) {
		/*
			 *   Implement the http routes here, with your new actions created
			 *   This file won't be updated, your code stays in this file.

			 *r = append(*r, workspaces.Module2Action{
				Method: "POST",
				Url:    "/license/from-plan/:uniqueId",
				Handlers: []gin.HandlerFunc{
					HttpActivateLicenseFromPlanId,
				},
				RequestEntity:  &LicenseFromPlanIdDto{},
				ResponseEntity: &LicenseEntity{},
			})

		*/
	}
}
