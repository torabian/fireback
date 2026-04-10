package fireback

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/migrations"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

var EverRunEntities []interface{} = []interface{}{
	&CapabilityEntity{},
	&CapabilityEntityPolyglot{},
}

func workspaceModuleCore(module *ModuleProvider) {

	module.ProvidePermissionHandler(

		ALL_CAPABILITY_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		items2 := []interface{}{}
		items2 = append(items2, EverRunEntities...)

		for _, item := range items2 {

			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", GetInterfaceName(item))
				return err
			}
		}

		return nil
	})

}

type FirebackModuleConfig struct{}

type X = func(query QueryDSL, done chan bool, read chan SocketReadChan) (chan []byte, error)

func FirebackModuleSetup(setup *FirebackModuleConfig) *ModuleProvider {

	module := &ModuleProvider{
		Name:        "fireback",
		Definitions: &Module3Definitions,
		Actions: [][]Module3Action{
			GetCapabilityModule3Actions(),
			FirebackCustomActions(),
		},
		EntityBundles: []EntityBundle{
			WebPushConfigEntityBundle,
		},

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *FirebackApp) error{
			func(g *gin.RouterGroup, x *FirebackApp) error {

				meta := EventBusSubscriptionActionMeta()
				g.GET(
					meta.URL,
					WithSocketAuthorization(EventBusSubscriptionSecurityModel),
					EventBusSubscriptionActionReactiveHandler(EventBusSubscriptionActionSig),
				)

				return nil
			},
		},
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	module.ProvideCliHandlers([]cli.Command{
		CapabilityCliFn(),
		PushNotificationCmd,
		CapabilitiesTreeActionDef.ToCli(),
	})

	workspaceModuleCore(module)

	return module
}
