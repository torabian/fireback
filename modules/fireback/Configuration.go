package fireback

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

/**
* Configuration generator
 */
type Config struct {
	// When true, the sessions (after authentication) would not return the token back in the response, and token will be only accessible via secure cookie.
	CookieAuthOnly bool `envconfig:"COOKIE_AUTH_ONLY" description:"When true, the sessions (after authentication) would not return the token back in the response, and token will be only accessible via secure cookie."`
	// If true, set's the environment behavior to production, and some functionality will be limited
	Production bool `envconfig:"PRODUCTION" description:"If true, set's the environment behavior to production, and some functionality will be limited"`
	// If true, prints the doctor report (config, environment urls, database info) when the http server starts via 'start'. Defaults to false - run 'fireback doctor' directly when you need it instead.
	ShowDoctor bool `envconfig:"SHOW_DOCTOR" description:"If true, prints the doctor report (config, environment urls, database info) when the http server starts via 'start'. Defaults to false - run 'fireback doctor' directly when you need it instead."`
	// Prefix all gorm tables with some string
	TablePrefix string `envconfig:"TABLE_PREFIX" description:"Prefix all gorm tables with some string"`
	// VAPID Web push notification public key
	VapidPublicKey string `envconfig:"VAPID_PUBLIC_KEY" description:"VAPID Web push notification public key"`
	// VAPID Web push notification private key
	VapidPrivateKey string `envconfig:"VAPID_PRIVATE_KEY" description:"VAPID Web push notification private key"`
	// Fireback supports generating tokens based on random short string, or jwt.
	TokenGenerationStrategy string `envconfig:"TOKEN_GENERATION_STRATEGY" description:"Fireback supports generating tokens based on random short string, or jwt."`
	// If tokenGenerationStrategy is set to jwt, then these secret will be used.
	JwtSecretKey string `envconfig:"JWT_SECRET_KEY" description:"If tokenGenerationStrategy is set to jwt, then these secret will be used."`
	// Runs the tasks server asyncq library when the http server starts. Useful for all in one applications to run everything in single instance
	WithTaskServer bool `envconfig:"WITH_TASK_SERVER" description:"Runs the tasks server asyncq library when the http server starts. Useful for all in one applications to run everything in single instance"`
	// Environment name, such as dev, prod, test, test-eu, etc...
	Name string `envconfig:"NAME" description:"Environment name, such as dev, prod, test, test-eu, etc..."`
	// SSL Certification location to server on http listener
	CertFile string `envconfig:"CERT_FILE" description:"SSL Certification location to server on http listener"`
	// SSL Certification key file
	KeyFile string `envconfig:"KEY_FILE" description:"SSL Certification key file"`
	// Database log level for SQL queries, used by GORM orm. Default it's silent. 'warn', 'error', 'info' are other options.
	DbLogLevel string `envconfig:"DB_LOG_LEVEL" description:"Database log level for SQL queries, used by GORM orm. Default it's silent. 'warn', 'error', 'info' are other options."`
	// If set to true, all http traffic will be redirected into https. Needs certFile and keyFile to be defined otherwise no effect
	UseSSL bool `envconfig:"USE_SSL" description:"If set to true, all http traffic will be redirected into https. Needs certFile and keyFile to be defined otherwise no effect"`
	// How the certificate used by useSSL is obtained: 'manual' (certFile/keyFile point at a certificate you already have), 'self-signed' (a locally generated certificate, persisted to certFile/keyFile so it survives restarts - not trusted by browsers), or 'letsencrypt' (an ACME certificate for acmeDomains, requested and renewed automatically via golang.org/x/crypto/acme/autocert while the server is running). Empty falls back to the legacy behavior: manual if certFile/keyFile are set, otherwise an ephemeral self-signed certificate regenerated on every start. Set by 'fireback ssl enable'.
	SslProvider string `envconfig:"SSL_PROVIDER" description:"How the certificate used by useSSL is obtained: 'manual' (certFile/keyFile point at a certificate you already have), 'self-signed' (a locally generated certificate, persisted to certFile/keyFile so it survives restarts - not trusted by browsers), or 'letsencrypt' (an ACME certificate for acmeDomains, requested and renewed automatically via golang.org/x/crypto/acme/autocert while the server is running). Empty falls back to the legacy behavior: manual if certFile/keyFile are set, otherwise an ephemeral self-signed certificate regenerated on every start. Set by 'fireback ssl enable'."`
	// Comma separated list of domains to request a Let's Encrypt certificate for, when sslProvider is 'letsencrypt'. Each domain must resolve to this server and have port 80 reachable for the ACME HTTP-01 challenge.
	AcmeDomains string `envconfig:"ACME_DOMAINS" description:"Comma separated list of domains to request a Let's Encrypt certificate for, when sslProvider is 'letsencrypt'. Each domain must resolve to this server and have port 80 reachable for the ACME HTTP-01 challenge."`
	// Contact email registered with Let's Encrypt for expiry notices, when sslProvider is 'letsencrypt'. Optional but recommended.
	AcmeEmail string `envconfig:"ACME_EMAIL" description:"Contact email registered with Let's Encrypt for expiry notices, when sslProvider is 'letsencrypt'. Optional but recommended."`
	// Directory where the ACME account key and issued certificates are cached, when sslProvider is 'letsencrypt'. Must persist across restarts so certificates aren't re-requested (and to stay under Let's Encrypt's rate limits).
	AcmeCacheDir string `envconfig:"ACME_CACHE_DIR" description:"Directory where the ACME account key and issued certificates are cached, when sslProvider is 'letsencrypt'. Must persist across restarts so certificates aren't re-requested (and to stay under Let's Encrypt's rate limits)."`
	// The single source of truth for the database connection: a full DSN (postgres keyword/value string, mysql go-sql-driver DSN, or - for sqlite - the database file path itself, or in-memory). See modules/fireback/dbdsn for reading/writing individual pieces (host/port/username/password/database/ssl) of this string; there is no separate set of DB_HOST/DB_PORT/etc fields to keep in sync with it.
	DbDsn string `envconfig:"DB_DSN" description:"The single source of truth for the database connection: a full DSN (postgres keyword/value string, mysql go-sql-driver DSN, or - for sqlite - the database file path itself, or in-memory). See modules/fireback/dbdsn for reading/writing individual pieces (host/port/username/password/database/ssl) of this string; there is no separate set of DB_HOST/DB_PORT/etc fields to keep in sync with it."`
	// Gin framework mode, which could be 'test', 'debug', 'release'
	GinMode string `envconfig:"GIN_MODE" description:"Gin framework mode, which could be 'test', 'debug', 'release'"`
	// This is the storage url which files will be uploaded to
	Storage string `envconfig:"STORAGE" description:"This is the storage url which files will be uploaded to"`
	// Database vendor name, such as sqlite, mysql, or any other supported database.
	DbVendor string `envconfig:"DB_VENDOR" description:"Database vendor name, such as sqlite, mysql, or any other supported database."`
	// Writes the logs instead of std out into these log files.
	StdOut string `envconfig:"STD_OUT" description:"Writes the logs instead of std out into these log files."`
	// This is the url (host and port) of a queue service. If not set, we use the internal queue system
	WorkerAddress string `envconfig:"WORKER_ADDRESS" description:"This is the url (host and port) of a queue service. If not set, we use the internal queue system"`
	// How many tasks worker can take concurrently
	WorkerConcurrency int `envconfig:"WORKER_CONCURRENCY" description:"How many tasks worker can take concurrently"`
	// Writes the errors instead of std err into these log files.
	StdErr string `envconfig:"STD_ERR" description:"Writes the errors instead of std err into these log files."`
	// Authorization token for cli apps, to access resoruces similar on http api
	CliToken string `envconfig:"CLI_TOKEN" description:"Authorization token for cli apps, to access resoruces similar on http api"`
	// Region, for example us or pl
	CliRegion string `envconfig:"CLI_REGION" description:"Region, for example us or pl"`
	// Language of the cli operations, for example en or pl
	CliLanguage string `envconfig:"CLI_LANGUAGE" description:"Language of the cli operations, for example en or pl"`
	// Selected workspace in the cli context.
	CliWorkspace string `envconfig:"CLI_WORKSPACE" description:"Selected workspace in the cli context."`
	// The port which application would be lifted
	Port int64 `envconfig:"PORT" description:"The port which application would be lifted"`
	// Application host which http server will be lifted
	Host string `envconfig:"HOST" description:"Application host which http server will be lifted"`
	// Used name for installing app as system service on macos installers
	MacIdentifier string `envconfig:"MAC_IDENTIFIER" description:"Used name for installing app as system service on macos installers"`
	// Used name for installing app as system service on ubuntu installers
	DebianIdentifier string `envconfig:"DEBIAN_IDENTIFIER" description:"Used name for installing app as system service on ubuntu installers"`
	// Used name for installing app as system service on windows installers
	WindowsIdentifier string `envconfig:"WINDOWS_IDENTIFIER" description:"Used name for installing app as system service on windows installers"`
}

// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{
	TokenGenerationStrategy: "random",
	WithTaskServer:          false,
	DbLogLevel:              "silent",
	AcmeCacheDir:            "./.fireback-acme-cache",
	DbVendor:                "sqlite",
	WorkerAddress:           "127.0.0.1:6379",
	WorkerConcurrency:       10,
	CliRegion:               "us",
	CliLanguage:             "en",
	Port:                    4500,
	Host:                    "localhost",
	MacIdentifier:           "fireback",
	DebianIdentifier:        "fireback",
	WindowsIdentifier:       "fireback",
}

func (x *Config) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

/*
*
You can call this function on first line of your main function.
This is different from fireback configuration (for now), you can
define config: in module3 file, similar to fields in entities,
and we generate the config struct and this function would read .env.local,
.env.prod, etc - depending on the ENV=xxx env variable (or, under a wasm
build, whatever env vars the host page already set via os.Setenv before
this ran - see emigo.HandleEnvVars's own doc comments in Config.go/
ConfigWasm.go for the two implementations this dispatches to).
*
*/
func LoadConfiguration() Config {
	emigo.HandleEnvVars(&config)
	return config
}
