package geo

import (
	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/workspaces"
)

/**
*	Override this function on GeoStateHttp.go,
*	In order to add your own http
**/
var AppendGeoStateRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoState(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoStateActionCreate)
}

func HttpExportStreamGeoState(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoStateActionExport)
}

func HttpQueryGeoStates(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoStateActionQuery)
}

func HttpGetOneGeoState(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoStateActionGetOne)
}

func HttpRemoveGeoState(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoStateActionRemove)
}

func HttpUpdateGeoState(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoStateActionUpdate)
}

func HttpBulkUpdateGeoState(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoStateActionBulkUpdate)
}

func GetGeoStateModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/geoStates",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_QUERY}),

				HttpQueryGeoStates,
			},
			Format:         "QUERY",
			Action:         GeoStateActionQuery,
			ResponseEntity: &[]GeoStateEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoStates/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_QUERY}),

				HttpExportStreamGeoState,
			},
			Format:         "QUERY",
			Action:         GeoStateActionExport,
			ResponseEntity: &[]GeoStateEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoState/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_QUERY}),

				HttpGetOneGeoState,
			},
			Format:         "GET_ONE",
			Action:         GeoStateActionGetOne,
			ResponseEntity: &GeoStateEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoState",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_CREATE}),
				HttpPostGeoState,
			},
			Action:         GeoStateActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoStateEntity{},
			ResponseEntity: &GeoStateEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoState",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_UPDATE}),
				HttpUpdateGeoState,
			},
			Action:         GeoStateActionUpdate,
			RequestEntity:  &GeoStateEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoStateEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoStates",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_UPDATE}),
				HttpBulkUpdateGeoState,
			},
			Action:         GeoStateActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoStateEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoStateEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoState",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOSTATE_DELETE}),
				HttpRemoveGeoState,
			},
			Action:         GeoStateActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoStateEntity{},
		},
	}

	// Append user defined functions
	AppendGeoStateRouter(&routes)

	return routes

}

func CreateGeoStateRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoStateModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoStateEntityJsonSchema, "geoState-http", "geo")
	workspaces.WriteEntitySchema("GeoState", GeoStateEntityJsonSchema, "geo")

	return httpRoutes
}
