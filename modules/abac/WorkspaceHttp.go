package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {

	AppendWorkspaceRouter = func(r *[]workspaces.Module3Action) {
		*r = append(*r,
			workspaces.Module3Action{
				Method:         "REACTIVE",
				ResponseEntity: &ReactiveSearchResultDto{},
				Out: &workspaces.Module3ActionBody{
					Dto: "ReactiveSearchResultDto",
				},
			},
			workspaces.Module3Action{
				Method:         "POST",
				ResponseEntity: &ImportRequestDto{},
				RequestEntity:  &ImportRequestDto{},
				Out: &workspaces.Module3ActionBody{
					Dto: "ImportRequestDto",
				},
				In: &workspaces.Module3ActionBody{
					Dto: "ImportRequestDto",
				},
			},
		)

	}
}
