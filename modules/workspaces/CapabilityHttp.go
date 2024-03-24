package workspaces

import (
	"github.com/gin-gonic/gin"
)

func init() {

	AppendCapabilityRouter = func(r *[]Module2Action) {

		*r = append(*r, Module2Action{
			Method: "GET",
			Url:    "/capabilitiesTree",
			Handlers: []gin.HandlerFunc{
				WithAuthorization([]PermissionInfo{PERM_ROOT_CAPABILITY_QUERY}),
				func(c *gin.Context) {
					HttpGetEntity(c, CapabilityActionGetTree)
				},
			},
			Action:         CapabilityActionGetTree,
			Format:         "GET_ONE",
			ResponseEntity: &CapabilitiesResult{},
		})

	}
}
