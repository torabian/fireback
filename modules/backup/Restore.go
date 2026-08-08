package backup

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// RestoreOptions describes a single restore: fetch the right base backup,
// then let Postgres replay WAL forward.
type RestoreOptions struct {
	// TargetTime, if set, restores to that specific instant (point-in-time
	// restore: undo something that happened after it). If nil, restores to
	// the latest available WAL instead - replays everything archived and
	// auto-promotes once none is left (the same mechanism Verify uses) -
	// for "recreate the database as fully up to date as possible" rather
	// than undoing a specific past mistake.
	TargetTime *time.Time
	TargetDir  string

	// Start, if true, starts Postgres against TargetDir after preparing it
	// and waits for recovery to finish (recovery.signal disappears once
	// Postgres promotes). This requires pg_ctl/postgres to be on PATH -
	// best-effort convenience, not required for the restore itself: an
	// operator can always start Postgres against TargetDir by hand.
	Start bool
	// Port Postgres should listen on when Start is true. Defaults to 5433
	// so a restore drill on the same host never collides with a live
	// primary on 5432.
	Port int
	// Promote, if true, sets recovery_target_action=promote so Postgres
	// exits recovery automatically once it reaches TargetTime. If false,
	// recovery_target_action=pause, leaving the cluster paused in
	// recovery at that instant so the operator can inspect data before
	// committing to the promotion.
	Promote bool
}

// ResolveBackupForTime returns the newest backup whose creation time is at
// or before target - the base wal-g needs to fetch before WAL replay can
// carry it the rest of the way to target. Returns an error if every known
// backup is newer than target (nothing to restore from) or none exist.
func ResolveBackupForTime(ctx context.Context, engine *Engine, target time.Time) (BackupInfo, error) {
	backups, err := engine.BackupList(ctx)
	if err != nil {
		return BackupInfo{}, err
	}

	var chosen *BackupInfo
	for i := range backups {
		b := &backups[i]
		if b.Time.After(target) {
			continue
		}
		if chosen == nil || b.Time.After(chosen.Time) {
			chosen = b
		}
	}
	if chosen == nil {
		return BackupInfo{}, fmt.Errorf("no backup found at or before %s (oldest available backup is newer than the requested target)", target.Format(time.RFC3339))
	}
	return *chosen, nil
}

// resolveRestoreBackup picks the base backup a restore should start from:
// the newest one at or before targetTime, or - if targetTime is nil -
// simply the newest backup available, to replay forward from as far as WAL
// allows.
func resolveRestoreBackup(ctx context.Context, engine *Engine, targetTime *time.Time) (BackupInfo, error) {
	if targetTime != nil {
		return ResolveBackupForTime(ctx, engine, *targetTime)
	}

	backups, err := engine.BackupList(ctx)
	if err != nil {
		return BackupInfo{}, err
	}
	if len(backups) == 0 {
		return BackupInfo{}, errors.New("no backups found")
	}
	newest := backups[0]
	for _, b := range backups {
		if b.Time.After(newest.Time) {
			newest = b
		}
	}
	return newest, nil
}

// Restore fetches the right base backup into opts.TargetDir and configures
// it for Postgres to replay WAL forward - to opts.TargetTime if set, or as
// far as available WAL allows otherwise. If opts.Start is set, it also
// starts Postgres and waits for recovery to finish.
func Restore(ctx context.Context, engine *Engine, opts RestoreOptions) (BackupInfo, error) {
	backup, err := resolveRestoreBackup(ctx, engine, opts.TargetTime)
	if err != nil {
		return BackupInfo{}, err
	}

	if opts.TargetTime != nil {
		fmt.Printf("restoring base backup %s (created %s) as the starting point for target time %s\n",
			backup.BackupName, backup.Time.Format(time.RFC3339), opts.TargetTime.Format(time.RFC3339))
	} else {
		fmt.Printf("restoring base backup %s (created %s), replaying all available WAL to bring the database fully up to date\n",
			backup.BackupName, backup.Time.Format(time.RFC3339))
	}

	if err := engine.BackupFetch(ctx, opts.TargetDir, backup.BackupName); err != nil {
		return backup, fmt.Errorf("fetching base backup: %w", err)
	}

	if err := writeRecoverySignal(opts.TargetDir); err != nil {
		return backup, err
	}

	if opts.TargetTime != nil {
		if err := writeRecoveryConfig(engine, opts); err != nil {
			return backup, err
		}
		fmt.Printf("base backup restored into %s, configured to recover to %s\n", opts.TargetDir, opts.TargetTime.Format(time.RFC3339))
	} else {
		if err := writeBareRestoreCommand(engine, opts.TargetDir); err != nil {
			return backup, err
		}
		fmt.Printf("base backup restored into %s, configured to replay all available WAL and auto-promote\n", opts.TargetDir)
	}

	if !opts.Start {
		fmt.Println("start Postgres against this data directory to begin WAL replay (pass --start to have this command do it and wait).")
		return backup, nil
	}

	port := opts.Port
	if port == 0 {
		port = 5433
	}
	if err := StartPostgresAndWait(ctx, opts.TargetDir, port, 5*time.Minute); err != nil {
		return backup, fmt.Errorf("starting postgres for recovery: %w", err)
	}
	if err := WaitRecoveryComplete(ctx, opts.TargetDir, 30*time.Minute); err != nil {
		return backup, err
	}
	fmt.Println("recovery complete")
	return backup, nil
}

