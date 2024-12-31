package workspaces

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InjectReactiveSearch(e *gin.Engine, app *FirebackApp) {
	CastRoutes([]Module2Action{
		{
			Method: "REACTIVE",
			Url:    "/reactiveSearch",
			Handlers: []gin.HandlerFunc{

				// @todo Doesn't search require a security anyway?
				WithSocketAuthorization(&SecurityModel{}, true),
				func(ctx *gin.Context) {
					HttpReactiveQuery(ctx,
						func(query QueryDSL, j chan bool, read chan map[string]interface{}) chan *ReactiveSearchResultDto {

							chanStream := make(chan *ReactiveSearchResultDto)

							go func() {
								// defer close(chanStream)
								fmt.Println("Search providers", app.SearchProviders)
								for _, handler := range app.SearchProviders {
									handler(query, chanStream)
								}

							}()

							return chanStream
						},
					)
				},
			},
			ResponseEntity: &ReactiveSearchResultDto{},
		},
	}, e)
}
