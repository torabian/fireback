package licenses

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {

	AppendLicensableProductRouter = func(r *[]fireback.Module3Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */
		*r = append(*r, fireback.Module3Action{
			Method: "POST",
			Url:    "/licensableProducts/generate",
			Handlers: []gin.HandlerFunc{
				func(ctx *gin.Context) {
					fireback.HttpPostEntity(ctx, ProductActionGenerate)
				},
			},
			RequestEntity:  &LicensableProductEntity{},
			ResponseEntity: &LicensableProductEntity{},
			In: &fireback.Module3ActionBody{
				Dto: "LicensableProductEntity",
			},
			Out: &fireback.Module3ActionBody{
				Dto: "LicensableProductEntity",
			},
		})

	}
}
