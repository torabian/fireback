package fireback

func GetCapabilitiesAction(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error) {
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

	return &CapabilityBrowseActionResponse{
		Payload: GResponseQuery(res, &QueryResultMeta{
			TotalItems: qrm.TotalItems,
			Cursor:     qrm.Cursor,
		}, query),
	}, nil
}

func CapabilityGetAction(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error) {
	_, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})

	if err != nil {
		return nil, err
	}

	res, err2 := CapabilityEntityActions.Get(GetDbRef(), c.Params.UniqueId)
	if err2 != nil {
		return nil, err
	}

	return &CapabilityGetActionResponse{
		Payload: GResponseSingleItem(res),
	}, nil

}

func CapabilityUpdateAction(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error) {
	_, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_UPDATE},
	})

	if err != nil {
		return nil, err
	}

	res, err2 := CapabilityEntityActions.Update(GetDbRef(), c.Params.UniqueId, c.Body)
	if err2 != nil {
		return nil, err2
	}

	return &CapabilityUpdateActionResponse{
		Payload: GResponseSingleItem(res),
	}, nil
}

func CapabilityAwareDeleteAction(c CapabilityAwareDeleteActionRequest) (*CapabilityAwareDeleteActionResponse, error) {
	_, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_DELETE},
	})

	if err != nil {
		return nil, err
	}

	if err2 := CapabilityEntityActions.AwareDelete(GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, err2
	}

	return &CapabilityAwareDeleteActionResponse{
		Payload: GResponseSingleItem(struct{}{}),
	}, nil
}
