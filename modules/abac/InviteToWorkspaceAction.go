package abac

import (
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func InviteToWorkspaceAction(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error) {
	queryPtr, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}
	query := *queryPtr

	req := c.Body

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	_, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(req.RoleId))
	if len(roleErrors) != 0 {
		return nil, &fireback.IError{
			Errors: roleErrors,
		}
	}

	userLocale := query.Language
	if strings.TrimSpace(req.TargetUserLocale) != "" {
		userLocale = req.TargetUserLocale
	}

	invite := WorkspaceInviteEntity{}

	invite.WorkspaceId = emigo.NullableOf(query.WorkspaceId)
	invite.UniqueId = fireback.UUID()
	invite.TargetUserLocale = userLocale

	if err := fireback.GetRef(query).Create(&invite).Error; err != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	if invite.Email != "" {
		if err7 := SendInviteEmail(query, &invite); err7 != nil {
			return nil, err7
		}
	}

	if invite.Phonenumber != "" {
		inviteBody := "You are invite " + invite.FirstName + " " + invite.LastName

		if _, err7 := GsmSendSMSUsingNotificationConfig(inviteBody, []string{invite.Phonenumber}); err7 != nil {
			return nil, fireback.GormErrorToIError(err7)
		}
	}

	return &InviteToWorkspaceActionResponse{
		Payload: fireback.GResponseSingleItem(invite),
	}, nil
}
