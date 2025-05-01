package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {

	AppendWorkspaceRouter = func(r *[]fireback.Module3Action) {
		*r = append(*r,
			fireback.Module3Action{
				Method:         "REACTIVE",
				ResponseEntity: &ReactiveSearchResultDto{},
				Out: &fireback.Module3ActionBody{
					Dto: "ReactiveSearchResultDto",
				},
			},
			fireback.Module3Action{
				Method:         "POST",
				ResponseEntity: &ImportRequestDto{},
				RequestEntity:  &ImportRequestDto{},
				Out: &fireback.Module3ActionBody{
					Dto: "ImportRequestDto",
				},
				In: &fireback.Module3ActionBody{
					Dto: "ImportRequestDto",
				},
			},
		)

	}
}
