package abac

import (
	"github.com/gin-gonic/gin"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

var GinWebServerInitHooks = []func(g *gin.RouterGroup, x *application.Application) error{
	func(g *gin.RouterGroup, x *application.Application) error {
		abacdefs.PassportMethodBrowseActionGin(g, PassportMethodBrowseAction)
		abacdefs.PassportMethodGetActionGin(g, PassportMethodGetAction)
		abacdefs.PassportMethodCreateActionGin(g, PassportMethodCreateAction)
		abacdefs.PassportMethodUpdateActionGin(g, PassportMethodUpdateAction)
		abacdefs.PassportMethodAwareDeletePreviewActionGin(g, PassportMethodAwareDeletePreviewAction)
		abacdefs.PassportMethodAwareDeleteActionGin(g, PassportMethodAwareDeleteAction)

		abacdefs.PublicJoinKeyBrowseActionGin(g, PublicJoinKeyBrowseAction)
		abacdefs.PublicJoinKeyGetActionGin(g, PublicJoinKeyGetAction)
		abacdefs.PublicJoinKeyCreateActionGin(g, PublicJoinKeyCreateAction)
		abacdefs.PublicJoinKeyUpdateActionGin(g, PublicJoinKeyUpdateAction)
		abacdefs.PublicJoinKeyAwareDeletePreviewActionGin(g, PublicJoinKeyAwareDeletePreviewAction)
		abacdefs.PublicJoinKeyAwareDeleteActionGin(g, PublicJoinKeyAwareDeleteAction)

		abacdefs.EmailConfirmationBrowseActionGin(g, EmailConfirmationBrowseAction)
		abacdefs.EmailConfirmationGetActionGin(g, EmailConfirmationGetAction)
		abacdefs.EmailConfirmationCreateActionGin(g, EmailConfirmationCreateAction)
		abacdefs.EmailConfirmationUpdateActionGin(g, EmailConfirmationUpdateAction)
		abacdefs.EmailConfirmationAwareDeletePreviewActionGin(g, EmailConfirmationAwareDeletePreviewAction)
		abacdefs.EmailConfirmationAwareDeleteActionGin(g, EmailConfirmationAwareDeleteAction)

		abacdefs.PhoneConfirmationBrowseActionGin(g, PhoneConfirmationBrowseAction)
		abacdefs.PhoneConfirmationGetActionGin(g, PhoneConfirmationGetAction)
		abacdefs.PhoneConfirmationCreateActionGin(g, PhoneConfirmationCreateAction)
		abacdefs.PhoneConfirmationUpdateActionGin(g, PhoneConfirmationUpdateAction)
		abacdefs.PhoneConfirmationAwareDeletePreviewActionGin(g, PhoneConfirmationAwareDeletePreviewAction)
		abacdefs.PhoneConfirmationAwareDeleteActionGin(g, PhoneConfirmationAwareDeleteAction)

		abacdefs.PublicAuthenticationBrowseActionGin(g, PublicAuthenticationBrowseAction)
		abacdefs.PublicAuthenticationGetActionGin(g, PublicAuthenticationGetAction)
		abacdefs.PublicAuthenticationCreateActionGin(g, PublicAuthenticationCreateAction)
		abacdefs.PublicAuthenticationUpdateActionGin(g, PublicAuthenticationUpdateAction)
		abacdefs.PublicAuthenticationAwareDeletePreviewActionGin(g, PublicAuthenticationAwareDeletePreviewAction)
		abacdefs.PublicAuthenticationAwareDeleteActionGin(g, PublicAuthenticationAwareDeleteAction)

		abacdefs.PassportBrowseActionGin(g, PassportBrowseAction)
		abacdefs.PassportGetActionGin(g, PassportGetAction)
		abacdefs.PassportCreateActionGin(g, PassportCreateAction)
		abacdefs.PassportUpdateActionGin(g, PassportUpdateAction)
		abacdefs.PassportAwareDeletePreviewActionGin(g, PassportAwareDeletePreviewAction)
		abacdefs.PassportAwareDeleteActionGin(g, PassportAwareDeleteAction)

		return nil
	},
}

var PassportProvidePermissionHandler = [][]application.PermissionInfo{
	ALL_PASSPORT_PERMISSIONS,
	ALL_PASSPORT_METHOD_PERMISSIONS,
	ALL_PUBLIC_JOIN_KEY_PERMISSIONS,
	ALL_ROLE_PERMISSIONS,
	ALL_USER_PERMISSIONS,
	ALL_PUBLIC_AUTHENTICATION_PERMISSIONS,
}

var PassportProvideEntityHandlers = func(dbref *gorm.DB) error {

	return dbref.AutoMigrate(
		&abacdefs.EmailConfirmationEntity{},
		&abacdefs.PhoneConfirmationEntity{},
		&abacdefs.PublicAuthenticationEntity{},
		&abacdefs.PassportEntity{},
		&abacdefs.PassportMethodEntity{},
		&abacdefs.PublicJoinKeyEntity{},
	)
}
