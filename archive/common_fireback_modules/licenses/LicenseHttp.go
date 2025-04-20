package licenses

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

func HttpActivateLicenseFromPlanId(c *gin.Context) {
	fireback.HttpPostEntity(c, LicenseActionFromPlanId)
}

func init() {

	AppendLicenseRouter = func(r *[]fireback.Module3Action) {

		*r = append(*r, fireback.Module3Action{
			Method: "POST",
			Url:    "/license/from-plan/:uniqueId",
			Handlers: []gin.HandlerFunc{
				HttpActivateLicenseFromPlanId,
			},
			RequestEntity:  &LicenseFromPlanIdDto{},
			ResponseEntity: &LicenseEntity{},
			In: &fireback.Module3ActionBody{
				Dto: "LicenseFromPlanIdDto",
			},
			Out: &fireback.Module3ActionBody{
				Dto: "LicenseEntity",
			},
		})

	}
}
