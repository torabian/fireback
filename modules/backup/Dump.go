// Dump.go implements the plain dump/restore path (backup dump / backup
// restore-dump): a single-shot pg_dump/mysqldump/sqlite snapshot of one
// named database, zipped and either written to disk or streamed straight
// through an io.Writer (a local file, an HTTP response, ...). This is
// deliberately separate from Engine.go/wal-g: wal-g gives continuous WAL
// archiving + point-in-time restore for Postgres only, while this gives the
// simpler "just get me a copy of this one database" workflow, for all three
// vendors fireback supports.
package backup

import (
	"archive/zip"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"       // registers the "mysql" database/sql driver
	_ "github.com/jackc/pgx/v5/stdlib"       // registers the "pgx" database/sql driver
	_ "github.com/ncruces/go-sqlite3/driver" // registers the "sqlite3" database/sql driver - same module as sqlitedriver.GetSQLiteDialector's gormlite wrapper, just database/sql instead of gorm
)

// ListDatabases returns every database visible on cfg's connection, for
// interactive selection when `backup dump`/`backup restore-dump` aren't
// given an explicit --database. Postgres and MySQL both expose this from
// any connected database (pg_database is a cluster-wide catalog;
// information_schema.schemata likewise isn't scoped to one schema), so
// there's no need to connect to a specific "admin" database first.
//
// sqlite has no notion of multiple databases on one connection - a sqlite
// "database" is just cfg.Database, the configured file path - so this
// always returns exactly that one entry for sqlite.
func ListDatabases(ctx context.Context, cfg *DumpConfig) ([]string, error) {
	switch cfg.Vendor {
	case VendorSqlite:
		return []string{cfg.Database}, nil

	case VendorPostgres:
		db, err := sql.Open("pgx", postgresDsn(cfg, cfg.Database))
		if err != nil {
			return nil, err
		}
		defer db.Close()

		rows, err := db.QueryContext(ctx, `select datname from pg_database where not datistemplate and datallowconn order by datname`)
		if err != nil {
			return nil, fmt.Errorf("listing postgres databases: %w", err)
		}
		defer rows.Close()
		return scanStrings(rows)

	case VendorMysql:
		db, err := sql.Open("mysql", mysqlDsn(cfg, ""))
		if err != nil {
			return nil, err
		}
		defer db.Close()

		rows, err := db.QueryContext(ctx, `select schema_name from information_schema.schemata where schema_name not in ('information_schema','mysql','performance_schema','sys') order by schema_name`)
		if err != nil {
			return nil, fmt.Errorf("listing mysql databases: %w", err)
		}
		defer rows.Close()
		return scanStrings(rows)

	default:
		return nil, fmt.Errorf("unsupported vendor %q", cfg.Vendor)
	}
}

func scanStrings(rows *sql.Rows) ([]string, error) {
	var out []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		out = append(out, name)
	}
	return out, rows.Err()
}

func postgresDsn(cfg *DumpConfig, database string) string {
	// An empty "password=" keyword/value pair (rather than omitting the
	// keyword entirely) trips up pgx.ParseConfig - confirmed directly: it
	// silently drops every keyword after it, including dbname, so the
	// connection falls back to Postgres's "no database given" default (the
	// OS username) instead of the one actually requested. Passwordless
	// local/trust-auth setups are common enough that this must be handled,
	// not just assumed away.
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, database)
	if cfg.Password != "" {
		dsn += " password=" + cfg.Password
	}
	return dsn
}

// mysqlDsn builds a go-sql-driver/mysql DSN. An empty database connects
// without selecting one (fine for information_schema queries).
func mysqlDsn(cfg *DumpConfig, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, database)
}

// dumpEntryName is the file this database's dump is stored under inside the
// output zip - also what restore-dump looks for on the way back in.
func dumpEntryName(vendor DumpVendor, database string) string {
	switch vendor {
	case VendorSqlite:
		return database + ".sqlite"
	default:
		return database + ".sql"
	}
}

