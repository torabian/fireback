package backup

import (
	"fmt"

	backupdefs "github.com/torabian/fireback/modules/backup/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/dbdsn"
)

// DumpVendor is which database engine `backup dump`/`backup restore-dump`
// talks to. Unlike ModuleConfig (wal-g, Postgres-only), the plain dump/
// restore path supports all three vendors fireback itself supports.
type DumpVendor string

const (
	VendorPostgres DumpVendor = "postgres"
	VendorMysql    DumpVendor = "mysql"
	VendorSqlite   DumpVendor = "sqlite"
)

// DumpConfig is everything Dump.go/RestoreDump.go need: the database
// connection to dump/restore against (parsed out of fireback's own DB_DSN
// via modules/fireback/dbdsn - unlike ModuleConfig, there's no separate
// PGHOST-style override here, since dump/restore is meant to run inside a
// fireback app that already knows its own database, not standalone on a
// bare data directory the way wal-g's push/restore must), plus the
// dump-specific knobs from Backup.emi.yml's config: block.
type DumpConfig struct {
	Vendor DumpVendor

	Host     string
	Port     string
	User     string
	Password string

	// Database is fireback's own configured default database (the dbname
	// piece of DB_DSN) - used as the fallback when dump/restore aren't
	// given an explicit --database, and as the only meaningful value for
	// sqlite (a sqlite "database" is just this file path, i.e. DB_DSN
	// itself; there is nothing to list/select).
	Database string

	DumpDir        string
	HashTTLSeconds int64

	PgDumpBin    string
	PsqlBin      string
	MysqldumpBin string
	MysqlBin     string
}

// LoadDumpConfig resolves DumpConfig from fireback's own configuration.
func LoadDumpConfig() (*DumpConfig, error) {
	fc := fireback.LoadConfiguration()
	bc := backupdefs.LoadConfiguration()

	var vendor DumpVendor
	switch fc.DbVendor {
	case "postgres":
		vendor = VendorPostgres
	case "mysql", "mariadb":
		vendor = VendorMysql
	case "sqlite", "":
		vendor = VendorSqlite
	default:
		return nil, fmt.Errorf("backup dump/restore-dump: unsupported DB_VENDOR %q (supported: postgres, mysql, mariadb, sqlite)", fc.DbVendor)
	}

	cfg := &DumpConfig{
		Vendor: vendor,

		DumpDir:        bc.DumpDir,
		HashTTLSeconds: bc.DumpHashTtlSeconds,

		PgDumpBin:    bc.PgDumpBin,
		PsqlBin:      bc.PsqlBin,
		MysqldumpBin: bc.MysqldumpBin,
		MysqlBin:     bc.MysqlBin,
	}

	if vendor == VendorSqlite {
		// sqlite's "dsn" is just the file path (or ":memory:") - see
		// fireback.GetDatabaseDsn.
		cfg.Database = fc.DbDsn
	} else {
		info, err := dbdsn.Parse(fc.DbVendor, fc.DbDsn)
		if err != nil {
			return nil, fmt.Errorf("backup dump/restore-dump: %w", err)
		}
		cfg.Host = info.Host
		cfg.Port = info.Port
		cfg.User = info.Username
		cfg.Password = info.Password
		cfg.Database = info.Database
	}

	return cfg, nil
}
