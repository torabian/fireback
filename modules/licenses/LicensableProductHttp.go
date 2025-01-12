package licenses

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
)

func init() {

	AppendLicensableProductRouter = func(r *[]workspaces.Module3Action) {
		/*
		 *   Implement the http routes here, with your new actions created
		 *   This file won't be updated, your code stays in this file.
		 */
		*r = append(*r, workspaces.Module3Action{
			Method: "POST",
			Url:    "/licensableProducts/generate",
			Handlers: []gin.HandlerFunc{
				func(ctx *gin.Context) {
					workspaces.HttpPostEntity(ctx, ProductActionGenerate)
				},
			},
			RequestEntity:  &LicensableProductEntity{},
			ResponseEntity: &LicensableProductEntity{},
			In: &workspaces.Module3ActionBody{
				Dto: "LicensableProductEntity",
			},
			Out: &workspaces.Module3ActionBody{
				Dto: "LicensableProductEntity",
			},
		})

	}
}
