package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/security"
)

// CreatePassportForUserAction lets a root administrator manually create a new email or
// phone passport for an existing user, with a password already set and the passport
// marked confirmed - for provisioning a working login when the user can't complete the
// normal signup/email-confirmation flow themselves (e.g. no mail server configured yet).
//
// Unlike the generic 'passport create' entity endpoint (PassportCreateAction), which can
// never set a password - PassportEntity.Password is tagged json:"-", so it's dropped on
// both encode and decode, by design, to keep raw passwords out of every ordinary
// entity-CRUD path - this hashes the given plaintext password the same way
// SetPassportPasswordAction does before persisting it.
func CreatePassportForUserAction(c abacdefs.CreatePassportForUserActionRequest) (*abacdefs.CreatePassportForUserActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_CREATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	if req.Type != PASSPORT_METHOD_EMAIL && req.Type != PASSPORT_METHOD_PHONE {
		return nil, &fireback.IError{
			HttpCode: 400,
			Errors: []*fireback.IErrorItem{
				{
					Location: "type",
					Message: &fireback.ErrorItem{
						"$":  "InvalidPassportType",
						"en": "type must be either 'email' or 'phone'.",
					},
				},
			},
		}
	}

	// Same minimum-length rule SetPassportPasswordAction/ChangePasswordAction
	// enforce - an admin creating a passport on someone's behalf shouldn't be able
	// to set a weaker password than the user could themselves.
	if len(req.Password) < 6 {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	if _, userErr := UserActions.GetOne(fireback.QueryDSL{UniqueId: req.UserId}); userErr != nil {
		return nil, &fireback.IError{
			HttpCode: 404,
			Errors: []*fireback.IErrorItem{
				{
					Location: "userId",
					Message: &fireback.ErrorItem{
						"$":  "UserNotFound",
						"en": "The given user does not exist.",
					},
				},
			},
		}
	}

	hashed, hashErr := security.HashPassword(req.Password)
	if hashErr != nil {
		return nil, fireback.CastToIError(hashErr)
	}

	created, createErr := PassportActions.Create(&abacdefs.PassportEntity{
		UniqueId:  fireback.UUID(),
		UserId:    emigo.NullableOf(req.UserId),
		Type:      req.Type,
		Value:     req.Value,
		Password:  hashed,
		Confirmed: emigo.NullableOf(true),
	}, *query)
	if createErr != nil {
		return nil, createErr
	}

	return &abacdefs.CreatePassportForUserActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.CreatePassportForUserActionRes{
			UniqueId: created.UniqueId,
			UserId:   req.UserId,
			Type:     created.Type,
			Value:    created.Value,
		}),
	}, nil
}
