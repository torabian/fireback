package abacdefs

import "encoding/json"

/**
* Configuration generator
 */
type Config struct {
	// How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go.
	OtpLockoutSeconds int64 `envconfig:"OTP_LOCKOUT_SECONDS" description:"How many seconds a passport value (email or phone) must wait before it can request another OTP/magic-link, once a request is already pending. See ClassicPassportRequestOtpActionImplementation.go."`
	// Base URL of the running self-service frontend (ui/packages/selfservice), used to build the recovery link sent by ClassicPassportRequestOtpActionImplementation.go and SendPassportResetEmailActionImplementation.go - the value gets /en/selfservice/reset-password?value=... appended directly, so include a trailing /# fragment here too if that frontend build uses a HashRouter (BUILD_VARIABLES.USE_HASH_ROUTER, the default for ui/src/apps/self-service). The default here matches this same binary's own embedded self-service build (see cmd/fireback/main.go's PublicFolders), mounted at /selfservice on this same origin - point this at wherever the self-service frontend is actually deployed in production, adjusting or omitting the /selfservice path segment and the /# fragment (HashRouter vs plain BrowserRouter) to match that deployment.
	SelfServiceBaseUrl string `envconfig:"SELF_SERVICE_BASE_URL" description:"Base URL of the running self-service frontend (ui/packages/selfservice), used to build the recovery link sent by ClassicPassportRequestOtpActionImplementation.go and SendPassportResetEmailActionImplementation.go - the value gets /en/selfservice/reset-password?value=... appended directly, so include a trailing /# fragment here too if that frontend build uses a HashRouter (BUILD_VARIABLES.USE_HASH_ROUTER, the default for ui/src/apps/self-service). The default here matches this same binary's own embedded self-service build (see cmd/fireback/main.go's PublicFolders), mounted at /selfservice on this same origin - point this at wherever the self-service frontend is actually deployed in production, adjusting or omitting the /selfservice path segment and the /# fragment (HashRouter vs plain BrowserRouter) to match that deployment."`
}

// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{
	OtpLockoutSeconds:  120,
	SelfServiceBaseUrl: "http://localhost:8888/selfservice/#",
}

func (x *Config) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
