package workspaces

func init() {

	AppendWorkspaceRouter = func(r *[]Module2Action) {
		*r = append(*r,
			Module2Action{
				Method:         "REACTIVE",
				Url:            "/reactiveSearch",
				Virtual:        true,
				ResponseEntity: &ReactiveSearchResultDto{},
			},
			Module2Action{
				Method:         "POST",
				Url:            "/backupImport",
				Virtual:        true,
				ResponseEntity: &ImportRequestDto{},
				RequestEntity:  &ImportRequestDto{},
			},
		)

	}
}
