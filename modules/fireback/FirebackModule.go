package fireback

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
				// The actual callback is extracted, in case you need to handle multiple handlers or customize, use it directly.

				g.GET(meta.URL,
					WithSocketAuthorization(EventBusSubscriptionSecurityModel),
					func(ctx *gin.Context) {
						EventBusSubscriptionActionDuplexGinHandler(ctx, func(ctx *EventBusSubscriptionActionSession) {
							done := make(chan bool)
							var query QueryDSL
							query = ExtractQueryDslFromGinContext(ctx.G)

							query.RawSocketConnection = ctx.Socket

							// Adapt incoming messages
							read := make(chan SocketReadChan)

							go func() {
								defer close(read)
								for msg := range ctx.In {
									read <- SocketReadChan{
										Data:  msg.Raw,
										Error: msg.Error,
									}
								}
							}()

							// Call your existing function
							out, _ := EventBusSubscriptionActionSig(query, done, read)

							// Forward outgoing messages
							for {
								select {
								case data, ok := <-out:
									if !ok {
										return
									}

									ctx.Out <- EventBusSubscriptionActionMessage{
										MessageType: websocket.TextMessage,
										Raw:         data,
									}

								case <-ctx.Done:
									close(done)
									return
								}
							}
						})
					})

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
