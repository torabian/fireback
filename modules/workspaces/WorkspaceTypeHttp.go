package workspaces

import (
	"github.com/gin-gonic/gin"
)

func init() {

	AppendWorkspaceTypeRouter = func(r *[]Module2Action) {

		*r = append(*r, Module2Action{
			Method: "GET",
			Url:    "/publicWorkspaceTypes",
			Handlers: []gin.HandlerFunc{
				// Intentionally available publicly
				func(ctx *gin.Context) {
					HttpQueryEntity(ctx, WorkspaceTypeActionPublicQuery)
				},
			},
			Action:         WorkspaceTypeActionPublicQuery,
			Format:         "QUERY",
			ResponseEntity: &[]WorkspaceTypeEntity{},
			Out: &Module2ActionBody{
				Entity: "WorkspaceTypeEntity",
			},
		})

	}
}
