package fireback

func init() {
	// Override the implementation with our actual code.
	ListCapabilitiesActionImp = ListCapabilitiesAction
}

func ListCapabilitiesAction(q QueryDSL) ([]string, *IError) {
	stream, _, err := CapabilityEntityStream(QueryDSL{ItemsPerPage: 20, WorkspaceId: ROOT_VAR})
	if err != nil {
		return nil, CastToIError(err)
	}
	keys := []string{}
	for items := range stream {
		for _, item := range items {
			keys = append(keys, item.UniqueId)
		}
	}
	return keys, nil
}
