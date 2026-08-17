//go:build !wasm

package fireback

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
		&cli.BoolFlag{
			Name:  "cookie-auth-only",
			Usage: "When true, the sessions (after authentication) would not return the token back in the response, and token will be only accessible via secure cookie.",
		},
		&cli.BoolFlag{
			Name:  "production",
			Usage: "If true, set's the environment behavior to production, and some functionality will be limited",
		},
		&cli.BoolFlag{
			Name:  "show-doctor",
			Usage: "If true, prints the doctor report (config, environment urls, database info) when the http server starts via 'start'. Defaults to false - run 'fireback doctor' directly when you need it instead.",
		},
		&cli.StringFlag{
			Name:  "table-prefix",
			Usage: "Prefix all gorm tables with some string",
		},
		&cli.StringFlag{
			Name:  "vapid-public-key",
			Usage: "VAPID Web push notification public key",
		},
		&cli.StringFlag{
			Name:  "vapid-private-key",
			Usage: "VAPID Web push notification private key",
		},
		&cli.StringFlag{
			Name:  "token-generation-strategy",
			Usage: "Fireback supports generating tokens based on random short string, or jwt.",
		},
		&cli.StringFlag{
			Name:  "jwt-secret-key",
			Usage: "If tokenGenerationStrategy is set to jwt, then these secret will be used.",
		},
		&cli.BoolFlag{
			Name:  "with-task-server",
			Usage: "Runs the tasks server asyncq library when the http server starts. Useful for all in one applications to run everything in single instance",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Environment name, such as dev, prod, test, test-eu, etc...",
		},
		&cli.StringFlag{
			Name:  "cert-file",
			Usage: "SSL Certification location to server on http listener",
		},
		&cli.StringFlag{
			Name:  "key-file",
			Usage: "SSL Certification key file",
		},
		&cli.StringFlag{
			Name:  "db-log-level",
			Usage: "Database log level for SQL queries, used by GORM orm. Default it's silent. 'warn', 'error', 'info' are other options.",
		},
		&cli.BoolFlag{
			Name:  "use-ssl",
			Usage: "If set to true, all http traffic will be redirected into https. Needs certFile and keyFile to be defined otherwise no effect",
		},
		&cli.StringFlag{
			Name:  "ssl-provider",
			Usage: "How the certificate used by useSSL is obtained: 'manual' (certFile/keyFile point at a certificate you already have), 'self-signed' (a locally generated certificate, persisted to certFile/keyFile so it survives restarts - not trusted by browsers), or 'letsencrypt' (an ACME certificate for acmeDomains, requested and renewed automatically via golang.org/x/crypto/acme/autocert while the server is running). Empty falls back to the legacy behavior: manual if certFile/keyFile are set, otherwise an ephemeral self-signed certificate regenerated on every start. Set by 'fireback ssl enable'.",
		},
		&cli.StringFlag{
			Name:  "acme-domains",
			Usage: "Comma separated list of domains to request a Let's Encrypt certificate for, when sslProvider is 'letsencrypt'. Each domain must resolve to this server and have port 80 reachable for the ACME HTTP-01 challenge.",
		},
		&cli.StringFlag{
			Name:  "acme-email",
			Usage: "Contact email registered with Let's Encrypt for expiry notices, when sslProvider is 'letsencrypt'. Optional but recommended.",
		},
		&cli.StringFlag{
			Name:  "acme-cache-dir",
			Usage: "Directory where the ACME account key and issued certificates are cached, when sslProvider is 'letsencrypt'. Must persist across restarts so certificates aren't re-requested (and to stay under Let's Encrypt's rate limits).",
		},
		&cli.StringFlag{
			Name:  "db-dsn",
			Usage: "The single source of truth for the database connection: a full DSN (postgres keyword/value string, mysql go-sql-driver DSN, or - for sqlite - the database file path itself, or in-memory). See modules/fireback/dbdsn for reading/writing individual pieces (host/port/username/password/database/ssl) of this string; there is no separate set of DB_HOST/DB_PORT/etc fields to keep in sync with it.",
		},
		&cli.StringFlag{
			Name:  "gin-mode",
			Usage: "Gin framework mode, which could be 'test', 'debug', 'release'",
		},
		&cli.StringFlag{
			Name:  "storage",
			Usage: "This is the storage url which files will be uploaded to",
		},
		&cli.StringFlag{
			Name:  "db-vendor",
			Usage: "Database vendor name, such as sqlite, mysql, or any other supported database.",
		},
		&cli.StringFlag{
			Name:  "std-out",
			Usage: "Writes the logs instead of std out into these log files.",
		},
		&cli.StringFlag{
			Name:  "worker-address",
			Usage: "This is the url (host and port) of a queue service. If not set, we use the internal queue system",
		},
		&cli.IntFlag{
			Name:  "worker-concurrency",
			Usage: "How many tasks worker can take concurrently",
		},
		&cli.StringFlag{
			Name:  "std-err",
			Usage: "Writes the errors instead of std err into these log files.",
		},
		&cli.StringFlag{
			Name:  "cli-token",
			Usage: "Authorization token for cli apps, to access resoruces similar on http api",
		},
		&cli.StringFlag{
			Name:  "cli-region",
			Usage: "Region, for example us or pl",
		},
		&cli.StringFlag{
			Name:  "cli-language",
			Usage: "Language of the cli operations, for example en or pl",
		},
		&cli.StringFlag{
			Name:  "cli-workspace",
			Usage: "Selected workspace in the cli context.",
		},
		&cli.Int64Flag{
			Name:  "port",
			Usage: "The port which application would be lifted",
		},
		&cli.StringFlag{
			Name:  "host",
			Usage: "Application host which http server will be lifted",
		},
		&cli.StringFlag{
			Name:  "mac-identifier",
			Usage: "Used name for installing app as system service on macos installers",
		},
		&cli.StringFlag{
			Name:  "debian-identifier",
			Usage: "Used name for installing app as system service on ubuntu installers",
		},
		&cli.StringFlag{
			Name:  "windows-identifier",
			Usage: "Used name for installing app as system service on windows installers",
		},
	}
}
func CastConfigFromCli(config *Config, c emigo.CliCastable) {
	if c.IsSet("cookie-auth-only") {
		config.CookieAuthOnly = c.Bool("cookie-auth-only")
	}
	if c.IsSet("production") {
		config.Production = c.Bool("production")
	}
	if c.IsSet("show-doctor") {
		config.ShowDoctor = c.Bool("show-doctor")
	}
	if c.IsSet("table-prefix") {
		config.TablePrefix = c.String("table-prefix")
	}
	if c.IsSet("vapid-public-key") {
		config.VapidPublicKey = c.String("vapid-public-key")
	}
	if c.IsSet("vapid-private-key") {
		config.VapidPrivateKey = c.String("vapid-private-key")
	}
	if c.IsSet("token-generation-strategy") {
		config.TokenGenerationStrategy = c.String("token-generation-strategy")
	}
	if c.IsSet("jwt-secret-key") {
		config.JwtSecretKey = c.String("jwt-secret-key")
	}
	if c.IsSet("with-task-server") {
		config.WithTaskServer = c.Bool("with-task-server")
	}
	if c.IsSet("name") {
		config.Name = c.String("name")
	}
	if c.IsSet("cert-file") {
		config.CertFile = c.String("cert-file")
	}
	if c.IsSet("key-file") {
		config.KeyFile = c.String("key-file")
	}
	if c.IsSet("db-log-level") {
		config.DbLogLevel = c.String("db-log-level")
	}
	if c.IsSet("use-ssl") {
		config.UseSSL = c.Bool("use-ssl")
	}
	if c.IsSet("ssl-provider") {
		config.SslProvider = c.String("ssl-provider")
	}
	if c.IsSet("acme-domains") {
		config.AcmeDomains = c.String("acme-domains")
	}
	if c.IsSet("acme-email") {
		config.AcmeEmail = c.String("acme-email")
	}
	if c.IsSet("acme-cache-dir") {
		config.AcmeCacheDir = c.String("acme-cache-dir")
	}
	if c.IsSet("db-dsn") {
		config.DbDsn = c.String("db-dsn")
	}
	if c.IsSet("gin-mode") {
		config.GinMode = c.String("gin-mode")
	}
	if c.IsSet("storage") {
		config.Storage = c.String("storage")
	}
	if c.IsSet("db-vendor") {
		config.DbVendor = c.String("db-vendor")
	}
	if c.IsSet("std-out") {
		config.StdOut = c.String("std-out")
	}
	if c.IsSet("worker-address") {
		config.WorkerAddress = c.String("worker-address")
	}
	if c.IsSet("worker-concurrency") {
		config.WorkerConcurrency = c.Int("worker-concurrency")
	}
	if c.IsSet("std-err") {
		config.StdErr = c.String("std-err")
	}
	if c.IsSet("cli-token") {
		config.CliToken = c.String("cli-token")
	}
	if c.IsSet("cli-region") {
		config.CliRegion = c.String("cli-region")
	}
	if c.IsSet("cli-language") {
		config.CliLanguage = c.String("cli-language")
	}
	if c.IsSet("cli-workspace") {
		config.CliWorkspace = c.String("cli-workspace")
	}
	if c.IsSet("port") {
		config.Port = c.Int64("port")
	}
	if c.IsSet("host") {
		config.Host = c.String("host")
	}
	if c.IsSet("mac-identifier") {
		config.MacIdentifier = c.String("mac-identifier")
	}
	if c.IsSet("debian-identifier") {
		config.DebianIdentifier = c.String("debian-identifier")
	}
	if c.IsSet("windows-identifier") {
		config.WindowsIdentifier = c.String("windows-identifier")
	}
}
func GetConfigCli() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "cookie-auth-only",
			Usage: "When true, the sessions (after authentication) would not return the token back in the response, and token will be only accessible via secure cookie. (bool)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CookieAuthOnly)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetBoolean(c, config.CookieAuthOnly, func(value bool) {
							config.CookieAuthOnly = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "production",
			Usage: "If true, set's the environment behavior to production, and some functionality will be limited (bool)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.Production)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetBoolean(c, config.Production, func(value bool) {
							config.Production = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "show-doctor",
			Usage: "If true, prints the doctor report (config, environment urls, database info) when the http server starts via 'start'. Defaults to false - run 'fireback doctor' directly when you need it instead. (bool)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.ShowDoctor)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetBoolean(c, config.ShowDoctor, func(value bool) {
							config.ShowDoctor = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "table-prefix",
			Usage: "Prefix all gorm tables with some string (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.TablePrefix)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.TablePrefix, func(value string) {
							config.TablePrefix = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "vapid-public-key",
			Usage: "VAPID Web push notification public key (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.VapidPublicKey)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.VapidPublicKey, func(value string) {
							config.VapidPublicKey = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "vapid-private-key",
			Usage: "VAPID Web push notification private key (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.VapidPrivateKey)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.VapidPrivateKey, func(value string) {
							config.VapidPrivateKey = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "token-generation-strategy",
			Usage: "Fireback supports generating tokens based on random short string, or jwt. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.TokenGenerationStrategy)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.TokenGenerationStrategy, func(value string) {
							config.TokenGenerationStrategy = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "jwt-secret-key",
			Usage: "If tokenGenerationStrategy is set to jwt, then these secret will be used. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.JwtSecretKey)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.JwtSecretKey, func(value string) {
							config.JwtSecretKey = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "with-task-server",
			Usage: "Runs the tasks server asyncq library when the http server starts. Useful for all in one applications to run everything in single instance (bool)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.WithTaskServer)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetBoolean(c, config.WithTaskServer, func(value bool) {
							config.WithTaskServer = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "name",
			Usage: "Environment name, such as dev, prod, test, test-eu, etc... (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.Name)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.Name, func(value string) {
							config.Name = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "cert-file",
			Usage: "SSL Certification location to server on http listener (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CertFile)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CertFile, func(value string) {
							config.CertFile = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "key-file",
			Usage: "SSL Certification key file (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.KeyFile)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.KeyFile, func(value string) {
							config.KeyFile = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "db-log-level",
			Usage: "Database log level for SQL queries, used by GORM orm. Default it's silent. 'warn', 'error', 'info' are other options. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DbLogLevel)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DbLogLevel, func(value string) {
							config.DbLogLevel = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "use-ssl",
			Usage: "If set to true, all http traffic will be redirected into https. Needs certFile and keyFile to be defined otherwise no effect (bool)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.UseSSL)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetBoolean(c, config.UseSSL, func(value bool) {
							config.UseSSL = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "ssl-provider",
			Usage: "How the certificate used by useSSL is obtained: 'manual' (certFile/keyFile point at a certificate you already have), 'self-signed' (a locally generated certificate, persisted to certFile/keyFile so it survives restarts - not trusted by browsers), or 'letsencrypt' (an ACME certificate for acmeDomains, requested and renewed automatically via golang.org/x/crypto/acme/autocert while the server is running). Empty falls back to the legacy behavior: manual if certFile/keyFile are set, otherwise an ephemeral self-signed certificate regenerated on every start. Set by 'fireback ssl enable'. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.SslProvider)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.SslProvider, func(value string) {
							config.SslProvider = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "acme-domains",
			Usage: "Comma separated list of domains to request a Let's Encrypt certificate for, when sslProvider is 'letsencrypt'. Each domain must resolve to this server and have port 80 reachable for the ACME HTTP-01 challenge. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.AcmeDomains)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.AcmeDomains, func(value string) {
							config.AcmeDomains = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "acme-email",
			Usage: "Contact email registered with Let's Encrypt for expiry notices, when sslProvider is 'letsencrypt'. Optional but recommended. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.AcmeEmail)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.AcmeEmail, func(value string) {
							config.AcmeEmail = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "acme-cache-dir",
			Usage: "Directory where the ACME account key and issued certificates are cached, when sslProvider is 'letsencrypt'. Must persist across restarts so certificates aren't re-requested (and to stay under Let's Encrypt's rate limits). (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.AcmeCacheDir)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.AcmeCacheDir, func(value string) {
							config.AcmeCacheDir = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "db-dsn",
			Usage: "The single source of truth for the database connection: a full DSN (postgres keyword/value string, mysql go-sql-driver DSN, or - for sqlite - the database file path itself, or in-memory). See modules/fireback/dbdsn for reading/writing individual pieces (host/port/username/password/database/ssl) of this string; there is no separate set of DB_HOST/DB_PORT/etc fields to keep in sync with it. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DbDsn)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DbDsn, func(value string) {
							config.DbDsn = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "gin-mode",
			Usage: "Gin framework mode, which could be 'test', 'debug', 'release' (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.GinMode)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.GinMode, func(value string) {
							config.GinMode = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "storage",
			Usage: "This is the storage url which files will be uploaded to (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.Storage)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.Storage, func(value string) {
							config.Storage = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "db-vendor",
			Usage: "Database vendor name, such as sqlite, mysql, or any other supported database. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DbVendor)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DbVendor, func(value string) {
							config.DbVendor = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "std-out",
			Usage: "Writes the logs instead of std out into these log files. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.StdOut)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.StdOut, func(value string) {
							config.StdOut = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "worker-address",
			Usage: "This is the url (host and port) of a queue service. If not set, we use the internal queue system (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.WorkerAddress)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.WorkerAddress, func(value string) {
							config.WorkerAddress = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "worker-concurrency",
			Usage: "How many tasks worker can take concurrently (int)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.WorkerConcurrency)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetInt(c, config.WorkerConcurrency, func(value int) {
							config.WorkerConcurrency = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "std-err",
			Usage: "Writes the errors instead of std err into these log files. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.StdErr)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.StdErr, func(value string) {
							config.StdErr = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "cli-token",
			Usage: "Authorization token for cli apps, to access resoruces similar on http api (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CliToken)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CliToken, func(value string) {
							config.CliToken = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "cli-region",
			Usage: "Region, for example us or pl (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CliRegion)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CliRegion, func(value string) {
							config.CliRegion = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "cli-language",
			Usage: "Language of the cli operations, for example en or pl (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CliLanguage)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CliLanguage, func(value string) {
							config.CliLanguage = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "cli-workspace",
			Usage: "Selected workspace in the cli context. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CliWorkspace)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CliWorkspace, func(value string) {
							config.CliWorkspace = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "port",
			Usage: "The port which application would be lifted (int64)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.Port)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetInt64(c, config.Port, func(value int64) {
							config.Port = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "host",
			Usage: "Application host which http server will be lifted (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.Host)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.Host, func(value string) {
							config.Host = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "mac-identifier",
			Usage: "Used name for installing app as system service on macos installers (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.MacIdentifier)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.MacIdentifier, func(value string) {
							config.MacIdentifier = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "debian-identifier",
			Usage: "Used name for installing app as system service on ubuntu installers (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DebianIdentifier)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DebianIdentifier, func(value string) {
							config.DebianIdentifier = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "windows-identifier",
			Usage: "Used name for installing app as system service on windows installers (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.WindowsIdentifier)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.WindowsIdentifier, func(value string) {
							config.WindowsIdentifier = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
	}
}
