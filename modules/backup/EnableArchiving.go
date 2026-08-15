// EnableArchiving.go implements `backup enable` - applying the settings
// README.md's own "Postgres-side configuration" section otherwise asks an
// operator to set by hand (wal_level, archive_mode, archive_command,
// archive_timeout) directly, via ALTER SYSTEM SET, instead of leaving it as
// a manual copy-paste-into-postgresql.conf step.
package backup

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// archiveCommandString builds the archive_command Postgres itself will exec
// as a plain shell command for every WAL segment it finishes - the
// archive_command mirror of restoreCommandString (Restore.go): the same
// approach (point at this same running binary via os.Executable(), with the
// hidden __walg-exec marker and WALG_FILE_PREFIX baked in inline - Postgres
// forks archive_command with its own environment, which won't already have
// WALG_FILE_PREFIX in it unless it's put there explicitly), just wal-push
// instead of wal-fetch.
func archiveCommandString(engine *Engine) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("resolving own executable for archive_command: %w", err)
	}
	return fmt.Sprintf("WALG_FILE_PREFIX=%s %s %s wal-push %%p", engine.cfg.FilePrefix, exePath, walgExecMarker), nil
}

// SettingChange records one Postgres GUC EnablePostgresArchiving actually
// changed - printed by `backup enable` so the operator sees exactly what
// moved and from what, not just "done". Pending is true for a
// PGC_POSTMASTER setting (wal_level/archive_mode) that's been written to
// postgresql.auto.conf but isn't live yet, pending a restart.
type SettingChange struct {
	Name    string
	Old     string
	New     string
	Pending bool
}

// pgLiteral quotes v as a SQL string literal for ALTER SYSTEM SET - the
// same single-quote-doubling escaping already used elsewhere in this module
// (Restore.go's escapeConfValue, Dump.go's VACUUM INTO quoting). ALTER
// SYSTEM SET has no way to bind a parameter; the value has to be a literal
// in the statement text, and a quoted string literal is accepted for every
// GUC type this sets (enum, string, integer alike).
func pgLiteral(v string) string {
	return "'" + strings.ReplaceAll(v, "'", "''") + "'"
}

// moduleConfigDsn builds a libpq keyword/value DSN from ModuleConfig, the
// same "omit password= entirely when empty" way Dump.go's postgresDsn does
// for DumpConfig - pgx.ParseConfig silently drops dbname (and everything
// after it) when password= is present but empty, confirmed for real - see
// the memory note on this from when Dump.go's own DSN builder hit it.
func moduleConfigDsn(cfg *ModuleConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgDatabase)
	if cfg.PgPassword != "" {
		dsn += " password=" + cfg.PgPassword
	}
	return dsn
}

func currentSetting(ctx context.Context, db *sql.DB, name string) (string, error) {
	var value string
	err := db.QueryRowContext(ctx, "select setting from pg_settings where name = $1", name).Scan(&value)
	return value, err
}

// fileSetting returns what name is set to in postgresql.auto.conf, per
// pg_file_settings (the parsed view of every config file Postgres loaded),
// and whether that value is actually applied to the running backend.
//
// This - not pg_settings.setting/SHOW - is the reliable way to check "is
// this GUC already configured the way I want", for two different reasons
// depending on the GUC:
//
//   - wal_level/archive_mode (PGC_POSTMASTER): a value written by ALTER
//     SYSTEM SET only takes effect on a full restart. Until then it shows
//     up here with applied=false, while pg_settings.setting still reports
//     the old live value - this is exactly how "already staged, pending
//     restart" is detected instead of re-issuing an identical ALTER SYSTEM
//     SET (and re-reporting it as a fresh change) every time this runs.
//   - archive_command (PGC_SIGHUP - takes effect on reload, no restart
//     needed): confirmed for real, Postgres's own SHOW/pg_settings report
//     archive_command as the literal string "(disabled)" whenever
//     archive_mode is off, regardless of what archive_command is actually
//     configured to - a documented quirk of its GUC show-hook, not a sign
//     the value didn't take. pg_file_settings.applied=true here means the
//     real underlying value genuinely is loaded, even while archive_mode
//     is still off pending its own restart.
func fileSetting(ctx context.Context, db *sql.DB, name string) (value string, applied bool, found bool, err error) {
	err = db.QueryRowContext(ctx, `
		select setting, applied from pg_file_settings
		where name = $1 and sourcefile like '%postgresql.auto.conf'
		order by seqno desc limit 1
	`, name).Scan(&value, &applied)
	if err == sql.ErrNoRows {
		return "", false, false, nil
	}
	if err != nil {
		return "", false, false, err
	}
	return value, applied, true, nil
}

