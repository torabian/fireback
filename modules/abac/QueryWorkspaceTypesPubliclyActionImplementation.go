package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func QueryWorkspaceTypesPubliclyAction(c abacdefs.QueryWorkspaceTypesPubliclyActionRequest) (*abacdefs.QueryWorkspaceTypesPubliclyActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}
	q := *query

	res, qrm, err2 := WorkspaceTypeActionPublicQuery(q)
	if err2 != nil {
		return nil, err2
	}

	return &abacdefs.QueryWorkspaceTypesPubliclyActionResponse{
		Payload: fireback.GResponseQuery(res, qrm, &q),
	}, nil
}
