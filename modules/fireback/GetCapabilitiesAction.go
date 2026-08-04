package fireback

func GetCapabilitiesAction(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error) {
	query, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})

	if err != nil {
		return nil, err
	}

	res, qrm, err2 := CapabilityActions.Query(*query)
	if err2 != nil {
		return nil, err
	}

	return &GetCapabilitiesActionResponse{
		Payload: GResponseQuery(res, qrm, query),
	}, nil
}
