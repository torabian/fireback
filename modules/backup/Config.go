// Package backup drives wal-g (https://github.com/wal-g/wal-g, imported
// directly as a go.mod dependency - see Exec.go) to provide Postgres
// physical backups (base backup + continuous WAL archiving) with true
// point-in-time restore, as a set of urfave/cli/v3 commands mountable into
// any fireback app - the same convention modules/storage uses for its own
// Cli.go.
//
// It deliberately does not reimplement anything wal-g already does
// correctly (delta backups, WAL push/fetch, retention). This package is the
// orchestration/policy layer: resolving connection config, cataloging what
// wal-g reports, and driving restore-to-a-point-in-time end to end.
package backup

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/torabian/fireback/modules/fireback"
)

// ModuleConfig holds everything Engine needs to invoke wal-g: the Postgres
// connection it should back up/restore, and where backup artifacts live.
type ModuleConfig struct {
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgDatabase string

	// FilePrefix is WALG_FILE_PREFIX - the local/mounted directory backups
	// and WAL segments are stored under. See the warning in
	// docker-compose.yml: this must be genuinely separate storage from the
	// database's own disk to serve as a real disaster-recovery plan.
	FilePrefix string

	CompressionMethod string // WALG_COMPRESSION_METHOD
	DeltaMaxSteps     string // WALG_DELTA_MAX_STEPS - how many deltas before a full backup is forced
	RetainFullBackups int    // how many full backups `backup prune` keeps
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// LoadModuleConfig resolves Postgres connection details from fireback's own
// configuration first (the same DB_HOST/DB_PORT/DB_USERNAME/DB_PASSWORD/
// DB_NAME pieces GetDatabaseDsn builds a connection string from), so this
// module never needs its own separate set of DB_* settings - then lets the
// standard libpq PGHOST/PGPORT/PGUSER/PGPASSWORD/PGDATABASE env vars
// override them, which matters when this runs as the standalone
// backup-cli binary pointed at a database with no app config nearby.
func LoadModuleConfig() (*ModuleConfig, error) {
	fc := fireback.LoadConfiguration()

	// Explicit PGHOST/PGDATABASE (the standalone backup-cli case, run
	// wherever wal-g itself would be) always wins over fireback's own
	// config, including its DbVendor - fireback defaults DbVendor to
	// sqlite when nothing else is configured, which must not reject a
	// caller that has plainly pointed us at a Postgres instance directly.
	explicitPg := os.Getenv("PGHOST") != "" || os.Getenv("PGDATABASE") != ""
	if !explicitPg && fc.DbVendor != "" && fc.DbVendor != "postgres" {
		return nil, fmt.Errorf("backup module only supports postgres, got DB_VENDOR=%q", fc.DbVendor)
	}

	cfg := &ModuleConfig{
		PgHost:     envOr("PGHOST", fc.DbHost),
		PgPort:     envOr("PGPORT", fmt.Sprintf("%v", fc.DbPort)),
		PgUser:     envOr("PGUSER", fc.DbUsername),
		PgPassword: envOr("PGPASSWORD", fc.DbPassword),
		PgDatabase: envOr("PGDATABASE", fc.DbName),

		FilePrefix:        os.Getenv("WALG_FILE_PREFIX"),
		CompressionMethod: envOr("WALG_COMPRESSION_METHOD", "lz4"),
		DeltaMaxSteps:     envOr("WALG_DELTA_MAX_STEPS", "6"),
		RetainFullBackups: 4,
	}

	if n := os.Getenv("BACKUP_RETAIN_FULL"); n != "" {
		v, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("BACKUP_RETAIN_FULL must be an integer, got %q: %w", n, err)
		}
		cfg.RetainFullBackups = v
	}

	if cfg.FilePrefix == "" {
		return nil, errors.New("WALG_FILE_PREFIX must be set to the local/mounted directory backups are stored under")
	}
	if cfg.PgHost == "" || cfg.PgDatabase == "" {
		return nil, errors.New("no Postgres connection found: set PGHOST/PGDATABASE (and PGUSER/PGPASSWORD/PGPORT as needed), or run this inside a fireback app configured for postgres")
	}

	return cfg, nil
}
