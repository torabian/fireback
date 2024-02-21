package geo

import (
	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/workspaces"
)

/**
*	Override this function on LocationDataHttp.go,
*	In order to add your own http
**/
var AppendLocationDataRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostLocationData(c *gin.Context) {
	workspaces.HttpPostEntity(c, LocationDataActionCreate)
}

func HttpExportStreamLocationData(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, LocationDataActionExport)
}

func HttpQueryLocationDatas(c *gin.Context) {
	workspaces.HttpQueryEntity(c, LocationDataActionQuery)
}

func HttpGetOneLocationData(c *gin.Context) {
	workspaces.HttpGetEntity(c, LocationDataActionGetOne)
}

func HttpRemoveLocationData(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, LocationDataActionRemove)
}

func HttpUpdateLocationData(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, LocationDataActionUpdate)
}

func HttpBulkUpdateLocationData(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, LocationDataActionBulkUpdate)
}

func GetLocationDataModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/locationDatas",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_QUERY}),

				HttpQueryLocationDatas,
			},
			Format:         "QUERY",
			Action:         LocationDataActionQuery,
			ResponseEntity: &[]LocationDataEntity{},
		},
		{
			Method: "GET",
			Url:    "/locationDatas/export",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_QUERY}),

				HttpExportStreamLocationData,
			},
			Format:         "QUERY",
			Action:         LocationDataActionExport,
			ResponseEntity: &[]LocationDataEntity{},
		},

		{
			Method: "GET",
			Url:    "/locationData/:uniqueId",
			Handlers: []gin.HandlerFunc{

				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_QUERY}),

				HttpGetOneLocationData,
			},
			Format:         "GET_ONE",
			Action:         LocationDataActionGetOne,
			ResponseEntity: &LocationDataEntity{},
		},

		{
			Method: "POST",
			Url:    "/locationData",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_CREATE}),
				HttpPostLocationData,
			},
			Action:         LocationDataActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &LocationDataEntity{},
			ResponseEntity: &LocationDataEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/locationData",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_UPDATE}),
				HttpUpdateLocationData,
			},
			Action:         LocationDataActionUpdate,
			RequestEntity:  &LocationDataEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &LocationDataEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/locationDatas",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_UPDATE}),
				HttpBulkUpdateLocationData,
			},
			Action:         LocationDataActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[LocationDataEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[LocationDataEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/locationData",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_LOCATIONDATA_DELETE}),
				HttpRemoveLocationData,
			},
			Action:         LocationDataActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &LocationDataEntity{},
		},
	}

	// Append user defined functions
	AppendLocationDataRouter(&routes)

	return routes

}

func CreateLocationDataRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetLocationDataModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, LocationDataEntityJsonSchema, "locationData-http", "geo")
	workspaces.WriteEntitySchema("LocationData", LocationDataEntityJsonSchema, "geo")

	return httpRoutes
}
