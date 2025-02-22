package workspaces

func init() {
	// Override the implementation with our actual code.
	InviteToWorkspaceActionImp = InviteToWorkspaceAction
}

func InviteToWorkspaceAction(req *WorkspaceInviteEntity, q QueryDSL) (*WorkspaceInviteEntity, *IError) {
	if err := WorkspaceInviteValidator(req, false); err != nil {
		return nil, err
	}

	var invite WorkspaceInviteEntity = WorkspaceInviteEntity{
		Value:       req.Value,
		WorkspaceId: &q.WorkspaceId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		RoleId:      req.RoleId,
		UniqueId:    UUID(),
	}

	if err := GetRef(q).Create(&invite).Error; err != nil {
		return &invite, GormErrorToIError(err)
	}

	// @todo: Detect the type of passport, and

	_, method := validatePassportType(*req.Value)

	if method == PASSPORT_METHOD_EMAIL {
		if err7 := SendInviteEmail(q, &invite); err7 != nil {
			return nil, err7
		}
	}
	if method == PASSPORT_METHOD_PHONE {
		inviteBody := "You are invite " + *invite.FirstName + " " + *invite.LastName
		if _, err7 := GsmSendSmsAction(&GsmSendSmsActionReqDto{ToNumber: req.Value, Body: &inviteBody}, q); err7 != nil {
			return nil, GormErrorToIError(err7)
		}
	}

	return &invite, nil
}
