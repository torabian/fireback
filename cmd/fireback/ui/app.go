package ui

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed all:*
var UI embed.FS

func Bootstrap(e *gin.Engine) {

	// e.GET("/", func(ctx *gin.Context) {
	// 	query := workspaces.TemplateQueryDSL(ctx)
	// 	query.Deep = true
	// 	workspaces.RenderTemplateToGin(ctx, "index.tpl", UI, gin.H{
	// 		"products": workspaces.QueryHelper[shop.ProductSubmissionEntity](shop.ProductSubmissionActionQuery, query),
	// 	})
	// })

	// e.GET("/products-inline", func(ctx *gin.Context) {
	// 	query := workspaces.TemplateQueryDSL(ctx)
	// 	query.Deep = true
	// 	workspaces.RenderTemplateToGin(ctx, "partials/products-inline.tpl", UI, workspaces.QueryHelper(shop.ProductSubmissionActionQuery, query))
	// })
}
