package widget

import "github.com/torabian/fireback/modules/fireback"

func init() {

	AppendWidgetAreaRouter = func(r *[]fireback.Module3Action) {
		/*
			 *   Implement the http routes here, with your new actions created
			 *   This file won't be updated, your code stays in this file.

			 *r = append(*r, fireback.Module3Action{
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
