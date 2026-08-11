package abacdefs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

/**
* Configuration generator
 */
type Config struct {
	// How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go.
	OtpLockoutSeconds int64 `envconfig:"OTP_LOCKOUT_SECONDS" description:"How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go."`
}

// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{
	OtpLockoutSeconds: 120,
}

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
func (x *Config) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
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
	}
}
func CastConfigFromCli(config *Config, c emigo.CliCastable) {
	if c.IsSet("otp-lockout-seconds") {
		config.OtpLockoutSeconds = c.Int64("otp-lockout-seconds")
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
	}
}
