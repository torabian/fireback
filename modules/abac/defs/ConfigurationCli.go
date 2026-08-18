//go:build !wasm

package abacdefs

import (
	"context"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

/*
*
You can call this function on first line of your main function.
This is different from fireback configuration (for now), you can
define config: in module3 file, similar to fields in entities,
and we generate the config struct and this function would read .env.local,
.env.prod, etc - depending on the ENV=xxx env variable.
*
*/
func LoadConfiguration() Config {
	emigo.HandleEnvVars(&config)
	return config
}
func (x *Config) Save(filepath string) error {
	return emigo.SaveEnvFile(x, filepath)
}
func GetConfigCliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.Int64Flag{
			Name:  "otp-lockout-seconds",
			Usage: "How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go.",
		},
		&cli.StringFlag{
			Name:  "self-service-base-url",
			Usage: "Base URL of the running self-service frontend (ui/packages/selfservice), used to build the recovery link sent by ClassicPassportRequestOtpActionImplementation.go and SendPassportResetEmailActionImplementation.go - the value gets /en/selfservice/reset-password?value=... appended directly, so include a trailing /# fragment here too if that frontend build uses a HashRouter (BUILD_VARIABLES.USE_HASH_ROUTER, the default for ui/src/apps/self-service). The default here matches this same binary's own embedded self-service build (see cmd/fireback/main.go's PublicFolders), mounted at /selfservice on this same origin - point this at wherever the self-service frontend is actually deployed in production, adjusting or omitting the /selfservice path segment and the /# fragment (HashRouter vs plain BrowserRouter) to match that deployment.",
		},
	}
}
func CastConfigFromCli(config *Config, c emigo.CliCastable) {
	if c.IsSet("otp-lockout-seconds") {
		config.OtpLockoutSeconds = c.Int64("otp-lockout-seconds")
	}
	if c.IsSet("self-service-base-url") {
		config.SelfServiceBaseUrl = c.String("self-service-base-url")
	}
}
func GetConfigCli() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "otp-lockout-seconds",
			Usage: "How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go. (int64)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.OtpLockoutSeconds)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetInt64(c, config.OtpLockoutSeconds, func(value int64) {
							config.OtpLockoutSeconds = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "self-service-base-url",
			Usage: "Base URL of the running self-service frontend (ui/packages/selfservice), used to build the recovery link sent by ClassicPassportRequestOtpActionImplementation.go and SendPassportResetEmailActionImplementation.go - the value gets /en/selfservice/reset-password?value=... appended directly, so include a trailing /# fragment here too if that frontend build uses a HashRouter (BUILD_VARIABLES.USE_HASH_ROUTER, the default for ui/src/apps/self-service). The default here matches this same binary's own embedded self-service build (see cmd/fireback/main.go's PublicFolders), mounted at /selfservice on this same origin - point this at wherever the self-service frontend is actually deployed in production, adjusting or omitting the /selfservice path segment and the /# fragment (HashRouter vs plain BrowserRouter) to match that deployment. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.SelfServiceBaseUrl)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.SelfServiceBaseUrl, func(value string) {
							config.SelfServiceBaseUrl = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
	}
}