// DumpDatabase streams a single-file zip containing a snapshot of database
// to w. Every vendor path is a genuine stream (pg_dump/mysqldump's stdout,
// or a temp file for sqlite's VACUUM INTO - see dumpSqlite) copied straight
// into the zip entry as it's produced - nothing about the *source* dump is
// buffered in memory, though archive/zip itself only needs w to be a plain
// io.Writer (no Seek required), which is what makes writing this directly
// to an *os.File or an http.ResponseWriter equally possible.
//
// See the package doc / README "Dumping a live database" section for what
// consistency guarantee each vendor gives a concurrently-written database:
// short version, all three give you an atomic, consistent snapshot as of
// the instant the dump started - concurrent writes are simply not in it,
// never partially in it.
func DumpDatabase(ctx context.Context, cfg *DumpConfig, database string, w io.Writer) error {
	zw := zip.NewWriter(w)

	entry, err := zw.Create(dumpEntryName(cfg.Vendor, database))
	if err != nil {
		return err
	}

	switch cfg.Vendor {
	case VendorPostgres:
		if err := dumpPostgres(ctx, cfg, database, entry); err != nil {
			return err
		}
	case VendorMysql:
		if err := dumpMysql(ctx, cfg, database, entry); err != nil {
			return err
		}
	case VendorSqlite:
		if err := dumpSqlite(ctx, cfg, entry); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported vendor %q", cfg.Vendor)
	}

	return zw.Close()
}

// dumpPostgres runs pg_dump in the plain-SQL format (-Fp, the default) so
// restore-dump can pipe it straight into psql with no pg_restore/custom-
// format dependency. --single-transaction on the *dump* side isn't a thing
// pg_dump needs - pg_dump already always runs inside one transaction
// (REPEATABLE READ) internally, which is what makes dumping a live database
// safe: it sees a consistent snapshot as of when the dump started, nothing
// written afterwards, and blocks nothing but concurrent DDL on the tables
// it's reading.
func dumpPostgres(ctx context.Context, cfg *DumpConfig, database string, dst io.Writer) error {
	cmd := exec.CommandContext(ctx, cfg.PgDumpBin,
		"-h", cfg.Host, "-p", cfg.Port, "-U", cfg.User,
		"-d", database, "--no-owner", "--no-privileges",
	)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+cfg.Password)
	cmd.Stdout = dst
	var stderr sinkBuffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pg_dump %s: %w: %s", database, err, stderr.String())
	}
	return nil
}

// dumpMysql runs mysqldump with --single-transaction, which is what makes
// dumping a live InnoDB database consistent: without it, a dump of a
// database under concurrent write load can be torn (earlier-dumped tables
// reflect an earlier state than later ones), breaking cross-table
// referential consistency even though no single statement failed.
func dumpMysql(ctx context.Context, cfg *DumpConfig, database string, dst io.Writer) error {
	cmd := exec.CommandContext(ctx, cfg.MysqldumpBin,
		"-h", cfg.Host, "-P", cfg.Port, "-u", cfg.User,
		"--single-transaction", "--routines", "--triggers",
		// mysqldump emits `SET @@GLOBAL.GTID_PURGED=...` by default on a
		// GTID-enabled server - confirmed for real: restoring that
		// statement into a database that already has its own
		// GTID_EXECUTED history fails outright ("gtid set must not
		// overlap"). GTID_PURGED is a server-wide setting anyway, not
		// something a single-database dump/restore should be touching.
		"--set-gtid-purged=OFF",
		database,
	)
	if cfg.Password != "" {
		cmd.Env = append(os.Environ(), "MYSQL_PWD="+cfg.Password)
	}
	cmd.Stdout = dst
	var stderr sinkBuffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mysqldump %s: %w: %s", database, err, stderr.String())
	}
	return nil
}

