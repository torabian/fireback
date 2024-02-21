package licenses

import (
	"github.com/gin-gonic/gin"
	"pixelplux.com/fireback/modules/workspaces"
)

func HttpActivateLicenseFromPlanId(c *gin.Context) {
	workspaces.HttpPostEntity(c, LicenseActionFromPlanId)
}

func init() {

	AppendLicenseRouter = func(r *[]workspaces.Module2Action) {

		*r = append(*r, workspaces.Module2Action{
			Method: "POST",
			Url:    "/license/from-plan/:uniqueId",
			Handlers: []gin.HandlerFunc{
				HttpActivateLicenseFromPlanId,
			},
			RequestEntity:  &LicenseFromPlanIdDto{},
			ResponseEntity: &LicenseEntity{},
		})

	}
}
