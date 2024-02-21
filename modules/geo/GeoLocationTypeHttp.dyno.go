package geo

import (
	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/workspaces"
)

/**
*	Override this function on GeoLocationTypeHttp.go,
*	In order to add your own http
**/
var AppendGeoLocationTypeRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoLocationType(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoLocationTypeActionCreate)
}

func HttpExportStreamGeoLocationType(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoLocationTypeActionExport)
}

func HttpQueryGeoLocationTypes(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoLocationTypeActionQuery)
}

func HttpGetOneGeoLocationType(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoLocationTypeActionGetOne)
}

func HttpRemoveGeoLocationType(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoLocationTypeActionRemove)
}

func HttpUpdateGeoLocationType(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoLocationTypeActionUpdate)
}

func HttpBulkUpdateGeoLocationType(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoLocationTypeActionBulkUpdate)
}

func GetGeoLocationTypeModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/geoLocationTypes",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_QUERY}),

				HttpQueryGeoLocationTypes,
			},
			Format:         "QUERY",
			Action:         GeoLocationTypeActionQuery,
			ResponseEntity: &[]GeoLocationTypeEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoLocationTypes/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_QUERY}),

				HttpExportStreamGeoLocationType,
			},
			Format:         "QUERY",
			Action:         GeoLocationTypeActionExport,
			ResponseEntity: &[]GeoLocationTypeEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoLocationType/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_QUERY}),

				HttpGetOneGeoLocationType,
			},
			Format:         "GET_ONE",
			Action:         GeoLocationTypeActionGetOne,
			ResponseEntity: &GeoLocationTypeEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoLocationType",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_CREATE}),
				HttpPostGeoLocationType,
			},
			Action:         GeoLocationTypeActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoLocationTypeEntity{},
			ResponseEntity: &GeoLocationTypeEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoLocationType",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_UPDATE}),
				HttpUpdateGeoLocationType,
			},
			Action:         GeoLocationTypeActionUpdate,
			RequestEntity:  &GeoLocationTypeEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoLocationTypeEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoLocationTypes",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_UPDATE}),
				HttpBulkUpdateGeoLocationType,
			},
			Action:         GeoLocationTypeActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoLocationTypeEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoLocationTypeEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoLocationType",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOLOCATIONTYPE_DELETE}),
				HttpRemoveGeoLocationType,
			},
			Action:         GeoLocationTypeActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoLocationTypeEntity{},
		},
	}

	// Append user defined functions
	AppendGeoLocationTypeRouter(&routes)

	return routes

}

func CreateGeoLocationTypeRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoLocationTypeModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoLocationTypeEntityJsonSchema, "geoLocationType-http", "geo")
	workspaces.WriteEntitySchema("GeoLocationType", GeoLocationTypeEntityJsonSchema, "geo")

	return httpRoutes
}
