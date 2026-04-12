package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPubliclyImpl = func(c QueryWorkspaceTypesPubliclyActionRequest, q fireback.QueryDSL) (*QueryWorkspaceTypesPubliclyActionResponse, error) {

		res, qrm, err := WorkspaceTypeActionPublicQuery(q)
		if err != nil {
			return nil, err
		}

		return &QueryWorkspaceTypesPubliclyActionResponse{
			Payload: fireback.GResponseQuery(res, qrm, &q),
		}, nil
	}
}