// EnablePostgresArchiving applies wal_level=replica, archive_mode=on,
// archive_command (pointed at this same binary's own `backup archive-push`,
// see archiveCommandString), and archive_timeout - via ALTER SYSTEM SET
// (writes postgresql.auto.conf) plus a reload. Settings already at the
// wanted value are left alone and not reported as changed.
//
// wal_level/archive_mode are PGC_POSTMASTER parameters - Postgres can only
// pick up a change to either on a full restart, which this function does
// not do itself (the returned needsRestart reports whether one is actually
// required): restarting a database out from under a caller without being
// asked to isn't something to do silently. archive_command/archive_timeout
// are PGC_SIGHUP and take effect from the reload alone.
//
// Requires connecting as a Postgres superuser (or, on Postgres 15+, a role
// specifically granted these parameters) - ALTER SYSTEM SET fails with a
// normal Postgres permission error otherwise, surfaced as-is.
func EnablePostgresArchiving(ctx context.Context, cfg *ModuleConfig, archiveTimeoutSeconds int) ([]SettingChange, bool, error) {
	engine := NewEngine(cfg)
	archiveCmd, err := archiveCommandString(engine)
	if err != nil {
		return nil, false, err
	}

	db, err := sql.Open("pgx", moduleConfigDsn(cfg))
	if err != nil {
		return nil, false, err
	}
	defer db.Close()

	wanted := []struct {
		name         string
		value        string
		needsRestart bool
	}{
		{"wal_level", "replica", true},
		{"archive_mode", "on", true},
		{"archive_command", archiveCmd, false},
		{"archive_timeout", fmt.Sprintf("%d", archiveTimeoutSeconds), false},
	}

	var changed []SettingChange
	var needsRestart bool

	for _, w := range wanted {
		current, err := currentSetting(ctx, db, w.name)
		if err != nil {
			return changed, needsRestart, fmt.Errorf("reading current %s: %w", w.name, err)
		}

		// pg_file_settings.applied, not pg_settings.setting/SHOW alone, is
		// what makes this reliable - see fileSetting's own doc comment for
		// the two different reasons plain live-value comparison can't be
		// trusted by itself for every setting here.
		staged, applied, found, err := fileSetting(ctx, db, w.name)
		if err != nil {
			return changed, needsRestart, fmt.Errorf("checking configured %s: %w", w.name, err)
		}

		switch {
		case current == w.value, found && applied && staged == w.value:
			continue // already fully in effect (live, or via a config source that isn't postgresql.auto.conf) - nothing to do, nothing to report

		case found && staged == w.value && !applied:
			// Staged from a previous run, not live yet - only possible
			// for a PGC_POSTMASTER setting (a PGC_SIGHUP one would have
			// already shown applied=true from its own reload).
			changed = append(changed, SettingChange{Name: w.name, Old: current, New: w.value, Pending: true})
			needsRestart = true

		default:
			if _, err := db.ExecContext(ctx, "ALTER SYSTEM SET "+w.name+" = "+pgLiteral(w.value)); err != nil {
				return changed, needsRestart, fmt.Errorf("ALTER SYSTEM SET %s: %w", w.name, err)
			}
			changed = append(changed, SettingChange{Name: w.name, Old: current, New: w.value})
			if w.needsRestart {
				needsRestart = true
			}
		}
	}

	if len(changed) > 0 {
		if _, err := db.ExecContext(ctx, "select pg_reload_conf()"); err != nil {
			return changed, needsRestart, fmt.Errorf("reloading config: %w", err)
		}
	}

	return changed, needsRestart, nil
}
