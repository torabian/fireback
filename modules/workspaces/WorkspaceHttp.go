package workspaces

func init() {

	AppendWorkspaceRouter = func(r *[]Module2Action) {
		*r = append(*r,
			Module2Action{
				Method:         "REACTIVE",
				ResponseEntity: &ReactiveSearchResultDto{},
				Out: &Module2ActionBody{
					Dto: "ReactiveSearchResultDto",
				},
			},
			Module2Action{
				Method:         "POST",
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
