package geo

import (
	"github.com/gin-gonic/gin"

	"pixelplux.com/fireback/modules/workspaces"
)

/**
*	Override this function on GeoCountryHttp.go,
*	In order to add your own http
**/
var AppendGeoCountryRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoCountry(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoCountryActionCreate)
}

func HttpExportStreamGeoCountry(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoCountryActionExport)
}

func HttpQueryGeoCountrys(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoCountryActionQuery)
}

func HttpGetOneGeoCountry(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoCountryActionGetOne)
}

func HttpRemoveGeoCountry(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoCountryActionRemove)
}

func HttpUpdateGeoCountry(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoCountryActionUpdate)
}

func HttpBulkUpdateGeoCountry(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoCountryActionBulkUpdate)
}

func GetGeoCountryModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/geoCountrys",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_QUERY}),

				HttpQueryGeoCountrys,
			},
			Format:         "QUERY",
			Action:         GeoCountryActionQuery,
			ResponseEntity: &[]GeoCountryEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoCountrys/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_QUERY}),

				HttpExportStreamGeoCountry,
			},
			Format:         "QUERY",
			Action:         GeoCountryActionExport,
			ResponseEntity: &[]GeoCountryEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoCountry/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_QUERY}),

				HttpGetOneGeoCountry,
			},
			Format:         "GET_ONE",
			Action:         GeoCountryActionGetOne,
			ResponseEntity: &GeoCountryEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoCountry",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_CREATE}),
				HttpPostGeoCountry,
			},
			Action:         GeoCountryActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoCountryEntity{},
			ResponseEntity: &GeoCountryEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoCountry",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_UPDATE}),
				HttpUpdateGeoCountry,
			},
			Action:         GeoCountryActionUpdate,
			RequestEntity:  &GeoCountryEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoCountryEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoCountrys",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_UPDATE}),
				HttpBulkUpdateGeoCountry,
			},
			Action:         GeoCountryActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoCountryEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoCountryEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoCountry",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOCOUNTRY_DELETE}),
				HttpRemoveGeoCountry,
			},
			Action:         GeoCountryActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoCountryEntity{},
		},
	}

	// Append user defined functions
	AppendGeoCountryRouter(&routes)

	return routes

}

func CreateGeoCountryRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoCountryModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoCountryEntityJsonSchema, "geoCountry-http", "geo")
	workspaces.WriteEntitySchema("GeoCountry", GeoCountryEntityJsonSchema, "geo")

	return httpRoutes
}
