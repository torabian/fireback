package worldtimezone

import (
	"github.com/gin-gonic/gin"

	"pixelplux.com/fireback/modules/workspaces"
)

/**
*	Override this function on TimezoneGroupHttp.go,
*	In order to add your own http
**/
var AppendTimezoneGroupRouter = func(r *[]workspaces.Module2Action) {}

func HttpPostTimezoneGroup(c *gin.Context) {
	workspaces.HttpPostEntity(c, TimezoneGroupActionCreate)
}

func HttpExportStreamTimezoneGroup(c *gin.Context) {
	workspaces.HttpStreamFileChannel(c, TimezoneGroupActionExport)
}

func HttpQueryTimezoneGroups(c *gin.Context) {
	workspaces.HttpQueryEntity(c, TimezoneGroupActionQuery)
}

func HttpGetOneTimezoneGroup(c *gin.Context) {
	workspaces.HttpGetEntity(c, TimezoneGroupActionGetOne)
}

func HttpRemoveTimezoneGroup(c *gin.Context) {
	workspaces.HttpRemoveEntity(c, TimezoneGroupActionRemove)
}

func HttpCreateTimezoneGroupUtcItems(
	c *gin.Context,
) {
	workspaces.HttpPostEntity(c, TimezoneGroupUtcItemsActionCreate)
}

func HttpUpdateTimezoneGroupUtcItems(
	c *gin.Context,
) {
	workspaces.HttpUpdateEntity(c, TimezoneGroupUtcItemsActionUpdate)
}

func HttpGetTimezoneGroupUtcItems(
	c *gin.Context,
) {
	workspaces.HttpGetEntity(c, TimezoneGroupUtcItemsActionGetOne)
}

func HttpUpdateTimezoneGroup(c *gin.Context) {
	workspaces.HttpUpdateEntity(c, TimezoneGroupActionUpdate)
}

func HttpBulkUpdateTimezoneGroup(c *gin.Context) {
	workspaces.HttpUpdateEntities(c, TimezoneGroupActionBulkUpdate)
}

func GetTimezoneGroupModule2Actions() []workspaces.Module2Action {

	routes := []workspaces.Module2Action{

		{
			Method: "GET",
			Url:    "/timezoneGroups",
			Handlers: []gin.HandlerFunc{

				/* Intentionlly allow to query all */

				HttpQueryTimezoneGroups,
			},
			Format:         "QUERY",
			Action:         TimezoneGroupActionQuery,
			ResponseEntity: &[]TimezoneGroupEntity{},
		},
		{
			Method: "GET",
			Url:    "/timezoneGroups/export",
			Handlers: []gin.HandlerFunc{

				/* Intentionlly allow to query all */

				HttpExportStreamTimezoneGroup,
			},
			Format:         "QUERY",
			Action:         TimezoneGroupActionExport,
			ResponseEntity: &[]TimezoneGroupEntity{},
		},

		{
			Method: "GET",
			Url:    "/timezoneGroup/:uniqueId",
			Handlers: []gin.HandlerFunc{

				/* Intentionlly allowed to query specific ones */

				HttpGetOneTimezoneGroup,
			},
			Format:         "GET_ONE",
			Action:         TimezoneGroupActionGetOne,
			ResponseEntity: &TimezoneGroupEntity{},
		},

		{
			Method: "POST",
			Url:    "/timezoneGroup",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_CREATE}),
				HttpPostTimezoneGroup,
			},
			Action:         TimezoneGroupActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &TimezoneGroupEntity{},
			ResponseEntity: &TimezoneGroupEntity{},
		},
		{
			Method: "PATCH",
			Url:    "/timezoneGroup",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_UPDATE}),
				HttpUpdateTimezoneGroup,
			},
			Action:         TimezoneGroupActionUpdate,
			RequestEntity:  &TimezoneGroupEntity{},
			Format:         "PATCH_ONE",
			ResponseEntity: &TimezoneGroupEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/timezoneGroup/:linkerId/utcItems/:uniqueId",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_UPDATE}),
				HttpUpdateTimezoneGroupUtcItems,
			},
			Action:         TimezoneGroupUtcItemsActionUpdate,
			Format:         "PATCH_ONE",
			RequestEntity:  &TimezoneGroupUtcItemsEntity{},
			ResponseEntity: &TimezoneGroupUtcItemsEntity{},
		},

		{
			Method: "GET",
			Url:    "/timezoneGroup/utcItems/:linkerId/:uniqueId",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{}),
				HttpGetTimezoneGroupUtcItems,
			},
			Action:         TimezoneGroupUtcItemsActionGetOne,
			Format:         "GET_ONE",
			ResponseEntity: &TimezoneGroupUtcItemsEntity{},
		},
		{
			Method: "POST",
			Url:    "/timezoneGroup/:linkerId/utcItems",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_CREATE}),
				HttpCreateTimezoneGroupUtcItems,
			},
			Action:         TimezoneGroupUtcItemsActionCreate,
			Format:         "POST_ONE",
			RequestEntity:  &TimezoneGroupUtcItemsEntity{},
			ResponseEntity: &TimezoneGroupUtcItemsEntity{},
		},

		{
			Method: "PATCH",
			Url:    "/timezoneGroups",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_UPDATE}),
				HttpBulkUpdateTimezoneGroup,
			},
			Action:         TimezoneGroupActionBulkUpdate,
			Format:         "PATCH_BULK",
			RequestEntity:  &workspaces.BulkRecordRequest[TimezoneGroupEntity]{},
			ResponseEntity: &workspaces.BulkRecordRequest[TimezoneGroupEntity]{},
		},
		{
			Method: "DELETE",
			Url:    "/timezoneGroup",
			Format: "DELETE_DSL",
			Handlers: []gin.HandlerFunc{
				workspaces.WithAuthorization([]string{PERM_ROOT_TIMEZONEGROUP_DELETE}),
				HttpRemoveTimezoneGroup,
			},
			Action:         TimezoneGroupActionRemove,
			RequestEntity:  &workspaces.DeleteRequest{},
			ResponseEntity: &workspaces.DeleteResponse{},
			TargetEntity:   &TimezoneGroupEntity{},
		},
	}

	// Append user defined functions
	AppendTimezoneGroupRouter(&routes)

	return routes

}

func CreateTimezoneGroupRouter(r *gin.Engine) []workspaces.Module2Action {

	httpRoutes := GetTimezoneGroupModule2Actions()

	workspaces.CastRoutes(httpRoutes, r)
	workspaces.WriteHttpInformationToFile(&httpRoutes, TimezoneGroupEntityJsonSchema, "timezoneGroup-http", "worldtimezone")
	workspaces.WriteEntitySchema("TimezoneGroup", TimezoneGroupEntityJsonSchema, "worldtimezone")

	return httpRoutes
}
