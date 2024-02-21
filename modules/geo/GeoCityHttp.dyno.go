package geo

import (
	"github.com/gin-gonic/gin"

	"pixelplux.com/fireback/modules/workspaces"
)

/**
*	Override this function on GeoCityHttp.go,
*	In order to add your own http
**/
var AppendGeoCityRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoCity(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoCityActionCreate)
}

func HttpExportStreamGeoCity(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoCityActionExport)
}

func HttpQueryGeoCitys(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoCityActionQuery)
}

func HttpGetOneGeoCity(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoCityActionGetOne)
}

func HttpRemoveGeoCity(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoCityActionRemove)
}

func HttpUpdateGeoCity(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoCityActionUpdate)
}

func HttpBulkUpdateGeoCity(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoCityActionBulkUpdate)
}

func GetGeoCityModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/geoCitys",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_QUERY}),

				HttpQueryGeoCitys,
			},
			Format:         "QUERY",
			Action:         GeoCityActionQuery,
			ResponseEntity: &[]GeoCityEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoCitys/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_QUERY}),

				HttpExportStreamGeoCity,
			},
			Format:         "QUERY",
			Action:         GeoCityActionExport,
			ResponseEntity: &[]GeoCityEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoCity/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_QUERY}),

				HttpGetOneGeoCity,
			},
			Format:         "GET_ONE",
			Action:         GeoCityActionGetOne,
			ResponseEntity: &GeoCityEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoCity",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_CREATE}),
				HttpPostGeoCity,
			},
			Action:         GeoCityActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoCityEntity{},
			ResponseEntity: &GeoCityEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoCity",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_UPDATE}),
				HttpUpdateGeoCity,
			},
			Action:         GeoCityActionUpdate,
			RequestEntity:  &GeoCityEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoCityEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoCitys",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_UPDATE}),
				HttpBulkUpdateGeoCity,
			},
			Action:         GeoCityActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoCityEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoCityEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoCity",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCITY_DELETE}),
				HttpRemoveGeoCity,
			},
			Action:         GeoCityActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoCityEntity{},
		},
	}

	// Append user defined functions
	AppendGeoCityRouter(&routes)

	return routes

}

func CreateGeoCityRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoCityModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoCityEntityJsonSchema, "geoCity-http", "geo")
	workspaces.WriteEntitySchema("GeoCity", GeoCityEntityJsonSchema, "geo")

	return httpRoutes
}
