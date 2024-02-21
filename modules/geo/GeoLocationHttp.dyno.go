package geo

import (
	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/workspaces"
)

/**
*	Override this function on GeoLocationHttp.go,
*	In order to add your own http
**/
var AppendGeoLocationRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoLocation(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoLocationActionCreate)
}

func HttpExportStreamGeoLocation(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoLocationActionExport)
}

func HttpQueryGeoLocations(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoLocationActionQuery)
}

func HttpCteQueryGeoLocations(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoLocationActionCteQuery)
}

func HttpGetOneGeoLocation(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoLocationActionGetOne)
}

func HttpRemoveGeoLocation(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoLocationActionRemove)
}

func HttpUpdateGeoLocation(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoLocationActionUpdate)
}

func HttpBulkUpdateGeoLocation(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoLocationActionBulkUpdate)
}

func GetGeoLocationModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/cteGeoLocations",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_QUERY}),

				HttpCteQueryGeoLocations,
			},
			Format:         "QUERY",
			Action:         GeoLocationActionCteQuery,
			ResponseEntity: &[]GeoLocationEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoLocations",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_QUERY}),

				HttpQueryGeoLocations,
			},
			Format:         "QUERY",
			Action:         GeoLocationActionQuery,
			ResponseEntity: &[]GeoLocationEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoLocations/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_QUERY}),

				HttpExportStreamGeoLocation,
			},
			Format:         "QUERY",
			Action:         GeoLocationActionExport,
			ResponseEntity: &[]GeoLocationEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoLocation/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_QUERY}),

				HttpGetOneGeoLocation,
			},
			Format:         "GET_ONE",
			Action:         GeoLocationActionGetOne,
			ResponseEntity: &GeoLocationEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoLocation",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_CREATE}),
				HttpPostGeoLocation,
			},
			Action:         GeoLocationActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoLocationEntity{},
			ResponseEntity: &GeoLocationEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoLocation",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_UPDATE}),
				HttpUpdateGeoLocation,
			},
			Action:         GeoLocationActionUpdate,
			RequestEntity:  &GeoLocationEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoLocationEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoLocations",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_UPDATE}),
				HttpBulkUpdateGeoLocation,
			},
			Action:         GeoLocationActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoLocationEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoLocationEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoLocation",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATION_DELETE}),
				HttpRemoveGeoLocation,
			},
			Action:         GeoLocationActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoLocationEntity{},
		},
	}

	// Append user defined functions
	AppendGeoLocationRouter(&routes)

	return routes

}

func CreateGeoLocationRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoLocationModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoLocationEntityJsonSchema, "geoLocation-http", "geo")
	workspaces.WriteEntitySchema("GeoLocation", GeoLocationEntityJsonSchema, "geo")

	return httpRoutes
}
