package backup

import (
	"context"
	"fmt"
	"os"
)

// Prune enforces retention by keeping only the newest retainFull full
// backups (plus whatever deltas/WAL wal-g determines are needed to restore
// from the oldest retained one forward), deleting everything older than
// that. Run this right after every successful `backup push` so local disk
// usage stays bounded regardless of how long the system has been running.
func Prune(ctx context.Context, engine *Engine, retainFull int) error {
	if retainFull <= 0 {
		return fmt.Errorf("retainFull must be positive, got %d", retainFull)
	}
	return engine.DeleteRetainFull(ctx, retainFull)
}

// Download produces a self-contained, restorable copy of backupName at
// destPath (created if it doesn't already exist) via Engine.CopyBackup -
// ready to be moved off-box (offsite copy, another environment) and
// restored from directly. Pass withHistory to bring along every WAL
// segment through "now" too, so the copy can be restored to any instant up
// to when this ran, not just to backupName's own completion.
func Download(ctx context.Context, engine *Engine, backupName, destPath string, withHistory bool) error {
	if err := os.MkdirAll(destPath, 0o755); err != nil {
		return err
	}
	return engine.CopyBackup(ctx, backupName, destPath, withHistory)
}
