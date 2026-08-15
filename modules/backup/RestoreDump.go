package backup

import (
	"archive/zip"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// RestoreDump reads a zip produced by DumpDatabase from r and loads it into
// database - for postgres/mysql, that's fireback.CREATE DATABASE'd
// automatically if it doesn't already exist yet (restoring is meant to
// re-create a database, not require one to already be sitting there empty);
// for sqlite it's the destination file path, created if missing.
//
// If database *does* already exist (postgres/mysql), or the sqlite file
// already exists, RestoreDump refuses and returns an error unless force is
// true - restoring blindly into something that already has data would
// silently merge/overwrite it, which is far more often a mistake than an
// intent. Pass force to restore into an existing target anyway (e.g. you
// already dropped/emptied it yourself, or you're intentionally overwriting
// a scratch/staging database) - RestoreDump does not itself drop or clean
// any existing objects first, so a real conflict (e.g. a table that already
// exists) still surfaces as a normal psql/mysql error either way.
//
// r only needs to be a plain io.Reader for the postgres/mysql paths (piped
// straight into psql/mysql's stdin as it's read out of the zip); sqlite's
// restore needs the whole file materialized on disk first since it's a
// direct file replacement, not a stream into a running server - see
// restoreSqlite.
func RestoreDump(ctx context.Context, cfg *DumpConfig, database string, r io.Reader, force bool) error {
	// Checked before spooling anything from r - no point staging a
	// potentially large incoming stream just to refuse it a moment later.
	switch cfg.Vendor {
	case VendorPostgres:
		if err := ensurePostgresTarget(ctx, cfg, database, force); err != nil {
			return err
		}
	case VendorMysql:
		if err := ensureMysqlTarget(ctx, cfg, database, force); err != nil {
			return err
		}
	case VendorSqlite:
		if err := ensureSqliteTarget(cfg, force); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported vendor %q", cfg.Vendor)
	}

	// archive/zip needs a ReaderAt + size (random access into the central
	// directory at the end of the file), which an arbitrary io.Reader
	// (e.g. an HTTP response body for --hash restores) doesn't give us -
	// so the incoming zip is always spooled to a temp file first. This is
	// restore, not dump: it's inherently a slower, rarer, operator-driven
	// path, so trading the ability to stream byte-for-byte for the
	// simplicity/correctness of using archive/zip normally is the right
	// call here (unlike DumpDatabase, which stays a true stream in both
	// directions).
	tmp, err := os.CreateTemp("", "backup-restore-*.zip")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	size, err := io.Copy(tmp, r)
	tmp.Close()
	if err != nil {
		return fmt.Errorf("staging incoming dump: %w", err)
	}

	zr, err := zip.OpenReader(tmpPath)
	if err != nil {
		return fmt.Errorf("reading dump as zip (%d bytes staged): %w", size, err)
	}
	defer zr.Close()

	if len(zr.File) != 1 {
		return fmt.Errorf("expected exactly one file inside the dump zip, found %d", len(zr.File))
	}
	entry := zr.File[0]

	rc, err := entry.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	switch cfg.Vendor {
	case VendorPostgres:
		return restorePostgres(ctx, cfg, database, rc)
	case VendorMysql:
		return restoreMysql(ctx, cfg, database, rc)
	case VendorSqlite:
		return restoreSqlite(cfg, rc)
	default:
		return fmt.Errorf("unsupported vendor %q", cfg.Vendor)
	}
}

// refuseExistingErr is the shared "not without --force" message shape for
// all three vendors, so the CLI's own error text is consistent regardless
// of which one triggered it.
func refuseExistingErr(kind, target string) error {
	return fmt.Errorf("%s %q already exists - restore-dump refuses to restore into an existing %s by default (it could overwrite or merge into existing data); pass --force to restore into it anyway, or choose a different --database", kind, target, kind)
}

// pgQuoteIdent quotes name as a Postgres identifier (double-quoted,
// embedded double quotes doubled) - needed since database is only ever
// interpolated into SQL here, never passed as a bind parameter (Postgres
// has no way to parameterize an identifier in CREATE DATABASE).
func pgQuoteIdent(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

// mysqlQuoteIdent is pgQuoteIdent's MySQL equivalent - backtick-quoted,
// embedded backticks doubled.
func mysqlQuoteIdent(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

// ensurePostgresTarget connects to cfg.Database (fireback's own already-
// configured, already-existing default database - not a hardcoded
// "postgres"/"template1" admin database, which isn't guaranteed to still
// exist) to check whether database exists, creating it if not, or refusing
// if it does and force is false.
func ensurePostgresTarget(ctx context.Context, cfg *DumpConfig, database string, force bool) error {
	db, err := sql.Open("pgx", postgresDsn(cfg, cfg.Database))
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	if err := db.QueryRowContext(ctx, `select exists(select 1 from pg_database where datname = $1)`, database).Scan(&exists); err != nil {
		return fmt.Errorf("checking whether database %s already exists: %w", database, err)
	}

	if exists {
		if !force {
			return refuseExistingErr("database", database)
		}
		return nil
	}

	if _, err := db.ExecContext(ctx, "CREATE DATABASE "+pgQuoteIdent(database)); err != nil {
		return fmt.Errorf("creating database %s: %w", database, err)
	}
	return nil
}

// ensureMysqlTarget is ensurePostgresTarget's MySQL/MariaDB equivalent -
// connects without selecting a schema (mysqlDsn(cfg, "") - same as
// ListDatabases uses) since CREATE DATABASE doesn't need one.
func ensureMysqlTarget(ctx context.Context, cfg *DumpConfig, database string, force bool) error {
	db, err := sql.Open("mysql", mysqlDsn(cfg, ""))
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	if err := db.QueryRowContext(ctx, `select exists(select 1 from information_schema.schemata where schema_name = ?)`, database).Scan(&exists); err != nil {
		return fmt.Errorf("checking whether database %s already exists: %w", database, err)
	}

	if exists {
		if !force {
			return refuseExistingErr("database", database)
		}
		return nil
	}

	if _, err := db.ExecContext(ctx, "CREATE DATABASE "+mysqlQuoteIdent(database)); err != nil {
		return fmt.Errorf("creating database %s: %w", database, err)
	}
	return nil
}

// ensureSqliteTarget applies the same "refuse unless --force" rule sqlite's
// destination file: a sqlite "database" is just a file, and an existing one
// at cfg.Database is just as much "this already has data" as an existing
// postgres/mysql database is.
func ensureSqliteTarget(cfg *DumpConfig, force bool) error {
	if cfg.Database == "" || cfg.Database == ":memory:" {
		return fmt.Errorf("sqlite database is %q, nowhere on disk to restore into", cfg.Database)
	}
	if _, err := os.Stat(cfg.Database); err == nil {
		if !force {
			return refuseExistingErr("file", cfg.Database)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func restorePostgres(ctx context.Context, cfg *DumpConfig, database string, src io.Reader) error {
	cmd := exec.CommandContext(ctx, cfg.PsqlBin,
		"-h", cfg.Host, "-p", cfg.Port, "-U", cfg.User,
		"-d", database, "-v", "ON_ERROR_STOP=1", "-q",
	)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+cfg.Password)
	cmd.Stdin = src
	var stderr sinkBuffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("psql restore into %s: %w: %s", database, err, stderr.String())
	}
	return nil
}

func restoreMysql(ctx context.Context, cfg *DumpConfig, database string, src io.Reader) error {
	cmd := exec.CommandContext(ctx, cfg.MysqlBin,
		"-h", cfg.Host, "-P", cfg.Port, "-u", cfg.User,
		database,
	)
	if cfg.Password != "" {
		cmd.Env = append(os.Environ(), "MYSQL_PWD="+cfg.Password)
	}
	cmd.Stdin = src
	var stderr sinkBuffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mysql restore into %s: %w: %s", database, err, stderr.String())
	}
	return nil
}

// restoreSqlite writes src to cfg.Database. ensureSqliteTarget has already
// refused (unless force) if the file existed, but a timestamped .bak copy
// is still made here whenever it does exist - a --force restore is still a
// destructive operation, and a typo'd --database/--file shouldn't be
// unrecoverable just because the operator explicitly opted into overwriting
// *something*.
//
// This does not (and cannot, from here) check whether some other process
// still has cfg.Database open - a concurrent writer holding it open across
// this file replacement is the one real hazard specific to sqlite restore,
// since sqlite has no server process to coordinate through. Stop whatever
// has it open first.
func restoreSqlite(cfg *DumpConfig, src io.Reader) error {
	if _, err := os.Stat(cfg.Database); err == nil {
		bak := fmt.Sprintf("%s.%s.bak", cfg.Database, time.Now().UTC().Format("20060102T150405Z"))
		if err := copyFile(cfg.Database, bak); err != nil {
			return fmt.Errorf("backing up existing %s to %s before restore: %w", cfg.Database, bak, err)
		}
	}

	if dir := filepath.Dir(cfg.Database); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	tmp := cfg.Database + ".restoring"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, src); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return err
	}

	// A quick sanity check that this is actually a sqlite database before
	// committing to the rename - opening it and running a cheap
	// integrity_check catches "this was actually the wrong file" before
	// it overwrites the real one.
	db, err := sql.Open("sqlite3", tmp)
	if err != nil {
		os.Remove(tmp)
		return err
	}
	_, execErr := db.ExecContext(context.Background(), "PRAGMA quick_check")
	db.Close()
	if execErr != nil {
		os.Remove(tmp)
		return fmt.Errorf("restored file failed sqlite quick_check, not committing: %w", execErr)
	}

	return os.Rename(tmp, cfg.Database)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
