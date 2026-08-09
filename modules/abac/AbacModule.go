package abac

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/abac/interfacetools"
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
func AbacCompleteModules() []*application.ModuleProvider {
	return []*application.ModuleProvider{
		WorkspaceModuleSetup(),
		NotificationModuleSetup(),
		messaging.ModuleSetup(),
		interfacetools.ModuleSetup(),
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
				QueryUserRoleWorkspacesActionGin(g, QueryUserRoleWorkspacesAction)
				InviteToWorkspaceActionGin(g, InviteToWorkspaceAction)
				UserInvitationsActionGin(g, UserInvitationsAction)
				SignoutActionGin(g, SignoutAction)
				OauthAuthenticateActionGin(g, OauthAuthenticateAction)
				AcceptInviteActionGin(g, AcceptInviteAction)
				ConfirmClassicPassportTotpActionGin(g, ConfirmClassicPassportTotpAction)
				ChangePasswordActionGin(g, ChangePasswordAction)
				UserPassportsActionGin(g, UserPassportsAction)
				CreateWorkspaceActionGin(g, CreateWorkspaceAction)
				ClassicPassportRequestOtpActionGin(g, ClassicPassportRequestOtpAction)
				ClassicPassportOtpActionGin(g, ClassicPassportOtpAction)
				CheckClassicPassportActionGin(g, CheckClassicPassportAction)
				ClassicSignupActionGin(g, ClassicSignupAction)
				ClassicSigninActionGin(g, ClassicSigninAction)
				QueryWorkspaceTypesPubliclyActionGin(g, QueryWorkspaceTypesPubliclyAction)
				CheckPassportMethodsActionGin(g, CheckPassportMethodsAction)
				OsLoginAuthenticateActionGin(g, OsLoginAuthenticateAction)
				// SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider moved to
				// modules/abac/messaging - see messaging.ModuleSetup. GsmSendSms stays
				// here since it depends on NotificationConfigEntity, which stays in abac.
				GsmSendSmsActionGin(g, GsmSendSmsAction)

				// TableViewSizing, TimezoneGroup, AppMenu (and /cte-app-menus) moved to
				// modules/abac/interfacetools - see interfacetools.ModuleSetup.

				PreferenceBrowseActionGin(g, PreferenceBrowseAction)
				PreferenceGetActionGin(g, PreferenceGetAction)
				PreferenceCreateActionGin(g, PreferenceCreateAction)
				PreferenceUpdateActionGin(g, PreferenceUpdateAction)
				PreferenceAwareDeletePreviewActionGin(g, PreferenceAwareDeletePreviewAction)
				PreferenceAwareDeleteActionGin(g, PreferenceAwareDeleteAction)

				UserProfileBrowseActionGin(g, UserProfileBrowseAction)
				UserProfileGetActionGin(g, UserProfileGetAction)
				UserProfileCreateActionGin(g, UserProfileCreateAction)
				UserProfileUpdateActionGin(g, UserProfileUpdateAction)
				UserProfileAwareDeletePreviewActionGin(g, UserProfileAwareDeletePreviewAction)
				UserProfileAwareDeleteActionGin(g, UserProfileAwareDeleteAction)

				PendingWorkspaceInviteBrowseActionGin(g, PendingWorkspaceInviteBrowseAction)
				PendingWorkspaceInviteGetActionGin(g, PendingWorkspaceInviteGetAction)
				PendingWorkspaceInviteCreateActionGin(g, PendingWorkspaceInviteCreateAction)
				PendingWorkspaceInviteUpdateActionGin(g, PendingWorkspaceInviteUpdateAction)
				PendingWorkspaceInviteAwareDeletePreviewActionGin(g, PendingWorkspaceInviteAwareDeletePreviewAction)
				PendingWorkspaceInviteAwareDeleteActionGin(g, PendingWorkspaceInviteAwareDeleteAction)

				TokenBrowseActionGin(g, TokenBrowseAction)
				TokenGetActionGin(g, TokenGetAction)
				TokenCreateActionGin(g, TokenCreateAction)
				TokenUpdateActionGin(g, TokenUpdateAction)
				TokenAwareDeletePreviewActionGin(g, TokenAwareDeletePreviewAction)
				TokenAwareDeleteActionGin(g, TokenAwareDeleteAction)

				WorkspaceInviteBrowseActionGin(g, WorkspaceInviteBrowseAction)
				WorkspaceInviteGetActionGin(g, WorkspaceInviteGetAction)
				WorkspaceInviteCreateActionGin(g, WorkspaceInviteCreateAction)
				WorkspaceInviteUpdateActionGin(g, WorkspaceInviteUpdateAction)
				WorkspaceInviteAwareDeletePreviewActionGin(g, WorkspaceInviteAwareDeletePreviewAction)
				WorkspaceInviteAwareDeleteActionGin(g, WorkspaceInviteAwareDeleteAction)

				UserWorkspaceBrowseActionGin(g, UserWorkspaceBrowseAction)
				UserWorkspaceGetActionGin(g, UserWorkspaceGetAction)
				UserWorkspaceCreateActionGin(g, UserWorkspaceCreateAction)
				UserWorkspaceUpdateActionGin(g, UserWorkspaceUpdateAction)
				UserWorkspaceAwareDeletePreviewActionGin(g, UserWorkspaceAwareDeletePreviewAction)
				UserWorkspaceAwareDeleteActionGin(g, UserWorkspaceAwareDeleteAction)

				WorkspaceRoleBrowseActionGin(g, WorkspaceRoleBrowseAction)
				WorkspaceRoleGetActionGin(g, WorkspaceRoleGetAction)
				WorkspaceRoleCreateActionGin(g, WorkspaceRoleCreateAction)
				WorkspaceRoleUpdateActionGin(g, WorkspaceRoleUpdateAction)
				WorkspaceRoleAwareDeletePreviewActionGin(g, WorkspaceRoleAwareDeletePreviewAction)
				WorkspaceRoleAwareDeleteActionGin(g, WorkspaceRoleAwareDeleteAction)

				RegionalContentBrowseActionGin(g, RegionalContentBrowseAction)
				RegionalContentGetActionGin(g, RegionalContentGetAction)
				RegionalContentCreateActionGin(g, RegionalContentCreateAction)
				RegionalContentUpdateActionGin(g, RegionalContentUpdateAction)
				RegionalContentAwareDeletePreviewActionGin(g, RegionalContentAwareDeletePreviewAction)
				RegionalContentAwareDeleteActionGin(g, RegionalContentAwareDeleteAction)

				CapabilityBrowseActionGin(g, GetCapabilitiesAction)
				CapabilityGetActionGin(g, CapabilityGetAction)
				CapabilityUpdateActionGin(g, CapabilityUpdateAction)
				CapabilityAwareDeleteActionGin(g, CapabilityAwareDeleteAction)
				CapabilitiesTreeActionGin(g, CapabilitiesTreeAction)

				WorkspaceConfigBrowseActionGin(g, WorkspaceConfigBrowseAction)
				WorkspaceConfigGetActionGin(g, WorkspaceConfigGetAction)
				WorkspaceConfigCreateActionGin(g, WorkspaceConfigCreateAction)
				WorkspaceConfigUpdateActionGin(g, WorkspaceConfigUpdateAction)
				WorkspaceConfigAwareDeletePreviewActionGin(g, WorkspaceConfigAwareDeletePreviewAction)
				WorkspaceConfigAwareDeleteActionGin(g, WorkspaceConfigAwareDeleteAction)
				WorkspaceConfigDistinctGetActionGin(g, WorkspaceConfigDistinctGetAction)
				WorkspaceConfigDistinctUpdateActionGin(g, WorkspaceConfigDistinctUpdateAction)

				RoleBrowseActionGin(g, RoleBrowseAction)
				RoleGetActionGin(g, RoleGetAction)
				RoleCreateActionGin(g, RoleCreateAction)
				RoleUpdateActionGin(g, RoleUpdateAction)
				RoleAwareDeletePreviewActionGin(g, RoleAwareDeletePreviewAction)
				RoleAwareDeleteActionGin(g, RoleAwareDeleteAction)

				WorkspaceTypeBrowseActionGin(g, WorkspaceTypeBrowseAction)
				WorkspaceTypeGetActionGin(g, WorkspaceTypeGetAction)
				WorkspaceTypeCreateActionGin(g, WorkspaceTypeCreateAction)
				WorkspaceTypeUpdateActionGin(g, WorkspaceTypeUpdateAction)
				WorkspaceTypeAwareDeletePreviewActionGin(g, WorkspaceTypeAwareDeletePreviewAction)
				WorkspaceTypeAwareDeleteActionGin(g, WorkspaceTypeAwareDeleteAction)

				UserBrowseActionGin(g, UserBrowseAction)
				UserGetActionGin(g, UserGetAction)
				UserCreateActionGin(g, UserCreateAction)
				UserUpdateActionGin(g, UserUpdateAction)
				UserAwareDeletePreviewActionGin(g, UserAwareDeletePreviewAction)
				UserAwareDeleteActionGin(g, UserAwareDeleteAction)

				WorkspaceBrowseActionGin(g, WorkspaceBrowseAction)
				WorkspaceGetActionGin(g, WorkspaceGetAction)
				WorkspaceCreateActionGin(g, WorkspaceCreateAction)
				WorkspaceUpdateActionGin(g, WorkspaceUpdateAction)
				WorkspaceAwareDeletePreviewActionGin(g, WorkspaceAwareDeletePreviewAction)
				WorkspaceAwareDeleteActionGin(g, WorkspaceAwareDeleteAction)

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
			&UserEntity{},
			&TokenEntity{},
			&PreferenceEntity{},
			&RoleEntity{},
			&WorkspaceEntity{},
			&WorkspaceInviteEntity{},
			&WorkspaceConfigEntity{},
			&WorkspaceTypeEntity{},
			&WorkspaceRoleEntity{},
			&UserWorkspaceEntity{},
			&RegionalContentEntity{},
			&UserProfileEntity{},
			&PendingWorkspaceInviteEntity{},
			&CapabilityEntity{},
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
	}

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
		SignoutActionCliHandler(SignoutAction),
		InviteToWorkspaceActionCliHandler(InviteToWorkspaceAction),
		ConfirmClassicPassportTotpActionCliHandler(ConfirmClassicPassportTotpAction),
		ChangePasswordActionCliHandler(ChangePasswordAction),
		CreateWorkspaceActionCliHandler(CreateWorkspaceAction),
		ClassicPassportRequestOtpActionCliHandler(ClassicPassportRequestOtpAction),
		ClassicPassportOtpActionCliHandler(ClassicPassportOtpAction),
		CheckClassicPassportActionCliHandler(CheckClassicPassportAction),
		ClassicSignupActionCliHandler(ClassicSignupAction),
		ClassicSigninActionCliHandler(ClassicSigninAction),
		// SendEmail/SendEmailWithProvider/GsmSendSmsWithProvider moved to
		// modules/abac/messaging - see messaging.ModuleSetup.
		GsmSendSmsActionCliHandler(GsmSendSmsAction),

		{
			Name:        "role",
			Description: "Actions related to roles, creation, browsing and other.",
			Commands: []*cli.Command{
				RoleBrowseActionCliHandler(RoleBrowseAction),
				RoleGetActionCliHandler(RoleGetAction),
				RoleCreateActionCliHandler(RoleCreateAction),
				RoleUpdateActionCliHandler(RoleUpdateAction),
				RoleAwareDeletePreviewActionCliHandler(RoleAwareDeletePreviewAction),
				RoleAwareDeleteActionCliHandler(RoleAwareDeleteAction),
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
				Name:  "internal",
				Usage: "Internal entities which are used for processes. Manipulating these requires deep internal knowledge",
				Commands: []*cli.Command{
					PublicJoinKeyBrowseActionCliHandler(PublicJoinKeyBrowseAction),
					PublicJoinKeyGetActionCliHandler(PublicJoinKeyGetAction),
					PublicJoinKeyCreateActionCliHandler(PublicJoinKeyCreateAction),
					PublicJoinKeyUpdateActionCliHandler(PublicJoinKeyUpdateAction),
					PublicJoinKeyAwareDeletePreviewActionCliHandler(PublicJoinKeyAwareDeletePreviewAction),
					PublicJoinKeyAwareDeleteActionCliHandler(PublicJoinKeyAwareDeleteAction),
					PublicAuthenticationBrowseActionCliHandler(PublicAuthenticationBrowseAction),
					PublicAuthenticationGetActionCliHandler(PublicAuthenticationGetAction),
					PublicAuthenticationCreateActionCliHandler(PublicAuthenticationCreateAction),
					PublicAuthenticationUpdateActionCliHandler(PublicAuthenticationUpdateAction),
					PublicAuthenticationAwareDeletePreviewActionCliHandler(PublicAuthenticationAwareDeletePreviewAction),
					PublicAuthenticationAwareDeleteActionCliHandler(PublicAuthenticationAwareDeleteAction),
					PreferenceBrowseActionCliHandler(PreferenceBrowseAction),
					PreferenceGetActionCliHandler(PreferenceGetAction),
					PreferenceCreateActionCliHandler(PreferenceCreateAction),
					PreferenceUpdateActionCliHandler(PreferenceUpdateAction),
					PreferenceAwareDeletePreviewActionCliHandler(PreferenceAwareDeletePreviewAction),
					PreferenceAwareDeleteActionCliHandler(PreferenceAwareDeleteAction),
					UserProfileBrowseActionCliHandler(UserProfileBrowseAction),
					UserProfileGetActionCliHandler(UserProfileGetAction),
					UserProfileCreateActionCliHandler(UserProfileCreateAction),
					UserProfileUpdateActionCliHandler(UserProfileUpdateAction),
					UserProfileAwareDeletePreviewActionCliHandler(UserProfileAwareDeletePreviewAction),
					UserProfileAwareDeleteActionCliHandler(UserProfileAwareDeleteAction),
					PendingWorkspaceInviteBrowseActionCliHandler(PendingWorkspaceInviteBrowseAction),
					PendingWorkspaceInviteGetActionCliHandler(PendingWorkspaceInviteGetAction),
					PendingWorkspaceInviteCreateActionCliHandler(PendingWorkspaceInviteCreateAction),
					PendingWorkspaceInviteUpdateActionCliHandler(PendingWorkspaceInviteUpdateAction),
					PendingWorkspaceInviteAwareDeletePreviewActionCliHandler(PendingWorkspaceInviteAwareDeletePreviewAction),
					PendingWorkspaceInviteAwareDeleteActionCliHandler(PendingWorkspaceInviteAwareDeleteAction),
					CapabilityBrowseActionCliHandler(GetCapabilitiesAction),
					CapabilityGetActionCliHandler(CapabilityGetAction),
					CapabilityUpdateActionCliHandler(CapabilityUpdateAction),
					CapabilityAwareDeleteActionCliHandler(CapabilityAwareDeleteAction),
					CapabilitiesTreeActionCliHandler(CapabilitiesTreeAction),
				},
			},
		},
		GetAbacActionsCli()...,
	),
}