func writeRecoverySignal(dataDir string) error {
	return os.WriteFile(filepath.Join(dataDir, "recovery.signal"), nil, 0o644)
}

// escapeConfValue doubles single quotes the way postgresql.conf's own
// single-quoted string syntax requires, so a path or command containing one
// doesn't break the config file.
func escapeConfValue(v string) string {
	return strings.ReplaceAll(v, "'", "''")
}

// restoreCommandString builds the restore_command Postgres itself will
// exec as a plain shell command during recovery. Postgres always forks+execs
// a command line - it has no way to call into our embedded wal-g directly -
// so this must name an actual executable on disk. There is still no
// separate wal-g binary anywhere: it points at this same running binary
// (resolved via os.Executable()) with the __walg-exec marker, exactly like
// Engine.command's own re-exec, which is what makes wal-fetch run when
// Postgres invokes this.
func restoreCommandString(engine *Engine) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("resolving own executable for restore_command: %w", err)
	}
	// WALG_FILE_PREFIX is set inline rather than relying on the ambient
	// environment of whatever process starts Postgres for this restore -
	// wal-fetch is a pure storage read, it needs no Postgres connection
	// details, only to know where backups live.
	return fmt.Sprintf("WALG_FILE_PREFIX=%s %s %s wal-fetch %%f %%p", engine.cfg.FilePrefix, exePath, walgExecMarker), nil
}

// writeRecoveryConfig appends restore_command/recovery_target_time/
// recovery_target_action to postgresql.auto.conf.
func writeRecoveryConfig(engine *Engine, opts RestoreOptions) error {
	action := "pause"
	if opts.Promote {
		action = "promote"
	}

	restoreCommand, err := restoreCommandString(engine)
	if err != nil {
		return err
	}

	lines := []string{
		"",
		"# added by nima backup restore",
		fmt.Sprintf("restore_command = '%s'", escapeConfValue(restoreCommand)),
		// .000000 (not .999999) so this always emits full microsecond
		// precision - found by actually restoring across a sub-second gap
		// between two transactions: the old format ("...05-07", no
		// fractional seconds at all) silently truncated the target down to
		// the whole second, which could land *before* a transaction that
		// happened earlier in the same second than the real target,
		// excluding it from the restore even though it should have been
		// included.
		fmt.Sprintf("recovery_target_time = '%s'", opts.TargetTime.UTC().Format("2006-01-02 15:04:05.000000-07")),
		fmt.Sprintf("recovery_target_action = '%s'", action),
		"",
	}

	path := filepath.Join(opts.TargetDir, "postgresql.auto.conf")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strings.Join(lines, "\n"))
	return err
}

// writeBareRestoreCommand configures dataDir to replay every WAL segment
// available and auto-promote once none are left, with no recovery target -
// used by Verify, which wants the broadest possible replay, not a stop at a
// specific instant.
func writeBareRestoreCommand(engine *Engine, dataDir string) error {
	restoreCommand, err := restoreCommandString(engine)
	if err != nil {
		return err
	}
	content := fmt.Sprintf("\n# added by nima backup verify\nrestore_command = '%s'\n", escapeConfValue(restoreCommand))

	path := filepath.Join(dataDir, "postgresql.auto.conf")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// StartPostgresAndWait starts Postgres against dataDir via pg_ctl on the
// given port and waits (bounded by timeout) for it to report ready. It
// disables archive_mode for this run - a recovering copy must never archive
// WAL into the same prefix its own base backups came from.
func StartPostgresAndWait(ctx context.Context, dataDir string, port int, timeout time.Duration) error {
	args := []string{
		"start", "-D", dataDir, "-w",
		"-t", fmt.Sprintf("%d", int(timeout.Seconds())),
		"-o", fmt.Sprintf("-p %d -c archive_mode=off", port),
	}
	cmd := exec.CommandContext(ctx, "pg_ctl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// StopPostgres stops Postgres running against dataDir.
func StopPostgres(ctx context.Context, dataDir string) error {
	cmd := exec.CommandContext(ctx, "pg_ctl", "stop", "-D", dataDir, "-m", "fast", "-w")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// WaitRecoveryComplete polls for Postgres to finish recovery: Postgres
// itself deletes recovery.signal once it promotes (recovery_target_action =
// promote), which is a simpler and more portable completion signal than
// authenticating a client connection to run pg_is_in_recovery().
func WaitRecoveryComplete(ctx context.Context, dataDir string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(filepath.Join(dataDir, "recovery.signal")); os.IsNotExist(err) {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return fmt.Errorf("timed out after %s waiting for recovery to complete in %s", timeout, dataDir)
}
