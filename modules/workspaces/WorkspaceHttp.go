package workspaces

func init() {

	AppendWorkspaceRouter = func(r *[]Module2Action) {
		*r = append(*r,
			Module2Action{
				Method:         "REACTIVE",
				Url:            "/reactiveSearch",
				Virtual:        true,
				ResponseEntity: &ReactiveSearchResultDto{},
				Out: &Module2ActionBody{
					Dto: "ReactiveSearchResultDto",
				},
			},
			Module2Action{
				Method:         "POST",
				Url:            "/backupImport",
				Virtual:        true,
				ResponseEntity: &ImportRequestDto{},
				RequestEntity:  &ImportRequestDto{},
				Out: &Module2ActionBody{
					Dto: "ImportRequestDto",
				},
				In: &Module2ActionBody{
					Dto: "ImportRequestDto",
				},
			},
		)

	}
}
