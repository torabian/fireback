package workspaces

func init() {

	AppendWorkspaceRouter = func(r *[]Module3Action) {
		*r = append(*r,
			Module3Action{
				Method:         "REACTIVE",
				ResponseEntity: &ReactiveSearchResultDto{},
				Out: &Module3ActionBody{
					Dto: "ReactiveSearchResultDto",
				},
			},
			Module3Action{
				Method:         "POST",
				ResponseEntity: &ImportRequestDto{},
				RequestEntity:  &ImportRequestDto{},
				Out: &Module3ActionBody{
					Dto: "ImportRequestDto",
				},
				In: &Module3ActionBody{
					Dto: "ImportRequestDto",
				},
			},
		)

	}
}
