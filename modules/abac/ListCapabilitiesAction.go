package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	ListCapabilitiesActionImp = ListCapabilitiesAction
}

func ListCapabilitiesAction(q workspaces.QueryDSL) ([]string, *workspaces.IError) {
	stream, _, err := CapabilityEntityStream(workspaces.QueryDSL{ItemsPerPage: 20, WorkspaceId: ROOT_VAR})
	if err != nil {
		return nil, workspaces.CastToIError(err)
	}
	keys := []string{}
	for items := range stream {
		for _, item := range items {
			keys = append(keys, item.UniqueId)
		}
	}
	return keys, nil
}
