package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	QueryWorkspaceTypesPubliclyActionImp = func(q fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto, *fireback.QueryResultMeta, *fireback.IError) {
		res, qrm, err := WorkspaceTypeActionPublicQuery(q)
		return res, qrm, fireback.CastToIError(err)
	}
}
