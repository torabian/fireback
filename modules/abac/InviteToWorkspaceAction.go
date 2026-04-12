package abac

import (
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	InviteToWorkspaceImpl = InviteToWorkspaceAction
}

func InviteToWorkspaceAction(c InviteToWorkspaceActionRequest, query fireback.QueryDSL) (*InviteToWorkspaceActionResponse, error) {
	req := c.Body

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	_, roleErrors := ValidateRoleAndItsExistence(fireback.NewString(req.RoleId))
	if len(roleErrors) != 0 {
		return nil, &fireback.IError{
			Errors: roleErrors,
		}
	}

	userLocale := query.Language
	if strings.TrimSpace(req.TargetUserLocale) != "" {
		userLocale = req.TargetUserLocale
	}

	forceEmailAddress := false
	if value, isSet := req.ForceEmailAddress.Get(); isSet && *value {
		forceEmailAddress = true
	}

	invite := WorkspaceInviteEntity{
		PublicKey:         req.PublicKey,
		CoverLetter:       req.CoverLetter,
		TargetUserLocale:  req.TargetUserLocale,
		Email:             req.Email,
		Phonenumber:       req.Phonenumber,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		ForceEmailAddress: fireback.NewBool(forceEmailAddress),
	}

	invite.WorkspaceId = fireback.NewString(query.WorkspaceId)
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
		if _, err7 := GsmSendSmsAction(&GsmSendSmsActionReqDto{
			ToNumber: invite.Phonenumber,
			Body:     inviteBody,
		}, query); err7 != nil {
			return nil, fireback.GormErrorToIError(err7)
		}
	}

	return &InviteToWorkspaceActionResponse{Payload: invite}, nil
}
