package fireback

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/migrations"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

var EverRunEntities []interface{} = []interface{}{}

type FirebackModuleConfig struct{}

// type X = func(query QueryDSL, done chan bool, read chan socket.SocketReadChan) (chan []byte, error)

func FirebackModuleSetup(setup *FirebackModuleConfig) *ModuleProvider {

	module := &ModuleProvider{
		Name:        "fireback",
		Definitions: &Module3Definitions,

		EntityBundles: []EntityBundle{},

		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *FirebackApp) error{
			func(g *gin.RouterGroup, x *FirebackApp) error {

				{
					meta := EventBusSubscriptionActionMeta()
					g.GET(meta.URL, EventBusSubscriptionActionReactiveHandler(EventBusSubscriptionActionSig))

				}

				{
					meta := ReactiveSearchActionMeta()
					g.GET(
						meta.URL,
						ReactiveSearchActionReactiveHandler(CreateReactiveSearchHanlder(x)),
					)
				}

				return nil
			},
		},
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	module.ProvideCliHandlers([]*cli.Command{
		&PushNotificationCmd,
	})

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

	return module
}
