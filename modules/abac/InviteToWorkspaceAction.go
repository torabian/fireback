package abac

import (
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	InviteToWorkspaceActionImp = InviteToWorkspaceAction
}

func InviteToWorkspaceAction(req *WorkspaceInviteEntity, q fireback.QueryDSL) (*WorkspaceInviteEntity, *fireback.IError) {
	if err := WorkspaceInviteValidator(req, false); err != nil {
		return nil, err
	}

	_, roleErrors := ValidateRoleAndItsExsitence(req.RoleId)
	if len(roleErrors) != 0 {
		return nil, &fireback.IError{
			Errors: roleErrors,
		}
	}

	userLocale := q.Language
	if strings.TrimSpace(req.TargetUserLocale) != "" {
		userLocale = req.TargetUserLocale
	}

	var invite WorkspaceInviteEntity = *req
	invite.WorkspaceId = fireback.NewString(q.WorkspaceId)
	invite.UniqueId = fireback.UUID()
	invite.TargetUserLocale = userLocale

	if err := fireback.GetRef(q).Create(&invite).Error; err != nil {
		return &invite, fireback.GormErrorToIError(err)
	}

	if invite.Email != "" {
		if err7 := SendInviteEmail(q, &invite); err7 != nil {
			return nil, err7
		}
	}

	if invite.Phonenumber != "" {
		inviteBody := "You are invite " + invite.FirstName + " " + invite.LastName
		if _, err7 := GsmSendSmsAction(&GsmSendSmsActionReqDto{
			ToNumber: invite.Phonenumber,
			Body:     inviteBody,
		}, q); err7 != nil {
			return nil, fireback.GormErrorToIError(err7)
		}
	}

	return &invite, nil
}
