package zayshop

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/shop"
	"github.com/torabian/fireback/modules/workspaces"
)

//go:embed all:*
var UI embed.FS

func Bootstrap(e *gin.Engine) {

	e.GET("/", func(ctx *gin.Context) {
		query := workspaces.TemplateQueryDSL(ctx)
		query.Deep = true
		workspaces.RenderTemplateToGin(ctx, "index.html", UI, gin.H{
			"products": workspaces.QueryHelper[shop.ProductSubmissionEntity](shop.ProductSubmissionActionQuery, query),
		})
	})

}
