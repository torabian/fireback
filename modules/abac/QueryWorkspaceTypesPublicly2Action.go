package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPublicly2Impl = func(c QueryWorkspaceTypesPublicly2ActionRequest, q fireback.QueryDSL) (*QueryWorkspaceTypesPublicly2ActionResponse, error) {

		res, qrm, err := WorkspaceTypeActionPublicQuery(q)
		if err != nil {
			return nil, err
		}

		return &QueryWorkspaceTypesPublicly2ActionResponse{
			Payload: fireback.GResponseQuery(res, qrm, &q),
		}, nil
	}
}
