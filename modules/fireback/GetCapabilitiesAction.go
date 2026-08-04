package fireback

func GetCapabilitiesAction(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error) {
	query, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})

	if err != nil {
		return nil, err
	}

	res, qrm, err2 := CapabilityEntityActions.Browse(GetDbRef(), CapabilityBrowseActionQuery{}, "")
	if err2 != nil {
		return nil, err
	}

	return &GetCapabilitiesActionResponse{
		Payload: GResponseQuery(res, &QueryResultMeta{
			TotalItems: qrm.TotalItems,
			Cursor:     qrm.Cursor,
		}, query),
	}, nil
}
