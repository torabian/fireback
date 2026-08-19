//go:build !wasm

package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/messaging"
	"github.com/torabian/fireback/modules/abac/migrations"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// Inject this into any project as a complete solution
//
// Deliberately does NOT include interfacetools.ModuleSetup() (unlike before) - abac only
// imports interfacetoolsdefs (see Menu.go), not the interfacetools package itself, so it
// has no ModuleSetup to call here. main.go is what actually constructs that module now,
// via interfacetools.ModuleSetup(&interfacetools.InterfaceToolsModuleConfig{ExtraAppMenus:
// abac.Menu}), appended alongside whatever AbacCompleteModules returns.
func AbacCompleteModules() []*application.ModuleProvider {
	return []*application.ModuleProvider{
		WorkspaceModuleSetup(),
		NotificationModuleSetup(),
		messaging.ModuleSetup(),
		PassportsModuleSetup(),
	}
}

func WorkspaceModuleSetup() *application.ModuleProvider {

	m := AbacEssentialModule()
	// Default Fireback authorization. You can Override this on microservices
	fireback.WithAuthorizationPure = WithAuthorizationPureDefault
	fireback.WithAuthorizationFn = WithAuthorizationFn
	fireback.AuthorizeRequest = AuthorizeRequest
	fireback.WithSocketAuthorization = WithSocketAuthorization
	fireback.MeetsAccessLevel = MeetsAccessLevel

	module := &application.ModuleProvider{
		Name:               "abac",
		OnEnvInit:          OnInitEnvHook,
		GoMigrateDirectory: &migrations.MigrationsFs,

		// Actions declared in Abac.emi.yml (moved out of AbacModule3.yml's old actions:
		// section) are wired directly here, the same way FirebackModuleSetup wires the
		// Capability* actions - rather than through the legacy Module3Action/Impl glue.
		GinWebServerInitHooks: m.AbacGinWebServerInitHooks,
	}

	module.ProvidePermissionHandler(m.Permissions...)
	module.ProvideEntityHandlers(m.MigrationFn)

	// TimezoneGroupSyncSeeders moved to modules/abac/interfacetools - see
	// interfacetools.ModuleSetup's own ProvideSeederImportHandler call.

	module.MigrationFunction = func(x *application.Application, db *gorm.DB) {
		SyncPermissionsInDatabase(x, db)
		// Insert-if-missing default regional content (email/sms OTP, invite-to-
		// workspace templates in en/fa/pl) - see RegionalContentDefaults.go's own
		// doc comment for why this replaced the old, never-actually-wired
		// modules/abac/seeders/RegionalContent/*.yml.
		SyncRegionalContentDefaults(db)
	}

	// Abac's own `config:` block (Abac.emi.yml) - folded into fireback's combined
	// "config list" listing (see modules/fireback/ConfigRegistry.go) alongside
	// fireback's own core config and every other registered module's, the same way
	// abac.Menu lets abac contribute AppMenu items without interfacetools importing it
	// back. LoadConfiguration() (not a cached snapshot) so it always reflects the
	// latest env/`.env`-loaded values, including ones changed at runtime via the
	// "abac config otp-lockout-seconds set" CLI command below.
	module.ProvideConfigHandler(func() interface{} {
		return abacdefs.LoadConfiguration()
	})

	module.AppendCli = func(x *application.Application) []*cli.Command {
		return []*cli.Command{
			GetMigrationCommand(x),
		}
	}

	module.ProvideCliHandlers([]*cli.Command{

		UserCliFn(),
		WorkspaceCliFn(),
		&MiscCli,
		// The actions below moved to Abac.emi.yml and don't have a home yet among the
		// entity-scoped cli groups above (AcceptInvite/UserInvitations live under UserCliFn,
		// QueryUserRoleWorkspaces/QueryWorkspaceTypesPublicly under WorkspaceCliFn,
		// OsLoginAuthenticate/CheckPassportMethods/UserPassports/OauthAuthenticate under
		// PassportCliFn - see UserEntity.go/WorkspaceCli.go/PassportCli.go).
		abacdefs.WhoamiActionCliHandler(WhoamiAction),
		abacdefs.AnalyticsOverviewActionCliHandler(AnalyticsOverviewAction),
		abacdefs.SignoutActionCliHandler(SignoutAction),
		abacdefs.InviteToWorkspaceActionCliHandler(InviteToWorkspaceAction),
		abacdefs.AddUserToWorkspaceActionCliHandler(AddUserToWorkspaceAction),
		abacdefs.RemoveUserFromWorkspaceActionCliHandler(RemoveUserFromWorkspaceAction),
		abacdefs.ChangeUserWorkspaceRoleActionCliHandler(ChangeUserWorkspaceRoleAction),
		abacdefs.CreatePassportForUserActionCliHandler(CreatePassportForUserAction),
		abacdefs.QueryWorkspaceRolesActionCliHandler(QueryWorkspaceRolesAction),
		abacdefs.ConfirmClassicPassportTotpActionCliHandler(ConfirmClassicPassportTotpAction),
		abacdefs.ChangePasswordActionCliHandler(ChangePasswordAction),
		abacdefs.CreateWorkspaceActionCliHandler(CreateWorkspaceAction),
		abacdefs.ClassicPassportRequestOtpActionCliHandler(ClassicPassportRequestOtpAction),
		abacdefs.ClassicPassportOtpActionCliHandler(ClassicPassportOtpAction),
		abacdefs.CheckClassicPassportActionCliHandler(CheckClassicPassportAction),
		abacdefs.ClassicSignupActionCliHandler(ClassicSignupAction),
		abacdefs.ClassicSigninActionCliHandler(ClassicSigninAction),
		abacdefs.GsmSendSmsActionCliHandler(GsmSendSmsAction),

		{
			Name:        "role",
			Description: "Actions related to roles, creation, browsing and other.",
			Commands: []*cli.Command{
				abacdefs.RoleBrowseActionCliHandler(RoleBrowseAction),
				abacdefs.RoleGetActionCliHandler(RoleGetAction),
				abacdefs.RoleCreateActionCliHandler(RoleCreateAction),
				abacdefs.RoleUpdateActionCliHandler(RoleUpdateAction),
				abacdefs.RoleAwareDeletePreviewActionCliHandler(RoleAwareDeletePreviewAction),
				abacdefs.RoleAwareDeleteActionCliHandler(RoleAwareDeleteAction),
			},
		},
	})

	module.ProvideCliHandlers([]*cli.Command{&AuthFlow, &AbacActions})

	return module
}

