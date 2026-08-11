package abac

import (
	"github.com/torabian/emi/emigo"
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// Menu is abac's own built-in sidebar/menu definitions - the "Root" admin section, the
// cloud/self-service "Workspace" section, and the "Personal" section - moved here (as
// plain Go data, hand-converted from the old fireback-manage-menu.yml/
// fireback-menu-cloud.yml/fireback-personal-menu.yml) from
// modules/abac/interfacetools/seeders/AppMenu, which only ever got imported on request
// via the generic "seeders" CLI command (fireback.SeederFromFSImport, a bare insert with
// no de-duplication - re-running it duplicated every row). abac deliberately only
// imports interfacetoolsdefs (plain generated types) here, not interfacetools itself
// (the module with the actual DB/migration logic) - main.go is what ties the two
// together, passing this slice into interfacetools.ModuleSetup's
// InterfaceToolsModuleConfig.ExtraAppMenus, which syncs it (interfacetools.SyncAppMenus)
// as part of interfacetools' own migration, every time migrations run, instead of
// requiring a separate manual step.
var Menu = []*interfacetoolsdefs.AppMenuEntity{
	// --- fireback-manage-menu.yml: the "Root" admin section ---
	{
		UniqueId:    "root-actions",
		Label:       complexes.TStringFrom(map[string]string{"en": "Root"}),
		WorkspaceId: emigo.NullableOf(fireback.USER_SYSTEM),
	},
	{
		UniqueId:     "root-capabilities",
		Label:        complexes.TStringFrom(map[string]string{"en": "Capabilities & Permissions", "fa": "سطوح دسترسی"}),
		Href:         "/manage/capabilities",
		Icon:         "ios-theme/icons/settings.svg",
		ParentId:     emigo.NullableOf("root-actions"),
		WorkspaceId:  emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId: emigo.NullableOf("root.manage.fireback.capability.query"),
	},
	{
		UniqueId:      "workspaces",
		Label:         complexes.TStringFrom(map[string]string{"en": "Workspaces", "fa": "تیم ها"}),
		Href:          "/manage/workspaces",
		Icon:          "common/workspace.svg",
		ActiveMatcher: "/workspaces/|workspace/new",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.workspace.query"),
	},
	{
		UniqueId:      "regional_content",
		Label:         complexes.TStringFrom(map[string]string{"en": "Regional Content", "fa": "متن ها و محتویات"}),
		Href:          "/manage/regional-contents",
		Icon:          "common/workspace.svg",
		ActiveMatcher: "/manage/regional-content",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.regional-content.query"),
	},
	{
		UniqueId:      "workspace_config",
		Label:         complexes.TStringFrom(map[string]string{"en": "Workspace Config", "fa": "تنظیمات تیم"}),
		Href:          "/manage/workspace-config",
		Icon:          "ios-theme/icons/settings.svg",
		ActiveMatcher: "/workspace-config",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.workspace-config.query"),
	},
	{
		UniqueId:      "email_sender",
		Label:         complexes.TStringFrom(map[string]string{"en": "Email Sender", "fa": "ارسال ایمیل"}),
		Href:          "/manage/email-senders",
		Icon:          "common/mail.svg",
		ActiveMatcher: "email-sender",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.email-sender.query"),
	},
	{
		UniqueId:      "workspace_types",
		Label:         complexes.TStringFrom(map[string]string{"en": "Workspace Types", "fa": "نوع تیم ها"}),
		Href:          "/manage/workspace-types",
		Icon:          "ios-theme/icons/settings.svg",
		ActiveMatcher: "workspace-type",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.workspace-type.query"),
	},
	{
		UniqueId:      "email_provider",
		Label:         complexes.TStringFrom(map[string]string{"en": "Email Provider", "fa": "ایمیل سرویس"}),
		Href:          "/manage/email-providers",
		Icon:          "common/emailprovider.svg",
		ActiveMatcher: "email-provider",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.email-provider.query"),
	},
	{
		UniqueId:      "gsm_provider",
		Label:         complexes.TStringFrom(map[string]string{"en": "gsm Provider", "fa": "سرویس پیامک"}),
		Href:          "/manage/gsm-providers",
		Icon:          "common/emailprovider.svg",
		ActiveMatcher: "gsm-provider",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.gsm-provider.query"),
	},
	{
		UniqueId:      "users",
		Label:         complexes.TStringFrom(map[string]string{"en": "Users", "fa": "کاربران"}),
		Href:          "/manage/users",
		Icon:          "common/user.svg",
		ActiveMatcher: "/user/",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.user.query"),
	},
	{
		UniqueId:      "passport-methods",
		Label:         complexes.TStringFrom(map[string]string{"en": "Passport Methods", "fa": "انواع پاسپورت"}),
		Href:          "/manage/passport-methods",
		Icon:          "common/user.svg",
		ActiveMatcher: "/passport-method/",
		ParentId:      emigo.NullableOf("root-actions"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.manage.abac.passport-method.query"),
	},

	// --- fireback-menu-cloud.yml: the "Workspace" self-service section ---
	{
		UniqueId:    "fireback",
		Label:       complexes.TStringFrom(map[string]string{"en": "Workspace", "fa": "ورک اسپیس"}),
		WorkspaceId: emigo.NullableOf(fireback.USER_SYSTEM),
	},
	{
		UniqueId:      "publicjoinkey",
		Label:         complexes.TStringFrom(map[string]string{"en": "Public join keys", "fa": "کلید عضویت"}),
		Href:          "/selfservice/public-join-keys",
		Icon:          "common/joinkey.svg",
		ActiveMatcher: "publicjoinkey",
		ParentId:      emigo.NullableOf("fireback"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.modules.abac.public-join-key.query"),
	},
	{
		UniqueId:      "roles",
		Label:         complexes.TStringFrom(map[string]string{"en": "Roles", "fa": "نقش ها"}),
		Href:          "/selfservice/roles",
		Icon:          "common/role.svg",
		ActiveMatcher: "/role/",
		ParentId:      emigo.NullableOf("fireback"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.modules.abac.role.query"),
	},
	{
		UniqueId:      "invites",
		Label:         complexes.TStringFrom(map[string]string{"en": "Invites", "fa": "دعوت نامه ها"}),
		Href:          "/selfservice/workspace-invites",
		Icon:          "common/workspaceinvite.svg",
		ActiveMatcher: "/workspace/invite(s)?",
		ParentId:      emigo.NullableOf("fireback"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
		CapabilityId:  emigo.NullableOf("root.modules.abac.workspace-invite.query"),
	},

	// --- fireback-personal-menu.yml: the "Personal" section ---
	{
		UniqueId:    "personal",
		Label:       complexes.TStringFrom(map[string]string{"en": "Personal", "fa": "شخصی سازی"}),
		WorkspaceId: emigo.NullableOf(fireback.USER_SYSTEM),
	},
	{
		UniqueId:      "settings",
		Label:         complexes.TStringFrom(map[string]string{"en": "Settings", "fa": "تنظیمات"}),
		Href:          "/settings",
		Icon:          "ios-theme/icons/settings.svg",
		ActiveMatcher: "/settings",
		ParentId:      emigo.NullableOf("personal"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
	},
	{
		UniqueId:      "selfservice",
		Label:         complexes.TStringFrom(map[string]string{"en": "Account & Profile", "fa": "حساب و پروفایل"}),
		Href:          "/selfservice",
		Icon:          "ios-theme/icons/settings.svg",
		ActiveMatcher: "/selfservice$",
		ParentId:      emigo.NullableOf("personal"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
	},
	{
		UniqueId:      "my_invitation",
		Label:         complexes.TStringFrom(map[string]string{"en": "My Invitations", "fa": "دعوت نامه های من"}),
		Href:          "/selfservice/user-invitations",
		Icon:          "common/workspaceinvite.svg",
		ActiveMatcher: "/invites/",
		ParentId:      emigo.NullableOf("personal"),
		WorkspaceId:   emigo.NullableOf(fireback.USER_SYSTEM),
	},
}
