package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// DisablePassportTotpAction force-disables TOTP (2FA) on a passport by clearing its
// secret and confirmation flag - root-only, for unblocking a user who lost their
// authenticator device. There is no matching "enable": turning TOTP on requires the
// user's own authenticator app to scan a fresh secret (see
// ConfirmClassicPassportTotpActionImplementation.go), which an administrator cannot do
// on their behalf.
func DisablePassportTotpAction(c abacdefs.DisablePassportTotpActionRequest) (*abacdefs.DisablePassportTotpActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_UPDATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}

	q := *query
	q.UniqueId = c.Body.UniqueId
	existing, err3 := PassportActions.GetOne(q)
	if err3 != nil {
		return nil, err3
	}

	wasEnabled := existing.TotpSecret != ""

	// A plain PassportActions.Update(&PassportEntity{TotpSecret: ""}) would not
	// actually clear the column: the generic UpdateEntity helper (CrudCoreActions.go)
	// goes through gorm's struct-based UpdateColumns, which skips any field left at
	// its Go zero value - "" is indistinguishable from "not provided" for a plain
	// (non-Nullable) string field. A raw column update with an explicit map sidesteps
	// that entirely, the standard gorm workaround for writing a real zero value.
	if err4 := fireback.GetDbRef().
		Model(&abacdefs.PassportEntity{}).
		Where("unique_id = ?", existing.UniqueId).
		Updates(map[string]interface{}{
			"totp_secret":    "",
			"totp_confirmed": false,
		}).Error; err4 != nil {
		return nil, fireback.GormErrorToIError(err4)
	}

	return &abacdefs.DisablePassportTotpActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.DisablePassportTotpActionRes{Disabled: wasEnabled}),
	}, nil
}
