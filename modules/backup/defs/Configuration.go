package backupdefs

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
	// Required - LoadModuleConfig refuses to start without it. Local/mounted directory wal-g stores base backups and WAL segments under. Must be genuinely separate storage from the database's own disk to serve as a real disaster-recovery plan - see README.md.
	FilePrefix string `envconfig:"WALG_FILE_PREFIX" description:"Required - LoadModuleConfig refuses to start without it. Local/mounted directory wal-g stores base backups and WAL segments under. Must be genuinely separate storage from the database's own disk to serve as a real disaster-recovery plan - see README.md."`
	// Compression method wal-g uses for base backups and WAL segments.
	CompressionMethod string `envconfig:"WALG_COMPRESSION_METHOD" description:"Compression method wal-g uses for base backups and WAL segments."`
	// How many delta backups wal-g chains before forcing a full backup.
	DeltaMaxSteps string `envconfig:"WALG_DELTA_MAX_STEPS" description:"How many delta backups wal-g chains before forcing a full backup."`
	// How many full backups the prune command keeps (plus what's needed to restore from the oldest of them forward).
	RetainFullBackups int64 `envconfig:"BACKUP_RETAIN_FULL" description:"How many full backups the prune command keeps (plus what's needed to restore from the oldest of them forward)."`
	// Directory the dump command writes its output zip into when neither --output nor --hash is given.
	DumpDir string `envconfig:"BACKUP_DUMP_DIR" description:"Directory the dump command writes its output zip into when neither --output nor --hash is given."`
	// How many seconds a dump --hash download link stays claimable before it expires unfetched. The hash is single-use - once fetched (or once it fails), it's disabled immediately regardless of this TTL.
	DumpHashTtlSeconds int64 `envconfig:"BACKUP_DUMP_HASH_TTL_SECONDS" description:"How many seconds a dump --hash download link stays claimable before it expires unfetched. The hash is single-use - once fetched (or once it fails), it's disabled immediately regardless of this TTL."`
	// pg_dump executable used by the dump command for postgres.
	PgDumpBin string `envconfig:"PG_DUMP_BIN" description:"pg_dump executable used by the dump command for postgres."`
	// psql executable used by the restore-dump command for postgres.
	PsqlBin string `envconfig:"PSQL_BIN" description:"psql executable used by the restore-dump command for postgres."`
	// mysqldump executable used by the dump command for mysql/mariadb.
	MysqldumpBin string `envconfig:"MYSQLDUMP_BIN" description:"mysqldump executable used by the dump command for mysql/mariadb."`
	// mysql client executable used by the restore-dump command for mysql/mariadb.
	MysqlBin string `envconfig:"MYSQL_BIN" description:"mysql client executable used by the restore-dump command for mysql/mariadb."`
}

// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{
	CompressionMethod:  "lz4",
	DeltaMaxSteps:      "6",
	RetainFullBackups:  4,
	DumpDir:            "./backups/dumps",
	DumpHashTtlSeconds: 1800,
	PgDumpBin:          "pg_dump",
	PsqlBin:            "psql",
	MysqldumpBin:       "mysqldump",
	MysqlBin:           "mysql",
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
		&cli.StringFlag{
			Name:  "file-prefix",
			Usage: "Required - LoadModuleConfig refuses to start without it. Local/mounted directory wal-g stores base backups and WAL segments under. Must be genuinely separate storage from the database's own disk to serve as a real disaster-recovery plan - see README.md.",
		},
		&cli.StringFlag{
			Name:  "compression-method",
			Usage: "Compression method wal-g uses for base backups and WAL segments.",
		},
		&cli.StringFlag{
			Name:  "delta-max-steps",
			Usage: "How many delta backups wal-g chains before forcing a full backup.",
		},
		&cli.Int64Flag{
			Name:  "retain-full-backups",
			Usage: "How many full backups the prune command keeps (plus what's needed to restore from the oldest of them forward).",
		},
		&cli.StringFlag{
			Name:  "dump-dir",
			Usage: "Directory the dump command writes its output zip into when neither --output nor --hash is given.",
		},
		&cli.Int64Flag{
			Name:  "dump-hash-ttl-seconds",
			Usage: "How many seconds a dump --hash download link stays claimable before it expires unfetched. The hash is single-use - once fetched (or once it fails), it's disabled immediately regardless of this TTL.",
		},
		&cli.StringFlag{
			Name:  "pg-dump-bin",
			Usage: "pg_dump executable used by the dump command for postgres.",
		},
		&cli.StringFlag{
			Name:  "psql-bin",
			Usage: "psql executable used by the restore-dump command for postgres.",
		},
		&cli.StringFlag{
			Name:  "mysqldump-bin",
			Usage: "mysqldump executable used by the dump command for mysql/mariadb.",
		},
		&cli.StringFlag{
			Name:  "mysql-bin",
			Usage: "mysql client executable used by the restore-dump command for mysql/mariadb.",
		},
	}
}
func CastConfigFromCli(config *Config, c emigo.CliCastable) {
	if c.IsSet("file-prefix") {
		config.FilePrefix = c.String("file-prefix")
	}
	if c.IsSet("compression-method") {
		config.CompressionMethod = c.String("compression-method")
	}
	if c.IsSet("delta-max-steps") {
		config.DeltaMaxSteps = c.String("delta-max-steps")
	}
	if c.IsSet("retain-full-backups") {
		config.RetainFullBackups = c.Int64("retain-full-backups")
	}
	if c.IsSet("dump-dir") {
		config.DumpDir = c.String("dump-dir")
	}
	if c.IsSet("dump-hash-ttl-seconds") {
		config.DumpHashTtlSeconds = c.Int64("dump-hash-ttl-seconds")
	}
	if c.IsSet("pg-dump-bin") {
		config.PgDumpBin = c.String("pg-dump-bin")
	}
	if c.IsSet("psql-bin") {
		config.PsqlBin = c.String("psql-bin")
	}
	if c.IsSet("mysqldump-bin") {
		config.MysqldumpBin = c.String("mysqldump-bin")
	}
	if c.IsSet("mysql-bin") {
		config.MysqlBin = c.String("mysql-bin")
	}
}
func GetConfigCli() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "file-prefix",
			Usage: "Required - LoadModuleConfig refuses to start without it. Local/mounted directory wal-g stores base backups and WAL segments under. Must be genuinely separate storage from the database's own disk to serve as a real disaster-recovery plan - see README.md. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.FilePrefix)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.FilePrefix, func(value string) {
							config.FilePrefix = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "compression-method",
			Usage: "Compression method wal-g uses for base backups and WAL segments. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.CompressionMethod)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.CompressionMethod, func(value string) {
							config.CompressionMethod = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "delta-max-steps",
			Usage: "How many delta backups wal-g chains before forcing a full backup. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DeltaMaxSteps)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DeltaMaxSteps, func(value string) {
							config.DeltaMaxSteps = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "retain-full-backups",
			Usage: "How many full backups the prune command keeps (plus what's needed to restore from the oldest of them forward). (int64)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.RetainFullBackups)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetInt64(c, config.RetainFullBackups, func(value int64) {
							config.RetainFullBackups = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "dump-dir",
			Usage: "Directory the dump command writes its output zip into when neither --output nor --hash is given. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DumpDir)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.DumpDir, func(value string) {
							config.DumpDir = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "dump-hash-ttl-seconds",
			Usage: "How many seconds a dump --hash download link stays claimable before it expires unfetched. The hash is single-use - once fetched (or once it fails), it's disabled immediately regardless of this TTL. (int64)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.DumpHashTtlSeconds)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetInt64(c, config.DumpHashTtlSeconds, func(value int64) {
							config.DumpHashTtlSeconds = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "pg-dump-bin",
			Usage: "pg_dump executable used by the dump command for postgres. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.PgDumpBin)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.PgDumpBin, func(value string) {
							config.PgDumpBin = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "psql-bin",
			Usage: "psql executable used by the restore-dump command for postgres. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.PsqlBin)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.PsqlBin, func(value string) {
							config.PsqlBin = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "mysqldump-bin",
			Usage: "mysqldump executable used by the dump command for mysql/mariadb. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.MysqldumpBin)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.MysqldumpBin, func(value string) {
							config.MysqldumpBin = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
		{
			Name:  "mysql-bin",
			Usage: "mysql client executable used by the restore-dump command for mysql/mariadb. (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.MysqlBin)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.MysqlBin, func(value string) {
							config.MysqlBin = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
	}
}
