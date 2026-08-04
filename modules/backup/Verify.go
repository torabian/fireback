package backup

import (
	"context"
	"fmt"
	"os"
	"time"
)

// Verify performs a full restore drill: fetches the latest backup into a
// scratch directory, replays every WAL segment available, and confirms
// Postgres reaches a promoted, running state. This is the single
// highest-value reliability check in this whole module - an unverified
// backup strategy is not a backup strategy, since silent corruption in a
// base backup or a gap in WAL archiving otherwise only surfaces the day you
// actually need to restore. Run it on a schedule (e.g. weekly).
//
// scratchDir is wiped before and after, regardless of outcome, so this must
// never be pointed at a real data directory.
func Verify(ctx context.Context, engine *Engine, scratchDir string, port int) error {
	if port == 0 {
		port = 5433
	}

	if err := os.RemoveAll(scratchDir); err != nil {
		return fmt.Errorf("clearing scratch dir: %w", err)
	}
	defer os.RemoveAll(scratchDir)

	backups, err := engine.BackupList(ctx)
	if err != nil {
		return err
	}
	if len(backups) == 0 {
		return fmt.Errorf("no backups found - nothing to verify")
	}

	if err := engine.BackupFetch(ctx, scratchDir, "LATEST"); err != nil {
		return fmt.Errorf("fetching latest backup: %w", err)
	}

	if err := writeRecoverySignal(scratchDir); err != nil {
		return err
	}
	if err := writeBareRestoreCommand(engine, scratchDir); err != nil {
		return err
	}

	if err := StartPostgresAndWait(ctx, scratchDir, port, 5*time.Minute); err != nil {
		return fmt.Errorf("starting postgres for verify: %w", err)
	}
	defer StopPostgres(ctx, scratchDir)

	if err := WaitRecoveryComplete(ctx, scratchDir, 30*time.Minute); err != nil {
		return err
	}

	return nil
}

// CheckResult is Check's monitoring-friendly verdict: OK plus, when false,
// every reason found - callers (backup check's CLI command) turn a false OK
// into a non-zero exit so cron/alerting can key off it.
type CheckResult struct {
	OK      bool
	Reasons []string
}

// Check runs fast, non-destructive health checks suitable for a monitoring
// cron job (unlike Verify, it never starts Postgres or touches disk beyond
// wal-g's own storage reads):
//
//  1. at least one backup exists, and the newest is no older than
//     maxBackupAge - catches a silently broken/uninstalled cron job driving
//     `backup push`.
//  2. `wal-g wal-verify integrity` reports OK - catches gaps in WAL
//     archiving (the classic silent Postgres backup failure: archive_command
//     erroring while Postgres just piles up WAL in pg_wal).
func Check(ctx context.Context, engine *Engine, maxBackupAge time.Duration) (CheckResult, error) {
	result := CheckResult{OK: true}

	backups, err := engine.BackupList(ctx)
	if err != nil {
		return result, err
	}
	if len(backups) == 0 {
		result.OK = false
		result.Reasons = append(result.Reasons, "no backups found in storage")
	} else {
		newest := backups[0].Time
		for _, b := range backups {
			if b.Time.After(newest) {
				newest = b.Time
			}
		}
		if age := time.Since(newest); age > maxBackupAge {
			result.OK = false
			result.Reasons = append(result.Reasons, fmt.Sprintf("newest backup is %s old (limit %s)", age.Round(time.Minute), maxBackupAge))
		}
	}

	var walVerify map[string]any
	if err := engine.RunJSON(ctx, &walVerify, "wal-verify", "integrity", "--json"); err != nil {
		result.OK = false
		result.Reasons = append(result.Reasons, fmt.Sprintf("wal-verify failed to run: %v", err))
	} else if integrity, ok := walVerify["integrity"].(map[string]any); ok {
		if status, _ := integrity["status"].(string); status != "OK" {
			result.OK = false
			result.Reasons = append(result.Reasons, fmt.Sprintf("wal-verify integrity status: %s", status))
		}
	}

	return result, nil
}