// dumpSqlite uses SQLite's own `VACUUM INTO` (SQLite >= 3.27) rather than a
// plain file copy: VACUUM INTO is atomic and safe against concurrent
// readers/writers on the source database (SQLite's own backup API
// guarantee), where copying the file bytes directly could race a concurrent
// writer mid-page. It writes to a real temp file first since VACUUM INTO's
// destination must be a path SQLite itself opens, not an arbitrary
// io.Writer - that temp file is then streamed into the zip entry and
// removed.
func dumpSqlite(ctx context.Context, cfg *DumpConfig, dst io.Writer) error {
	if cfg.Database == "" || cfg.Database == ":memory:" {
		return fmt.Errorf("sqlite database is %q, nothing on disk to dump", cfg.Database)
	}

	tmp, err := os.CreateTemp("", "backup-dump-sqlite-*.sqlite")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	tmp.Close()
	os.Remove(tmpPath) // VACUUM INTO refuses to write to a file that already exists
	defer os.Remove(tmpPath)

	db, err := sql.Open("sqlite3", cfg.Database)
	if err != nil {
		return fmt.Errorf("opening sqlite database %s: %w", cfg.Database, err)
	}
	defer db.Close()

	// SQL string-literal quoting (single quotes, doubled to escape), not
	// Go's %q - VACUUM INTO takes its destination as a SQL string literal.
	quoted := "'" + strings.ReplaceAll(tmpPath, "'", "''") + "'"
	if _, err := db.ExecContext(ctx, "VACUUM INTO "+quoted); err != nil {
		return fmt.Errorf("VACUUM INTO %s: %w", cfg.Database, err)
	}

	f, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(dst, f)
	return err
}

// DefaultDumpFilename is the suggested/default output filename for a dump
// of database - <database>_<YYYYMMDD>.zip, so re-running `backup dump`
// against the same database on different days doesn't silently clobber a
// previous day's file the way a fixed <database>.zip name would.
func DefaultDumpFilename(database string) string {
	return fmt.Sprintf("%s_%s.zip", database, time.Now().Format("20060102"))
}

// DefaultDumpPath makes sure cfg.DumpDir exists and returns
// <dump-dir>/<DefaultDumpFilename(database)> - the path a new dump for
// database is written to when neither --output nor an interactively-chosen
// path is given.
func DefaultDumpPath(cfg *DumpConfig, database string) (string, error) {
	dir := cfg.DumpDir
	if dir == "" {
		dir = "."
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("creating dump dir %s: %w", dir, err)
	}
	return filepath.Join(dir, DefaultDumpFilename(database)), nil
}

// EstimateDatabaseSize returns a rough, best-effort byte estimate of what
// database would dump to - shown to the operator before an interactive
// `backup dump` actually runs, alongside available disk space (see
// DiskSpace.go), so a run that's clearly going to fail partway through
// (or clearly won't fit) is visible up front instead of discovered
// mid-transfer.
//
// This is deliberately a rough estimate, not a prediction of the exact
// output size: postgres/mysql report the *on-disk* size of the database
// (tables + indexes + overhead), while the actual dump is usually smaller
// (no index structures, often compressible) - callers should present it as
// an upper-bound-ish figure, not an exact number. sqlite is the one exact
// case: VACUUM INTO produces a compacted copy no larger than the source
// file, so the current file size is already a safe upper bound.
func EstimateDatabaseSize(ctx context.Context, cfg *DumpConfig, database string) (int64, error) {
	switch cfg.Vendor {
	case VendorPostgres:
		db, err := sql.Open("pgx", postgresDsn(cfg, cfg.Database))
		if err != nil {
			return 0, err
		}
		defer db.Close()

		var size int64
		if err := db.QueryRowContext(ctx, `select pg_database_size($1)`, database).Scan(&size); err != nil {
			return 0, fmt.Errorf("estimating size of %s: %w", database, err)
		}
		return size, nil

	case VendorMysql:
		db, err := sql.Open("mysql", mysqlDsn(cfg, ""))
		if err != nil {
			return 0, err
		}
		defer db.Close()

		var size sql.NullInt64
		err = db.QueryRowContext(ctx,
			`select sum(data_length + index_length) from information_schema.tables where table_schema = ?`,
			database,
		).Scan(&size)
		if err != nil {
			return 0, fmt.Errorf("estimating size of %s: %w", database, err)
		}
		return size.Int64, nil // NULL (no tables yet) scans as 0, which is a fine estimate

	case VendorSqlite:
		info, err := os.Stat(cfg.Database)
		if err != nil {
			return 0, fmt.Errorf("statting %s: %w", cfg.Database, err)
		}
		return info.Size(), nil

	default:
		return 0, fmt.Errorf("unsupported vendor %q", cfg.Vendor)
	}
}

// sinkBuffer is a tiny bytes.Buffer stand-in kept local to this file so
// stderr from pg_dump/mysqldump can be captured for the error message
// without pulling in bytes.Buffer's full API footprint here.
type sinkBuffer struct{ data []byte }

func (b *sinkBuffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}
func (b *sinkBuffer) String() string { return string(b.data) }
