package workspaces

func init() {

	AppendBackupTableMetaRouter = func(r *[]Module2Action) {

		// *r = append(*r, Module2Action{
		// 	Method: "GET",
		// 	Url:    "/backup/export",
		// 	Handlers: []gin.HandlerFunc{

		// 		WithAuthorization([]string{}),

		// 		HttpExportStreamControlSheet,
		// 	},
		// 	Format:         "QUERY",
		// 	Action:         ControlSheetActionExport,
		// 	ResponseEntity: &[]ControlSheetEntity{},
		// })

	}
}