var AbacActions cli.Command = cli.Command{
	Name:  "abac",
	Usage: "All actions which are available for abac module",
	Commands: append(
		[]*cli.Command{
			{
				// Per-field get/set for Abac.emi.yml's `config:` block - abacdefs.GetConfigCli()
				// is otherwise identical in shape to fireback's own top-level "config" command
				// (ConfigCommand, CoreUtils.go), just nested here under "abac" to avoid a name
				// collision between the two at the CLI root.
				Name:     "config",
				Usage:    "Set of tools to configure the abac module",
				Commands: abacdefs.GetConfigCli(),
			},
			{
				Name:  "internal",
				Usage: "Internal entities which are used for processes. Manipulating these requires deep internal knowledge",
				Commands: []*cli.Command{
					abacdefs.PublicJoinKeyBrowseActionCliHandler(PublicJoinKeyBrowseAction),
					abacdefs.PublicJoinKeyGetActionCliHandler(PublicJoinKeyGetAction),
					abacdefs.PublicJoinKeyCreateActionCliHandler(PublicJoinKeyCreateAction),
					abacdefs.PublicJoinKeyUpdateActionCliHandler(PublicJoinKeyUpdateAction),
					abacdefs.PublicJoinKeyAwareDeletePreviewActionCliHandler(PublicJoinKeyAwareDeletePreviewAction),
					abacdefs.PublicJoinKeyAwareDeleteActionCliHandler(PublicJoinKeyAwareDeleteAction),
					abacdefs.PublicAuthenticationBrowseActionCliHandler(PublicAuthenticationBrowseAction),
					abacdefs.PublicAuthenticationGetActionCliHandler(PublicAuthenticationGetAction),
					abacdefs.PublicAuthenticationCreateActionCliHandler(PublicAuthenticationCreateAction),
					abacdefs.PublicAuthenticationUpdateActionCliHandler(PublicAuthenticationUpdateAction),
					abacdefs.PublicAuthenticationAwareDeletePreviewActionCliHandler(PublicAuthenticationAwareDeletePreviewAction),
					abacdefs.PublicAuthenticationAwareDeleteActionCliHandler(PublicAuthenticationAwareDeleteAction),
					abacdefs.PreferenceBrowseActionCliHandler(PreferenceBrowseAction),
					abacdefs.PreferenceGetActionCliHandler(PreferenceGetAction),
					abacdefs.PreferenceCreateActionCliHandler(PreferenceCreateAction),
					abacdefs.PreferenceUpdateActionCliHandler(PreferenceUpdateAction),
					abacdefs.PreferenceAwareDeletePreviewActionCliHandler(PreferenceAwareDeletePreviewAction),
					abacdefs.PreferenceAwareDeleteActionCliHandler(PreferenceAwareDeleteAction),
					abacdefs.UserProfileBrowseActionCliHandler(UserProfileBrowseAction),
					abacdefs.UserProfileGetActionCliHandler(UserProfileGetAction),
					abacdefs.UserProfileCreateActionCliHandler(UserProfileCreateAction),
					abacdefs.UserProfileUpdateActionCliHandler(UserProfileUpdateAction),
					abacdefs.UserProfileAwareDeletePreviewActionCliHandler(UserProfileAwareDeletePreviewAction),
					abacdefs.UserProfileAwareDeleteActionCliHandler(UserProfileAwareDeleteAction),
					abacdefs.PendingWorkspaceInviteBrowseActionCliHandler(PendingWorkspaceInviteBrowseAction),
					abacdefs.PendingWorkspaceInviteGetActionCliHandler(PendingWorkspaceInviteGetAction),
					abacdefs.PendingWorkspaceInviteCreateActionCliHandler(PendingWorkspaceInviteCreateAction),
					abacdefs.PendingWorkspaceInviteUpdateActionCliHandler(PendingWorkspaceInviteUpdateAction),
					abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionCliHandler(PendingWorkspaceInviteAwareDeletePreviewAction),
					abacdefs.PendingWorkspaceInviteAwareDeleteActionCliHandler(PendingWorkspaceInviteAwareDeleteAction),
					abacdefs.CapabilityBrowseActionCliHandler(GetCapabilitiesAction),
					abacdefs.CapabilityGetActionCliHandler(CapabilityGetAction),
					abacdefs.CapabilityCreateActionCliHandler(CapabilityCreateAction),
					abacdefs.CapabilityUpdateActionCliHandler(CapabilityUpdateAction),
					abacdefs.CapabilityAwareDeleteActionCliHandler(CapabilityAwareDeleteAction),
					abacdefs.CapabilitiesTreeActionCliHandler(CapabilitiesTreeAction),
				},
			},
		},
		GetAbacActionsCli()...,
	),
}
