package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

func PassportsModuleSetup() *application.ModuleProvider {
	module := &application.ModuleProvider{

		// it must write on the workspaces instead
		Name: "abac",

		// passportMethod/publicJoinKey/emailConfirmation/phoneConfirmation moved from
		// AbacModule3.yml's old entities: section to Abac.emi.yml, so they're wired
		// directly here now, the same way FirebackModuleSetup wires Capability* actions.
		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *application.Application) error{
			func(g *gin.RouterGroup, x *application.Application) error {
				PassportMethodBrowseActionGin(g, PassportMethodBrowseAction)
				PassportMethodGetActionGin(g, PassportMethodGetAction)
				PassportMethodCreateActionGin(g, PassportMethodCreateAction)
				PassportMethodUpdateActionGin(g, PassportMethodUpdateAction)
				PassportMethodAwareDeletePreviewActionGin(g, PassportMethodAwareDeletePreviewAction)
				PassportMethodAwareDeleteActionGin(g, PassportMethodAwareDeleteAction)

				PublicJoinKeyBrowseActionGin(g, PublicJoinKeyBrowseAction)
				PublicJoinKeyGetActionGin(g, PublicJoinKeyGetAction)
				PublicJoinKeyCreateActionGin(g, PublicJoinKeyCreateAction)
				PublicJoinKeyUpdateActionGin(g, PublicJoinKeyUpdateAction)
				PublicJoinKeyAwareDeletePreviewActionGin(g, PublicJoinKeyAwareDeletePreviewAction)
				PublicJoinKeyAwareDeleteActionGin(g, PublicJoinKeyAwareDeleteAction)

				EmailConfirmationBrowseActionGin(g, EmailConfirmationBrowseAction)
				EmailConfirmationGetActionGin(g, EmailConfirmationGetAction)
				EmailConfirmationCreateActionGin(g, EmailConfirmationCreateAction)
				EmailConfirmationUpdateActionGin(g, EmailConfirmationUpdateAction)
				EmailConfirmationAwareDeletePreviewActionGin(g, EmailConfirmationAwareDeletePreviewAction)
				EmailConfirmationAwareDeleteActionGin(g, EmailConfirmationAwareDeleteAction)

				PhoneConfirmationBrowseActionGin(g, PhoneConfirmationBrowseAction)
				PhoneConfirmationGetActionGin(g, PhoneConfirmationGetAction)
				PhoneConfirmationCreateActionGin(g, PhoneConfirmationCreateAction)
				PhoneConfirmationUpdateActionGin(g, PhoneConfirmationUpdateAction)
				PhoneConfirmationAwareDeletePreviewActionGin(g, PhoneConfirmationAwareDeletePreviewAction)
				PhoneConfirmationAwareDeleteActionGin(g, PhoneConfirmationAwareDeleteAction)

				PublicAuthenticationBrowseActionGin(g, PublicAuthenticationBrowseAction)
				PublicAuthenticationGetActionGin(g, PublicAuthenticationGetAction)
				PublicAuthenticationCreateActionGin(g, PublicAuthenticationCreateAction)
				PublicAuthenticationUpdateActionGin(g, PublicAuthenticationUpdateAction)
				PublicAuthenticationAwareDeletePreviewActionGin(g, PublicAuthenticationAwareDeletePreviewAction)
				PublicAuthenticationAwareDeleteActionGin(g, PublicAuthenticationAwareDeleteAction)

				PassportBrowseActionGin(g, PassportBrowseAction)
				PassportGetActionGin(g, PassportGetAction)
				PassportCreateActionGin(g, PassportCreateAction)
				PassportUpdateActionGin(g, PassportUpdateAction)
				PassportAwareDeletePreviewActionGin(g, PassportAwareDeletePreviewAction)
				PassportAwareDeleteActionGin(g, PassportAwareDeleteAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_PASSPORT_PERMISSIONS,
		ALL_PASSPORT_METHOD_PERMISSIONS,
		ALL_PUBLIC_JOIN_KEY_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
		ALL_PUBLIC_AUTHENTICATION_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		return dbref.AutoMigrate(
			&EmailConfirmationEntity{},
			&PhoneConfirmationEntity{},
			&PublicAuthenticationEntity{},
			&PassportEntity{},
			&PassportMethodEntity{},
			&PublicJoinKeyEntity{},
		)
	})

	module.ProvideCliHandlers([]*cli.Command{
		&PassportCli,
	})

	return module
}
