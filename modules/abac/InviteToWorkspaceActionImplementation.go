package abac

import (
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

func InviteToWorkspaceAction(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error) {
	queryPtr, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE},
	})
	if err != nil {
		return nil, err
	}
	query := *queryPtr

	req := c.Body

	// Only validate FirstName/LastName/RoleId here, not the whole request struct:
	// req.Workspace is an emigo.One[WorkspaceEntity] that nothing ever populates (the
	// invite's workspace always comes from the resolved query context/Workspace-id
	// header, set below) - but go-playground's validator recurses into nested structs by
	// default, so validating &req directly also enforced WorkspaceEntity's OWN
	// validate:"required" tags ("name"/"typeId") against the always-zero-valued entity
	// sitting in .Item. That made this the actual root cause of "invitation sending is
	// broken": literally every call to this action, no matter what was posted, failed
	// with a bogus "workspace.item.name/typeId is required" 403 before ever reaching any
	// of the validation/logic below.
	if err := fireback.CommonStructValidatorPointer(&struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		RoleId    string `validate:"required"`
	}{FirstName: req.FirstName, LastName: req.LastName, RoleId: req.RoleId}, false); err != nil {
		return nil, err
	}

	_, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(req.RoleId))
	if len(roleErrors) != 0 {
		return nil, &fireback.IError{
			Errors: roleErrors,
		}
	}

	// The invite has to actually reach the invitee somehow - without either a contact
	// method, the row would still get created below, but both the SendInviteEmail and
	// GsmSendSMSUsingNotificationConfig calls further down are gated on
	// invite.Email/invite.Phonenumber being non-empty, so nobody would ever be notified.
	if strings.TrimSpace(req.Email) == "" && strings.TrimSpace(req.Phonenumber) == "" {
		return nil, &fireback.IError{
			Errors: []*fireback.IErrorItem{
				{
					Location: "email",
					Message: &fireback.ErrorItem{
						"$":  "InviteRequiresEmailOrPhone",
						"en": "Provide at least an email address or a phone number to send the invitation to.",
					},
				},
			},
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

	// The rest of the submitted invitation - previously dropped entirely, which meant
	// every invite was saved (and, below, "sent") with a blank name/email/phone/role.
	invite.PublicKey = req.PublicKey
	invite.CoverLetter = req.CoverLetter
	invite.Email = req.Email
	invite.Phonenumber = req.Phonenumber
	invite.FirstName = req.FirstName
	invite.LastName = req.LastName
	invite.ForceEmailAddress = req.ForceEmailAddress
	invite.ForcePhoneNumber = req.ForcePhoneNumber
	invite.RoleId = emigo.NullableOf(req.RoleId)

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

		// GsmSendSMSUsingNotificationConfig already returns a *fireback.IError - wrapping
		// it again through GormErrorToIError (which expects a plain/gorm error) discarded
		// its real message and status in favor of a generic one.
		if _, err7 := GsmSendSMSUsingNotificationConfig(inviteBody, []string{invite.Phonenumber}); err7 != nil {
			return nil, err7
		}
	}

	return &InviteToWorkspaceActionResponse{
		Payload: fireback.GResponseSingleItem(invite),
	}, nil
}
