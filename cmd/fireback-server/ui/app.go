package ui

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/shop"
	"github.com/torabian/fireback/modules/workspaces"
)

// import (
// 	"embed"

// 	"github.com/gin-gonic/gin"
// 	"github.com/torabian/fireback/modules/shop"
// 	"github.com/torabian/fireback/modules/workspaces"
// )

//go:embed all:*
var UI embed.FS

func QueryHelper[T any](fn workspaces.QueryableAction[T], query workspaces.QueryDSL) gin.H {
	products, qrm, err := fn(query)
	return gin.H{
		"items": products,
		"qrm":   qrm,
		"err":   err,
	}
}

func Bootstrap(e *gin.Engine) {

	e.GET("/", func(ctx *gin.Context) {
		query := workspaces.TemplateQueryDSL(ctx)
		query.Deep = true
		workspaces.RenderTemplateToGin(ctx, "index.tpl", UI, gin.H{
			"products": QueryHelper[shop.ProductSubmissionEntity](shop.ProductSubmissionActionQuery, query),
		})
	})

	e.GET("/products-inline", func(ctx *gin.Context) {
		query := workspaces.TemplateQueryDSL(ctx)
		query.Deep = true
		workspaces.RenderTemplateToGin(ctx, "partials/products-inline.tpl", UI, QueryHelper(shop.ProductSubmissionActionQuery, query))
	})
}
