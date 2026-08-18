package abac

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/messaging"
	"github.com/torabian/fireback/modules/abac/migrations"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

type MicroserviceSetupConfig struct {
	AuthorizationResolver WithAuthorizationPureImpl
}

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
		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				abacdefs.WhoamiActionGin(g, WhoamiAction)
				abacdefs.QueryUserRoleWorkspacesActionGin(g, QueryUserRoleWorkspacesAction)
				abacdefs.InviteToWorkspaceActionGin(g, InviteToWorkspaceAction)
				abacdefs.AddUserToWorkspaceActionGin(g, AddUserToWorkspaceAction)
				abacdefs.RemoveUserFromWorkspaceActionGin(g, RemoveUserFromWorkspaceAction)
				abacdefs.ChangeUserWorkspaceRoleActionGin(g, ChangeUserWorkspaceRoleAction)
				abacdefs.CreatePassportForUserActionGin(g, CreatePassportForUserAction)
				abacdefs.QueryWorkspaceRolesActionGin(g, QueryWorkspaceRolesAction)
				abacdefs.UserInvitationsActionGin(g, UserInvitationsAction)
				abacdefs.SignoutActionGin(g, SignoutAction)
				abacdefs.OauthAuthenticateActionGin(g, OauthAuthenticateAction)
				abacdefs.AcceptInviteActionGin(g, AcceptInviteAction)
				abacdefs.ConfirmClassicPassportTotpActionGin(g, ConfirmClassicPassportTotpAction)
				abacdefs.ChangePasswordActionGin(g, ChangePasswordAction)
				abacdefs.SetPassportPasswordActionGin(g, SetPassportPasswordAction)
				abacdefs.DisablePassportTotpActionGin(g, DisablePassportTotpAction)
				abacdefs.SendPassportResetEmailActionGin(g, SendPassportResetEmailAction)
				abacdefs.CompletePassportPasswordResetActionGin(g, CompletePassportPasswordResetAction)
				abacdefs.UserPassportsActionGin(g, UserPassportsAction)
				abacdefs.CreateWorkspaceActionGin(g, CreateWorkspaceAction)
				abacdefs.ClassicPassportRequestOtpActionGin(g, ClassicPassportRequestOtpAction)
				abacdefs.ClassicPassportOtpActionGin(g, ClassicPassportOtpAction)
				abacdefs.CheckClassicPassportActionGin(g, CheckClassicPassportAction)
				abacdefs.ClassicSignupActionGin(g, ClassicSignupAction)
				abacdefs.ClassicSigninActionGin(g, ClassicSigninAction)
				abacdefs.QueryWorkspaceTypesPubliclyActionGin(g, QueryWorkspaceTypesPubliclyAction)
				abacdefs.CheckPassportMethodsActionGin(g, CheckPassportMethodsAction)
				abacdefs.OsLoginAuthenticateActionGin(g, OsLoginAuthenticateAction)
				// SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider moved to
				// modules/abac/messaging - see messaging.ModuleSetup. GsmSendSms stays
				// here since it depends on NotificationConfigEntity, which stays in abac.
				abacdefs.GsmSendSmsActionGin(g, GsmSendSmsAction)

				// TableViewSizing, TimezoneGroup, AppMenu (and /cte-app-menus) moved to
				// modules/abac/interfacetools - see interfacetools.ModuleSetup.

				abacdefs.PreferenceBrowseActionGin(g, PreferenceBrowseAction)
				abacdefs.PreferenceGetActionGin(g, PreferenceGetAction)
				abacdefs.PreferenceCreateActionGin(g, PreferenceCreateAction)
				abacdefs.PreferenceUpdateActionGin(g, PreferenceUpdateAction)
				abacdefs.PreferenceAwareDeletePreviewActionGin(g, PreferenceAwareDeletePreviewAction)
				abacdefs.PreferenceAwareDeleteActionGin(g, PreferenceAwareDeleteAction)

				abacdefs.UserProfileBrowseActionGin(g, UserProfileBrowseAction)
				abacdefs.UserProfileGetActionGin(g, UserProfileGetAction)
				abacdefs.UserProfileCreateActionGin(g, UserProfileCreateAction)
				abacdefs.UserProfileUpdateActionGin(g, UserProfileUpdateAction)
				abacdefs.UserProfileAwareDeletePreviewActionGin(g, UserProfileAwareDeletePreviewAction)
				abacdefs.UserProfileAwareDeleteActionGin(g, UserProfileAwareDeleteAction)

				abacdefs.PendingWorkspaceInviteBrowseActionGin(g, PendingWorkspaceInviteBrowseAction)
				abacdefs.PendingWorkspaceInviteGetActionGin(g, PendingWorkspaceInviteGetAction)
				abacdefs.PendingWorkspaceInviteCreateActionGin(g, PendingWorkspaceInviteCreateAction)
				abacdefs.PendingWorkspaceInviteUpdateActionGin(g, PendingWorkspaceInviteUpdateAction)
				abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionGin(g, PendingWorkspaceInviteAwareDeletePreviewAction)
				abacdefs.PendingWorkspaceInviteAwareDeleteActionGin(g, PendingWorkspaceInviteAwareDeleteAction)

				abacdefs.TokenBrowseActionGin(g, TokenBrowseAction)
				abacdefs.TokenGetActionGin(g, TokenGetAction)
				abacdefs.TokenCreateActionGin(g, TokenCreateAction)
				abacdefs.TokenUpdateActionGin(g, TokenUpdateAction)
				abacdefs.TokenAwareDeletePreviewActionGin(g, TokenAwareDeletePreviewAction)
				abacdefs.TokenAwareDeleteActionGin(g, TokenAwareDeleteAction)

				abacdefs.WorkspaceInviteBrowseActionGin(g, WorkspaceInviteBrowseAction)
				abacdefs.WorkspaceInviteGetActionGin(g, WorkspaceInviteGetAction)
				abacdefs.WorkspaceInviteCreateActionGin(g, WorkspaceInviteCreateAction)
				abacdefs.WorkspaceInviteUpdateActionGin(g, WorkspaceInviteUpdateAction)
				abacdefs.WorkspaceInviteAwareDeletePreviewActionGin(g, WorkspaceInviteAwareDeletePreviewAction)
				abacdefs.WorkspaceInviteAwareDeleteActionGin(g, WorkspaceInviteAwareDeleteAction)

				abacdefs.UserWorkspaceBrowseActionGin(g, UserWorkspaceBrowseAction)
				abacdefs.UserWorkspaceGetActionGin(g, UserWorkspaceGetAction)
				abacdefs.UserWorkspaceCreateActionGin(g, UserWorkspaceCreateAction)
				abacdefs.UserWorkspaceUpdateActionGin(g, UserWorkspaceUpdateAction)
				abacdefs.UserWorkspaceAwareDeletePreviewActionGin(g, UserWorkspaceAwareDeletePreviewAction)
				abacdefs.UserWorkspaceAwareDeleteActionGin(g, UserWorkspaceAwareDeleteAction)

				abacdefs.WorkspaceRoleBrowseActionGin(g, WorkspaceRoleBrowseAction)
				abacdefs.WorkspaceRoleGetActionGin(g, WorkspaceRoleGetAction)
				abacdefs.WorkspaceRoleCreateActionGin(g, WorkspaceRoleCreateAction)
				abacdefs.WorkspaceRoleUpdateActionGin(g, WorkspaceRoleUpdateAction)
				abacdefs.WorkspaceRoleAwareDeletePreviewActionGin(g, WorkspaceRoleAwareDeletePreviewAction)
				abacdefs.WorkspaceRoleAwareDeleteActionGin(g, WorkspaceRoleAwareDeleteAction)

				abacdefs.RegionalContentBrowseActionGin(g, RegionalContentBrowseAction)
				abacdefs.RegionalContentGetActionGin(g, RegionalContentGetAction)
				abacdefs.RegionalContentCreateActionGin(g, RegionalContentCreateAction)
				abacdefs.RegionalContentUpdateActionGin(g, RegionalContentUpdateAction)
				abacdefs.RegionalContentAwareDeletePreviewActionGin(g, RegionalContentAwareDeletePreviewAction)
				abacdefs.RegionalContentAwareDeleteActionGin(g, RegionalContentAwareDeleteAction)

				abacdefs.CapabilityBrowseActionGin(g, GetCapabilitiesAction)
				abacdefs.CapabilityGetActionGin(g, CapabilityGetAction)
				abacdefs.CapabilityCreateActionGin(g, CapabilityCreateAction)
				abacdefs.CapabilityUpdateActionGin(g, CapabilityUpdateAction)
				abacdefs.CapabilityAwareDeletePreviewActionGin(g, CapabilityAwareDeletePreviewAction)
				abacdefs.CapabilityAwareDeleteActionGin(g, CapabilityAwareDeleteAction)
				abacdefs.CapabilitiesTreeActionGin(g, CapabilitiesTreeAction)

				abacdefs.WorkspaceConfigBrowseActionGin(g, WorkspaceConfigBrowseAction)
				abacdefs.WorkspaceConfigGetActionGin(g, WorkspaceConfigGetAction)
				abacdefs.WorkspaceConfigCreateActionGin(g, WorkspaceConfigCreateAction)
				abacdefs.WorkspaceConfigUpdateActionGin(g, WorkspaceConfigUpdateAction)
				abacdefs.WorkspaceConfigAwareDeletePreviewActionGin(g, WorkspaceConfigAwareDeletePreviewAction)
				abacdefs.WorkspaceConfigAwareDeleteActionGin(g, WorkspaceConfigAwareDeleteAction)
				abacdefs.WorkspaceConfigDistinctGetActionGin(g, WorkspaceConfigDistinctGetAction)
				abacdefs.WorkspaceConfigDistinctUpdateActionGin(g, WorkspaceConfigDistinctUpdateAction)

				abacdefs.RoleBrowseActionGin(g, RoleBrowseAction)
				abacdefs.RoleGetActionGin(g, RoleGetAction)
				abacdefs.RoleCreateActionGin(g, RoleCreateAction)
				abacdefs.RoleUpdateActionGin(g, RoleUpdateAction)
				abacdefs.RoleAwareDeletePreviewActionGin(g, RoleAwareDeletePreviewAction)
				abacdefs.RoleAwareDeleteActionGin(g, RoleAwareDeleteAction)

				abacdefs.WorkspaceTypeBrowseActionGin(g, WorkspaceTypeBrowseAction)
				abacdefs.WorkspaceTypeGetActionGin(g, WorkspaceTypeGetAction)
				abacdefs.WorkspaceTypeCreateActionGin(g, WorkspaceTypeCreateAction)
				abacdefs.WorkspaceTypeUpdateActionGin(g, WorkspaceTypeUpdateAction)
				abacdefs.WorkspaceTypeAwareDeletePreviewActionGin(g, WorkspaceTypeAwareDeletePreviewAction)
				abacdefs.WorkspaceTypeAwareDeleteActionGin(g, WorkspaceTypeAwareDeleteAction)

				abacdefs.UserBrowseActionGin(g, UserBrowseAction)
				abacdefs.UserGetActionGin(g, UserGetAction)
				abacdefs.UserCreateActionGin(g, UserCreateAction)
				abacdefs.UserUpdateActionGin(g, UserUpdateAction)
				abacdefs.UserAwareDeletePreviewActionGin(g, UserAwareDeletePreviewAction)
				abacdefs.UserAwareDeleteActionGin(g, UserAwareDeleteAction)

				abacdefs.WorkspaceBrowseActionGin(g, WorkspaceBrowseAction)
				abacdefs.WorkspaceGetActionGin(g, WorkspaceGetAction)
				abacdefs.WorkspaceCreateActionGin(g, WorkspaceCreateAction)
				abacdefs.WorkspaceUpdateActionGin(g, WorkspaceUpdateAction)
				abacdefs.WorkspaceAwareDeletePreviewActionGin(g, WorkspaceAwareDeletePreviewAction)
				abacdefs.WorkspaceAwareDeleteActionGin(g, WorkspaceAwareDeleteAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_WORKSPACE_CONFIG_PERMISSIONS,
		ALL_WORKSPACE_TYPE_PERMISSIONS,
		// ALL_EMAIL_SENDER_PERMISSIONS/ALL_EMAIL_PROVIDER_PERMISSIONS/ALL_GSM_PROVIDER_PERMISSIONS
		// moved to modules/abac/messaging - already registered via messaging.ModuleSetup's
		// own ProvidePermissionHandler call (see AbacCompleteModules), no need to repeat here.
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
		ALL_WORKSPACE_INVITE_PERMISSIONS,
		// ALL_TABLE_VIEW_SIZING_PERMISSIONS/ALL_APP_MENU_PERMISSIONS/ALL_TIMEZONE_GROUP_PERMISSIONS
		// moved to modules/abac/interfacetools - already registered via
		// interfacetools.ModuleSetup's own ProvidePermissionHandler call (see
		// AbacCompleteModules), no need to repeat here.
		ALL_REGIONAL_CONTENT_PERMISSIONS,
		ALL_USER_WORKSPACE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_WORKSPACE_ROLE_PERMISSIONS,
		ALL_WORKSPACE_PERMISSIONS,
		ALL_PERM_ABAC_MODULE,
		ALL_PREFERENCE_PERMISSIONS,
		ALL_USER_PROFILE_PERMISSIONS,
		ALL_PENDING_WORKSPACE_INVITE_PERMISSIONS,
		ALL_TOKEN_PERMISSIONS,
		ALL_CAPABILITY_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		items := []interface{}{
			&abacdefs.UserEntity{},
			&abacdefs.TokenEntity{},
			&abacdefs.PreferenceEntity{},
			&abacdefs.RoleEntity{},
			&abacdefs.WorkspaceEntity{},
			&abacdefs.WorkspaceInviteEntity{},
			&abacdefs.WorkspaceConfigEntity{},
			&abacdefs.WorkspaceTypeEntity{},
			&abacdefs.WorkspaceRoleEntity{},
			&abacdefs.UserWorkspaceEntity{},
			&abacdefs.RegionalContentEntity{},
			&abacdefs.UserProfileEntity{},
			&abacdefs.PendingWorkspaceInviteEntity{},
			&abacdefs.CapabilityEntity{},
		}

		items2 := []interface{}{}
		items2 = append(items2, items...)

		for _, item := range items2 {

			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", reflect.ValueOf(item).Elem().String())
				return err
			}
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
	})

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
