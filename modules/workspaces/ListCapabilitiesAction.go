package workspaces

func init() {
	// Override the implementation with our actual code.
	ListCapabilitiesActionImp = ListCapabilitiesAction
}

func ListCapabilitiesAction(q QueryDSL) ([]string, *IError) {
	q.ItemsPerPage = 99999
	items, _, err := CapabilityActionQuery(q)
	if err != nil {
		return nil, CastToIError(err)
	}
	keys := []string{}
	for _, item := range items {
		keys = append(keys, item.UniqueId)
	}
	return keys, nil
}
