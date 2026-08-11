package abac

import (
	"errors"
	"log"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

// regionalContentDefault is one entry RegionalContentDefaults ensures exists - a plain
// Go literal (rather than the old modules/abac/seeders/RegionalContent/*.yml, which was
// never actually wired to anything - ProvideSeederImportHandler was never called for
// this module, so that file only ever sat there unused) so it's regenerated the same
// way abac.Menu (Menu.go) supplies AppMenu defaults: no separate opt-in "seeders"
// command to remember to run.
type regionalContentDefault struct {
	UniqueId   string
	Region     string
	LanguageId string
	KeyGroup   string
	Title      string
	Content    string
}

// RegionalContentDefaults is every default en/fa/pl template this module ships out of
// the box, one row per {keyGroup, language} pair, all region "any" (a specific
// deployment can add a region-scoped override later - see Abac.emi.yml's own comment on
// why region is part of regional_content_index's uniqueness). SyncRegionalContentDefaults
// inserts whichever of these don't already exist by UniqueId.
//
// EMAIL_OTP/SMS_OTP's en/fa content matches the original modules/abac/seeders/
// RegionalContent/regional-content-seeder.yml verbatim (now removed) - pl is new.
// INVITE_TO_WORKSPACE is a new keyGroup (see Abac.emi.yml) - note it isn't actually
// consumed by SendInviteEmail (WorkspaceActionUpdate.go) yet, which still reads its
// content from NotificationConfigEntity.InviteToWorkspaceContent (a separate, older,
// non-templated mechanism) rather than resolving a RegionalContent row the way
// QuickGetOtpMessage does for OTP - these rows exist so WorkspaceConfig's own
// "Invite To Workspace Content" picker (a RegionalContent select, see
// WorkspaceConfigEditForm.tsx) has real, sensible options, ahead of that wiring.
var RegionalContentDefaults = []regionalContentDefault{
	{
		UniqueId:   "EN_EMAIL_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "en",
		KeyGroup:   string(EMAIL_OTP),
		Content:    "Your Otp code is {{ .Otp }}",
	},
	{
		UniqueId:   "FA_EMAIL_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "fa",
		KeyGroup:   string(EMAIL_OTP),
		Content:    "کد فعال سازی شما {{ .Otp }} هست",
	},
	{
		UniqueId:   "PL_EMAIL_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "pl",
		KeyGroup:   string(EMAIL_OTP),
		Content:    "Twój kod jednorazowy to {{ .Otp }}",
	},
	{
		UniqueId:   "EN_SMS_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "en",
		KeyGroup:   string(SMS_OTP),
		Content:    "Your Otp code is {{ .Otp }}",
	},
	{
		UniqueId:   "FA_SMS_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "fa",
		KeyGroup:   string(SMS_OTP),
		Content:    "کد فعال سازی شما {{ .Otp }} هست",
	},
	{
		UniqueId:   "PL_SMS_OTP_DEFAULT",
		Region:     "any",
		LanguageId: "pl",
		KeyGroup:   string(SMS_OTP),
		Content:    "Twój kod jednorazowy to {{ .Otp }}",
	},
	{
		UniqueId:   "EN_INVITE_TO_WORKSPACE_DEFAULT",
		Region:     "any",
		LanguageId: "en",
		KeyGroup:   "INVITE_TO_WORKSPACE",
		Title:      "You're invited to join {{.WorkspaceName}}",
		Content:    "You have been invited to join {{.WorkspaceName}} as {{.RoleName}}. Click the link below to accept:\n{{.InviteUrl}}",
	},
	{
		UniqueId:   "FA_INVITE_TO_WORKSPACE_DEFAULT",
		Region:     "any",
		LanguageId: "fa",
		KeyGroup:   "INVITE_TO_WORKSPACE",
		Title:      "شما به {{.WorkspaceName}} دعوت شده‌اید",
		Content:    "شما به عضویت در {{.WorkspaceName}} با نقش {{.RoleName}} دعوت شده‌اید. برای پذیرفتن دعوت روی لینک زیر کلیک کنید:\n{{.InviteUrl}}",
	},
	{
		UniqueId:   "PL_INVITE_TO_WORKSPACE_DEFAULT",
		Region:     "any",
		LanguageId: "pl",
		KeyGroup:   "INVITE_TO_WORKSPACE",
		Title:      "Zostałeś zaproszony do {{.WorkspaceName}}",
		Content:    "Zostałeś zaproszony do {{.WorkspaceName}} jako {{.RoleName}}. Kliknij poniższy link, aby zaakceptować zaproszenie:\n{{.InviteUrl}}",
	},
}

// SyncRegionalContentDefaults inserts each entry of RegionalContentDefaults whose
// UniqueId doesn't already exist. Existing rows are left completely untouched - even
// if their content/region/etc no longer matches this list - since an admin may have
// already customized them; this only ever needs to provide a sane starting point,
// never to re-assert itself over a deployment's own edits (unlike
// interfacetools.SyncAppMenu, which deliberately does delete-then-recreate for
// AppMenu - app menu items aren't meant to be hand-edited the way a template's wording
// would be, so the two intentionally use different sync strategies).
func SyncRegionalContentDefaults(db *gorm.DB) {
	for _, item := range RegionalContentDefaults {
		var existing abacdefs.RegionalContentEntity
		err := db.Where(&abacdefs.RegionalContentEntity{UniqueId: item.UniqueId}).First(&existing).Error
		if err == nil {
			continue // already there - leave it alone, might have been modified since.
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("regionalContent defaults: checking", item.UniqueId, "failed:", err)
			continue
		}

		entity := &abacdefs.RegionalContentEntity{
			UniqueId:    item.UniqueId,
			Region:      item.Region,
			LanguageId:  item.LanguageId,
			KeyGroup:    item.KeyGroup,
			Title:       item.Title,
			Content:     item.Content,
			WorkspaceId: emigo.NullableOf(fireback.USER_SYSTEM),
		}
		if err := db.Create(entity).Error; err != nil {
			log.Println("regionalContent defaults: creating", item.UniqueId, "failed:", err)
		}
	}
}
