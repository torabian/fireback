package geo

import (
	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/workspaces"
)

/**
*	Override this function on GeoProvinceHttp.go,
*	In order to add your own http
**/
var AppendGeoProvinceRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostGeoProvince(c *gin.Context) {
	workspaces.HttpPostEntity(c, GeoProvinceActionCreate)
}

func HttpExportStreamGeoProvince(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, GeoProvinceActionExport)
}

func HttpQueryGeoProvinces(c *gin.Context) {
	workspaces.HttpQueryEntity(c, GeoProvinceActionQuery)
}

func HttpGetOneGeoProvince(c *gin.Context) {
	workspaces.HttpGetEntity(c, GeoProvinceActionGetOne)
}

func HttpRemoveGeoProvince(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, GeoProvinceActionRemove)
}

func HttpUpdateGeoProvince(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, GeoProvinceActionUpdate)
}

func HttpBulkUpdateGeoProvince(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, GeoProvinceActionBulkUpdate)
}

func GetGeoProvinceModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/geoProvinces",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_QUERY}),

				HttpQueryGeoProvinces,
			},
			Format:         "QUERY",
			Action:         GeoProvinceActionQuery,
			ResponseEntity: &[]GeoProvinceEntity{},
		},
		{
			Method: "GET",
			Url:    "/geoProvinces/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_QUERY}),

				HttpExportStreamGeoProvince,
			},
			Format:         "QUERY",
			Action:         GeoProvinceActionExport,
			ResponseEntity: &[]GeoProvinceEntity{},
		},

		{
			Method: "GET",
			Url:    "/geoProvince/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_QUERY}),

				HttpGetOneGeoProvince,
			},
			Format:         "GET_ONE",
			Action:         GeoProvinceActionGetOne,
			ResponseEntity: &GeoProvinceEntity{},
		},

		{
			Method: "POST",
			Url:    "/geoProvince",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_CREATE}),
				HttpPostGeoProvince,
			},
			Action:         GeoProvinceActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &GeoProvinceEntity{},
			ResponseEntity: &GeoProvinceEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/geoProvince",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_UPDATE}),
				HttpUpdateGeoProvince,
			},
			Action:         GeoProvinceActionUpdate,
			RequestEntity:  &GeoProvinceEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &GeoProvinceEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/geoProvinces",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_UPDATE}),
				HttpBulkUpdateGeoProvince,
			},
			Action:         GeoProvinceActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[GeoProvinceEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[GeoProvinceEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/geoProvince",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_GEOPROVINCE_DELETE}),
				HttpRemoveGeoProvince,
			},
			Action:         GeoProvinceActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &GeoProvinceEntity{},
		},
	}

	// Append user defined functions
	AppendGeoProvinceRouter(&routes)

	return routes

}

func CreateGeoProvinceRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetGeoProvinceModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, GeoProvinceEntityJsonSchema, "geoProvince-http", "geo")
	workspaces.WriteEntitySchema("GeoProvince", GeoProvinceEntityJsonSchema, "geo")

	return httpRoutes
}
